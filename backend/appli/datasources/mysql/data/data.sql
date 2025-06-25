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
(1, 'ZenTask - Task Manager App', 'Application de gestion de tâches simple et intuitive pour Mac/PC', 29.99, 1000, 'https://example.com/images/zentask.png', 'https://downloads.example.com/zentask.zip', 0, '2024-06-01 10:00:00'),
(2, 'CyberShield Pro VPN', 'Solution de VPN hautement sécurisée pour une navigation anonyme', 49.99, 1000, 'https://example.com/images/cybershield.png', 'https://downloads.example.com/cybershield.exe', 0, '2024-05-15 14:30:00'),
(3, 'The Digital Nomad Guide', 'Ebook pour vivre et travailler à distance efficacement', 14.99, 1000, 'https://example.com/images/nomadguide.png', 'https://downloads.example.com/nomadguide.pdf', 0, '2024-03-10 09:00:00'),
(4, 'Startup Marketing Playbook', 'Stratégies de croissance pour startups en démarrage', 19.99, 1000, 'https://example.com/images/marketingbook.png', 'https://downloads.example.com/startupplaybook.epub', 0, '2024-02-20 16:45:00'),
(5, 'Clean UI Kit - Figma', 'UI kit moderne pour applications SaaS', 9.99, 1000, 'https://example.com/images/uikit.png', 'https://downloads.example.com/cleanuikit.fig', 0, '2024-01-05 08:00:00'),
(6, 'iPhone Mockups PSD Pack', 'Mockups professionnels pour iPhone 14 Pro', 5.99, 1000, 'https://example.com/images/mockups.png', 'https://downloads.example.com/iphonemockups.zip', 0, '2024-04-10 11:20:00'),
(7, 'CodeStarter - IDE Cloud', 'Environnement de développement cloud pour développeurs web', 39.99, 1000, 'https://example.com/images/codestarter.png', 'https://downloads.example.com/codestarter.dmg', 0, '2024-04-25 09:45:00'),
(8, 'AI Resume Builder', 'Générateur de CV intelligent basé sur l’IA', 24.99, 1000, 'https://example.com/images/resumebuilder.png', 'https://downloads.example.com/airesume.app', 0, '2024-06-10 10:30:00'),
(9, 'Freelancer Invoice Generator', 'App de facturation simple pour freelances', 11.99, 1000, 'https://example.com/images/invoicegen.png', 'https://downloads.example.com/invoicegen.zip', 0, '2024-03-18 15:15:00'),
(10, 'Minimal Dashboard Template', 'Template HTML/CSS pour tableau de bord moderne', 7.99, 1000, 'https://example.com/images/dashboard.png', 'https://downloads.example.com/dashboard-template.zip', 0, '2024-02-12 08:30:00'),
(11, 'CyberSec Essentials Ebook', 'Introduction à la cybersécurité moderne', 17.99, 1000, 'https://example.com/images/cybersec.png', 'https://downloads.example.com/cybersec.pdf', 0, '2023-12-01 11:10:00'),
(12, 'React Starter Pack', 'Kit de démarrage pour applications React', 12.99, 1000, 'https://example.com/images/reactstarter.png', 'https://downloads.example.com/react-starter.zip', 0, '2024-05-02 10:00:00'),
(13, 'Video LUTs Pack Vol.1', 'Pack de filtres professionnels pour éditeurs vidéo', 8.49, 1000, 'https://example.com/images/luts.png', 'https://downloads.example.com/videoluts.zip', 0, '2024-03-03 12:20:00'),
(14, 'Dark Mode Icons', 'Pack d’icônes minimalistes en mode sombre', 4.99, 1000, 'https://example.com/images/icons.png', 'https://downloads.example.com/darkicons.zip', 0, '2024-05-07 14:10:00'),
(15, 'Podcast Audio FX Pack', 'Effets sonores libres pour podcasts', 6.49, 1000, 'https://example.com/images/audiofx.png', 'https://downloads.example.com/podcastfx.zip', 0, '2024-04-01 10:00:00'),
(16, 'SEO Audit Toolkit', 'Outils pour auditer un site web en SEO', 34.99, 1000, 'https://example.com/images/seo.png', 'https://downloads.example.com/seoaudit.zip', 0, '2024-03-22 17:00:00'),
(17, 'Notion Templates for Freelancers', 'Templates Notion pour gestion de projets freelances', 3.99, 1000, 'https://example.com/images/notion.png', 'https://downloads.example.com/notionfreelance.zip', 0, '2024-06-01 11:00:00'),
(18, 'Flutter UI Components', 'Composants UI pour applications Flutter', 10.99, 1000, 'https://example.com/images/flutter.png', 'https://downloads.example.com/flutterui.zip', 0, '2024-02-05 09:00:00'),
(19, 'Excel Formulas Bible', 'Guide des formules Excel avancées', 13.49, 1000, 'https://example.com/images/excelbible.png', 'https://downloads.example.com/excelbible.pdf', 0, '2024-01-30 07:45:00'),
(20, 'JavaScript Mastery Course', 'Cours vidéo complet pour devenir expert JS', 59.99, 1000, 'https://example.com/images/jsmastery.png', 'https://downloads.example.com/jsmastery.mp4', 0, '2024-05-29 16:00:00');



