CREATE TABLE IF NOT EXISTS `items_category` (
  `ic_id` int NOT NULL AUTO_INCREMENT,
  `ic_name` varchar(50) NOT NULL,
  `ic_description` text,
  `ic_picture_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ic_id`),
  UNIQUE KEY `ic_ic_name` (`ic_name`) USING BTREE
)