CREATE TABLE IF NOT EXISTS `desert_figure` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `full_name` varchar(255) UNIQUE,
  `first_name` varchar(255),
  `last_name` varchar(255),
  `type` int NOT NULL,
  `date_of_birth` timestamp,
  `date_of_death` timestamp,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp,
  `created_by` BIGINT NOT NULL,
  CONSTRAINT fk_df_user FOREIGN KEY (`created_by`) REFERENCES `user` (`id`)
);
