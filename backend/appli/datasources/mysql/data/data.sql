-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               8.0.30 - MySQL Community Server - GPL
-- Server OS:                    Win64
-- HeidiSQL Version:             12.8.0.6908
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- Dumping data for table shop-mysql.app_settings: ~0 rows (approximately)

-- Dumping data for table shop-mysql.discount: ~0 rows (approximately)

-- Dumping data for table shop-mysql.discount_link: ~0 rows (approximately)

-- Dumping data for table shop-mysql.groups_privileges: ~46 rows (approximately)
INSERT INTO `groups_privileges` (`gp_id`, `gp_path`) VALUES
	(37, 'APP/SETTINGS_CREATE'),
	(39, 'APP/SETTINGS_DELETE'),
	(36, 'APP/SETTINGS_READ'),
	(38, 'APP/SETTINGS_UPDATE'),
	(52, 'DASHBOARD_READ'),
	(41, 'DISCOUNT_CREATE'),
	(43, 'DISCOUNT_DELETE'),
	(40, 'DISCOUNT_READ'),
	(42, 'DISCOUNT_UPDATE'),
	(45, 'DISCOUNT/LINK_CREATE'),
	(47, 'DISCOUNT/LINK_DELETE'),
	(44, 'DISCOUNT/LINK_READ'),
	(46, 'DISCOUNT/LINK_UPDATE'),
	(53, 'GROUPS/USERS_READ'),
	(8, 'GROUPS/USERS/LINK_READ'),
	(7, 'GROUPS/USERS/LINK_UPDATE'),
	(11, 'ITEMS_CREATE'),
	(12, 'ITEMS_DELETE'),
	(9, 'ITEMS_READ'),
	(10, 'ITEMS_UPDATE'),
	(14, 'ITEMS/CATEGORY_CREATE'),
	(16, 'ITEMS/CATEGORY_DELETE'),
	(49, 'ITEMS/CATEGORY_READ'),
	(15, 'ITEMS/CATEGORY_UPDATE'),
	(22, 'ITEMS/CATEGORY/LINK_CREATE'),
	(21, 'ITEMS/CATEGORY/LINK_READ'),
	(23, 'ITEMS/CATEGORY/LINK_UPDATE'),
	(34, 'ITEMS/REVIEWS_CREATE'),
	(35, 'ITEMS/REVIEWS_DELETE'),
	(33, 'ITEMS/REVIEWS_READ'),
	(17, 'ITEMS/SUB/CATEGORY_CREATE'),
	(20, 'ITEMS/SUB/CATEGORY_DELETE'),
	(48, 'ITEMS/SUB/CATEGORY_READ'),
	(18, 'ITEMS/SUB/CATEGORY_UPDATE'),
	(24, 'ITEMS/SUB/CATEGORY/LINK_CREATE'),
	(25, 'ITEMS/SUB/CATEGORY/LINK_READ'),
	(26, 'ITEMS/SUB/CATEGORY/LINK_UPDATE'),
	(4, 'USERS_DELETE'),
	(3, 'USERS_READ'),
	(5, 'USERS_UPDATE'),
	(31, 'USERS/BASKET_CREATE'),
	(32, 'USERS/BASKET_DELETE'),
	(30, 'USERS/BASKET_READ'),
	(29, 'USERS/WISHLIST_CREATE'),
	(28, 'USERS/WISHLIST_DELETE'),
	(27, 'USERS/WISHLIST_READ');

