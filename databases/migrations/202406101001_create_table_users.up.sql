CREATE TABLE users (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  country_calling_code VARCHAR(255) NOT NULL,
  phone_number VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,

  gender_id TINYINT DEFAULT 1,
  birth_date DATE DEFAULT (CURRENT_DATE),
  height TINYINT unsigned DEFAULT 0,
  about_me LONGTEXT DEFAULT NULL,

  status_id VARCHAR(255) DEFAULT "1",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by VARCHAR(255) NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by VARCHAR(255) NULL,
  PRIMARY KEY (id),
  INDEX index_email (email)
);