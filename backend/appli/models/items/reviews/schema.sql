CREATE TABLE IF NOT EXISTS `items_reviews` (
  `ir_id` int NOT NULL AUTO_INCREMENT,
  `ir_user_id` int NOT NULL,
  `ir_items_id` int NOT NULL,
  `ir_comments` text NOT NULL,
  `ir_stars` int NOT NULL DEFAULT '5',
  PRIMARY KEY (`ir_id`),
  KEY `FK_items_reviews_items` (`ir_items_id`),
  KEY `FK_items_reviews_users` (`ir_user_id`),
  CONSTRAINT `FK_items_reviews_items` FOREIGN KEY (`ir_items_id`) REFERENCES `items` (`i_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_items_reviews_users` FOREIGN KEY (`ir_user_id`) REFERENCES `users` (`u_id`) ON DELETE CASCADE ON UPDATE CASCADE
)