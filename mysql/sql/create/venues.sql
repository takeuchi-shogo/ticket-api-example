CREATE TABLE IF NOT EXISTS `users` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`post_code` varchar(50) NOT NULL,
	`address` varchar(255) NOT NULL,
	`how_to_access` text NULL,
	`capacity` int UNSIGNED NOT NULL,
	`parking_space` int UNSIGNED NOT NULL,
	`parking_description` text NULL,
	`site_url` varchar(255) NULL,
	`contact_tel` varchar(50) NULL,
	`contact_email` varchar(50) NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
