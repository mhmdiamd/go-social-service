package infrafiber

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/NooBeeID/go-logging/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mhmdiamd/go-social-service/external/jwt"
	"github.com/mhmdiamd/go-social-service/infra/response"
	"github.com/mhmdiamd/go-social-service/internal/config"
	infralog "github.com/mhmdiamd/go-social-service/internal/log"
)
 
func LogTrace() fiber.Handler {
  return func (c *fiber.Ctx) error {
    ctx := c.UserContext()

    // Initiate Structure
    now := time.Now()
    traceId := uuid.New()
    c.Set("X-Trace-ID", traceId.String())

    data := map[logger.LogKey]interface{}{
      logger.TRACER_ID : traceId,
      logger.METHOD: c.Route().Method,
      logger.PATH: string(c.Context().URI().Path()),
    }

    ctx = context.WithValue(ctx, logger.DATA, data)

    // get Request
    infralog.Log.Infof(ctx, "incoming request")

    c.SetUserContext(ctx)
    err := c.Next()

    // Finish request
    data[logger.RESPONSE_TIME] = time.Since(now).Milliseconds()
    data[logger.RESPONSE_TYPE] = "ms"

    httpStatusCode := c.Response().Header.StatusCode()
    if httpStatusCode >= 200 && httpStatusCode <= 299 {
      ctx = context.WithValue(ctx, logger.DATA, data)
      infralog.Log.Infof(ctx, "success")
    }else {
      respBody := c.Response().Body()
      data["response_body"] = fmt.Sprintf("%s", respBody)

      ctx = context.WithValue(ctx, logger.DATA, data)
      infralog.Log.Errorf(ctx, "error")
    }

    return err
  }
}

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("authorization")

		if authorization == "" {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		bearer := strings.Split(authorization, "Bearer ")
		if len(bearer) != 2 {
			log.Println("Invalid Token")
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		// Decrypt token
		token := bearer[1]
		publicId, err := jwt.ValidateToken(token, config.Cfg.App.Encryption.JWTSecret)
		if err != nil {
			return NewResponse(
				WithError(response.ErrorUnauthorized),
			).Send(c)
		}

		c.Locals("PUBLIC_ID", publicId)

		return c.Next()
	}
}
