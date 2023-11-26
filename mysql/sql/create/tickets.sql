CREATE TABLE IF NOT EXISTS `tickets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`event_id` int UNSIGNED NOT NULL
	`title` varchar(255) NOT NULL,
	`note` text NULL,
	`sale_type` varchar(255) NOT NULL,
	`start_at` int UNSIGNED NOT NULL,
	`finish_at` int UNSIGNED NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
