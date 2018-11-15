-- MySQL dump 10.13  Distrib 5.7.19, for linux-glibc2.12 (x86_64)
--
-- Host: 127.0.0.1    Database: poow
-- ------------------------------------------------------
-- Server version	5.7.19-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `hosts`
--

DROP TABLE IF EXISTS `hosts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `hosts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `is_valid` tinyint(4) NOT NULL DEFAULT '1' COMMENT '该host是否可用: 0.否, 1.是',
  `is_dedicate` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否专用host: 0.否, 1.是',
  `host` varchar(50) NOT NULL COMMENT 'host',
  `running_task_count` int(11) NOT NULL DEFAULT '0' COMMENT '该host上有多少个任务再运行',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_host` (`host`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='命令可以使用哪些host';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `program_hosts`
--

DROP TABLE IF EXISTS `program_hosts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `program_hosts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `program_id` bigint(20) NOT NULL COMMENT '关联的host id',
  `host_id` bigint(20) NOT NULL COMMENT '关联的host id',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_program_host_id` (`program_id`,`host_id`),
  KEY `idx_host_id` (`host_id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='命令需要在哪个host上运行';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `programs`
--

DROP TABLE IF EXISTS `programs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `programs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `title` varchar(150) NOT NULL COMMENT '程序是干什么的',
  `file_name` varchar(150) NOT NULL COMMENT '需要执行的程序名称',
  `have_dedicate` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否有专用host: 0.否, 1.是',
  `params` varchar(500) NOT NULL DEFAULT '' COMMENT '执行命令的参数',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_title` (`title`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='有哪些命令可以使用';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `tasks`
--

DROP TABLE IF EXISTS `tasks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tasks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `program_id` bigint(20) NOT NULL COMMENT '关联命令ID',
  `task_uuid` varchar(30) NOT NULL COMMENT '任务的UUID',
  `host` varchar(50) NOT NULL COMMENT '任务运行在哪个host上',
  `file_name` varchar(150) NOT NULL COMMENT '需要执行的程序名称',
  `params` varchar(500) NOT NULL DEFAULT '' COMMENT '执行命令的参数',
  `pid` bigint(20) NOT NULL DEFAULT '0' COMMENT '父任务id, 关联自己',
  `log_path` varchar(255) NOT NULL DEFAULT '' COMMENT '保存了一些实时需要持久化的信息',
  `notify_info` varchar(500) NOT NULL DEFAULT '' COMMENT '会实时读该字段的信息, 一般外部其他程序可以通过修改这个来和任务进行通讯.',
  `real_info` varchar(500) NOT NULL DEFAULT '' COMMENT '保存了一些实时需要持久化的信息',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '任务状态: 1.执行成功, 2.执行中, 3.执行失败',
  PRIMARY KEY (`id`),
  UNIQUE KEY `udx_task_uuid` (`task_uuid`),
  KEY `idx_programs_id` (`program_id`),
  KEY `idx_pid` (`pid`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='运行的任务';
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-11-15 19:55:58
