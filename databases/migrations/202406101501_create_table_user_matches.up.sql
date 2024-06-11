CREATE TABLE user_matches (
  id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  display_user_id VARCHAR(255) NOT NULL,

  status_id VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255) NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255) NULL,
  PRIMARY KEY (id),
  INDEX index_user_id (user_id),
  INDEX index_display_user_id (display_user_id)
);