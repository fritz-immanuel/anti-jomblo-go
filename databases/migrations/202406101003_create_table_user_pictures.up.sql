CREATE TABLE user_pictures (
  id VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  img_url VARCHAR(255) DEFAULT "",
  is_main TINYINT DEFAULT 0,
  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by INT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by INT NULL,
  PRIMARY KEY (id),
  INDEX index_user_id (user_id)
);