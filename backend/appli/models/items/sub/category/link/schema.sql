CREATE TABLE IF NOT EXISTS `items_sub_category_link` (
  `iscl_id` int NOT NULL AUTO_INCREMENT,
  `iscl_items_id` int NOT NULL,
  `iscl_sub_category_id` int NOT NULL,
  PRIMARY KEY (`iscl_id`) USING BTREE,
  UNIQUE KEY `sub_category_id` (`iscl_items_id`,`iscl_sub_category_id`) USING BTREE,
  UNIQUE KEY `iscl_items_id` (`iscl_items_id`),
  KEY `FK__items_sub_category` (`iscl_sub_category_id`),
  CONSTRAINT `FK__items_items` FOREIGN KEY (`iscl_items_id`) REFERENCES `items` (`i_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK__items_sub_category` FOREIGN KEY (`iscl_sub_category_id`) REFERENCES `items_sub_category` (`isc_id`) ON DELETE CASCADE ON UPDATE CASCADE
)