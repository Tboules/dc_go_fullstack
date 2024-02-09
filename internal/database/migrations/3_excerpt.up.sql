CREATE TABLE `excerpt` IF NOT EXISTS (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `body` varchar(255) NOT NULL,
  `type` int NOT NULL,
  `reference_title` varchar(255),
  `reference_page` integer,
  `reference_url` varchar(255),
  `desert_figure` BIGINT NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp,
  `created_by` BIGINT NOT NULL
);
