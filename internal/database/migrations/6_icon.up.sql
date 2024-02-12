CREATE TABLE IF NOT EXISTS `icon` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `url` varchar(2048) NOT NULL,
  `description` varchar(255),
  `created_by` BIGINT NOT NULL,
  `desert_figure` BIGINT NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp,
  CONSTRAINT fk_icon_user FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  CONSTRAINT fk_icon_df FOREIGN KEY (`desert_figure`) REFERENCES `desert_figure` (`id`)
);
