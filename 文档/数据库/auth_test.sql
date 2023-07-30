/*
 Navicat Premium Data Transfer

 Source Server         : 本地-Mysql8
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.2.200:3306
 Source Schema         : dev_admin

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 30/07/2023 09:08:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_test
-- ----------------------------
DROP TABLE IF EXISTS `auth_test`;
CREATE TABLE `auth_test`  (
  `testId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限菜单ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `testName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `testIcon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `testUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `idPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `extraData` json NULL COMMENT '额外数据。（json格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0待处理 1已处理 2驳回 3哈哈 4呵呵 5嘿嘿',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话号码',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别：0未知 1男 2女',
  `imgList` json NOT NULL COMMENT '图片集',
  `video` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '视频',
  `video_list` json NOT NULL COMMENT '视频集',
  PRIMARY KEY (`testId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_test
-- ----------------------------
INSERT INTO `auth_test` VALUES (1, 1, 0, '主页', 'AutoiconEpHomeFilled', '/', 1, '0-1', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 0, 0, '2023-07-01 23:05:22', '2023-06-09 12:03:30', 0, NULL, NULL, '', '', '', '', 0, '[\"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230624/1687601922719_8091.gif?w=471&h=373\"]', 'https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688223859500_2322.mp4', '[\"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688223859500_2322.mp4\"]');
INSERT INTO `auth_test` VALUES (29, 0, 0, '12314', '', '', 0, '', NULL, 50, 0, '2023-07-01 23:24:18', '2023-07-01 23:24:18', 0, '', NULL, 'e10adc3949ba59abbe56e057f20f883e', '', 'https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225054842_4556.gif', '', 0, '[\"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225046484_7375.gif\", \"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225048766_4467.gif\"]', 'https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225038504_8366.mp4', '[\"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225041427_9597.mp4\", \"https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230701/1688225044074_6458.mp4\"]');

SET FOREIGN_KEY_CHECKS = 1;
