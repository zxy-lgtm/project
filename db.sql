/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`teaches` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `teaches`;

/*Table structure for table `students` */

DROP TABLE IF EXISTS `students`;

CREATE TABLE `students` (
  `id` bigint(20) NOT NULL  COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '姓名',
  `school` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '学校',
  `major` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '专业',
  `grade` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '年级',
  `tel` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系人电话',
  `email`varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系人邮箱',
  `type` tinyint DEFAULT NULL COMMENT '授课方式',
  `gender` tinyint DEFAULT NULL COMMENT '性别',
  `notes` text NULLCOLLATE utf8_unicode_ci DEFAULT NULL COMMENT `自我介绍`,
  `style` varchar NULLCOLLATE utf8_unicode_ci DEFAULT NULL COMMENT `教学风格`,
  `money` int COLLATE utf8_unicode_ci DEFAULT NULL COMMENT `教师设定的预约咨询价格`,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` int COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '评价等级',
  `all` text COLLATE utf8_unicode_ci DEFAULT NULL COMMENT `评价总结`,
  `c1` int NULL COMMENT '心理素质',
  `c2` int NULL COMMENT '仪态仪表',
  `c3` int NULL COMMENT '语言表达',
  `c4` int NULL COMMENT '思维品质',
  `c5` int NULL COMMENT '教学设计',
  `c6` int NULL COMMENT '教学实验',
  `c7` int NULL COMMENT `教学评价`,
  `w1` text NULL COMMENT '心理素质',
  `w2` text NULL COMMENT '仪态仪表',
  `w3` text NULL COMMENT '语言表达',
  `w4` text NULL COMMENT '思维品质',
  `w5` text NULL COMMENT '教学设计',
  `w6` text NULL COMMENT '教学实验',
  `w7` text NULL COMMENT `教学评价`,
  `person_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `time`;

CREATE TABLE `time` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '科目名字和年级',
  `time` int COLLATE utf8_unicode_ci DEFAULT NULL COMMENT `时长`，
  `person_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `award`;

CREATE TABLE `award` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '名字',
  `person_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `teach`;

CREATE TABLE `teach` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '名字',
  `time` varchar(30) NULLCOLLATE utf8_unicode_ci DEFAULT NULL COMMENT '教学时长比如2019.9-2020.4',
  `result` varchar NULLCOLLATE utf8_unicode_ci DEFAULT NULL COMMENT '教学成果',
  `person_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `subject`;

CREATE TABLE `subject` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '名字',
  `tips` varchar(30) NULLCOLLATE utf8_unicode_ci DEFAULT NULL COMMENT '学生对学科的要求或者教师对教授学科的描述',
  `person_id` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


/*Table structure for table `patriarchs` */

DROP TABLE IF EXISTS `patriarchs`;

CREATE TABLE `patriarchs` (
  `id` bigint(20) NOT NULL  COMMENT '主键ID',
  `name_p` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '家长姓名',
  `identity_card_p` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '家长身份证号码',
  `identity_card_c` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '孩子身份证号码',
  `school` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '学校',
  `name_c` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '孩子姓名',
  `grade` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '年级',
  `tel` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系人电话',
  `email`varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系人邮箱',
  `score_situation` text COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '成绩概况',
  `gender_p` tinyint DEFAULT NULL COMMENT '家长性别',
  `subject_id` int  NOT NULL  COMMENT '需求科目id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;


DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '姓名',
  `identity_card` varchar(50) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '身份证号码',
  `email`varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '联系人邮箱',
  `gender` tinyint DEFAULT NULL COMMENT '性别',
  `password` varchar(30) NOT NULL,
  `account` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `classes`;

CREATE TABLE `classes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `student_name` varchar(15) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '教师姓名',
  `name`varchar(30) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '课程名称',
  `tips` tinyint DEFAULT NULL COMMENT '类型',
  `status` tinyint DEFAULT NULL COMMENT '免费或收费',
  `tips2` tinyint DEFAULT NULL COMMENT '课程数量和时长',
  `money` int DEFAULT NULL COMMENT '收费',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;



--insert into `students`(`name`,`identity_card`,`school`,`major`,`gender`,`grade`,`tel`,`subject`,`type`,`award_situation`,`email`) values ("m","202020202202022020","华中师范大学","教育",1,"2020","110110110","数学",1,"无1111111111111打我v而外国分   ￥v2v3rg3ggvw    2gvv","123@qq.com");
--insert into `patriarchs`(`name_p`,`identity_card_p`,`name_c`,`identity_card_c`,`school`,`gender_p`,`grade`,`tel`,`score_situation`,`email`) values ("王","20202020220202202024","王小明","2020202022020220202424","华xiao学",1,"2020","110110110","无1111111111111打我v而外国分   ￥v2v3rg3ggvw    2gvv","123@qq.com");
--insert into `users`(`name`,`identity_card`,`gender`,`email`) values ("m","202020202202022020",1,"123@qq.com");


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;