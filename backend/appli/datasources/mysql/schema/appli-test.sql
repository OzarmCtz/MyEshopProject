-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.1.0.6537
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for shop-mysql-test
CREATE DATABASE IF NOT EXISTS `shop-mysql-test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `shop-mysql-test`;

-- Dumping structure for table shop-mysql-test.app_settings
CREATE TABLE IF NOT EXISTS `app_settings` (
  `as_id` int NOT NULL AUTO_INCREMENT,
  `as_key` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `as_value` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `as_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `as_last_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`as_id`),
  UNIQUE KEY `as_key_uq` (`as_key`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.discount
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
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.discount_link
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.groups_privileges
CREATE TABLE IF NOT EXISTS `groups_privileges` (
  `gp_id` int NOT NULL AUTO_INCREMENT,
  `gp_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`gp_id`),
  UNIQUE KEY `path` (`gp_path`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.groups_privileges_link
CREATE TABLE IF NOT EXISTS `groups_privileges_link` (
  `gpl_id` int NOT NULL AUTO_INCREMENT,
  `gpl_groups_users_id` int NOT NULL,
  `gpl_groups_privileges_id` int NOT NULL,
  PRIMARY KEY (`gpl_id`),
  UNIQUE KEY `groups_users_id` (`gpl_groups_users_id`,`gpl_groups_privileges_id`) USING BTREE,
  KEY `FK_groups_privileges_link_groups_privileges` (`gpl_groups_privileges_id`),
  CONSTRAINT `FK_groups_privileges_link_groups_privileges` FOREIGN KEY (`gpl_groups_privileges_id`) REFERENCES `groups_privileges` (`gp_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_groups_privileges_link_groups_users` FOREIGN KEY (`gpl_groups_users_id`) REFERENCES `groups_users` (`gu_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=87 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.groups_users
CREATE TABLE IF NOT EXISTS `groups_users` (
  `gu_id` int NOT NULL AUTO_INCREMENT,
  `gu_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `gu_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  PRIMARY KEY (`gu_id`),
  UNIQUE KEY `name` (`gu_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.groups_users_link
CREATE TABLE IF NOT EXISTS `groups_users_link` (
  `gul_id` int NOT NULL AUTO_INCREMENT,
  `gul_user_id` int NOT NULL,
  `gul_group_id` int NOT NULL,
  PRIMARY KEY (`gul_id`),
  UNIQUE KEY `gul_user_id_gul_group_id` (`gul_user_id`,`gul_group_id`) USING BTREE,
  KEY `FK_groups_users_link_groups_users` (`gul_group_id`),
  CONSTRAINT `FK_groups_users_link_groups_users` FOREIGN KEY (`gul_group_id`) REFERENCES `groups_users` (`gu_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK_groups_users_link_users` FOREIGN KEY (`gul_user_id`) REFERENCES `users` (`u_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items
CREATE TABLE IF NOT EXISTS `items` (
  `i_id` int NOT NULL AUTO_INCREMENT,
  `i_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `i_description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `i_price` decimal(10,2) NOT NULL,
  `i_quantity` int DEFAULT '0',
  `i_picture_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `i_file_path` varchar(255) DEFAULT NULL,
  `i_is_disabled` tinyint(1) NOT NULL,
  `i_release_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`i_id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_category
CREATE TABLE IF NOT EXISTS `items_category` (
  `ic_id` int NOT NULL AUTO_INCREMENT,
  `ic_name` varchar(50) NOT NULL,
  `ic_description` text,
  `ic_picture_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`ic_id`),
  UNIQUE KEY `ic_ic_name` (`ic_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_category_link
CREATE TABLE IF NOT EXISTS `items_category_link` (
  `icl_id` int NOT NULL AUTO_INCREMENT,
  `icl_items_sub_category_id` int NOT NULL,
  `icl_items_category_id` int NOT NULL,
  PRIMARY KEY (`icl_id`),
  UNIQUE KEY `icl_uq` (`icl_items_category_id`,`icl_items_sub_category_id`) USING BTREE,
  UNIQUE KEY `uq_icl_items_sub_category` (`icl_items_sub_category_id`),
  CONSTRAINT `FK__items_category` FOREIGN KEY (`icl_items_category_id`) REFERENCES `items_category` (`ic_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK__sub_category_items` FOREIGN KEY (`icl_items_sub_category_id`) REFERENCES `items_sub_category` (`isc_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_purshased
CREATE TABLE IF NOT EXISTS `items_purshased` (
  `ip` int NOT NULL AUTO_INCREMENT,
  `ip_items_id` int DEFAULT NULL,
  `ip_transactions_id` int NOT NULL,
  `p_quantity` int NOT NULL,
  `p_unit_price` int NOT NULL,
  PRIMARY KEY (`ip`),
  KEY `FK__items` (`ip_items_id`),
  KEY `FK__transactions` (`ip_transactions_id`),
  CONSTRAINT `FK__items` FOREIGN KEY (`ip_items_id`) REFERENCES `items` (`i_id`) ON DELETE SET NULL ON UPDATE SET NULL,
  CONSTRAINT `FK__transactions` FOREIGN KEY (`ip_transactions_id`) REFERENCES `transactions` (`t_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_reviews
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
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_sub_category
CREATE TABLE IF NOT EXISTS `items_sub_category` (
  `isc_id` int NOT NULL AUTO_INCREMENT,
  `isc_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isc_description` text,
  `isc_picture_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`isc_id`),
  UNIQUE KEY `isc_name` (`isc_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.items_sub_category_link
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
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.transactions
CREATE TABLE IF NOT EXISTS `transactions` (
  `t_id` int NOT NULL AUTO_INCREMENT,
  `t_user_id` int DEFAULT NULL,
  `t_date` datetime DEFAULT CURRENT_TIMESTAMP,
  `t_delivery_adress` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `t_paiement_method` varchar(255) NOT NULL,
  `t_amount` int NOT NULL,
  PRIMARY KEY (`t_id`),
  KEY `FK__users` (`t_user_id`),
  CONSTRAINT `FK__users` FOREIGN KEY (`t_user_id`) REFERENCES `users` (`u_id`) ON DELETE SET NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.users
CREATE TABLE IF NOT EXISTS `users` (
  `u_id` int NOT NULL AUTO_INCREMENT,
  `u_email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `u_register_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `u_is_disabled` tinyint(1) NOT NULL,
  PRIMARY KEY (`u_id`),
  UNIQUE KEY `email` (`u_email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.users_basket
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
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

-- Dumping structure for table shop-mysql-test.users_wishlist
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
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Data exporting was unselected.

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
