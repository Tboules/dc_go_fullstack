CREATE TABLE `icon` IF NOT EXISTS (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `url` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_by` BIGINT NOT NULL,
  `desert_figure` BIGINT NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp
);
