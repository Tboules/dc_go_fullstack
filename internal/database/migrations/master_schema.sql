CREATE TABLE `user` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `email_verified` boolean DEFAULT false,
  `image` varchar(255)
);

CREATE TABLE `desert_figure` (
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

CREATE TABLE `excerpt` (
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

CREATE TABLE `tag` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) UNIQUE NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` BIGINT NOT NULL
);

CREATE TABLE `excerpt_tag` (
  `excerpt_id` BIGINT NOT NULL,
  `tag_id` BIGINT NOT NULL,
  PRIMARY KEY (`excerpt_id`, `tag_id`)
);

CREATE TABLE `icon` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `url` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_by` BIGINT NOT NULL,
  `desert_figure` BIGINT NOT NULL,
  `date_added` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_updated` timestamp
);

CREATE INDEX `excerpt_index_0` ON `excerpt` (`created_by`);

CREATE INDEX `excerpt_index_1` ON `excerpt` (`desert_figure`);

CREATE INDEX `excerpt_tag_index_2` ON `excerpt_tag` (`tag_id`);

CREATE INDEX `icon_index_3` ON `icon` (`desert_figure`);


ALTER TABLE `excerpt` ADD FOREIGN KEY (`desert_figure`) REFERENCES `desert_figure` (`id`);

ALTER TABLE `excerpt` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`);

ALTER TABLE `tag` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`);

ALTER TABLE `excerpt_tag` ADD FOREIGN KEY (`excerpt_id`) REFERENCES `excerpt` (`id`);

ALTER TABLE `excerpt_tag` ADD FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`);

ALTER TABLE `icon` ADD FOREIGN KEY (`created_by`) REFERENCES `user` (`id`);

ALTER TABLE `icon` ADD FOREIGN KEY (`desert_figure`) REFERENCES `desert_figure` (`id`);
