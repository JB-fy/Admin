/*
 Navicat Premium Data Transfer

 Source Server         : 本地Mysql-8.0.30
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : 192.168.200.210:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 05/12/2022 04:26:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `actionId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限操作ID',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID（主要用于归类，方便查看。否则可以不要）',
  `actionName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `actionCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标识（代码中用于判断权限）',
  `pidPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`) USING BTREE,
  UNIQUE INDEX `actionCode`(`actionCode` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `actionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限操作ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`, `sceneId`) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `menuId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限菜单ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `menuName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `pidPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `extraData` json NULL COMMENT '额外数据。（json格式：{\"title（多语言时设置，未设置以menuName返回）\": {\"语言标识\":\"标题\",...},\"icon\": \"图标\",\"url\": \"链接地址\",...}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menuId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, 1, 0, '主页', '', 0, '{\"url\": \"/\", \"icon\": \"AutoiconEpHomeFilled\", \"title\": {\"en\": \"homepage\", \"zh-cn\": \"主页\"}}', 0, 0, '2022-11-03 23:30:30', '2022-09-17 23:46:20');
INSERT INTO `auth_menu` VALUES (2, 1, 0, '权限管理', '', 0, '{\"icon\": \"AutoiconEpLock\", \"title\": {\"en\": \"authManage \", \"zh-cn\": \"权限管理\"}}', 100, 0, '2022-11-04 00:38:28', '2022-09-17 23:49:51');
INSERT INTO `auth_menu` VALUES (3, 1, 2, '管理员', '', 0, '{\"url\": \"/platform/admin\", \"icon\": \"AutoiconEpUserFilled\", \"title\": {\"en\": \"admin \", \"zh-cn\": \"管理员\"}}', 50, 0, '2022-12-05 22:46:33', '2022-09-17 23:49:53');
INSERT INTO `auth_menu` VALUES (4, 1, 2, '菜单', '', 0, '{\"url\": \"/auth/menu\", \"icon\": \"AutoiconEpMenu\", \"title\": {\"en\": \"menu \", \"zh-cn\": \"菜单\"}}', 100, 0, '2022-12-05 22:46:38', '2022-09-17 23:49:54');
INSERT INTO `auth_menu` VALUES (5, 1, 2, '角色', '', 0, '{\"url\": \"/auth/role\", \"icon\": \"AutoiconEpView\", \"title\": {\"en\": \"role \", \"zh-cn\": \"角色\"}}', 50, 0, '2022-12-05 22:46:42', '2022-09-17 23:49:55');
INSERT INTO `auth_menu` VALUES (6, 1, 0, '系统管理', '', 0, '{\"icon\": \"AutoiconEpPlatform\", \"title\": {\"en\": \"systemManage \", \"zh-cn\": \"系统管理\"}}', 99, 0, '2022-11-04 00:38:16', '2022-09-17 23:49:56');
INSERT INTO `auth_menu` VALUES (7, 1, 6, '系统配置', '', 0, '{\"url\": \"/platform/config\", \"icon\": \"AutoiconEpSetting\", \"title\": {\"en\": \"systemConfig \", \"zh-cn\": \"系统配置\"}}', 50, 0, '2022-12-05 22:46:47', '2022-09-17 23:49:57');
INSERT INTO `auth_menu` VALUES (8, 1, 0, '日志管理', '', 0, '{\"icon\": \"AutoiconEpDataAnalysis\", \"title\": {\"en\": \"logManage \", \"zh-cn\": \"日志管理\"}}', 95, 0, '2022-11-04 00:36:00', '2022-09-17 23:49:59');
INSERT INTO `auth_menu` VALUES (9, 1, 8, '请求日志', '', 0, '{\"url\": \"/log/request\", \"icon\": \"AutoiconEpReading\", \"title\": {\"en\": \"requestLog \", \"zh-cn\": \"请求日志\"}}', 50, 0, '2022-12-05 22:46:59', '2022-09-17 23:50:06');
INSERT INTO `auth_menu` VALUES (10, 1, 2, '场景', '', 0, '{\"url\": \"/auth/scene\", \"icon\": \"AutoiconEpFlag\", \"title\": {\"en\": \"scene \", \"zh-cn\": \"场景\"}}', 100, 0, '2022-12-05 22:47:03', '2022-09-17 23:49:54');
INSERT INTO `auth_menu` VALUES (11, 1, 0, '第三方网站', '', 0, '{\"url\": \"/thirdUrl?url=https://cn.vuejs.org/api/\", \"icon\": \"AutoiconEpChromeFilled\", \"title\": {\"en\": \"thridWebsite \", \"zh-cn\": \"第三方网站\"}}', 100, 0, '2022-11-03 01:11:33', '2022-10-17 00:25:42');
INSERT INTO `auth_menu` VALUES (12, 1, 0, '第三方网站（新标签）', '', 0, '{\"url\": \"https://www.baidu.com/\", \"icon\": \"AutoiconEpChromeFilled\", \"title\": {\"en\": \"test \", \"zh-cn\": \"第三方网站（新标签）\"}}', 100, 0, '2022-11-03 01:11:36', '2022-10-17 00:25:42');
INSERT INTO `auth_menu` VALUES (13, 1, 0, '第三方网站1', '', 0, '{\"url\": \"/thirdUrl?url=https://element-plus.gitee.io/zh-CN/\", \"icon\": \"AutoiconEpElementPlus\", \"title\": {\"en\": \"thridWebsite1 \", \"zh-cn\": \"第三方网站1\"}}', 100, 0, '2022-11-03 23:45:56', '2022-10-17 00:25:42');

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role`  (
  `roleId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限角色ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID',
  `tableId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）',
  `roleName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `tableId`(`tableId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色ID',
  `adminId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '平台管理员ID',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `adminId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `adminId`(`adminId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色，系统管理员关联表（系统管理员包含哪些角色）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_action`;
CREATE TABLE `auth_role_rel_to_action`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色ID',
  `actionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限操作ID',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `actionId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色，权限操作关联表（角色包含哪些操作）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_menu`;
CREATE TABLE `auth_role_rel_to_menu`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色ID',
  `menuId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限菜单ID',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `menuId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `menuId`(`menuId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色，权限菜单关联表（角色包含哪些菜单）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_scene`;
CREATE TABLE `auth_scene`  (
  `sceneId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限场景ID',
  `sceneCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识（代码中用于识别调用接口的所在场景，做对应的身份鉴定及权力鉴定。如已在代码中使用，不建议更改）',
  `sceneName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneConfig` json NULL COMMENT '配置（内容自定义。json格式：{\"alg\": \"算法\",\"key\": \"密钥\",\"expTime\": \"签名有效时间\",...}）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sceneId`) USING BTREE,
  UNIQUE INDEX `sceneCode`(`sceneCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 114 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES (1, 'platformAdmin', '系统后台', '{\"signKey\": \"www.admin.com_platform\", \"signType\": \"HS256\", \"expireTime\": 14400}', 0, '2022-12-05 22:17:32', '2022-09-17 23:13:53');
INSERT INTO `auth_scene` VALUES (2, 'systemAdmin', '系统后台', '{\"signKey\": \"www.admin.com_system\", \"signType\": \"HS256\", \"expireTime\": 14400}', 0, '2022-12-05 22:20:03', '2022-11-17 23:51:32');

-- ----------------------------
-- Table structure for log_request
-- ----------------------------
DROP TABLE IF EXISTS `log_request`;
CREATE TABLE `log_request`  (
  `logId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '请求日志ID',
  `requestUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求地址',
  `requestData` json NULL COMMENT '请求数据',
  `requestHeader` json NULL COMMENT '请求头',
  `responseBody` json NULL COMMENT '响应体',
  `runTime` decimal(8, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '运行时间（单位：毫秒）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`logId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统日志-请求表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_request
-- ----------------------------

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `platform_admin`;
CREATE TABLE `platform_admin`  (
  `adminId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话号码',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO `platform_admin` VALUES (1, 'admin', NULL, 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', '', 0, '2022-09-04 22:54:09', '2022-09-04 22:53:41');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config`  (
  `configId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `configKey` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置项Key',
  `configValue` varchar(15000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `configKey`(`configKey` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
