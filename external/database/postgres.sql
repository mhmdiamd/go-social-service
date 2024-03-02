CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE auth (
  id SERIAL PRIMARY KEY,
  public_id VARCHAR(100) NOT NULL,
  name VARCHAR(255),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  gender GENDER,
  no_tlp CHAR(14),
  address TEXT,
  subdistrict_id INT,

  user_otp_public_id VARCHAR(100) REFERENCES user_otp(public_id) ON DELETE CASCADE,

  created_at timestamp default now(),
  updated_at timestamp default now()
)

CREATE TABLE user_otp (
  id SERIAL PRIMARY KEY,
  public_id VARCHAR(100) NOT NULL,
  otp CHAR(4) NOT NULL,
  email VARCHAR(255) NOT NULL,
  is_active INT NOT NULL,
  expired_at timestamp,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

CREATE TABLE category_community (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at timestamp default now(),
  updated_at timestamp default now()
);

CREATE TYPE gender_event_demographics AS ENUM ('male', 'female', 'all');

CREATE TABLE event_demographics (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  gender gender_event_demographics NOT NULL,
  graduation VARCHAR(255),
  start_age INT NOT NULL,
  end_age INT NOT NULL,
  created_at timestamp default now(),
  updated_at timestamp default now()
);


