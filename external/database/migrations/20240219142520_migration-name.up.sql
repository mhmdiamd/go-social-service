CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE auth (
  id SERIAL PRIMARY KEy,
  public_id VARCHAR(100) NOT NULL,
  name VARCHAR(255),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  gender GENDER,
  no_tlp CHAR(14),
  address TEXT,
  id_subdistrict INT,
  created_at timestamp default now(),
  updated_at timestamp default now()
)

