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

 Date: 13/01/2023 00:16:00
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
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`actionId`) USING BTREE,
  UNIQUE INDEX `actionCode`(`actionCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

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

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `menuId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '权限菜单ID',
  `sceneId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `menuName` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `pidPath` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `extraData` json NULL COMMENT '额外数据。（json格式：{\"title（多语言时设置，未设置以menuName返回）\": {\"语言标识\":\"标题\",...},\"icon\": \"图标\",\"url\": \"链接地址\",...}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menuId`) USING BTREE,
  INDEX `sceneId`(`sceneId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, 1, 0, '主页', 1, '0-1', '{\"url\": \"/\", \"icon\": \"AutoiconEpHomeFilled\", \"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}', 0, 0, '2023-01-12 22:20:27', '2022-12-25 23:28:45');
INSERT INTO `auth_menu` VALUES (2, 1, 0, '权限管理', 2, '0-2', '{\"icon\": \"AutoiconEpLock\", \"title\": {\"en\": \"Auth Manage \", \"zh-cn\": \"权限管理\"}}', 100, 0, '2023-01-12 22:13:54', '2022-12-25 23:31:28');
INSERT INTO `auth_menu` VALUES (3, 1, 2, '场景', 2, '0-2-3', '{\"url\": \"/auth/scene\", \"icon\": \"AutoiconEpFlag\", \"title\": {\"en\": \"Scene \", \"zh-cn\": \"场景\"}}', 100, 0, '2023-01-12 22:07:31', '2022-12-25 23:37:07');
INSERT INTO `auth_menu` VALUES (4, 1, 2, '操作', 2, '0-2-4', '{\"url\": \"/auth/action\", \"icon\": \"AutoiconEpCoordinate\", \"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}', 90, 0, '2023-01-12 22:08:31', '2022-12-26 21:12:00');
INSERT INTO `auth_menu` VALUES (5, 1, 2, '菜单', 2, '0-2-5', '{\"url\": \"/auth/menu\", \"icon\": \"AutoiconEpMenu\", \"title\": {\"en\": \"Menu \", \"zh-cn\": \"菜单\"}}', 80, 0, '2023-01-12 22:08:32', '2022-12-25 23:35:48');
INSERT INTO `auth_menu` VALUES (6, 1, 2, '角色', 2, '0-2-6', '{\"url\": \"/auth/role\", \"icon\": \"AutoiconEpView\", \"title\": {\"en\": \"Role \", \"zh-cn\": \"角色\"}}', 70, 0, '2023-01-12 22:08:34', '2022-12-25 23:35:57');
INSERT INTO `auth_menu` VALUES (7, 1, 2, '平台管理员', 2, '0-2-7', '{\"url\": \"/platform/admin\", \"icon\": \"AutoiconEpUserFilled\", \"title\": {\"en\": \"Platform Admin \", \"zh-cn\": \"平台管理员\"}}', 60, 0, '2023-01-12 22:08:35', '2022-12-25 23:35:16');
INSERT INTO `auth_menu` VALUES (8, 1, 0, '系统管理', 1, '0-8', '{\"icon\": \"AutoiconEpPlatform\", \"title\": {\"en\": \"System Manage \", \"zh-cn\": \"系统管理\"}}', 90, 0, '2023-01-12 22:30:25', '2022-12-25 23:36:04');
INSERT INTO `auth_menu` VALUES (9, 1, 8, '配置中心', 2, '0-8-9', '{\"icon\": \"AutoiconEpSetting\", \"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}', 100, 0, '2023-01-12 22:30:32', '2023-01-12 22:16:29');
INSERT INTO `auth_menu` VALUES (10, 1, 9, '平台配置', 3, '0-8-9-10', '{\"url\": \"/platform/config\", \"title\": {\"en\": \"Platform Config \", \"zh-cn\": \"平台配置\"}}', 50, 0, '2023-01-12 22:30:46', '2022-12-25 23:36:33');
INSERT INTO `auth_menu` VALUES (11, 1, 0, '日志管理', 1, '0-11', '{\"icon\": \"AutoiconEpDataAnalysis\", \"title\": {\"en\": \"Log Manage \", \"zh-cn\": \"日志管理\"}}', 80, 0, '2023-01-12 22:30:51', '2022-12-25 23:36:38');
INSERT INTO `auth_menu` VALUES (12, 1, 11, '请求日志', 2, '0-11-12', '{\"url\": \"/log/request\", \"icon\": \"AutoiconEpReading\", \"title\": {\"en\": \"Request Log \", \"zh-cn\": \"请求日志\"}}', 50, 0, '2023-01-12 22:30:57', '2022-12-25 23:36:45');

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

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
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`roleId`, `menuId`) USING BTREE,
  INDEX `roleId`(`roleId` ASC) USING BTREE,
  INDEX `menuId`(`menuId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限角色，权限菜单关联表（角色包含哪些菜单）' ROW_FORMAT = DYNAMIC;

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
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`sceneId`) USING BTREE,
  UNIQUE INDEX `sceneCode`(`sceneCode` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

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
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`logId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '系统日志-请求表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_request
-- ----------------------------
INSERT INTO `log_request` VALUES (1, 'http://0.0.0.0:20080/login/info', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/platform/admin\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"2\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '[]', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"info\": {\"phone\": null, \"avatar\": \"https://gamemt.oss-cn-hangzhou.aliyuncs.com/common/2023/01/12/204205_2094_16735273287104264.gif\", \"account\": \"admin\", \"adminId\": 1, \"nickname\": \"超级管理员\", \"createTime\": \"2022-09-04 22:53:41\", \"updateTime\": \"2023-01-12 20:42:10\"}}}', 19.230, '2023-01-13 00:14:37', '2023-01-13 00:14:37');
INSERT INTO `log_request` VALUES (2, 'http://0.0.0.0:20080/login/menuTree', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/platform/admin\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"2\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '[]', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"tree\": [{\"pid\": 0, \"url\": \"/\", \"icon\": \"AutoiconEpHomeFilled\", \"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}, \"menuId\": 1, \"children\": [], \"menuName\": \"主页\"}, {\"pid\": 0, \"url\": \"\", \"icon\": \"AutoiconEpDataAnalysis\", \"title\": {\"en\": \"Log Manage \", \"zh-cn\": \"日志管理\"}, \"menuId\": 11, \"children\": [{\"pid\": 11, \"url\": \"/log/request\", \"icon\": \"AutoiconEpReading\", \"title\": {\"en\": \"Request Log \", \"zh-cn\": \"请求日志\"}, \"menuId\": 12, \"children\": [], \"menuName\": \"请求日志\"}], \"menuName\": \"日志管理\"}, {\"pid\": 0, \"url\": \"\", \"icon\": \"AutoiconEpPlatform\", \"title\": {\"en\": \"System Manage \", \"zh-cn\": \"系统管理\"}, \"menuId\": 8, \"children\": [{\"pid\": 8, \"url\": \"\", \"icon\": \"AutoiconEpSetting\", \"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}, \"menuId\": 9, \"children\": [{\"pid\": 9, \"url\": \"/platform/config\", \"icon\": \"\", \"title\": {\"en\": \"Platform Config \", \"zh-cn\": \"平台配置\"}, \"menuId\": 10, \"children\": [], \"menuName\": \"平台配置\"}], \"menuName\": \"配置中心\"}], \"menuName\": \"系统管理\"}, {\"pid\": 0, \"url\": \"\", \"icon\": \"AutoiconEpLock\", \"title\": {\"en\": \"Auth Manage \", \"zh-cn\": \"权限管理\"}, \"menuId\": 2, \"children\": [{\"pid\": 2, \"url\": \"/platform/admin\", \"icon\": \"AutoiconEpUserFilled\", \"title\": {\"en\": \"Platform Admin \", \"zh-cn\": \"平台管理员\"}, \"menuId\": 7, \"children\": [], \"menuName\": \"平台管理员\"}, {\"pid\": 2, \"url\": \"/auth/role\", \"icon\": \"AutoiconEpView\", \"title\": {\"en\": \"Role \", \"zh-cn\": \"角色\"}, \"menuId\": 6, \"children\": [], \"menuName\": \"角色\"}, {\"pid\": 2, \"url\": \"/auth/menu\", \"icon\": \"AutoiconEpMenu\", \"title\": {\"en\": \"Menu \", \"zh-cn\": \"菜单\"}, \"menuId\": 5, \"children\": [], \"menuName\": \"菜单\"}, {\"pid\": 2, \"url\": \"/auth/action\", \"icon\": \"AutoiconEpCoordinate\", \"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}, \"menuId\": 4, \"children\": [], \"menuName\": \"操作\"}, {\"pid\": 2, \"url\": \"/auth/scene\", \"icon\": \"AutoiconEpFlag\", \"title\": {\"en\": \"Scene \", \"zh-cn\": \"场景\"}, \"menuId\": 3, \"children\": [], \"menuName\": \"场景\"}], \"menuName\": \"权限管理\"}]}}', 6.038, '2023-01-13 00:14:37', '2023-01-13 00:14:37');
INSERT INTO `log_request` VALUES (3, 'http://0.0.0.0:20080/platform/admin/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/platform/admin\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 1, \"phone\": null, \"avatar\": \"https://gamemt.oss-cn-hangzhou.aliyuncs.com/common/2023/01/12/204205_2094_16735273287104264.gif\", \"isStop\": 0, \"account\": \"admin\", \"adminId\": 1, \"nickname\": \"超级管理员\", \"createTime\": \"2022-09-04 22:53:41\", \"updateTime\": \"2023-01-12 20:42:10\"}], \"count\": 1}}', 8.220, '2023-01-13 00:14:37', '2023-01-13 00:14:37');
INSERT INTO `log_request` VALUES (4, 'http://0.0.0.0:20080/auth/scene/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/scene\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 1, \"isStop\": 0, \"sceneId\": 1, \"sceneCode\": \"platformAdmin\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-09-17 23:13:53\", \"updateTime\": \"2022-12-06 22:51:16\", \"sceneConfig\": \"{\\\"signKey\\\": \\\"www.admin.com_platform\\\", \\\"signType\\\": \\\"HS256\\\", \\\"expireTime\\\": 14400}\"}], \"count\": 1}}', 9.260, '2023-01-13 00:14:40', '2023-01-13 00:14:40');
INSERT INTO `log_request` VALUES (5, 'http://0.0.0.0:20080/auth/action/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/action\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 20, \"isStop\": 0, \"remark\": \"\", \"actionId\": 20, \"actionCode\": \"platformAdminDelete\", \"actionName\": \"平台管理员-删除\", \"createTime\": \"2022-12-26 23:00:24\", \"updateTime\": \"2022-12-26 23:00:24\"}, {\"id\": 19, \"isStop\": 0, \"remark\": \"\", \"actionId\": 19, \"actionCode\": \"platformAdminUpdate\", \"actionName\": \"平台管理员-编辑\", \"createTime\": \"2022-12-26 23:00:12\", \"updateTime\": \"2022-12-26 23:00:12\"}, {\"id\": 18, \"isStop\": 0, \"remark\": \"\", \"actionId\": 18, \"actionCode\": \"platformAdminCreate\", \"actionName\": \"平台管理员-新增\", \"createTime\": \"2022-12-26 22:59:56\", \"updateTime\": \"2022-12-26 22:59:56\"}, {\"id\": 17, \"isStop\": 0, \"remark\": \"\", \"actionId\": 17, \"actionCode\": \"platformAdminLook\", \"actionName\": \"平台管理员-查看\", \"createTime\": \"2022-12-26 22:58:33\", \"updateTime\": \"2022-12-26 22:58:33\"}, {\"id\": 16, \"isStop\": 0, \"remark\": \"\", \"actionId\": 16, \"actionCode\": \"authRoleDelete\", \"actionName\": \"权限角色-删除\", \"createTime\": \"2022-12-26 22:50:29\", \"updateTime\": \"2022-12-26 22:50:29\"}, {\"id\": 15, \"isStop\": 0, \"remark\": \"\", \"actionId\": 15, \"actionCode\": \"authRoleUpdate\", \"actionName\": \"权限角色-编辑\", \"createTime\": \"2022-12-26 22:50:21\", \"updateTime\": \"2022-12-26 22:50:21\"}, {\"id\": 14, \"isStop\": 0, \"remark\": \"\", \"actionId\": 14, \"actionCode\": \"authRoleCreate\", \"actionName\": \"权限角色-新增\", \"createTime\": \"2022-12-26 22:50:11\", \"updateTime\": \"2022-12-26 22:50:11\"}, {\"id\": 13, \"isStop\": 0, \"remark\": \"\", \"actionId\": 13, \"actionCode\": \"authRoleLook\", \"actionName\": \"权限角色-查看\", \"createTime\": \"2022-12-26 22:49:59\", \"updateTime\": \"2022-12-26 22:49:59\"}, {\"id\": 12, \"isStop\": 0, \"remark\": \"\", \"actionId\": 12, \"actionCode\": \"authMenuDelete\", \"actionName\": \"权限菜单-删除\", \"createTime\": \"2022-12-26 22:49:16\", \"updateTime\": \"2022-12-26 22:49:16\"}, {\"id\": 11, \"isStop\": 0, \"remark\": \"\", \"actionId\": 11, \"actionCode\": \"authMenuUpdate\", \"actionName\": \"权限菜单-编辑\", \"createTime\": \"2022-12-26 22:49:07\", \"updateTime\": \"2022-12-26 22:49:07\"}, {\"id\": 10, \"isStop\": 0, \"remark\": \"\", \"actionId\": 10, \"actionCode\": \"authMenuCreate\", \"actionName\": \"权限菜单-新增\", \"createTime\": \"2022-12-26 22:48:57\", \"updateTime\": \"2022-12-26 22:48:57\"}, {\"id\": 9, \"isStop\": 0, \"remark\": \"\", \"actionId\": 9, \"actionCode\": \"authMenuLook\", \"actionName\": \"权限菜单-查看\", \"createTime\": \"2022-12-26 22:48:42\", \"updateTime\": \"2022-12-26 22:48:42\"}, {\"id\": 8, \"isStop\": 0, \"remark\": \"\", \"actionId\": 8, \"actionCode\": \"authActionDelete\", \"actionName\": \"权限操作-删除\", \"createTime\": \"2022-12-26 22:47:43\", \"updateTime\": \"2022-12-26 22:47:43\"}, {\"id\": 7, \"isStop\": 0, \"remark\": \"\", \"actionId\": 7, \"actionCode\": \"authActionUpdate\", \"actionName\": \"权限操作-编辑\", \"createTime\": \"2022-12-26 22:47:32\", \"updateTime\": \"2022-12-26 22:47:32\"}, {\"id\": 6, \"isStop\": 0, \"remark\": \"\", \"actionId\": 6, \"actionCode\": \"authActionCreate\", \"actionName\": \"权限操作-新增\", \"createTime\": \"2022-12-26 22:47:20\", \"updateTime\": \"2022-12-26 22:47:20\"}, {\"id\": 5, \"isStop\": 0, \"remark\": \"\", \"actionId\": 5, \"actionCode\": \"authActionLook\", \"actionName\": \"权限操作-查看\", \"createTime\": \"2022-12-26 22:47:05\", \"updateTime\": \"2022-12-26 22:47:05\"}, {\"id\": 4, \"isStop\": 0, \"remark\": \"\", \"actionId\": 4, \"actionCode\": \"authSceneDelete\", \"actionName\": \"权限场景-删除\", \"createTime\": \"2022-12-26 21:19:17\", \"updateTime\": \"2022-12-26 21:19:17\"}, {\"id\": 3, \"isStop\": 0, \"remark\": \"\", \"actionId\": 3, \"actionCode\": \"authSceneUpdate\", \"actionName\": \"权限场景-编辑\", \"createTime\": \"2022-12-26 21:18:52\", \"updateTime\": \"2022-12-26 21:18:52\"}, {\"id\": 2, \"isStop\": 0, \"remark\": \"\", \"actionId\": 2, \"actionCode\": \"authSceneCreate\", \"actionName\": \"权限场景-新增\", \"createTime\": \"2022-12-26 21:18:28\", \"updateTime\": \"2022-12-26 21:18:28\"}, {\"id\": 1, \"isStop\": 0, \"remark\": \"\", \"actionId\": 1, \"actionCode\": \"authSceneLook\", \"actionName\": \"权限场景-查看\", \"createTime\": \"2022-12-26 21:17:51\", \"updateTime\": \"2022-12-26 21:17:51\"}], \"count\": 20}}', 9.317, '2023-01-13 00:14:41', '2023-01-13 00:14:41');
INSERT INTO `log_request` VALUES (6, 'http://0.0.0.0:20080/auth/menu/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/menu\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 12, \"pid\": 11, \"sort\": 50, \"level\": 2, \"isStop\": 0, \"menuId\": 12, \"pidPath\": \"0-11-12\", \"sceneId\": 1, \"menuName\": \"请求日志\", \"extraData\": \"{\\\"url\\\": \\\"/log/request\\\", \\\"icon\\\": \\\"AutoiconEpReading\\\", \\\"title\\\": {\\\"en\\\": \\\"Request Log \\\", \\\"zh-cn\\\": \\\"请求日志\\\"}}\", \"pMenuName\": \"日志管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:36:45\", \"updateTime\": \"2023-01-12 22:30:57\"}, {\"id\": 11, \"pid\": 0, \"sort\": 80, \"level\": 1, \"isStop\": 0, \"menuId\": 11, \"pidPath\": \"0-11\", \"sceneId\": 1, \"menuName\": \"日志管理\", \"extraData\": \"{\\\"icon\\\": \\\"AutoiconEpDataAnalysis\\\", \\\"title\\\": {\\\"en\\\": \\\"Log Manage \\\", \\\"zh-cn\\\": \\\"日志管理\\\"}}\", \"pMenuName\": null, \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:36:38\", \"updateTime\": \"2023-01-12 22:30:51\"}, {\"id\": 10, \"pid\": 9, \"sort\": 50, \"level\": 3, \"isStop\": 0, \"menuId\": 10, \"pidPath\": \"0-8-9-10\", \"sceneId\": 1, \"menuName\": \"平台配置\", \"extraData\": \"{\\\"url\\\": \\\"/platform/config\\\", \\\"title\\\": {\\\"en\\\": \\\"Platform Config \\\", \\\"zh-cn\\\": \\\"平台配置\\\"}}\", \"pMenuName\": \"配置中心\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:36:33\", \"updateTime\": \"2023-01-12 22:30:46\"}, {\"id\": 9, \"pid\": 8, \"sort\": 100, \"level\": 2, \"isStop\": 0, \"menuId\": 9, \"pidPath\": \"0-8-9\", \"sceneId\": 1, \"menuName\": \"配置中心\", \"extraData\": \"{\\\"icon\\\": \\\"AutoiconEpSetting\\\", \\\"title\\\": {\\\"en\\\": \\\"Config Center\\\", \\\"zh-cn\\\": \\\"配置中心\\\"}}\", \"pMenuName\": \"系统管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2023-01-12 22:16:29\", \"updateTime\": \"2023-01-12 22:30:32\"}, {\"id\": 8, \"pid\": 0, \"sort\": 90, \"level\": 1, \"isStop\": 0, \"menuId\": 8, \"pidPath\": \"0-8\", \"sceneId\": 1, \"menuName\": \"系统管理\", \"extraData\": \"{\\\"icon\\\": \\\"AutoiconEpPlatform\\\", \\\"title\\\": {\\\"en\\\": \\\"System Manage \\\", \\\"zh-cn\\\": \\\"系统管理\\\"}}\", \"pMenuName\": null, \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:36:04\", \"updateTime\": \"2023-01-12 22:30:25\"}, {\"id\": 7, \"pid\": 2, \"sort\": 60, \"level\": 2, \"isStop\": 0, \"menuId\": 7, \"pidPath\": \"0-2-7\", \"sceneId\": 1, \"menuName\": \"平台管理员\", \"extraData\": \"{\\\"url\\\": \\\"/platform/admin\\\", \\\"icon\\\": \\\"AutoiconEpUserFilled\\\", \\\"title\\\": {\\\"en\\\": \\\"Platform Admin \\\", \\\"zh-cn\\\": \\\"平台管理员\\\"}}\", \"pMenuName\": \"权限管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:35:16\", \"updateTime\": \"2023-01-12 22:08:35\"}, {\"id\": 6, \"pid\": 2, \"sort\": 70, \"level\": 2, \"isStop\": 0, \"menuId\": 6, \"pidPath\": \"0-2-6\", \"sceneId\": 1, \"menuName\": \"角色\", \"extraData\": \"{\\\"url\\\": \\\"/auth/role\\\", \\\"icon\\\": \\\"AutoiconEpView\\\", \\\"title\\\": {\\\"en\\\": \\\"Role \\\", \\\"zh-cn\\\": \\\"角色\\\"}}\", \"pMenuName\": \"权限管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:35:57\", \"updateTime\": \"2023-01-12 22:08:34\"}, {\"id\": 5, \"pid\": 2, \"sort\": 80, \"level\": 2, \"isStop\": 0, \"menuId\": 5, \"pidPath\": \"0-2-5\", \"sceneId\": 1, \"menuName\": \"菜单\", \"extraData\": \"{\\\"url\\\": \\\"/auth/menu\\\", \\\"icon\\\": \\\"AutoiconEpMenu\\\", \\\"title\\\": {\\\"en\\\": \\\"Menu \\\", \\\"zh-cn\\\": \\\"菜单\\\"}}\", \"pMenuName\": \"权限管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:35:48\", \"updateTime\": \"2023-01-12 22:08:32\"}, {\"id\": 4, \"pid\": 2, \"sort\": 90, \"level\": 2, \"isStop\": 0, \"menuId\": 4, \"pidPath\": \"0-2-4\", \"sceneId\": 1, \"menuName\": \"操作\", \"extraData\": \"{\\\"url\\\": \\\"/auth/action\\\", \\\"icon\\\": \\\"AutoiconEpCoordinate\\\", \\\"title\\\": {\\\"en\\\": \\\"Action\\\", \\\"zh-cn\\\": \\\"操作\\\"}}\", \"pMenuName\": \"权限管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-26 21:12:00\", \"updateTime\": \"2023-01-12 22:08:31\"}, {\"id\": 3, \"pid\": 2, \"sort\": 100, \"level\": 2, \"isStop\": 0, \"menuId\": 3, \"pidPath\": \"0-2-3\", \"sceneId\": 1, \"menuName\": \"场景\", \"extraData\": \"{\\\"url\\\": \\\"/auth/scene\\\", \\\"icon\\\": \\\"AutoiconEpFlag\\\", \\\"title\\\": {\\\"en\\\": \\\"Scene \\\", \\\"zh-cn\\\": \\\"场景\\\"}}\", \"pMenuName\": \"权限管理\", \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:37:07\", \"updateTime\": \"2023-01-12 22:07:31\"}, {\"id\": 2, \"pid\": 0, \"sort\": 100, \"level\": 2, \"isStop\": 0, \"menuId\": 2, \"pidPath\": \"0-2\", \"sceneId\": 1, \"menuName\": \"权限管理\", \"extraData\": \"{\\\"icon\\\": \\\"AutoiconEpLock\\\", \\\"title\\\": {\\\"en\\\": \\\"Auth Manage \\\", \\\"zh-cn\\\": \\\"权限管理\\\"}}\", \"pMenuName\": null, \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:31:28\", \"updateTime\": \"2023-01-12 22:13:54\"}, {\"id\": 1, \"pid\": 0, \"sort\": 0, \"level\": 1, \"isStop\": 0, \"menuId\": 1, \"pidPath\": \"0-1\", \"sceneId\": 1, \"menuName\": \"主页\", \"extraData\": \"{\\\"url\\\": \\\"/\\\", \\\"icon\\\": \\\"AutoiconEpHomeFilled\\\", \\\"title\\\": {\\\"en\\\": \\\"Homepage\\\", \\\"zh-cn\\\": \\\"主页\\\"}}\", \"pMenuName\": null, \"sceneName\": \"平台后台\", \"createTime\": \"2022-12-25 23:28:45\", \"updateTime\": \"2023-01-12 22:20:27\"}], \"count\": 12}}', 19.512, '2023-01-13 00:14:42', '2023-01-13 00:14:42');
INSERT INTO `log_request` VALUES (7, 'http://0.0.0.0:20080/auth/role/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [], \"count\": 0}}', 6.208, '2023-01-13 00:14:43', '2023-01-13 00:14:43');
INSERT INTO `log_request` VALUES (8, 'http://0.0.0.0:20080/auth/scene/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"81\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [\"id\", \"sceneName\"], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 1, \"sceneName\": \"平台后台\"}], \"count\": 1}}', 9.625, '2023-01-13 00:14:53', '2023-01-13 00:14:53');
INSERT INTO `log_request` VALUES (9, 'http://0.0.0.0:20080/auth/menu/tree', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"90\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [\"id\", \"menuName\"], \"limit\": 0, \"order\": {\"id\": \"desc\"}, \"where\": {\"sceneId\": 1}}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"tree\": [{\"id\": 1, \"pid\": 0, \"menuId\": 1, \"children\": [], \"menuName\": \"主页\"}, {\"id\": 11, \"pid\": 0, \"menuId\": 11, \"children\": [{\"id\": 12, \"pid\": 11, \"menuId\": 12, \"children\": [], \"menuName\": \"请求日志\"}], \"menuName\": \"日志管理\"}, {\"id\": 8, \"pid\": 0, \"menuId\": 8, \"children\": [{\"id\": 9, \"pid\": 8, \"menuId\": 9, \"children\": [{\"id\": 10, \"pid\": 9, \"menuId\": 10, \"children\": [], \"menuName\": \"平台配置\"}], \"menuName\": \"配置中心\"}], \"menuName\": \"系统管理\"}, {\"id\": 2, \"pid\": 0, \"menuId\": 2, \"children\": [{\"id\": 7, \"pid\": 2, \"menuId\": 7, \"children\": [], \"menuName\": \"平台管理员\"}, {\"id\": 6, \"pid\": 2, \"menuId\": 6, \"children\": [], \"menuName\": \"角色\"}, {\"id\": 5, \"pid\": 2, \"menuId\": 5, \"children\": [], \"menuName\": \"菜单\"}, {\"id\": 4, \"pid\": 2, \"menuId\": 4, \"children\": [], \"menuName\": \"操作\"}, {\"id\": 3, \"pid\": 2, \"menuId\": 3, \"children\": [], \"menuName\": \"场景\"}], \"menuName\": \"权限管理\"}]}}', 7.377, '2023-01-13 00:14:54', '2023-01-13 00:14:54');
INSERT INTO `log_request` VALUES (10, 'http://0.0.0.0:20080/auth/action/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"92\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [\"id\", \"actionName\"], \"limit\": 0, \"order\": {\"id\": \"desc\"}, \"where\": {\"sceneId\": 1}}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 20, \"actionName\": \"平台管理员-删除\"}, {\"id\": 19, \"actionName\": \"平台管理员-编辑\"}, {\"id\": 18, \"actionName\": \"平台管理员-新增\"}, {\"id\": 17, \"actionName\": \"平台管理员-查看\"}, {\"id\": 16, \"actionName\": \"权限角色-删除\"}, {\"id\": 15, \"actionName\": \"权限角色-编辑\"}, {\"id\": 14, \"actionName\": \"权限角色-新增\"}, {\"id\": 13, \"actionName\": \"权限角色-查看\"}, {\"id\": 12, \"actionName\": \"权限菜单-删除\"}, {\"id\": 11, \"actionName\": \"权限菜单-编辑\"}, {\"id\": 10, \"actionName\": \"权限菜单-新增\"}, {\"id\": 9, \"actionName\": \"权限菜单-查看\"}, {\"id\": 8, \"actionName\": \"权限操作-删除\"}, {\"id\": 7, \"actionName\": \"权限操作-编辑\"}, {\"id\": 6, \"actionName\": \"权限操作-新增\"}, {\"id\": 5, \"actionName\": \"权限操作-查看\"}, {\"id\": 4, \"actionName\": \"权限场景-删除\"}, {\"id\": 3, \"actionName\": \"权限场景-编辑\"}, {\"id\": 2, \"actionName\": \"权限场景-新增\"}, {\"id\": 1, \"actionName\": \"权限场景-查看\"}], \"count\": 20}}', 60.183, '2023-01-13 00:14:54', '2023-01-13 00:14:54');
INSERT INTO `log_request` VALUES (11, 'http://0.0.0.0:20080/auth/role/create', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"171\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"sort\": 50, \"isStop\": 0, \"sceneId\": 1, \"roleName\": \"超级管理员\", \"menuIdArr\": [1, 11, 12, 8, 9, 10, 2, 7, 6, 5, 4, 3], \"actionIdArr\": [20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1]}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": []}', 52.380, '2023-01-13 00:15:39', '2023-01-13 00:15:39');
INSERT INTO `log_request` VALUES (12, 'http://0.0.0.0:20080/auth/role/list', '{\"host\": [\"0.0.0.0:20080\"], \"scene\": [\"platformAdmin\"], \"accept\": [\"application/json, text/plain, */*\"], \"origin\": [\"http://192.168.200.200:5173\"], \"referer\": [\"http://192.168.200.200:5173/view/admin/platform/auth/role\"], \"language\": [\"zh-cn\"], \"connection\": [\"close\"], \"user-agent\": [\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.76\"], \"content-type\": [\"application/json\"], \"content-length\": [\"65\"], \"accept-encoding\": [\"gzip, deflate\"], \"accept-language\": [\"zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6\"], \"platformadmintoken\": [\"eyJzaWduVHlwZSI6IkhTMjU2IiwidHlwZSI6IkpXVCJ9.eyJpZCI6MSwiZXhwaXJlVGltZSI6MTY3MzU1NDMyN30.8n1MlXXGL-T-bJGB51tH6QbVpWSBXPW_ewe2DgZHv0Y\"]}', '{\"page\": 1, \"field\": [], \"limit\": 20, \"order\": {\"id\": \"desc\"}, \"where\": []}', '{\"msg\": \"成功\", \"code\": \"00000000\", \"data\": {\"list\": [{\"id\": 1, \"isStop\": 0, \"roleId\": 1, \"sceneId\": 1, \"tableId\": 0, \"roleName\": \"超级管理员\", \"sceneName\": \"平台后台\", \"createTime\": \"2023-01-13 00:15:39\", \"updateTime\": \"2023-01-13 00:15:39\"}], \"count\": 1}}', 8.872, '2023-01-13 00:15:39', '2023-01-13 00:15:39');

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '平台管理员表' ROW_FORMAT = DYNAMIC;

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
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `configKey`(`configKey` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '平台配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
