package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
)

type TimeoutConfig struct {
	ProcessTimeout time.Duration
	SkipPaths      []string
}

func Timeout(config *TimeoutConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		path := c.Path()
		for _, pattern := range config.SkipPaths {
			if path == pattern {
				return c.Next()
			}
			if strings.HasSuffix(pattern, "/*") {
				base := strings.TrimSuffix(pattern, "/*")
				if after, ok := strings.CutPrefix(path, base+"/"); ok {
					remainingPath := after
					if !strings.Contains(remainingPath, "/") {
						return c.Next()
					}
				}
			}
			if strings.HasSuffix(pattern, "/**") {
				base := strings.TrimSuffix(pattern, "/**")
				if strings.HasPrefix(path, base+"/") || path == base {
					return c.Next()
				}
			}
		}
		ctx, cancel := context.WithTimeout(c.UserContext(), config.ProcessTimeout)
		defer cancel()
		c.SetUserContext(ctx)
		done := make(chan error, 1)
		go func() {
			done <- c.Next()
		}()
		select {
		case <-ctx.Done():
			return fiber.NewError(fiber.StatusRequestTimeout, error_usecase.ProcessTimeout)
		case err := <-done:
			return err
		}
	}
}

type TimeoutWrapperConfig[T any] struct {
	FiberCtx       *fiber.Ctx
	ExpectedStatus int
	Fn             func(ctx context.Context) (*T, error)
}

func TimeoutWrapper[T any](config *TimeoutWrapperConfig[T]) error {
	ctx := config.FiberCtx.UserContext()
	respChannel := make(chan *T, 1)
	errChannel := make(chan error, 1)
	go func(ctx context.Context) {
		resp, err := config.Fn(ctx)
		if err != nil {
			errChannel <- err
			return
		}
		respChannel <- resp
	}(ctx)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChannel:
		return err
	case resp := <-respChannel:
		return config.FiberCtx.Status(config.ExpectedStatus).JSON(resp)
	}
}
