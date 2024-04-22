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

 Date: 22/04/2024 18:21:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `action_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作ID',
  `action_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `action_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`action_id`) USING BTREE,
  UNIQUE INDEX `action_code`(`action_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 25 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES (1, '权限管理-场景-查看', 'authSceneLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (2, '权限管理-场景-新增', 'authSceneCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (3, '权限管理-场景-编辑', 'authSceneUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (4, '权限管理-场景-删除', 'authSceneDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (5, '权限管理-操作-查看', 'authActionLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (6, '权限管理-操作-新增', 'authActionCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (7, '权限管理-操作-编辑', 'authActionUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (8, '权限管理-操作-删除', 'authActionDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (9, '权限管理-菜单-查看', 'authMenuLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (10, '权限管理-菜单-新增', 'authMenuCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (11, '权限管理-菜单-编辑', 'authMenuUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (12, '权限管理-菜单-删除', 'authMenuDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (13, '权限管理-角色-查看', 'authRoleLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (14, '权限管理-角色-新增', 'authRoleCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (15, '权限管理-角色-编辑', 'authRoleUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (16, '权限管理-角色-删除', 'authRoleDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (17, '权限管理-平台管理员-查看', 'platformAdminLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (18, '权限管理-平台管理员-新增', 'platformAdminCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (19, '权限管理-平台管理员-编辑', 'platformAdminUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (20, '权限管理-平台管理员-删除', 'platformAdminDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (21, '平台配置-查看', 'platformConfigLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (22, '平台配置-保存', 'platformConfigSave', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (23, '用户管理-用户-查看', 'userLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action` VALUES (24, '用户管理-用户-编辑', 'userUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `action_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`action_id`, `scene_id`) USING BTREE,
  INDEX `action_id`(`action_id` ASC) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO `auth_action_rel_to_scene` VALUES (1, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (2, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (3, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (4, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (5, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (6, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (7, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (8, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (9, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (10, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (11, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (12, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (13, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (14, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (15, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (16, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (17, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (18, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (19, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (20, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (21, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (22, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (23, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_action_rel_to_scene` VALUES (24, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `menu_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `id_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '层级路径',
  `menu_icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}',
  `menu_url` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `extra_data` json NULL COMMENT '额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值。从小到大排序，默认50，范围0-100',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, '主页', 1, 0, 1, '0-1', 'autoicon-ep-home-filled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 0, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (2, '权限管理', 1, 0, 1, '0-2', 'autoicon-ep-lock', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 90, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (3, '场景', 1, 2, 2, '0-2-3', 'autoicon-ep-flag', '/auth/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"场景\"}}}', 100, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (4, '操作', 1, 2, 2, '0-2-4', 'autoicon-ep-coordinate', '/auth/action', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 90, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (5, '菜单', 1, 2, 2, '0-2-5', 'autoicon-ep-menu', '/auth/menu', '{\"i18n\": {\"title\": {\"en\": \"Menu\", \"zh-cn\": \"菜单\"}}}', 80, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (6, '角色', 1, 2, 2, '0-2-6', 'autoicon-ep-view', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 70, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (7, '平台管理员', 1, 2, 2, '0-2-7', 'vant-manager-o', '/platform/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"平台管理员\"}}}', 60, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (8, '系统管理', 1, 0, 1, '0-8', 'autoicon-ep-platform', '', '{\"i18n\": {\"title\": {\"en\": \"System Manage\", \"zh-cn\": \"系统管理\"}}}', 85, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (9, '配置中心', 1, 8, 2, '0-8-9', 'autoicon-ep-setting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 100, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (10, '平台配置', 1, 9, 3, '0-8-9-10', '', '/platform/config/platform', '{\"i18n\": {\"title\": {\"en\": \"Platform Config\", \"zh-cn\": \"平台配置\"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (11, '插件配置', 1, 9, 3, '0-8-9-11', '', '/platform/config/plugin', '{\"i18n\": {\"title\": {\"en\": \"Plugin Config\", \"zh-cn\": \"插件配置\"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (12, '用户管理', 1, 0, 1, '0-12', 'vant-friends', '', '{\"i18n\": {\"title\": {\"en\": \"User Manage\", \"zh-cn\": \"用户管理\"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_menu` VALUES (13, '用户', 1, 12, 2, '0-12-13', 'vant-user-o', '/user/user', '{\"i18n\": {\"title\": {\"en\": \"User\", \"zh-cn\": \"用户\"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role`  (
  `role_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `table_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE,
  INDEX `table_id`(`table_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `admin_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`role_id`, `admin_id`) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE,
  INDEX `admin_id`(`admin_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，系统管理员关联表（系统管理员包含哪些角色）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_action`;
CREATE TABLE `auth_role_rel_to_action`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `action_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`role_id`, `action_id`) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE,
  INDEX `action_id`(`action_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限操作关联表（角色包含哪些操作）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_menu`;
CREATE TABLE `auth_role_rel_to_menu`  (
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menu_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE,
  INDEX `menu_id`(`menu_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色，权限菜单关联表（角色包含哪些菜单）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_scene`;
CREATE TABLE `auth_scene`  (
  `scene_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '场景ID',
  `scene_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `scene_config` json NOT NULL COMMENT '配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{\"signType\": \"算法\",\"signKey\": \"密钥\",\"expireTime\": 过期时间,...}',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`scene_id`) USING BTREE,
  UNIQUE INDEX `scene_code`(`scene_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES (1, '平台后台', 'platform', '{\"signKey\": \"www.admin.com_platform\", \"signType\": \"HS256\", \"expireTime\": 14400}', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `auth_scene` VALUES (2, 'APP', 'app', '{\"signKey\": \"www.admin.com_app\", \"signType\": \"HS256\", \"expireTime\": 604800}', '23:59:59', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `platform_admin`;
CREATE TABLE `platform_admin`  (
  `admin_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`admin_id`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO `platform_admin` VALUES (1, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB', '超级管理员', 'http://JB.Admin.com/common/20240106/1704522339892_31917913.png', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config`  (
  `config_key` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置Key',
  `config_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`config_key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------
INSERT INTO `platform_config` VALUES ('idCardOfAliyunAppcode', 'appcode', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('idCardOfAliyunHost', 'http://idcard.market.alicloudapi.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('idCardOfAliyunPath', '/lianzhuo/idcard', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('idCardType', 'idCardOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfAliAppId', 'appId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfAliNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfAli', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfAliOpAppId', 'opAppId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfAliPrivateKey', '****************', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfAliPublicKey', '****************', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxApiV3Key', '********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxAppId', 'appId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxMchid', 'mchId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfWx', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxPrivateKey', '-----BEGIN RSA PRIVATE KEY-----\n****************************************************************\n-----END RSA PRIVATE KEY-----', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('payOfWxSerialNo', '********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxAndroidAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxAndroidSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxHost', 'https://api.tpns.tencent.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxIosAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxIosSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxMacOSAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushOfTxMacOSSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('pushType', 'pushOfTx', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsOfAliyunAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsOfAliyunAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsOfAliyunEndpoint', 'dysmsapi.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsOfAliyunSignName', 'JB Admin', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsOfAliyunTemplateCode', 'SMS_********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('smsType', 'smsOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssBucket', 'bucket', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssCallbackUrl', 'http://JB.Admin.com/upload/notify/uploadOfAliyunOss', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssEndpoint', 'sts.cn-hangzhou.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssHost', 'https://oss-cn-hangzhou.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfAliyunOssRoleArn', 'acs:ram::********:role/********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfLocalFileSaveDir', '../public/', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfLocalFileUrlPrefix', 'http://JB.Admin.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfLocalSignKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadOfLocalUrl', 'http://JB.Admin.com/upload/upload', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('uploadType', 'uploadOfLocal', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('vodOfAliyunAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('vodOfAliyunAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('vodOfAliyunEndpoint', 'sts.cn-shanghai.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('vodOfAliyunRoleArn', 'acs:ram::********:role/********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO `platform_config` VALUES ('vodType', 'vodOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS `platform_server`;
CREATE TABLE `platform_server`  (
  `server_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '服务器ID',
  `network_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公网IP',
  `local_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内网IP',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`server_id`) USING BTREE,
  UNIQUE INDEX `network_ip`(`network_ip` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台服务器表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密盐',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别：0未设置 1男 2女',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `address` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细地址',
  `id_card_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证姓名',
  `id_card_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
