CREATE TABLE IF NOT EXISTS `user_mail_logs` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`email` varchar(50) NULL,
	`user_id` INT(10) NULL,
	`user_book_ticket_id` INT(10) NULL,
	`mail_type` varchar(50) NOT NULL,
	`is_send` tinyint(1) NOT NULL,
	`error_message` text NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
