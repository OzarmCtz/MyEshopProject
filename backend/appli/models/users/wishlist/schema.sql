CREATE TABLE IF NOT EXISTS `users_wishlist` (
  `wl_id` int NOT NULL AUTO_INCREMENT,
  `wl_user_id` int NOT NULL,
  `wl_items_id` int DEFAULT NULL,
  `wl_times_added` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`wl_id`),
  UNIQUE KEY `UQ_wl_user_id_wl_items_id` (`wl_user_id`,`wl_items_id`),
  KEY `FK_wish_list_items` (`wl_items_id`),
  CONSTRAINT `FK_wish_list_items` FOREIGN KEY (`wl_items_id`) REFERENCES `items` (`i_id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `FK_wish_list_users` FOREIGN KEY (`wl_user_id`) REFERENCES `users` (`u_id`) ON DELETE CASCADE ON UPDATE CASCADE
)