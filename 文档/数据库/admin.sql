/*
 Navicat Premium Data Transfer

 Source Server         : 本地-Mysql8
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.2.200:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 11/05/2023 18:30:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `actionId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限操作ID',
  `actionName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `actionCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标识（代码中用于判断权限）',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`) USING BTREE,
  UNIQUE INDEX `actionCode`(`actionCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES (1, '权限场景-查看', 'authSceneLook', '', 0, '2022-12-26 21:17:51', '2022-12-26 21:17:51');
INSERT INTO `auth_action` VALUES (2, '权限场景-新增', 'authSceneCreate', '', 0, '2022-12-26 21:18:28', '2022-12-26 21:18:28');
INSERT INTO `auth_action` VALUES (3, '权限场景-编辑', 'authSceneUpdate', '', 0, '2022-12-26 21:18:52', '2022-12-26 21:18:52');
INSERT INTO `auth_action` VALUES (4, '权限场景-删除', 'authSceneDelete', '', 0, '2022-12-26 21:19:17', '2022-12-26 21:19:17');
INSERT INTO `auth_action` VALUES (5, '权限操作-查看', 'authActionLook', '', 0, '2022-12-26 22:47:05', '2022-12-26 22:47:05');
INSERT INTO `auth_action` VALUES (6, '权限操作-新增', 'authActionCreate', '', 0, '2022-12-26 22:47:20', '2022-12-26 22:47:20');
INSERT INTO `auth_action` VALUES (7, '权限操作-编辑', 'authActionUpdate', '', 0, '2022-12-26 22:47:32', '2022-12-26 22:47:32');
INSERT INTO `auth_action` VALUES (8, '权限操作-删除', 'authActionDelete', '', 0, '2022-12-26 22:47:43', '2022-12-26 22:47:43');
INSERT INTO `auth_action` VALUES (9, '权限菜单-查看', 'authMenuLook', '', 0, '2022-12-26 22:48:42', '2022-12-26 22:48:42');
INSERT INTO `auth_action` VALUES (10, '权限菜单-新增', 'authMenuCreate', '', 0, '2022-12-26 22:48:57', '2022-12-26 22:48:57');
INSERT INTO `auth_action` VALUES (11, '权限菜单-编辑', 'authMenuUpdate', '', 0, '2022-12-26 22:49:07', '2022-12-26 22:49:07');
INSERT INTO `auth_action` VALUES (12, '权限菜单-删除', 'authMenuDelete', '', 0, '2022-12-26 22:49:16', '2022-12-26 22:49:16');
INSERT INTO `auth_action` VALUES (13, '权限角色-查看', 'authRoleLook', '', 0, '2022-12-26 22:49:59', '2022-12-26 22:49:59');
INSERT INTO `auth_action` VALUES (14, '权限角色-新增', 'authRoleCreate', '', 0, '2022-12-26 22:50:11', '2022-12-26 22:50:11');
INSERT INTO `auth_action` VALUES (15, '权限角色-编辑', 'authRoleUpdate', '', 0, '2022-12-26 22:50:21', '2022-12-26 22:50:21');
INSERT INTO `auth_action` VALUES (16, '权限角色-删除', 'authRoleDelete', '', 0, '2022-12-26 22:50:29', '2022-12-26 22:50:29');
INSERT INTO `auth_action` VALUES (17, '平台管理员-查看', 'platformAdminLook', '', 0, '2022-12-26 22:58:33', '2022-12-26 22:58:33');
INSERT INTO `auth_action` VALUES (18, '平台管理员-新增', 'platformAdminCreate', '', 0, '2022-12-26 22:59:56', '2022-12-26 22:59:56');
INSERT INTO `auth_action` VALUES (19, '平台管理员-编辑', 'platformAdminUpdate', '', 0, '2022-12-26 23:00:12', '2022-12-26 23:00:12');
INSERT INTO `auth_action` VALUES (20, '平台管理员-删除', 'platformAdminDelete', '', 0, '2022-12-26 23:00:24', '2022-12-26 23:00:24');
INSERT INTO `auth_action` VALUES (21, '平台配置-查看', 'platformConfigLook', '', 0, '2023-02-12 23:38:19', '2023-02-12 23:38:14');
INSERT INTO `auth_action` VALUES (22, '平台配置-保存', 'platformConfigSave', '', 0, '2023-02-12 23:38:35', '2023-02-12 23:38:35');
INSERT INTO `auth_action` VALUES (23, '请求日志-查看', 'logRequestLook', '', 0, '2023-02-12 23:40:06', '2023-02-12 23:40:06');
INSERT INTO `auth_action` VALUES (24, '服务器-查看', 'platformServerLook', '', 0, '2023-02-12 23:40:39', '2023-02-12 23:40:39');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `actionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限操作ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`, `sceneId`) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO `auth_action_rel_to_scene` VALUES (1, 1, '2022-12-26 21:17:51', '2022-12-26 21:17:51');
INSERT INTO `auth_action_rel_to_scene` VALUES (2, 1, '2022-12-26 21:18:28', '2022-12-26 21:18:28');
INSERT INTO `auth_action_rel_to_scene` VALUES (3, 1, '2022-12-26 21:18:52', '2022-12-26 21:18:52');
INSERT INTO `auth_action_rel_to_scene` VALUES (4, 1, '2022-12-26 21:19:17', '2022-12-26 21:19:17');
INSERT INTO `auth_action_rel_to_scene` VALUES (5, 1, '2022-12-26 22:47:05', '2022-12-26 22:47:05');
INSERT INTO `auth_action_rel_to_scene` VALUES (6, 1, '2022-12-26 22:47:20', '2022-12-26 22:47:20');
INSERT INTO `auth_action_rel_to_scene` VALUES (7, 1, '2022-12-26 22:47:32', '2022-12-26 22:47:32');
INSERT INTO `auth_action_rel_to_scene` VALUES (8, 1, '2022-12-26 22:47:43', '2022-12-26 22:47:43');
INSERT INTO `auth_action_rel_to_scene` VALUES (9, 1, '2022-12-26 22:48:42', '2022-12-26 22:48:42');
INSERT INTO `auth_action_rel_to_scene` VALUES (10, 1, '2022-12-26 22:48:57', '2022-12-26 22:48:57');
INSERT INTO `auth_action_rel_to_scene` VALUES (11, 1, '2022-12-26 22:49:07', '2022-12-26 22:49:07');
INSERT INTO `auth_action_rel_to_scene` VALUES (12, 1, '2022-12-26 22:49:16', '2022-12-26 22:49:16');
INSERT INTO `auth_action_rel_to_scene` VALUES (13, 1, '2022-12-26 22:49:59', '2022-12-26 22:49:59');
INSERT INTO `auth_action_rel_to_scene` VALUES (14, 1, '2022-12-26 22:50:11', '2022-12-26 22:50:11');
INSERT INTO `auth_action_rel_to_scene` VALUES (15, 1, '2022-12-26 22:50:21', '2022-12-26 22:50:21');
INSERT INTO `auth_action_rel_to_scene` VALUES (16, 1, '2022-12-26 22:50:29', '2022-12-26 22:50:29');
INSERT INTO `auth_action_rel_to_scene` VALUES (17, 1, '2022-12-26 22:58:33', '2022-12-26 22:58:33');
INSERT INTO `auth_action_rel_to_scene` VALUES (18, 1, '2022-12-26 22:59:56', '2022-12-26 22:59:56');
INSERT INTO `auth_action_rel_to_scene` VALUES (19, 1, '2022-12-26 23:00:12', '2022-12-26 23:00:12');
INSERT INTO `auth_action_rel_to_scene` VALUES (20, 1, '2022-12-26 23:00:24', '2022-12-26 23:00:24');
INSERT INTO `auth_action_rel_to_scene` VALUES (21, 1, '2023-02-12 23:38:14', '2023-02-12 23:38:14');
INSERT INTO `auth_action_rel_to_scene` VALUES (22, 1, '2023-02-12 23:38:35', '2023-02-12 23:38:35');
INSERT INTO `auth_action_rel_to_scene` VALUES (23, 1, '2023-02-12 23:40:06', '2023-02-12 23:40:06');
INSERT INTO `auth_action_rel_to_scene` VALUES (24, 1, '2023-02-12 23:40:39', '2023-02-12 23:40:39');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `menuId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限菜单ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `menuName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `menuIcon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menuUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `pidPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `extraData` json NULL COMMENT '额外数据。（json格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menuId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, 1, 0, '主页', 'AutoiconEpHomeFilled', '/', 1, '0-1', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 0, 0, '2023-05-11 17:40:12', '2022-12-25 23:28:45');
INSERT INTO `auth_menu` VALUES (2, 1, 0, '权限管理', 'AutoiconEpLock', '', 2, '0-2', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage \", \"zh-cn\": \"权限管理\"}}}', 100, 0, '2023-05-11 17:40:21', '2022-12-25 23:31:28');
INSERT INTO `auth_menu` VALUES (3, 1, 2, '场景', 'AutoiconEpFlag', '/auth/scene', 2, '0-2-3', '{\"i18n\": {\"title\": {\"en\": \"Scene \", \"zh-cn\": \"场景\"}}}', 100, 0, '2023-05-11 17:40:27', '2022-12-25 23:37:07');
INSERT INTO `auth_menu` VALUES (4, 1, 2, '操作', 'AutoiconEpCoordinate', '/auth/action', 2, '0-2-4', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 90, 0, '2023-05-11 17:40:37', '2022-12-26 21:12:00');
INSERT INTO `auth_menu` VALUES (5, 1, 2, '菜单', 'AutoiconEpMenu', '/auth/menu', 2, '0-2-5', '{\"i18n\": {\"title\": {\"en\": \"Menu \", \"zh-cn\": \"菜单\"}}}', 80, 0, '2023-05-11 17:40:45', '2022-12-25 23:35:48');
INSERT INTO `auth_menu` VALUES (6, 1, 2, '角色', 'AutoiconEpView', '/auth/role', 2, '0-2-6', '{\"i18n\": {\"title\": {\"en\": \"Role \", \"zh-cn\": \"角色\"}}}', 70, 0, '2023-05-11 17:41:55', '2022-12-25 23:35:57');
INSERT INTO `auth_menu` VALUES (7, 1, 2, '平台管理员', 'AutoiconEpUserFilled', '/platform/admin', 2, '0-2-7', '{\"i18n\": {\"title\": {\"en\": \"Platform Admin \", \"zh-cn\": \"平台管理员\"}}}', 60, 0, '2023-05-11 17:42:04', '2022-12-25 23:35:16');
INSERT INTO `auth_menu` VALUES (8, 1, 0, '系统管理', 'AutoiconEpPlatform', '', 1, '0-8', '{\"i18n\": {\"title\": {\"en\": \"System Manage \", \"zh-cn\": \"系统管理\"}}}', 90, 0, '2023-05-11 17:42:09', '2022-12-25 23:36:04');
INSERT INTO `auth_menu` VALUES (9, 1, 8, '配置中心', 'AutoiconEpSetting', '', 2, '0-8-9', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 100, 0, '2023-05-11 17:42:15', '2023-01-12 22:16:29');
INSERT INTO `auth_menu` VALUES (10, 1, 9, '平台配置', '', '/platform/config', 3, '0-8-9-10', '{\"i18n\": {\"title\": {\"en\": \"Platform Config \", \"zh-cn\": \"平台配置\"}}}', 50, 0, '2023-05-11 17:42:22', '2022-12-25 23:36:33');
INSERT INTO `auth_menu` VALUES (11, 1, 0, '日志管理', 'AutoiconEpDataAnalysis', '', 1, '0-11', '{\"i18n\": {\"title\": {\"en\": \"Log Manage \", \"zh-cn\": \"日志管理\"}}}', 80, 0, '2023-05-11 17:42:29', '2022-12-25 23:36:38');
INSERT INTO `auth_menu` VALUES (12, 1, 11, '请求日志', 'AutoiconEpReading', '/log/request', 2, '0-11-12', '{\"i18n\": {\"title\": {\"en\": \"Request Log \", \"zh-cn\": \"请求日志\"}}}', 50, 0, '2023-05-11 17:42:36', '2022-12-25 23:36:45');
INSERT INTO `auth_menu` VALUES (13, 1, 8, '服务器', 'AutoiconEpCpu', '/platform/server', 2, '0-8-13', '{\"i18n\": {\"title\": {\"en\": \"Server \", \"zh-cn\": \"服务器\"}}}', 100, 0, '2023-05-11 18:23:25', '2023-02-12 23:33:42');

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
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `tableId`(`tableId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES (1, 1, 0, '超级管理员', 0, '2023-01-13 00:15:39', '2023-01-13 00:15:39');

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色ID',
  `adminId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '平台管理员ID',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `adminId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `adminId`(`adminId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，系统管理员关联表（系统管理员包含哪些角色）' ROW_FORMAT = DYNAMIC;

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
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `actionId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限操作关联表（角色包含哪些操作）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------
INSERT INTO `auth_role_rel_to_action` VALUES (1, 1, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 2, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 3, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 4, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 5, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 6, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 7, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 8, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 9, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 10, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 11, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 12, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 13, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 14, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 15, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 16, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 17, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 18, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 19, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 20, '2023-01-13 00:15:39', '2023-01-13 00:15:39');

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_menu`;
CREATE TABLE `auth_role_rel_to_menu`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限角色ID',
  `menuId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限菜单ID',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `menuId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `menuId`(`menuId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限菜单关联表（角色包含哪些菜单）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 1, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 2, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 3, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 4, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 5, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 6, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 7, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 8, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 9, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 10, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 11, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 12, '2023-01-13 00:15:39', '2023-01-13 00:15:39');

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
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sceneId`) USING BTREE,
  UNIQUE INDEX `sceneCode`(`sceneCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES (1, 'platformAdmin', '平台后台', '{\"signKey\": \"www.admin.com_platform\", \"signType\": \"HS256\", \"expireTime\": 14400}', 0, '2022-12-06 22:51:16', '2022-09-17 23:13:53');

-- ----------------------------
-- Table structure for log_request
-- ----------------------------
DROP TABLE IF EXISTS `log_request`;
CREATE TABLE `log_request`  (
  `logId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '请求日志ID',
  `requestUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '请求地址',
  `requestHeader` json NULL COMMENT '请求头',
  `requestData` json NULL COMMENT '请求数据',
  `responseBody` json NULL COMMENT '响应体',
  `runTime` decimal(8, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '运行时间（单位：毫秒）',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`logId`, `createAt`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '系统日志-请求表' ROW_FORMAT = DYNAMIC;

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
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO `platform_admin` VALUES (1, 'admin', NULL, 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'https://gamemt.oss-cn-hangzhou.aliyuncs.com/common/2023/01/12/204205_2094_16735273287104264.gif', 0, '2023-01-12 20:42:10', '2022-09-04 22:53:41');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config`  (
  `configId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `configKey` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置项Key',
  `configValue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置项值（设置大点。以后可能需要保存富文本内容，如公司简介或协议等等）',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `configKey`(`configKey` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS `platform_server`;
CREATE TABLE `platform_server`  (
  `serverId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '服务器ID',
  `networkIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公网IP',
  `localIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内网IP',
  `updateAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`serverId`) USING BTREE,
  UNIQUE INDEX `networkIp`(`networkIp` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台服务器表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_server
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
