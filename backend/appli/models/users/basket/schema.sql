CREATE TABLE IF NOT EXISTS `users_basket` (
  `ub_id` int NOT NULL AUTO_INCREMENT,
  `ub_user_id` int NOT NULL,
  `ub_items_id` int NOT NULL,
  `ub_time_added` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `ub_quantity` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`ub_id`),
  UNIQUE KEY `UQ_ub_user_id_ub_items_id` (`ub_user_id`,`ub_items_id`),
  KEY `FK_users_basket_items` (`ub_items_id`),
  CONSTRAINT `FK_users_basket_items` FOREIGN KEY (`ub_items_id`) REFERENCES `items` (`i_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_users_basket_users` FOREIGN KEY (`ub_user_id`) REFERENCES `users` (`u_id`) ON DELETE CASCADE ON UPDATE CASCADE
) 