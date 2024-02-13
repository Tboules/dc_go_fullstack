CREATE TABLE IF NOT EXISTS `user` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `provider_id` varchar(255) NOT NULL,
  `name` varchar(255),
  `email` varchar(255) UNIQUE NOT NULL,
  `image` varchar(2048)
); 
