CREATE TYPE EVENT_POSITION AS ENUM ('admin', 'staff', 'member');

CREATE TABLE event_commite (
  id SERIAL PRIMARY KEY,
  user_public_id varchar(100) NOT NULL,         
  event_public_id varchar(100) NOT NULL,         
  position EVENT_POSITION NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  user_snapshot jsonb
);
