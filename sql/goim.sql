/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : localhost:3306
 Source Schema         : goim

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 08/06/2020 10:29:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for contacts
-- ----------------------------
DROP TABLE IF EXISTS `contacts`;
CREATE TABLE `contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `uid` bigint(20) NOT NULL COMMENT '用户id',
  `interactive_uid` bigint(20) NOT NULL COMMENT '互动人id',
  `status` tinyint(2) NOT NULL COMMENT '聊天状态(0正常，1删除)',
  `message_id` bigint(20) NOT NULL COMMENT '最新消息id',
  `is_read` tinyint(2) NOT NULL COMMENT '是否已读(0未读，1已读)',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid_interactive` (`uid`,`interactive_uid`) USING BTREE,
  KEY `uid_status` (`uid`,`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='最近联系人';

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `poster` bigint(20) NOT NULL COMMENT '发送人id',
  `receiver` bigint(20) NOT NULL COMMENT '接收人id',
  `message_type` tinyint(2) NOT NULL COMMENT '消息类型(0文字，1图片)',
  `message` text NOT NULL COMMENT '消息内容',
  `dateline` bigint(20) NOT NULL COMMENT '创建时间',
  `seq` bigint(20) NOT NULL COMMENT '序列号',
  PRIMARY KEY (`id`),
  KEY `poster_receiver_seq` (`poster`,`receiver`,`seq`) USING BTREE
  KEY `seq` (`seq`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='消息';

SET FOREIGN_KEY_CHECKS = 1;
