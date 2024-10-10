CREATE TABLE IF NOT EXISTS `discount` (
  `d_id` int NOT NULL AUTO_INCREMENT,
  `d_code` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `d_description` text,
  `d_start_time` timestamp NULL DEFAULT NULL,
  `d_end_time` timestamp NULL DEFAULT NULL,
  `d_zone_time` varchar(10) DEFAULT NULL,
  `d_is_disabled` tinyint(1) NOT NULL DEFAULT '0',
  `d_price_type` varchar(50) NOT NULL,
  `d_value` int NOT NULL,
  PRIMARY KEY (`d_id`),
  UNIQUE KEY `uq_dcode` (`d_code`)
) 