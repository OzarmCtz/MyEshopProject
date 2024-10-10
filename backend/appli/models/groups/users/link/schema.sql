CREATE TABLE IF NOT EXISTS `groups_users_link` (
  `gul_id` int NOT NULL AUTO_INCREMENT,
  `gul_user_id` int NOT NULL,
  `gul_group_id` int NOT NULL,
  PRIMARY KEY (`gul_id`),
  UNIQUE KEY `gul_user_id_gul_group_id` (`gul_user_id`,`gul_group_id`) USING BTREE,
  KEY `FK_groups_users_link_groups_users` (`gul_group_id`),
  CONSTRAINT `FK_groups_users_link_groups_users` FOREIGN KEY (`gul_group_id`) REFERENCES `groups_users` (`gu_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_groups_users_link_users` FOREIGN KEY (`gul_user_id`) REFERENCES `users` (`u_id`) ON DELETE CASCADE ON UPDATE CASCADE
)