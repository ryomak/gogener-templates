CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `rand_id` varchar(255) NOT NULL,
  `auth_id` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `last_name_kana` varchar(64) NOT NULL,
  `gender`  enum('male', 'female', 'none') NOT NULL,
  `birth_day` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `rand_id` (`rand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `maintenance` (
  `name` varchar(255) NOT NULL,
  `is_active` tinyint(1) NOT NULL DEFAULT 0,
  `message` text NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
