CREATE TABLE users ( 
  id BIGSERIAL PRIMARY KEY, 
  email VARCHAR(80) UNIQUE, 
  password VARCHAR(80), 
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP
);

CREATE TABLE user_profiles( 
  id BIGSERIAL PRIMARY KEY, 
  user_id BIGINT REFERENCES users(id) ON DELETE CASCADE, 
  fullname VARCHAR(80),
  picture TEXT,
  created_at TIMESTAMP DEFAULT NOW(), 
  update_at TIMESTAMP DEFAULT NOW()
);
