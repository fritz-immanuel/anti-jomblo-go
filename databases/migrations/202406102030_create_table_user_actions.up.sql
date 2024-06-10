CREATE TABLE user_actions (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  user_id VARCHAR(255) DEFAULT NULL,
  table_name VARCHAR(200) DEFAULT NULL,
  action VARCHAR(100) DEFAULT NULL,
  action_value INT DEFAULT 0,
  ref_id VARCHAR(255) DEFAULT '0',
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);