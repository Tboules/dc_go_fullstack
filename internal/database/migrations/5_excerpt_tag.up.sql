CREATE TABLE IF NOT EXISTS `excerpt_tag` (
  `excerpt_id` BIGINT NOT NULL,
  `tag_id` BIGINT NOT NULL,
  PRIMARY KEY (`excerpt_id`, `tag_id`),
  CONSTRAINT fk_et_excerpt FOREIGN KEY (`excerpt_id`) REFERENCES `excerpt` (`id`),
  CONSTRAINT fk_et_tag FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`)
);
