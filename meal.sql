/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : meal

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2019-10-23 10:57:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rms_advise
-- ----------------------------
DROP TABLE IF EXISTS `rms_advise`;
CREATE TABLE `rms_advise` (
  `id` int(11) NOT NULL,
  `advise` varchar(1024) DEFAULT NULL,
  `tag_id` int(11) DEFAULT NULL,
  `score` int(4) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of rms_advise
-- ----------------------------

-- ----------------------------
-- Table structure for rms_advise_tag
-- ----------------------------
DROP TABLE IF EXISTS `rms_advise_tag`;
CREATE TABLE `rms_advise_tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_name` varchar(64) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of rms_advise_tag
-- ----------------------------

-- ----------------------------
-- Table structure for rms_backend_user
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user`;
CREATE TABLE `rms_backend_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_backend_user
-- ----------------------------
INSERT INTO `rms_backend_user` VALUES ('1', 'szh', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1', '1', '13754338419', '206632394@qq.com', '/static/upload/微信图片_20190112122216.jpg');
INSERT INTO `rms_backend_user` VALUES ('3', '张三', 'zhangsan', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '', '', '');
INSERT INTO `rms_backend_user` VALUES ('5', '李四', 'lisi', 'e10adc3949ba59abbe56e057f20f883e', '0', '0', '', '', '');

-- ----------------------------
-- Table structure for rms_backend_user_rms_roles
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user_rms_roles`;
CREATE TABLE `rms_backend_user_rms_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rms_backend_user_id` int(11) NOT NULL,
  `rms_role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_backend_user_rms_roles
-- ----------------------------

-- ----------------------------
-- Table structure for rms_course
-- ----------------------------
DROP TABLE IF EXISTS `rms_course`;
CREATE TABLE `rms_course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `short_name` varchar(8) NOT NULL DEFAULT '',
  `price` double NOT NULL DEFAULT '0',
  `real_price` double NOT NULL DEFAULT '0',
  `img` varchar(256) NOT NULL DEFAULT '',
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `creator_id` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_course
-- ----------------------------
INSERT INTO `rms_course` VALUES ('1', '2020考研政治精讲1', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('2', '2020考研政治精讲2', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('3', '2020考研政治精讲3', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('4', '2020考研政治精讲3', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('5', '2020考研政治精讲4', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('6', '2020考研政治精讲5', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('7', '2020考研政治精讲6', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('8', '2020考研政治精讲7', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('9', '2020考研政治精讲8', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('10', '2020考研政治精讲9', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('11', '2020考研政治精讲10', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('12', '2020考研政治精讲11', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('13', '2020考研政治精讲12', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('14', '2020考研政治精讲13', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('15', '2020考研政治精讲13', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('16', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('17', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('18', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('19', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('20', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('21', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('22', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('23', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('24', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('25', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('26', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('27', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('28', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('29', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('30', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('31', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('32', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('33', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('34', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');
INSERT INTO `rms_course` VALUES ('35', '2020考研政治精讲15', '2020考研', '100', '0', '', '2019-04-30 16:00:00', '2019-05-30 16:00:00', '100', '1');

-- ----------------------------
-- Table structure for rms_day_meal
-- ----------------------------
DROP TABLE IF EXISTS `rms_day_meal`;
CREATE TABLE `rms_day_meal` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(3) DEFAULT '0' COMMENT '0 早饭 1 中饭 2 晚饭 3 外卖',
  `meal_id` int(11) NOT NULL COMMENT '菜单id',
  `meal_date` int(11) DEFAULT NULL COMMENT '日期 tian ',
  `out_nums` int(11) DEFAULT '0' COMMENT '销量',
  `seq` tinyint(3) DEFAULT NULL,
  `time` int(11) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `order_date` (`meal_date`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='每日菜单';

-- ----------------------------
-- Records of rms_day_meal
-- ----------------------------
INSERT INTO `rms_day_meal` VALUES ('1', '0', '6', '1569888000', '0', '0', '1571563473');
INSERT INTO `rms_day_meal` VALUES ('2', '0', '5', '1569888000', '0', '0', '1571563473');
INSERT INTO `rms_day_meal` VALUES ('3', '0', '6', '1570060800', '0', '0', '1571563588');
INSERT INTO `rms_day_meal` VALUES ('4', '0', '5', '1570060800', '0', '0', '1571563588');
INSERT INTO `rms_day_meal` VALUES ('5', '0', '6', '1569974400', '0', '0', '1571563645');
INSERT INTO `rms_day_meal` VALUES ('6', '0', '5', '1569974400', '0', '0', '1571563645');
INSERT INTO `rms_day_meal` VALUES ('7', '0', '8', '1571616000', '0', '0', '1571577786');
INSERT INTO `rms_day_meal` VALUES ('8', '1', '9', '1571616000', '0', '0', '1571578270');
INSERT INTO `rms_day_meal` VALUES ('9', '2', '8', '1571500800', '0', '0', '1571578685');

-- ----------------------------
-- Table structure for rms_meal
-- ----------------------------
DROP TABLE IF EXISTS `rms_meal`;
CREATE TABLE `rms_meal` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `meal_name` varchar(64) NOT NULL,
  `meal_img` varchar(128) DEFAULT NULL,
  `meal_desc` text,
  `price` varchar(20) DEFAULT NULL,
  `score` int(5) DEFAULT NULL,
  `score_list` varchar(512) DEFAULT NULL,
  `seq` tinyint(4) DEFAULT NULL,
  `time` int(11) DEFAULT '0',
  `meal_type_id` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `score` (`score`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='菜单';

-- ----------------------------
-- Records of rms_meal
-- ----------------------------
INSERT INTO `rms_meal` VALUES ('4', '红烧肉', '/static/upload/201249aid28sb25dsbiizs.jpg', '', null, '0', '', '111', '1571450873', '1');
INSERT INTO `rms_meal` VALUES ('5', '花菜', '/static/upload/微信图片_20190112122216.jpg', '', null, '0', '', '12', '1571535437', '1');
INSERT INTO `rms_meal` VALUES ('6', '包心菜', '/static/upload/201509031357525d68e.jpg', '', null, '0', '', '123', '1571535456', '1');
INSERT INTO `rms_meal` VALUES ('7', '馒头', '/static/upload/201509031357525d68e.jpg', '', null, '0', '', '12', '1571577734', '1');
INSERT INTO `rms_meal` VALUES ('8', '包子', '/static/upload/201509031357525d68e.jpg', '', null, '0', '', '123', '1571577742', '1');
INSERT INTO `rms_meal` VALUES ('9', '牛堡', '/static/upload/201249aid28sb25dsbiizs.jpg', '', null, '0', '', '10', '1571578174', '1');

-- ----------------------------
-- Table structure for rms_meal_type
-- ----------------------------
DROP TABLE IF EXISTS `rms_meal_type`;
CREATE TABLE `rms_meal_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of rms_meal_type
-- ----------------------------
INSERT INTO `rms_meal_type` VALUES ('1', '测试', '1571721401');
INSERT INTO `rms_meal_type` VALUES ('2', '蔬菜1', '1571723169');

-- ----------------------------
-- Table structure for rms_resource
-- ----------------------------
DROP TABLE IF EXISTS `rms_resource`;
CREATE TABLE `rms_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `parent_id` int(11) DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=47 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_resource
-- ----------------------------
INSERT INTO `rms_resource` VALUES ('7', '1', '权限管理', '8', '100', 'fa fa-balance-scale', '');
INSERT INTO `rms_resource` VALUES ('8', '0', '系统菜单', null, '200', '', '');
INSERT INTO `rms_resource` VALUES ('9', '1', '资源管理', '7', '100', '', 'ResourceController.Index');
INSERT INTO `rms_resource` VALUES ('12', '1', '角色管理', '7', '100', '', 'RoleController.Index');
INSERT INTO `rms_resource` VALUES ('13', '1', '用户管理', '7', '100', '', 'BackendUserController.Index');
INSERT INTO `rms_resource` VALUES ('14', '1', '系统管理', '8', '90', 'fa fa-gears', '');
INSERT INTO `rms_resource` VALUES ('21', '0', '业务菜单', null, '170', '', '');
INSERT INTO `rms_resource` VALUES ('23', '1', '日志管理(空)', '14', '100', '', '');
INSERT INTO `rms_resource` VALUES ('25', '2', '编辑', '9', '100', 'fa fa-pencil', 'ResourceController.Edit');
INSERT INTO `rms_resource` VALUES ('26', '2', '编辑', '13', '100', 'fa fa-pencil', 'BackendUserController.Edit');
INSERT INTO `rms_resource` VALUES ('27', '2', '删除', '9', '100', 'fa fa-trash', 'ResourceController.Delete');
INSERT INTO `rms_resource` VALUES ('29', '2', '删除', '13', '100', 'fa fa-trash', 'BackendUserController.Delete');
INSERT INTO `rms_resource` VALUES ('30', '2', '编辑', '12', '100', 'fa fa-pencil', 'RoleController.Edit');
INSERT INTO `rms_resource` VALUES ('31', '2', '删除', '12', '100', 'fa fa-trash', 'RoleController.Delete');
INSERT INTO `rms_resource` VALUES ('32', '2', '分配资源', '12', '100', 'fa fa-th', 'RoleController.Allocate');
INSERT INTO `rms_resource` VALUES ('35', '1', ' 首页', null, '100', 'fa fa-dashboard', 'HomeController.Index');
INSERT INTO `rms_resource` VALUES ('38', '1', '菜谱管理', '21', '100', 'fa fa-book', 'MealController.Index');
INSERT INTO `rms_resource` VALUES ('39', '2', '删除', '38', '100', '', 'MealController.Delete,:id,1');
INSERT INTO `rms_resource` VALUES ('40', '1', '添加菜谱', '38', '100', '', 'MealController.Index');
INSERT INTO `rms_resource` VALUES ('41', '1', '每日菜单', '38', '100', '', 'DailyMealController.Index');
INSERT INTO `rms_resource` VALUES ('42', '1', '菜单分类', '38', '100', '', 'MealTypeController.Index');
INSERT INTO `rms_resource` VALUES ('43', '1', '外卖订单', '38', '100', '', 'MealUserOrderController.Index');
INSERT INTO `rms_resource` VALUES ('44', '1', '意见管理', '38', '100', '', 'MealAdviseController.Index');
INSERT INTO `rms_resource` VALUES ('45', '1', '次日统计', '38', '100', '', 'MealUserCalcController.Index');
INSERT INTO `rms_resource` VALUES ('46', '1', '用户管理', '38', '100', '', 'MealUserController.Index');

-- ----------------------------
-- Table structure for rms_role
-- ----------------------------
DROP TABLE IF EXISTS `rms_role`;
CREATE TABLE `rms_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role
-- ----------------------------
INSERT INTO `rms_role` VALUES ('22', '超级管理员', '20');
INSERT INTO `rms_role` VALUES ('24', '角色管理员', '10');
INSERT INTO `rms_role` VALUES ('25', '课程资源管理员', '5');

-- ----------------------------
-- Table structure for rms_role_backenduser_rel
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_backenduser_rel`;
CREATE TABLE `rms_role_backenduser_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role_backenduser_rel
-- ----------------------------
INSERT INTO `rms_role_backenduser_rel` VALUES ('66', '25', '3', '2017-12-19 06:40:34');
INSERT INTO `rms_role_backenduser_rel` VALUES ('67', '22', '1', '2019-10-18 05:11:06');

-- ----------------------------
-- Table structure for rms_role_resource_rel
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_resource_rel`;
CREATE TABLE `rms_role_resource_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=667 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

-- ----------------------------
-- Records of rms_role_resource_rel
-- ----------------------------
INSERT INTO `rms_role_resource_rel` VALUES ('448', '24', '8', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('449', '24', '14', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('450', '24', '23', '2017-12-19 06:40:16');
INSERT INTO `rms_role_resource_rel` VALUES ('451', '25', '21', '2019-05-11 13:57:37');
INSERT INTO `rms_role_resource_rel` VALUES ('642', '22', '35', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('643', '22', '21', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('644', '22', '38', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('645', '22', '39', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('646', '22', '40', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('647', '22', '41', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('648', '22', '42', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('649', '22', '43', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('650', '22', '44', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('651', '22', '45', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('652', '22', '46', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('653', '22', '8', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('654', '22', '14', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('655', '22', '23', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('656', '22', '7', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('657', '22', '9', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('658', '22', '25', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('659', '22', '27', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('660', '22', '12', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('661', '22', '30', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('662', '22', '31', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('663', '22', '32', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('664', '22', '13', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('665', '22', '26', '2019-10-22 16:01:59');
INSERT INTO `rms_role_resource_rel` VALUES ('666', '22', '29', '2019-10-22 16:01:59');

-- ----------------------------
-- Table structure for rms_user
-- ----------------------------
DROP TABLE IF EXISTS `rms_user`;
CREATE TABLE `rms_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `open_id` varchar(64) DEFAULT NULL,
  `name` varchar(64) DEFAULT NULL,
  `phone` bigint(11) DEFAULT NULL,
  `nick_name` varchar(64) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  `session_key` varchar(255) NOT NULL DEFAULT '',
  `access_token` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of rms_user
-- ----------------------------
INSERT INTO `rms_user` VALUES ('1', '12312', '12312', '213123', '12312', null, '', '');

-- ----------------------------
-- Table structure for rms_user_calc_history
-- ----------------------------
DROP TABLE IF EXISTS `rms_user_calc_history`;
CREATE TABLE `rms_user_calc_history` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `meal_date` int(11) DEFAULT NULL,
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`meal_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='次日就餐历史';

-- ----------------------------
-- Records of rms_user_calc_history
-- ----------------------------

-- ----------------------------
-- Table structure for rms_user_meal_calc
-- ----------------------------
DROP TABLE IF EXISTS `rms_user_meal_calc`;
CREATE TABLE `rms_user_meal_calc` (
  `meal_date` bigint(11) NOT NULL DEFAULT '0',
  `meal_total` int(11) DEFAULT '0' COMMENT '总投票',
  `meal_nums` int(11) DEFAULT '0',
  PRIMARY KEY (`meal_date`),
  KEY `meal_date` (`meal_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='次日用餐统计';

-- ----------------------------
-- Records of rms_user_meal_calc
-- ----------------------------
INSERT INTO `rms_user_meal_calc` VALUES ('1571587200', null, '4');
INSERT INTO `rms_user_meal_calc` VALUES ('1571673600', null, '2');

-- ----------------------------
-- Table structure for rms_user_order
-- ----------------------------
DROP TABLE IF EXISTS `rms_user_order`;
CREATE TABLE `rms_user_order` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `type` tinyint(2) NOT NULL DEFAULT '0',
  `meal_ids` varchar(512) NOT NULL,
  `meal_date` int(11) NOT NULL DEFAULT '0',
  `meal_code` varchar(64) DEFAULT NULL COMMENT '取餐码',
  `total` varchar(255) DEFAULT NULL,
  `status` tinyint(3) DEFAULT '0' COMMENT '订单状态',
  `time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `meal_date` (`meal_date`,`meal_code`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户点餐表';

-- ----------------------------
-- Records of rms_user_order
-- ----------------------------
INSERT INTO `rms_user_order` VALUES ('1', '1', '0', '1,2', '0', '2313', null, '1', null);
