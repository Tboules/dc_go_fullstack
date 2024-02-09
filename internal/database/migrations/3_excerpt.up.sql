CREATE TABLE IF NOT EXISTS `excerpt` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `body` varchar(255) NOT NULL,
  `type` int NOT NULL,
  `reference_title` varchar(255),
  `reference_page` integer,
  `reference_url` varchar(255),
  `desert_figure` BIGINT NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp,
  `created_by` BIGINT NOT NULL,
  CONSTRAINT fk_excerpt_df FOREIGN KEY (`desert_figure`) REFERENCES `desert_figure` (`id`),
  CONSTRAINT fk_excerpt_user FOREIGN KEY (`created_by`) REFERENCES `user` (`id`)
);
