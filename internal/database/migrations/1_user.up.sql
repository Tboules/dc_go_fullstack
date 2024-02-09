CREATE TABLE `user` IF NOT EXISTS (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) UNIQUE NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `email_verified` boolean DEFAULT false,
  `image` varchar(255)
);

