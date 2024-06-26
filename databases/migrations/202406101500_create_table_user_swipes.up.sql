CREATE TABLE user_swipes (
  id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  display_user_id VARCHAR(255) NOT NULL,
  action_id TINYINT DEFAULT 0,

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255) NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255) NULL,
  PRIMARY KEY (id),
  INDEX index_user_id (user_id)
);