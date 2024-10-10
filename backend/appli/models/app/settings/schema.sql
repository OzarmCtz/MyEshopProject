CREATE TABLE IF NOT EXISTS `app_settings` (
  `as_id` int NOT NULL AUTO_INCREMENT,
  `as_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `as_value` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `as_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `as_last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`as_id`),
  UNIQUE KEY `as_key_uq` (`as_key`)
) 