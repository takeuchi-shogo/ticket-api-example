CREATE TABLE IF NOT EXISTS `payment_by_credit_cards` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_book_ticket_id` INT(10) UNSIGNED NOT NULL,
	`payment_id` varchar(255) NOT NULL,
	`is_valid` tinyint(1) NULL,
	`expire_at` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
