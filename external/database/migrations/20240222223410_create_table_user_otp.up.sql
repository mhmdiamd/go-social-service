CREATE TABLE user_otp (
  id SERIAL PRIMARY KEY,
  otp CHAR(4) NOT NULL,
  email VARCHAR(255) NOT NULL,
  is_active INT NOT NULL,
  expired_at timestamp,
  created_at timestamp default now(),
  updated_at timestamp default now()
);
