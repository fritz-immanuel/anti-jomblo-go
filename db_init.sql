CREATE DATABASE IF NOT EXISTS anti_jomblo;

CREATE TABLE IF NOT EXISTS schema_migrations (
  version INT DEFAULT 0,
  dirty VARCHAR(5) DEFAULT "",
  PRIMARY KEY (version)
);

TRUNCATE schema_migrations;
INSERT INTO schema_migrations (version, dirty) VALUES (0, 0);

CREATE TABLE IF NOT EXISTS status (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);