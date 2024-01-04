CREATE TABLE IF NOT EXISTS `register_emails` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`email` varchar(50) NOT NULL,
	`token` varchar(50) NOT NULL,
	`pin_code` varchar(10) NOT NULL,
	`is_valid` tinyint(1) NULL,
	`expire_at` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
