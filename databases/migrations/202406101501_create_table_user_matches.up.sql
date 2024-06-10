CREATE TABLE user_matches (
  id VARCHAR(255) DEFAULT UUID() NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  display_user_id VARCHAR(255) NOT NULL,

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by INT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by INT NULL,
  PRIMARY KEY (id),
  INDEX index_user_id (user_id)
);