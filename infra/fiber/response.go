package infrafiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mhmdiamd/go-social-service/infra/response"
)

type Response struct {
  HttpCode int `json:"-"`
  Success bool `json:"success"`
  Message string `json:"message"`
  Payload interface{} `json:"payload"`
  Query interface{} `json:"query"`
  Error string `json:"error,omitempty"`
  ErrorCode string `json:"error_code,omitempty"`
}

func (r Response) Send(ctx *fiber.Ctx) error {
  return ctx.Status(r.HttpCode).JSON(r)
}

func NewResponse(params ...func(*Response) *Response) Response {
  resp := Response{
    Success: true,
  }

  for _, param := range params {
    param(&resp)
  }

  return resp
}

func WithPayload(payload interface{}) func(*Response) *Response {
  return func(r *Response) *Response {
    r.Payload = payload
    return r
  }
}

func WithMessage(message string) func(*Response) *Response {
  return func(r *Response) *Response {
    r.Message = message
    return r
  }
}

func WithHttpCode(httpCode int) func(*Response) *Response {
  return func(r *Response) *Response {
    r.HttpCode = httpCode
    return r
  }
}

func WithQuery(query interface{}) func(*Response) *Response {
  return func(r *Response) *Response {
    r.Query = query
    return r
  }
}

func WithError(err error) func(*Response) *Response {
  return func(r *Response) *Response {
    r.Success = false

    myErr, ok := err.(response.Error)

    if !ok {
      myErr = response.ErrorGeneral
    }

    r.Error = myErr.Message
    r.ErrorCode = myErr.Code
    r.HttpCode = myErr.HttpCode

    return r
  }
}