-- Dumping data for table shop-mysql.items_category: ~2 rows (approximately)
INSERT INTO `items_category` (`ic_id`, `ic_name`, `ic_description`, `ic_picture_url`) VALUES
(1, 'Software', 'Applications logicielles pour Windows, Mac et Linux', 'https://example.com/images/software.png'),
(2, 'Ebooks', 'Livres numériques en PDF, EPUB ou Kindle', 'https://example.com/images/ebooks.png'),
(3, 'Design Assets', 'Templates, icônes, UI kits, mockups...', 'https://example.com/images/design_assets.png'),
(4, 'Audio / Vidéo', 'Effets audio, LUTs, sons libres de droit', 'https://example.com/images/audio_video.png'),
(5, 'Développement', 'Templates, starters, outils pour développeurs', 'https://example.com/images/devtools.png'),
(6, 'Formation', 'Cours vidéo, livres, guides interactifs', 'https://example.com/images/education.png');



-- Dumping data for table shop-mysql.items_category_link: ~3 rows (approximately)
-- Chaque sous-catégorie liée à UNE SEULE catégorie
INSERT INTO `items_category_link` (`icl_id`, `icl_items_sub_category_id`, `icl_items_category_id`) VALUES
(1, 1, 1),  -- Productivity Tools → Software
(2, 2, 1),  -- Antivirus & Sécurité → Software
(3, 3, 2),  -- Fiction → Ebooks
(4, 4, 2),  -- Business & Marketing → Ebooks
(5, 5, 3),  -- UI Kits → Design Assets
(6, 6, 3),  -- Mockups → Design Assets
(7, 7, 4),  -- Audio FX → Audio / Vidéo
(8, 8, 5),  -- Frontend Templates → Développement
(9, 9, 6);  -- Cours Vidéo → Formation




-- Dumping data for table shop-mysql.items_purshased: ~0 rows (approximately)

-- Dumping data for table shop-mysql.items_reviews: ~0 rows (approximately)

-- Dumping data for table shop-mysql.items_sub_category: ~3 rows (approximately)
INSERT INTO `items_sub_category` (`isc_id`, `isc_name`, `isc_description`, `isc_picture_url`) VALUES
(1, 'Productivity Tools', 'Logiciels de gestion de tâches, notes, calendrier, etc.', 'https://example.com/images/productivity.png'),
(2, 'Antivirus & Sécurité', 'Logiciels de cybersécurité et VPNs', 'https://example.com/images/security.png'),
(3, 'Fiction', 'Romans, nouvelles et récits narratifs', 'https://example.com/images/fiction.png'),
(4, 'Business & Marketing', 'Guides pour entrepreneurs et marketeurs', 'https://example.com/images/marketing.png'),
(5, 'UI Kits', 'Bibliothèques de composants UI', 'https://example.com/images/uikits.png'),
(6, 'Mockups', 'Templates PSD pour prévisualisation de produits', 'https://example.com/images/mockups.png'),
(7, 'Notion Templates', 'Templates prêts à l’emploi pour Notion', 'https://example.com/images/notion.png'),
(8, 'Frontend Templates', 'Templates HTML/CSS/React pour interfaces web', 'https://example.com/images/frontend.png'),
(9, 'Cours Vidéo', 'Cours en ligne sous format vidéo', 'https://example.com/images/courses.png'),
(10, 'Audio FX', 'Effets audio libres de droit', 'https://example.com/images/audiofx.png');



-- Dumping data for table shop-mysql.items_sub_category_link: ~1 rows (approximately)
-- Chaque item lié à UNE SEULE sous-catégorie
INSERT INTO `items_sub_category_link` (`iscl_id`, `iscl_items_id`, `iscl_sub_category_id`) VALUES
(1, 1, 1),  -- ZenTask → Productivity Tools
(2, 2, 2),  -- CyberShield → Antivirus & Sécurité
(3, 3, 3),  -- Nomad Guide → Fiction
(4, 4, 4),  -- Startup Playbook → Business
(5, 5, 5),  -- Clean UI Kit → UI Kits
(6, 6, 6),  -- iPhone Mockups → Mockups
(7, 7, 7),  -- Podcast FX → Audio FX
(8, 8, 8),  -- React Starter Pack → Frontend Templates
(9, 9, 8),  -- Flutter UI → Frontend Templates
(10, 10, 9),-- JS Mastery → Cours Vidéo
(11, 11, 1),-- Freelancer Invoice → Productivity Tools
(12, 12, 2),-- SEO Toolkit → Antivirus & Sécurité
(13, 13, 5),-- Dashboard Template → UI Kits
(14, 14, 5),-- Dark Mode Icons → UI Kits
(15, 15, 7),-- Video LUTs → Audio FX
(16, 16, 4),-- CyberSec Essentials → Business
(17, 17, 4);-- Excel Bible → Business




-- Dumping data for table shop-mysql.transactions: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users_basket: ~0 rows (approximately)

-- Dumping data for table shop-mysql.users_wishlist: ~0 rows (approximately)

/*!40103 SET TIME_ZONE=IFNULL(@OLD_TIME_ZONE, 'system') */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
