CREATE DATABASE IF NOT EXISTS anti_jomblo;

CREATE TABLE IF NOT EXISTS schema_migrations (
  version INT DEFAULT 0,
  dirty TINYINT DEFAULT 0
);

TRUNCATE schema_migrations;
INSERT INTO schema_migrations (version, dirty) VALUES (0, 0);