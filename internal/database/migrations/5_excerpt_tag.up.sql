CREATE TABLE `excerpt_tag` IF NOT EXISTS (
  `excerpt_id` BIGINT NOT NULL,
  `tag_id` BIGINT NOT NULL,
  PRIMARY KEY (`excerpt_id`, `tag_id`)
);
