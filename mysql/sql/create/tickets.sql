CREATE TABLE IF NOT EXISTS `tickets` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`event_id` int UNSIGNED NOT NULL,
	`title` varchar(255) NOT NULL,
	`venue_id` int UNSIGNED NULL,
	`note` text NULL,
	`sale_type` varchar(255) NOT NULL,
	`start_at` int UNSIGNED NOT NULL,
	`finish_at` int UNSIGNED NOT NULL,
	`lottery_at` int UNSIGNED NOT NULL,
	`is_payment_by_credit_card` tinyint(1) NOT NULL,
	`is_payment_by_convenience` tinyint(1) NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
