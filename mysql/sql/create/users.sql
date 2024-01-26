CREATE TABLE IF NOT EXISTS `users` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`first_name` varchar(255) NOT NULL,
	`last_name` varchar(255) NOT NULL,
	`display_name` varchar(255) NULL,
	`screen_name` varchar(50) NOT NULL,
	`email` varchar(255) NOT NULL,
	`tel` varchar(50) NULL,
	`password` varchar(255) NOT NULL,
	`post_code` varchar(50) NULL,
	`prefecture` varchar(50) NULL,
	`city` varchar(255) NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	UNIQUE (screen_name),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
