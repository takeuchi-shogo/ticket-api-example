CREATE TABLE IF NOT EXISTS `user_book_tickets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`book_id` varchar(50) NOT NULL,
	`user_id` int UNSIGNED NOT NULL,
	`event_id` int UNSIGNED NOT NULL,
	`ticket_id` int UNSIGNED NOT NULL,
	`ticket_item_id` int UNSIGNED NOT NULL,
	`payment_method` varchar(50) NOT NULL,
	`number_of_tickets` int UNSIGNED NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
