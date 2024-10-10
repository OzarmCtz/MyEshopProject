CREATE TABLE IF NOT EXISTS `items_category_link` (
  `icl_id` int NOT NULL AUTO_INCREMENT,
  `icl_items_sub_category_id` int NOT NULL,
  `icl_items_category_id` int NOT NULL,
  PRIMARY KEY (`icl_id`),
  UNIQUE KEY `icl_uq` (`icl_items_category_id`,`icl_items_sub_category_id`) USING BTREE,
  UNIQUE KEY `uq_icl_items_sub_category` (`icl_items_sub_category_id`),
  CONSTRAINT `FK__items_category` FOREIGN KEY (`icl_items_category_id`) REFERENCES `items_category` (`ic_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK__sub_category_items` FOREIGN KEY (`icl_items_sub_category_id`) REFERENCES `items_sub_category` (`isc_id`) ON DELETE CASCADE ON UPDATE CASCADE
)