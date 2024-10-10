CREATE TABLE IF NOT EXISTS `users` (
  `u_id` int NOT NULL AUTO_INCREMENT,
  `u_email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_register_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `u_is_disabled` tinyint(1) NOT NULL,
  PRIMARY KEY (`u_id`),
  UNIQUE KEY `email` (`u_email`) USING BTREE
)