CREATE TABLE IF NOT EXISTS `groups_users` (
  `gu_id` int NOT NULL AUTO_INCREMENT,
  `gu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `gu_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`gu_id`),
  UNIQUE KEY `name` (`gu_name`) USING BTREE
)