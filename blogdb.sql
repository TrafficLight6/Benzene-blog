# Host: 127.0.0.1  (Version: 5.7.26)
# Date: 2023-07-19 15:57:19
# Generator: MySQL-Front 5.3  (Build 4.234)

/*!40101 SET NAMES utf8 */;

#
# Structure for table "blog_email_code"
#

DROP TABLE IF EXISTS `blog_email_code`;
CREATE TABLE `blog_email_code` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL DEFAULT '',
  `code` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

#
# Structure for table "blog_page"
#

DROP TABLE IF EXISTS `blog_page`;
CREATE TABLE `blog_page` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '',
  `main` longtext NOT NULL,
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

#
# Structure for table "blog_token"
#

DROP TABLE IF EXISTS `blog_token`;
CREATE TABLE `blog_token` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `main_token` varchar(255) NOT NULL DEFAULT '',
  `time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

#
# Structure for table "blog_user"
#

DROP TABLE IF EXISTS `blog_user`;
CREATE TABLE `blog_user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'sha256',
  `allowadd` varchar(255) NOT NULL DEFAULT 'false',
  `allowpost` varchar(255) NOT NULL DEFAULT 'true',
  `admin` varchar(255) NOT NULL DEFAULT 'false',
  `email` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`Id`),
  KEY `username` (`username`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;
