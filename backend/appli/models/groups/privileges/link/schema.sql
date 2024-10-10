CREATE TABLE IF NOT EXISTS `groups_privileges_link` (
  `gpl_id` int NOT NULL AUTO_INCREMENT,
  `gpl_groups_users_id` int NOT NULL,
  `gpl_groups_privileges_id` int NOT NULL,
  PRIMARY KEY (`gpl_id`),
  UNIQUE KEY `groups_users_id` (`gpl_groups_users_id`,`gpl_groups_privileges_id`) USING BTREE,
  KEY `FK_groups_privileges_link_groups_privileges` (`gpl_groups_privileges_id`),
  CONSTRAINT `FK_groups_privileges_link_groups_privileges` FOREIGN KEY (`gpl_groups_privileges_id`) REFERENCES `groups_privileges` (`gp_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_groups_privileges_link_groups_users` FOREIGN KEY (`gpl_groups_users_id`) REFERENCES `groups_users` (`gu_id`) ON DELETE CASCADE ON UPDATE CASCADE
)