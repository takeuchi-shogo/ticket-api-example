CREATE TABLE IF NOT EXISTS `user_has_tickets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL,
	`user_book_ticket_id` int UNSIGNED NOT NULL,
	`seat_id` int UNSIGNED NULL,
	`ticket_status` varchar(50) NOT NULL,
	`reference_number` int UNSIGNED NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
