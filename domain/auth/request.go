package auth

type RegisterRequestPayload struct {
  Email string
  Password string
}

type LoginRequestPayload struct {
  Email string
  Password string
}