-- Dumping data for table shop-mysql.groups_privileges_link: ~81 rows (approximately)
INSERT INTO `groups_privileges_link` (`gpl_id`, `gpl_groups_users_id`, `gpl_groups_privileges_id`) VALUES
	(7, 9, 3),
	(9, 9, 4),
	(10, 9, 5),
	(17, 9, 7),
	(21, 9, 8),
	(23, 9, 9),
	(25, 9, 10),
	(27, 9, 11),
	(30, 9, 12),
	(34, 9, 14),
	(37, 9, 15),
	(38, 9, 16),
	(40, 9, 17),
	(42, 9, 18),
	(44, 9, 20),
	(47, 9, 21),
	(49, 9, 22),
	(51, 9, 23),
	(53, 9, 24),
	(55, 9, 25),
	(56, 9, 26),
	(65, 9, 33),
	(67, 9, 35),
	(73, 9, 36),
	(70, 9, 37),
	(72, 9, 38),
	(71, 9, 39),
	(75, 9, 40),
	(76, 9, 41),
	(77, 9, 42),
	(78, 9, 43),
	(82, 9, 44),
	(84, 9, 45),
	(85, 9, 46),
	(86, 9, 47),
	(88, 9, 48),
	(90, 9, 49),
	(91, 9, 52),
	(94, 9, 53),
	(8, 10, 3),
	(16, 10, 4),
	(14, 10, 5),
	(18, 10, 7),
	(19, 10, 8),
	(22, 10, 9),
	(26, 10, 10),
	(28, 10, 11),
	(29, 10, 12),
	(33, 10, 14),
	(35, 10, 15),
	(36, 10, 16),
	(39, 10, 17),
	(43, 10, 18),
	(45, 10, 20),
	(46, 10, 21),
	(48, 10, 22),
	(50, 10, 23),
	(52, 10, 24),
	(54, 10, 25),
	(57, 10, 26),
	(64, 10, 33),
	(68, 10, 35),
	(74, 10, 36),
	(79, 10, 40),
	(83, 10, 44),
	(87, 10, 48),
	(89, 10, 49),
	(92, 10, 52),
	(93, 10, 53),
	(15, 11, 3),
	(12, 11, 4),
	(13, 11, 5),
	(24, 11, 9),
	(58, 11, 27),
	(59, 11, 28),
	(60, 11, 29),
	(61, 11, 30),
	(62, 11, 31),
	(63, 11, 32),
	(66, 11, 34),
	(69, 11, 35);

-- Dumping data for table shop-mysql.groups_users: ~3 rows (approximately)
INSERT INTO `groups_users` (`gu_id`, `gu_name`, `gu_description`) VALUES
	(9, 'SUPERADMIN_STATUS', NULL),
	(10, 'ADMIN_STATUS', NULL),
	(11, 'CLIENT_STATUS', NULL);

-- Dumping data for table shop-mysql.groups_users_link: ~0 rows (approximately)

-- Dumping data for table shop-mysql.items: ~1 rows (approximately)
INSERT INTO `items` (`i_id`, `i_title`, `i_description`, `i_price`, `i_quantity`, `i_picture_url`, `i_file_path`, `i_is_disabled`, `i_release_date`) VALUES
	(36, 'Form Contact Template', 'An simple form contact using HTML and css', 20.00, 100, 'https://picture.com/kdlM', 'https://privnote.com/jdjqQ', 1, '2024-06-04 19:44:47');

-- Dumping data for table shop-mysql.items_category: ~2 rows (approximately)
INSERT INTO `items_category` (`ic_id`, `ic_name`, `ic_description`, `ic_picture_url`) VALUES
	(82, 'Software', 'Software Category', 'http://random.pucture.com'),
	(84, 'Online Course', 'Online Course Category', 'https://picture.com/kdlM');

-- Dumping data for table shop-mysql.items_category_link: ~3 rows (approximately)
INSERT INTO `items_category_link` (`icl_id`, `icl_items_sub_category_id`, `icl_items_category_id`) VALUES
	(33, 50, 82),
	(34, 51, 82),
	(37, 54, 84);

-- Dumping data for table shop-mysql.items_purshased: ~0 rows (approximately)

-- Dumping data for table shop-mysql.items_reviews: ~0 rows (approximately)

-- Dumping data for table shop-mysql.items_sub_category: ~3 rows (approximately)
INSERT INTO `items_sub_category` (`isc_id`, `isc_name`, `isc_description`, `isc_picture_url`) VALUES
	(50, 'Productivity Applications', 'Software designed to help users perform tasks efficiently, such as office suites, project management tools, and note-taking apps.', 'http://random.pucture.com'),
	(51, 'Security Software', 'Applications that protect computers and networks from viruses, malware, and unauthorized access, including antivirus programs, anti-malware tools, and VPNs.', 'http://random.pucture.com'),
	(54, 'HTML Form', 'HTML Form sub category', 'https://picture.com/kdlM');

-- Dumping data for table shop-mysql.items_sub_category_link: ~1 rows (approximately)
INSERT INTO `items_sub_category_link` (`iscl_id`, `iscl_items_id`, `iscl_sub_category_id`) VALUES
	(34, 36, 50);

-- Dumping data for table shop-mysql.transactions: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users_basket: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users_wishlist: ~0 rows (approximately)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
