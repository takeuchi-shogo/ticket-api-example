CREATE TABLE IF NOT EXISTS `credit_cards` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` int UNSIGNED NOT NULL,
	`customer_id` int UNSIGNED NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
