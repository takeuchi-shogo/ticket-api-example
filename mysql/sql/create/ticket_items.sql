CREATE TABLE IF NOT EXISTS `ticket_items` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`event_id` int UNSIGNED NOT NULL,
	`title` varchar(255) NOT NULL,
	`issuing_number` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
