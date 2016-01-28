CREATE DATABASE IF NOT EXISTS `dcsolusion` DEFAULT CHARACTER SET utf8;

USE `dcsolusion`;

DROP TABLE IF EXISTS `admin_user`;

CREATE TABLE `admin_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(32) NOT NULL COMMENT 'encryption by md5',
  PRIMARY KEY (`id`),
  UNIQUE KEY `USERNAME` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
