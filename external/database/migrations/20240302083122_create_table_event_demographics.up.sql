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
