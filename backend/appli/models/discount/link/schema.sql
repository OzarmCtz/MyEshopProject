CREATE TABLE IF NOT EXISTS `discount_link` (
  `dl_id` int NOT NULL AUTO_INCREMENT,
  `dl_discount_id` int NOT NULL,
  `dl_items_id` int DEFAULT NULL,
  `dl_items_sub_category` int DEFAULT NULL,
  `dl_items_category` int DEFAULT NULL,
  PRIMARY KEY (`dl_id`),
  KEY `FK__discount` (`dl_discount_id`),
  KEY `FK_items_discount` (`dl_items_id`),
  KEY `FK__sub_category_items_discount` (`dl_items_sub_category`),
  KEY `FK__category_items_discount` (`dl_items_category`),
  CONSTRAINT `FK__category_items_discount` FOREIGN KEY (`dl_items_category`) REFERENCES `items_category` (`ic_id`) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT `FK__discount` FOREIGN KEY (`dl_discount_id`) REFERENCES `discount` (`d_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK__sub_category_items_discount` FOREIGN KEY (`dl_items_sub_category`) REFERENCES `items_sub_category` (`isc_id`) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT `FK_items_discount` FOREIGN KEY (`dl_items_id`) REFERENCES `items` (`i_id`) ON DELETE SET NULL ON UPDATE SET NULL
) 