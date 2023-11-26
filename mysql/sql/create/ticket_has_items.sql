CREATE TABLE IF NOT EXISTS `ticket_has_items` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`event_id` int UNSIGNED NOT NULL,
	`ticket_id` int UNSIGNED NOT NULL,
	`ticket_item_id` int UNSIGNED NOT NULL,
	`amount` decimal(10,3) UNSIGNED NOT NULL,
	`remarks` text NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
