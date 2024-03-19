CREATE TYPE community_member_role AS ENUM ('member', 'admin', 'owner');

CREATE TABLE community_members (
  id SERIAL PRIMARY KEY,
  user_public_id VARCHAR(255) NOT NULL,
  community_id INT NOT NULL,
  role community_member_role NOT NULL,
  nik VARCHAR(255),
  photoKTP VARCHAR(255),
  is_active INT DEFAULT 0,
  created_at timestamp default now(),
  updated_at timestamp default now()
);
