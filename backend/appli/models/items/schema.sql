CREATE TABLE IF NOT EXISTS `items` (
  `i_id` int NOT NULL AUTO_INCREMENT,
  `i_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `i_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `i_price` decimal(10,2) NOT NULL,
  `i_quantity` int DEFAULT '0',
  `i_picture_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `i_file_path` varchar(255) DEFAULT NULL,
  `i_is_disabled` tinyint(1) NOT NULL,
  `i_release_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`i_id`)
) 