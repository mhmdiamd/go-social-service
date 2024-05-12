CREATE TABLE events (
  id SERIAL PRIMARY KEY,
  event_demographics_id INT NOT NULL, 
  name varchar(255) NOT NULL,         
  description varchar(100),         
  address TEXT ,         
  start_at timestamp NOT NULL,
  end_at timestamp NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  event_demographic_snapshot jsonb
);
