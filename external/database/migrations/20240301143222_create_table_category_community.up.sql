CREATE TABLE category_cummunity (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created_at timestamp default now(),
  updated_at timestamp default now()
);
