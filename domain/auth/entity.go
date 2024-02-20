package auth

import "time"

type Gender string

const (
  GENDER_Male Gender = "male"
  GENDER_Female Gender = "female"
)

type AuthEntity struct {
  Id int `db:"id"`
  PublicId string `db:"public_id"`
  Nik string `db:"nik"`
  Name string `db:"name"`
  Email string `db:"email"`
  Password string `db:"password"`
  Gender Gender `db:"gender"`

  CreatedAt time.Time `db:"created_at"`
  UpdatedAt time.Time `db:"updated_at"`
}
