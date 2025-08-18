package middleware

import (
	"errors"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/little-tonii/gofiber-base/internal/domain/persistence"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
	"gorm.io/gorm"
)

var (
	USER   = "user"
	CLAIMS = "claims"
)

var (
	ACCESS_TOKEN_KEY_ID   = "access"
	REFRESH_TOKEN_KEY_ID  = "refresh"
	VERIFY_TOKEN_KEY_ID   = "verify"
	PASSWORD_TOKEN_KEY_ID = "password"
)

var (
	ACCESS_TOKEN_SECRET   = "access_token_secret"
	REFRESH_TOKEN_SECRET  = "refresh_token_secret"
	VERIFY_TOKEN_SECRET   = "verify_token_secret"
	PASSWORD_TOKEN_SECRET = "password_token_secret"
)

type Claims struct {
	UserId       string `json:"user_id"`
	UserEmail    string `json:"user_email"`
	TokenVersion int64  `json:"token_version"`
	UserRole     string `json:"user_role"`
	Kid          string `json:"kid"`
	jwt.RegisteredClaims
}

type AuthGuardConfig struct {
	UserPersis persistence.UserPersistence
	AllowPaths []string
	Secrets    map[string]string
}

func AuthGuard(config *AuthGuardConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			path := c.Path()
			for _, pattern := range config.AllowPaths {
				if path == pattern {
					return true
				}
				if strings.HasSuffix(pattern, "/*") {
					base := strings.TrimSuffix(pattern, "/*")
					if after, ok := strings.CutPrefix(path, base+"/"); ok {
						remainingPath := after
						if !strings.Contains(remainingPath, "/") {
							return true
						}
					}
				}
				if strings.HasSuffix(pattern, "/**") {
					base := strings.TrimSuffix(pattern, "/**")
					if strings.HasPrefix(path, base+"/") || path == base {
						return true
					}
				}
			}
			return false
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Locals(USER).(*jwt.Token)
			claims := token.Claims.(*Claims)
			userId, err := uuid.Parse(claims.UserId)
			if err != nil {
				return fiber.NewError(fiber.StatusUnauthorized, error_usecase.Unauthorized)
			}
			_, err = config.UserPersis.FindById(c.Context(), userId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return fiber.NewError(fiber.StatusUnauthorized, error_usecase.Unauthorized)
				}
				return fiber.NewError(fiber.StatusInternalServerError, error_usecase.InternalServerError)
			}
			// if user.TokenVersion != claims.TokenVersion {
			// 	return fiber.NewError(fiber.StatusUnauthorized, error_usecase.Unauthorized)
			// }
			c.Locals(CLAIMS, claims)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, error_usecase.Unauthorized)
		},
		SigningKeys: map[string]jwtware.SigningKey{
			ACCESS_TOKEN_KEY_ID:   {Key: []byte(config.Secrets[ACCESS_TOKEN_SECRET])},
			REFRESH_TOKEN_KEY_ID:  {Key: []byte(config.Secrets[REFRESH_TOKEN_SECRET])},
			VERIFY_TOKEN_KEY_ID:   {Key: []byte(config.Secrets[VERIFY_TOKEN_SECRET])},
			PASSWORD_TOKEN_KEY_ID: {Key: []byte(config.Secrets[PASSWORD_TOKEN_SECRET])},
		},
		ContextKey:         USER,
		Claims:             &Claims{},
		TokenLookup:        "header:Authorization",
		TokenProcessorFunc: nil,
		AuthScheme:         "Bearer",
	})
}
