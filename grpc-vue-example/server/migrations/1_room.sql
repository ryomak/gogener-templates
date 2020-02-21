-- +migrate Up
CREATE TABLE `rooms` (
  `name` varchar(255) NOT NULL COMMENT 'name',
  `created` timestamp NOT NULL DEFAULT NOW() COMMENT 'when created',
  `updated` timestamp NOT NULL DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP COMMENT 'when last updated',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='list of rooms';

-- +migrate Down
DROP TABLE rooms;