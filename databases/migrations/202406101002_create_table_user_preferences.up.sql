CREATE TABLE user_preferences (
  id VARCHAR(255) DEFAULT UUID() NOT NULL,
  user_id VARCHAR(255) NOT NULL,

  search_distance TINYINT DEFAULT 50,
  school_name VARCHAR(255) DEFAULT "",
  drinking_frequency_id TINYINT DEFAULT 1,
  smoking_frequency_id TINYINT DEFAULT 1,
  workout_frequency_id TINYINT DEFAULT 1,
  pet_id TINYINT DEFAULT 1,

  communication_style_id TINYINT DEFAULT 1,
  love_language_id TINYINT DEFAULT 1,
  education_level_id TINYINT DEFAULT 1,
  zodiac_id TINYINT DEFAULT 1,
  
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_by INT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by INT NULL,
  PRIMARY KEY (id),
  INDEX index_user_id (user_id)
);