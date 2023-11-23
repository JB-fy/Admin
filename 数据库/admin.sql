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

 Date: 23/11/2023 17:27:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `actionId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作ID',
  `actionName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `actionCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`) USING BTREE,
  UNIQUE INDEX `actionCode`(`actionCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES (1, '权限场景-查看', 'authSceneLook', '', 0, '2023-07-01 15:38:45', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (2, '权限场景-新增', 'authSceneCreate', '', 0, '2023-07-01 15:38:46', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (3, '权限场景-编辑', 'authSceneUpdate', '', 0, '2023-07-01 15:38:49', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (4, '权限场景-删除', 'authSceneDelete', '', 0, '2023-07-01 15:38:50', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (5, '权限操作-查看', 'authActionLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (6, '权限操作-新增', 'authActionCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (7, '权限操作-编辑', 'authActionUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (8, '权限操作-删除', 'authActionDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (9, '权限菜单-查看', 'authMenuLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (10, '权限菜单-新增', 'authMenuCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (11, '权限菜单-编辑', 'authMenuUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (12, '权限菜单-删除', 'authMenuDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (13, '权限角色-查看', 'authRoleLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (14, '权限角色-新增', 'authRoleCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (15, '权限角色-编辑', 'authRoleUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (16, '权限角色-删除', 'authRoleDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (17, '平台管理员-查看', 'platformAdminLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (18, '平台管理员-新增', 'platformAdminCreate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (19, '平台管理员-编辑', 'platformAdminUpdate', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (20, '平台管理员-删除', 'platformAdminDelete', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (21, '平台配置-查看', 'platformConfigLook', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (22, '平台配置-保存', 'platformConfigSave', '', 0, '2023-06-09 12:03:29', '2023-06-09 12:03:30');
INSERT INTO `auth_action` VALUES (23, '用户-查看', 'userLook', '', 0, '2023-10-14 16:03:55', '2023-10-14 15:32:37');
INSERT INTO `auth_action` VALUES (24, '用户-编辑', 'userUpdate', '', 0, '2023-10-14 16:03:59', '2023-10-14 15:32:37');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `actionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`, `sceneId`) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO `auth_action_rel_to_scene` VALUES (1, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (2, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (3, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (4, 1, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (5, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (6, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (7, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (8, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (9, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (10, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (11, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (12, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (13, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (14, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (15, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (16, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (17, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (18, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (19, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (20, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (21, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (22, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_action_rel_to_scene` VALUES (23, 1, '2023-10-14 15:32:37', '2023-10-14 15:32:37');
INSERT INTO `auth_action_rel_to_scene` VALUES (24, 1, '2023-10-14 15:32:37', '2023-10-14 15:32:37');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `menuId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menuName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `idPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `menuIcon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menuUrl` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `extraData` json NULL COMMENT '额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值。从小到大排序，默认50，范围0-100',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menuId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, '主页', 1, 0, 1, '0-1', 'AutoiconEpHomeFilled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 0, 0, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (2, '权限管理', 1, 0, 1, '0-2', 'AutoiconEpLock', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 90, 0, '2023-09-20 23:02:07', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (3, '场景', 1, 2, 2, '0-2-3', 'AutoiconEpFlag', '/auth/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"场景\"}}}', 100, 0, '2023-06-30 17:51:35', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (4, '操作', 1, 2, 2, '0-2-4', 'AutoiconEpCoordinate', '/auth/action', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 90, 0, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (5, '菜单', 1, 2, 2, '0-2-5', 'AutoiconEpMenu', '/auth/menu', '{\"i18n\": {\"title\": {\"en\": \"Menu\", \"zh-cn\": \"菜单\"}}}', 80, 0, '2023-09-20 23:02:10', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (6, '角色', 1, 2, 2, '0-2-6', 'AutoiconEpView', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 70, 0, '2023-09-20 23:02:11', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (7, '平台管理员', 1, 2, 2, '0-2-7', 'Vant-manager-o', '/platform/admin', '{\"i18n\": {\"title\": {\"en\": \"Platform Admin\", \"zh-cn\": \"平台管理员\"}}}', 60, 0, '2023-10-14 16:18:16', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (8, '系统管理', 1, 0, 1, '0-8', 'AutoiconEpPlatform', '', '{\"i18n\": {\"title\": {\"en\": \"System Manage\", \"zh-cn\": \"系统管理\"}}}', 85, 0, '2023-09-20 23:02:14', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (9, '配置中心', 1, 8, 2, '0-8-9', 'AutoiconEpSetting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 100, 0, '2023-06-24 17:30:46', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (10, '平台配置', 1, 9, 3, '0-8-9-10', '', '/platform/config/platform', '{\"i18n\": {\"title\": {\"en\": \"Platform Config\", \"zh-cn\": \"平台配置\"}}}', 50, 0, '2023-11-22 09:31:42', '2023-06-09 12:03:30');
INSERT INTO `auth_menu` VALUES (11, '插件配置', 1, 9, 3, '0-8-9-11', '', '/platform/config/plugin', '{\"i18n\": {\"title\": {\"en\": \"Plugin Config\", \"zh-cn\": \"插件配置\"}}}', 50, 0, '2023-11-22 09:31:45', '2023-11-21 11:08:28');
INSERT INTO `auth_menu` VALUES (12, '用户管理', 1, 0, 1, '0-12', 'Vant-friends', '', '{\"i18n\": {\"title\": {\"en\": \"User Manage\", \"zh-cn\": \"用户管理\"}}}', 50, 0, '2023-11-21 11:10:29', '2023-10-14 15:32:37');
INSERT INTO `auth_menu` VALUES (13, '用户', 1, 12, 2, '0-12-13', 'Vant-user-o', '/user/user', '{\"i18n\": {\"title\": {\"en\": \"User\", \"zh-cn\": \"用户\"}}}', 50, 0, '2023-11-21 11:10:24', '2023-10-14 15:32:37');

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role`  (
  `roleId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `roleName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `tableId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `tableId`(`tableId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES (1, '超级管理员', 1, 0, 0, '2023-10-14 16:22:51', '2023-06-09 12:03:30');

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `adminId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
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
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `actionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `actionId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `actionId`(`actionId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限操作关联表（角色包含哪些操作）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------
INSERT INTO `auth_role_rel_to_action` VALUES (1, 1, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 2, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 3, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 4, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 5, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 6, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 7, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 8, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 9, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 10, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 11, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 12, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 13, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 14, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 15, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 16, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 17, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 18, '2023-06-11 23:44:22', '2023-06-11 23:44:22');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 19, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 20, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 21, '2023-06-11 14:52:39', '2023-06-11 14:52:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 22, '2023-06-11 14:52:39', '2023-06-11 14:52:39');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 23, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
INSERT INTO `auth_role_rel_to_action` VALUES (1, 24, '2023-10-14 16:22:51', '2023-10-14 16:22:51');

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_menu`;
CREATE TABLE `auth_role_rel_to_menu`  (
  `roleId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menuId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `menuId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `menuId`(`menuId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限菜单关联表（角色包含哪些菜单）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 1, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 2, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 3, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 4, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 5, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 6, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 7, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 8, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 9, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 10, '2023-06-09 12:03:30', '2023-06-09 12:03:30');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 11, '2023-10-14 16:22:51', '2023-10-14 16:22:51');
INSERT INTO `auth_role_rel_to_menu` VALUES (1, 12, '2023-10-14 16:22:51', '2023-10-14 16:22:51');

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_scene`;
CREATE TABLE `auth_scene`  (
  `sceneId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '场景ID',
  `sceneName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `sceneCode` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `sceneConfig` json NOT NULL COMMENT '配置。JSON格式：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sceneId`) USING BTREE,
  UNIQUE INDEX `sceneCode`(`sceneCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES (1, '平台后台', 'platform', '{\"signKey\": \"www.admin.com_platform\", \"signType\": \"HS256\", \"expireTime\": 14400}', 0, '2023-06-14 21:52:16', '2023-06-09 12:03:30');
INSERT INTO `auth_scene` VALUES (2, 'APP', 'app', '{\"signKey\": \"www.admin.com_app\", \"signType\": \"HS256\", \"expireTime\": 604800}', 0, '2023-10-21 17:21:15', '2023-10-21 17:21:15');

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `platform_admin`;
CREATE TABLE `platform_admin`  (
  `adminId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO `platform_admin` VALUES (1, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB', '超级管理员', 'http://www.admin.com/common/20230920/1695222477127_79698554.png', 0, '2023-09-20 23:09:17', '2023-06-09 12:03:30');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config`  (
  `configId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配置ID',
  `configKey` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置Key',
  `configValue` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `configKey`(`configKey` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------
INSERT INTO `platform_config` VALUES (1, 'uploadType', 'local', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (2, 'localUploadUrl', 'http://www.admin.com/upload/upload', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (3, 'localUploadSignKey', '123456', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (4, 'localUploadFileSaveDir', '../public/', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (5, 'localUploadFileUrlPrefix', 'http://www.admin.com', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (6, 'aliyunOssHost', 'https://oss-cn-hangzhou.aliyuncs.com', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (7, 'aliyunOssBucket', 'bucket', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (8, 'aliyunOssAccessKeyId', 'accessKeyId', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (9, 'aliyunOssAccessKeySecret', 'accessKeySecret', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (10, 'aliyunOssCallbackUrl', 'https://www.xxxx.com/upload/notify', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (11, 'aliyunOssEndpoint', 'sts.cn-hangzhou.aliyuncs.com', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (12, 'aliyunOssRoleArn', 'acs:ram::xxxxxxxxxxxxxxxx:role/aliyunosstokengeneratorrole', '2023-11-23 17:20:44', '2023-11-22 09:45:17');
INSERT INTO `platform_config` VALUES (13, 'smsType', 'aliyunSms', '2023-11-22 09:49:39', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (14, 'aliyunSmsAccessKeyId', 'accessKeyId', '2023-11-22 09:49:41', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (15, 'aliyunSmsAccessKeySecret', 'accessKeySecret', '2023-11-22 09:49:43', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (16, 'aliyunSmsEndpoint', 'dysmsapi.aliyuncs.com', '2023-11-22 09:49:53', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (17, 'aliyunSmsSignName', 'JB Admin', '2023-11-22 09:50:09', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (18, 'aliyunSmsTemplateCode', 'SMS_xxxxxxxx', '2023-11-22 09:50:07', '2023-11-22 09:49:18');
INSERT INTO `platform_config` VALUES (19, 'idCardType', 'aliyunIdCard', '2023-11-22 09:50:29', '2023-11-22 09:50:23');
INSERT INTO `platform_config` VALUES (20, 'aliyunIdCardHost', 'http://idcard.market.alicloudapi.com', '2023-11-22 09:50:33', '2023-11-22 09:50:23');
INSERT INTO `platform_config` VALUES (21, 'aliyunIdCardPath', '/lianzhuo/idcard', '2023-11-22 09:50:35', '2023-11-22 09:50:23');
INSERT INTO `platform_config` VALUES (22, 'aliyunIdCardAppcode', 'appcode', '2023-11-22 09:50:37', '2023-11-22 09:50:23');
INSERT INTO `platform_config` VALUES (23, 'pushType', 'txTpns', '2023-11-23 17:26:12', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (24, 'txTpnsHost', 'https://api.tpns.tencent.com', '2023-11-23 17:26:14', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (25, 'txTpnsAccessIDOfAndroid', '150xxxx000', '2023-11-23 17:26:16', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (26, 'txTpnsSecretKeyOfAndroid', '80e5xxxxxxxxxxxxxxxxxxxxxxxx0000', '2023-11-23 17:26:17', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (27, 'txTpnsAccessIDOfIos', '150xxxx001', '2023-11-23 17:26:19', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (28, 'txTpnsSecretKeyOfIos', '80e5xxxxxxxxxxxxxxxxxxxxxxxx0001', '2023-11-23 17:26:22', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (29, 'txTpnsAccessIDOfMacOS', '150xxxx002', '2023-11-23 17:26:24', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (30, 'txTpnsSecretKeyOfMacOS', '80e5xxxxxxxxxxxxxxxxxxxxxxxx0002', '2023-11-23 17:26:26', '2023-11-23 17:23:50');
INSERT INTO `platform_config` VALUES (31, 'vodType', 'aliyunVod', '2023-11-23 17:26:29', '2023-11-23 17:25:36');
INSERT INTO `platform_config` VALUES (32, 'aliyunVodAccessKeyId', 'accessKeyId', '2023-11-23 17:26:33', '2023-11-23 17:25:36');
INSERT INTO `platform_config` VALUES (33, 'aliyunVodAccessKeySecret', 'accessKeySecret', '2023-11-23 17:26:35', '2023-11-23 17:25:36');
INSERT INTO `platform_config` VALUES (34, 'aliyunVodEndpoint', 'sts.cn-shanghai.aliyuncs.com', '2023-11-23 17:26:37', '2023-11-23 17:25:36');
INSERT INTO `platform_config` VALUES (35, 'aliyunVodRoleArn', 'acs:ram::xxxxxxxxxxxxxxxx:role/aliyunvodtokengeneratorrole', '2023-11-23 17:26:40', '2023-11-23 17:25:36');

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS `platform_server`;
CREATE TABLE `platform_server`  (
  `serverId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '服务器ID',
  `networkIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公网IP',
  `localIp` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内网IP',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`serverId`) USING BTREE,
  UNIQUE INDEX `networkIp`(`networkIp` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台服务器表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别：0未设置 1男 2女',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `address` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `idCardName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证姓名',
  `idCardNo` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updatedAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createdAt` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
