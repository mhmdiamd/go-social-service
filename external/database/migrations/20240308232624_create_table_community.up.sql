CREATE TABLE community (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  descrption TEXT ,
  logo VARCHAR(255),
  categories_id INT,
  external_categories VARCHAR(255),
  created_at timestamp default now(),
  updated_at timestamp default now()
);
