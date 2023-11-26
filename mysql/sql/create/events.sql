CREATE TABLE IF NOT EXISTS `events` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`organizer_id` int UNSIGNED NOT NULL,
	`venue_id` int NULL,
	`title` varchar(255) NOT NULL,
	`performance_period` varchar(255) NOT NULL,
	`event_type` varchar(50) NOT NULL,
	`show_time` int UNSIGNED NOT NULL,
	`opening_time` int UNSIGNED NOT NULL,
	`description` text NULL,
	`note` text NULL,
	`is_private` tinyint(1) NOT NULL,
	`created_at` int UNSIGNED NOT NULL,
	`updated_at` int UNSIGNED NOT NULL,
	`deleted_at` int UNSIGNED NULL,
	UNIQUE (screen_name),
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
