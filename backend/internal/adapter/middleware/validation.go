package middleware

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/schema"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
)

var (
	VALIDATED_BODY   = "validated_body"
	VALIDATED_FORM   = "validated_form"
	VALIDATED_PARAMS = "validated_params"
	VALIDATED_QUERY  = "validated_query"
	validate         = validator.New()
	decoder          = schema.NewDecoder()
)

func BodyValidation(body any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := reflect.New(reflect.TypeOf(body).Elem()).Interface()
		if err := c.BodyParser(req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, error_usecase.InvalidRequestBody)
		}
		if err := validate.Struct(req); err != nil {
			if verrs, ok := err.(validator.ValidationErrors); ok {
				messages := make([]string, 0, len(verrs))
				for _, verr := range verrs {
					messages = append(messages, validationErrorToText(verr))
				}
				return fiber.NewError(fiber.StatusBadRequest, strings.Join(messages, ","))
			}
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		c.Locals(VALIDATED_BODY, req)
		return c.Next()
	}
}

func ParamsValidation(params any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := reflect.New(reflect.TypeOf(params).Elem()).Interface()
		if err := c.ParamsParser(req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, error_usecase.InvalidParams)
		}
		if err := validate.Struct(req); err != nil {
			if verrs, ok := err.(validator.ValidationErrors); ok {
				messages := make([]string, 0, len(verrs))
				for _, verr := range verrs {
					messages = append(messages, validationErrorToText(verr))
				}
				return fiber.NewError(fiber.StatusBadRequest, strings.Join(messages, ","))
			}
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		c.Locals(VALIDATED_PARAMS, req)
		return c.Next()
	}
}

func QueryValidation(query any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := reflect.New(reflect.TypeOf(query).Elem()).Interface()
		if err := c.QueryParser(req); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, error_usecase.InvalidQueryParams)
		}
		if err := validate.Struct(req); err != nil {
			if verrs, ok := err.(validator.ValidationErrors); ok {
				messages := make([]string, 0, len(verrs))
				for _, verr := range verrs {
					messages = append(messages, validationErrorToText(verr))
				}
				return fiber.NewError(fiber.StatusBadRequest, strings.Join(messages, ","))
			}
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		c.Locals(VALIDATED_QUERY, req)
		return c.Next()
	}
}

func FormValidation(form any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := reflect.New(reflect.TypeOf(form).Elem()).Interface()
		multipartForm, err := c.MultipartForm()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, error_usecase.InvalidFormData)
		}
		if err := decoder.Decode(req, multipartForm.Value); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, error_usecase.InvalidFormData)
		}
		if err := validate.Struct(req); err != nil {
			if verrs, ok := err.(validator.ValidationErrors); ok {
				messages := make([]string, 0, len(verrs))
				for _, verr := range verrs {
					messages = append(messages, validationErrorToText(verr))
				}
				return fiber.NewError(fiber.StatusBadRequest, strings.Join(messages, ","))
			}
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		c.Locals(VALIDATED_FORM, req)
		return c.Next()
	}
}

func validationErrorToText(fe validator.FieldError) string {
	field := strings.ToLower(fe.Field())
	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		if fe.Kind() == reflect.Slice {
			return field + " must have at least " + fe.Param() + " item(s)"
		}
		return field + " must be at least " + fe.Param() + " characters"
	case "len":
		return field + " must be exactly " + fe.Param() + " characters"
	case "gte":
		return field + " must be greater than or equal to " + fe.Param()
	case "lte":
		return field + " must be less than or equal to " + fe.Param()
	case "gt":
		return field + " must be greater than " + fe.Param()
	case "lt":
		return field + " must be less than " + fe.Param()
	case "oneof":
		return field + " must be one of [" + fe.Param() + "]"
	case "numeric":
		return field + " must be a number"
	}
	return field + " is invalid"
}
