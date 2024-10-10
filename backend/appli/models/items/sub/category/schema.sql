CREATE TABLE IF NOT EXISTS `items_sub_category` (
  `isc_id` int NOT NULL AUTO_INCREMENT,
  `isc_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isc_description` text,
  `isc_picture_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`isc_id`),
  UNIQUE KEY `isc_name` (`isc_name`) USING BTREE
) 