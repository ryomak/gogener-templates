-- +migrate Up
CREATE TABLE `users` (
  `user_name` varchar(255) NOT NULL COMMENT 'user_name',
  `room_name` varchar(255) NOT NULL COMMENT 'room_name',
  `latitude` FLOAT NOT NULL COMMENT 'latitude',
  `longitude` FLOAT NOT NULL COMMENT 'longitude',
  `created` timestamp NOT NULL DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  PRIMARY KEY (`user_name`,`room_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='list of users';

-- +migrate Down
DROP TABLE users;