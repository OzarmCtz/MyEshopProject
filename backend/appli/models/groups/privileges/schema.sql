CREATE TABLE IF NOT EXISTS `groups_privileges` (
  `gp_id` int NOT NULL AUTO_INCREMENT,
  `gp_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`gp_id`),
  UNIQUE KEY `path` (`gp_path`) USING BTREE
) 