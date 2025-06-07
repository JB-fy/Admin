/*
 Navicat Premium Dump SQL

 Source Server         : Mysql
 Source Server Type    : MySQL
 Source Server Version : 90300 (9.3.0)
 Source Host           : 192.168.0.200:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 90300 (9.3.0)
 File Encoding         : 65001

 Date: 08/06/2025 01:33:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `app_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT 'APPID',
  `app_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `app_config` json NULL COMMENT '配置。JSON格式，需要时设置',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`app_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'APP表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app
-- ----------------------------

-- ----------------------------
-- Table structure for app_pkg
-- ----------------------------
DROP TABLE IF EXISTS `app_pkg`;
CREATE TABLE `app_pkg`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `pkg_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '安装包ID',
  `app_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT 'APPID',
  `pkg_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0安卓 1苹果 2PC',
  `pkg_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '包名',
  `pkg_file` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '安装包',
  `ver_no` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
  `ver_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '版本名称',
  `ver_intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '版本介绍',
  `extra_config` json NULL COMMENT '额外配置。JSON格式，需要时设置',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_force_prev` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关',
  PRIMARY KEY (`pkg_id`) USING BTREE,
  INDEX `app_id`(`app_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'APP安装包表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_pkg
-- ----------------------------

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `action_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作ID',
  `action_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`action_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appCreate', '系统管理-APP-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appDelete', '系统管理-APP-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgCreate', '系统管理-APP管理-安装包-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgDelete', '系统管理-APP管理-安装包-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgRead', '系统管理-APP管理-安装包-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgUpdate', '系统管理-APP管理-安装包-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appRead', '系统管理-APP-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appUpdate', '系统管理-APP-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionCreate', '权限管理-操作-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionDelete', '权限管理-操作-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionRead', '权限管理-操作-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionUpdate', '权限管理-操作-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuCreate', '权限管理-菜单-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuDelete', '权限管理-菜单-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuRead', '权限管理-菜单-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuUpdate', '权限管理-菜单-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleCreate', '权限管理-角色-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleDelete', '权限管理-角色-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleRead', '权限管理-角色-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleUpdate', '权限管理-角色-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneCreate', '权限管理-场景-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneDelete', '权限管理-场景-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneRead', '权限管理-场景-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneUpdate', '权限管理-场景-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminCreate', '权限管理-机构管理员-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminDelete', '权限管理-机构管理员-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminRead', '权限管理-机构管理员-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminUpdate', '权限管理-机构管理员-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigCommonRead', '应用配置-常用-查看', '只能读取机构配置表中的某些配置。对应前端页面：配置中心-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigCommonSave', '应用配置-常用-保存', '只能保存机构配置表中的某些配置。对应前端页面：配置中心-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigRead', '配置中心-查看', '可任意读取机构配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigSave', '配置中心-保存', '可任意保存机构配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgCreate', '机构管理-机构-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgDelete', '机构管理-机构-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgRead', '机构管理-机构-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgUpdate', '机构管理-机构-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelCreate', '系统管理-配置中心-支付管理-支付通道-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelDelete', '系统管理-配置中心-支付管理-支付通道-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelRead', '系统管理-配置中心-支付管理-支付通道-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelUpdate', '系统管理-配置中心-支付管理-支付通道-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payCreate', '系统管理-配置中心-支付管理-支付配置-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payDelete', '系统管理-配置中心-支付管理-支付配置-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payRead', '系统管理-配置中心-支付管理-支付配置-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneCreate', '系统管理-配置中心-支付管理-支付场景-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneDelete', '系统管理-配置中心-支付管理-支付场景-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneRead', '系统管理-配置中心-支付管理-支付场景-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneUpdate', '系统管理-配置中心-支付管理-支付场景-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payUpdate', '系统管理-配置中心-支付管理-支付配置-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminCreate', '权限管理-平台管理员-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminDelete', '权限管理-平台管理员-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminRead', '权限管理-平台管理员-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminUpdate', '权限管理-平台管理员-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigCommonRead', '应用配置-常用-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigCommonSave', '应用配置-常用-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigEmailRead', '插件配置-邮箱-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigEmailSave', '插件配置-邮箱-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigIdCardRead', '插件配置-实名认证-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigIdCardSave', '插件配置-实名认证-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigOneClickRead', '插件配置-一键登录-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigOneClickSave', '插件配置-一键登录-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigPushRead', '插件配置-推送-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigPushSave', '插件配置-推送-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigRead', '平台配置-查看', '可任意读取平台配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSave', '平台配置-保存', '可任意保存平台配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSmsRead', '插件配置-短信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSmsSave', '插件配置-短信-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigVodRead', '插件配置-视频点播-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigVodSave', '插件配置-视频点播-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigWxRead', '插件配置-微信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigWxSave', '插件配置-微信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadCreate', '系统管理-配置中心-上传配置-新增', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadDelete', '系统管理-配置中心-上传配置-删除', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadRead', '系统管理-配置中心-上传配置-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadUpdate', '系统管理-配置中心-上传配置-编辑', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'usersRead', '用户管理-用户-查看', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'usersUpdate', '用户管理-用户-编辑', '');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `action_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作ID',
  `scene_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '场景ID',
  PRIMARY KEY (`action_id`, `scene_id`) USING BTREE,
  INDEX `action_id`(`action_id` ASC) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2025-05-28 01:45:24', '2025-05-28 01:45:24', 'appPkgCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2025-05-28 01:45:24', '2025-05-28 01:45:24', 'appPkgDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2025-05-28 01:45:24', '2025-05-28 01:45:24', 'appPkgRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2025-05-28 01:45:24', '2025-05-28 01:45:24', 'appPkgUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleCreate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleDelete', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleRead', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleUpdate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneCreate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneDelete', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneRead', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneUpdate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminCreate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminDelete', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminRead', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminUpdate', 'org');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigCommonRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigCommonSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigEmailRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigEmailSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigIdCardRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigIdCardSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigOneClickRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigOneClickSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigPushRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigPushSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSmsRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSmsSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigVodRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigVodSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigWxRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigWxSave', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadCreate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadDelete', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadUpdate', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'usersRead', 'platform');
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'usersUpdate', 'platform');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `menu_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
  `menu_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '场景ID',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `is_leaf` tinyint UNSIGNED NOT NULL DEFAULT 1 COMMENT '叶子：0否 1是',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `id_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT 'ID路径',
  `name_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '名称路径',
  `menu_icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `menu_url` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `extra_data` json NULL COMMENT '额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 100 COMMENT '排序值。从大到小排序',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '主页', 'platform', 0, 1, 1, '0-1', '-主页', 'autoicon-ep-home-filled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 255);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理', 'platform', 0, 0, 1, '0-2', '-权限管理', 'autoicon-ep-lock', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '场景', 'platform', 2, 1, 2, '0-2-3', '-权限管理-场景', 'autoicon-ep-flag', '/auth/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"场景\"}}}', 0);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '操作', 'platform', 2, 1, 2, '0-2-4', '-权限管理-操作', 'autoicon-ep-coordinate', '/auth/action', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '菜单', 'platform', 2, 1, 2, '0-2-5', '-权限管理-菜单', 'autoicon-ep-menu', '/auth/menu', '{\"i18n\": {\"title\": {\"en\": \"Menu\", \"zh-cn\": \"菜单\"}}}', 30);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '角色', 'platform', 2, 1, 2, '0-2-6', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 40);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '平台管理员', 'platform', 2, 1, 2, '0-2-7', '-权限管理-平台管理员', 'autoicon-ep-avatar', '/platform/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"平台管理员\"}}}', 50);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '系统管理', 'platform', 0, 0, 1, '0-8', '-系统管理', 'autoicon-ep-platform', '', '{\"i18n\": {\"title\": {\"en\": \"System Manage\", \"zh-cn\": \"系统管理\"}}}', 20);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '配置中心', 'platform', 8, 0, 2, '0-8-9', '-系统管理-配置中心', 'autoicon-ep-setting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 0);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '上传配置', 'platform', 9, 1, 3, '0-8-9-10', '-系统管理-配置中心-上传配置', 'autoicon-ep-upload', '/upload/upload', '{\"i18n\": {\"title\": {\"en\": \"Upload\", \"zh-cn\": \"上传配置\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '支付管理', 'platform', 9, 0, 3, '0-8-9-11', '-系统管理-配置中心-支付管理', 'autoicon-ep-coin', '', '{\"i18n\": {\"title\": {\"en\": \"Pay Manage\", \"zh-cn\": \"支付管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '支付配置', 'platform', 11, 1, 4, '0-8-9-11-12', '-系统管理-配置中心-支付管理-支付配置', 'autoicon-ep-money', '/pay/pay', '{\"i18n\": {\"title\": {\"en\": \"Pay\", \"zh-cn\": \"支付配置\"}}}', 50);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '支付场景', 'platform', 11, 1, 4, '0-8-9-11-13', '-系统管理-配置中心-支付管理-支付场景', 'autoicon-ep-guide', '/pay/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"支付场景\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '支付通道', 'platform', 11, 1, 4, '0-8-9-11-14', '-系统管理-配置中心-支付管理-支付通道', 'autoicon-ep-connection', '/pay/channel', '{\"i18n\": {\"title\": {\"en\": \"Channel\", \"zh-cn\": \"支付通道\"}}}', 150);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '插件配置', 'platform', 9, 1, 3, '0-8-9-15', '-系统管理-配置中心-插件配置', 'autoicon-ep-ticket', '/platform/config/plugin', '{\"i18n\": {\"title\": {\"en\": \"Plugin Config\", \"zh-cn\": \"插件配置\"}}}', 150);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '应用配置', 'platform', 9, 1, 3, '0-8-9-16', '-系统管理-配置中心-应用配置', 'autoicon-ep-set-up', '/platform/config/app', '{\"i18n\": {\"title\": {\"en\": \"APP Config\", \"zh-cn\": \"应用配置\"}}}', 200);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, 'APP管理', 'platform', 8, 0, 2, '0-8-17', '-系统管理-APP管理', 'autoicon-ep-suitcase-line', '', '{\"i18n\": {\"title\": {\"en\": \"APP Manage\", \"zh-cn\": \"APP管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, 'APP', 'platform', 17, 1, 3, '0-8-17-18', '-系统管理-APP管理-APP', 'autoicon-ep-apple', '/app/app', '{\"i18n\": {\"title\": {\"en\": \"App\", \"zh-cn\": \"APP\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '安装包', 'platform', 17, 1, 3, '0-8-17-19', '-系统管理-APP管理-安装包', 'autoicon-ep-box', '/app/pkg', '{\"i18n\": {\"title\": {\"en\": \"Pkg\", \"zh-cn\": \"安装包\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '用户管理', 'platform', 0, 0, 1, '0-20', '-用户管理', 'autoicon-ep-user-filled', '', '{\"i18n\": {\"title\": {\"en\": \"User Manage\", \"zh-cn\": \"用户管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '用户', 'platform', 20, 1, 2, '0-20-21', '-用户管理-用户', 'autoicon-ep-user', '/users/users', '{\"i18n\": {\"title\": {\"en\": \"Users\", \"zh-cn\": \"用户\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '机构管理', 'platform', 0, 0, 1, '0-22', '-机构管理', 'autoicon-ep-office-building', '', '{\"i18n\": {\"title\": {\"en\": \"Org Manage\", \"zh-cn\": \"机构管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '机构', 'platform', 22, 1, 2, '0-22-23', '-机构管理-机构', 'autoicon-ep-school', '/org/org', '{\"i18n\": {\"title\": {\"en\": \"Org\", \"zh-cn\": \"机构\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 24, '机构管理员', 'platform', 2, 1, 2, '0-2-24', '-权限管理-机构管理员', 'autoicon-ep-avatar', '/org/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"机构管理员\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 25, '主页', 'org', 0, 1, 1, '0-25', '-主页', 'autoicon-ep-home-filled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 255);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 26, '权限管理', 'org', 0, 0, 1, '0-26', '-权限管理', 'autoicon-ep-menu', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 27, '角色', 'org', 26, 1, 2, '0-26-27', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 40);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 28, '管理员', 'org', 26, 1, 2, '0-26-28', '-权限管理-管理员', 'autoicon-ep-avatar', '/org/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"管理员\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 29, '配置中心', 'org', 0, 0, 1, '0-29', '-配置中心', 'autoicon-ep-setting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 20);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 30, '应用配置', 'org', 29, 1, 2, '0-29-30', '-配置中心-应用配置', 'autoicon-ep-set-up', '/org/config/app', '{\"i18n\": {\"title\": {\"en\": \"APP Config\", \"zh-cn\": \"应用配置\"}}}', 200);

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `role_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '场景ID',
  `rel_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联ID。0表示平台创建，其它值根据scene_id对应不同表',
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE,
  INDEX `rel_id`(`rel_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限角色表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_org_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_org_admin`;
CREATE TABLE `auth_role_rel_of_org_admin`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `admin_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`admin_id`, `role_id`) USING BTREE,
  INDEX `admin_id`(`admin_id` ASC) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '机构管理员，权限角色关联表（机构管理员包含哪些角色）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_of_org_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_of_platform_admin`;
CREATE TABLE `auth_role_rel_of_platform_admin`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `admin_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '管理员ID',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  PRIMARY KEY (`admin_id`, `role_id`) USING BTREE,
  INDEX `admin_id`(`admin_id` ASC) USING BTREE,
  INDEX `role_id`(`role_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台管理员，权限角色关联表（平台管理员包含哪些角色）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_role_rel_to_action`;
CREATE TABLE `auth_role_rel_to_action`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `action_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作ID',
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
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `role_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menu_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '菜单ID',
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
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `scene_id` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '场景ID',
  `scene_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_config` json NULL COMMENT '配置。JSON格式，根据场景设置',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`scene_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'app', 'APP', '{\"token_config\": {\"is_ip\": 0, \"is_unique\": 0, \"sign_type\": \"HS256\", \"token_type\": 0, \"active_time\": 0, \"expire_time\": 604800, \"private_key\": \"任意字符串\"}}', '');
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'org', '机构后台', '{\"token_config\": {\"is_ip\": 1, \"is_unique\": 1, \"sign_type\": \"RS256\", \"public_key\": \"-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0nT4zSS2O+2yib+gdBm0\\nW0kU2AL6felZFp5U6B/ySeJnM8UE+fZwytgPuND5Z07khxtHR/YYH7huPGZ/fAgh\\nZHmNE5K5phMI5eETwPjk3RDeyYAyOosrKr3SAjjEQxJBISvBEillH4bKjaa4WF5/\\n+nGcp7f4e49caW/CfwuC2ZVrvySCPf1lR8o7/4Zz/hWUwgsEd/crR7ojgt+rbPeE\\n7+Cz11sZUZaMipTqsU3RVhbwtMyLdkos6KsYW7TZEK0VYt94/1XJQBUEjtCdDpS7\\n0XNF8ENpQPtuQdYE6/y+Jku8T9pqQQq/SOL6uPgsn6zJQ41u/l2AhG0i5GxYD86C\\n5wIDAQAB\\n-----END PUBLIC KEY-----\", \"token_type\": 0, \"active_time\": 3600, \"expire_time\": 14400, \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDSdPjNJLY77bKJ\\nv6B0GbRbSRTYAvp96VkWnlToH/JJ4mczxQT59nDK2A+40PlnTuSHG0dH9hgfuG48\\nZn98CCFkeY0TkrmmEwjl4RPA+OTdEN7JgDI6iysqvdICOMRDEkEhK8ESKWUfhsqN\\nprhYXn/6cZynt/h7j1xpb8J/C4LZlWu/JII9/WVHyjv/hnP+FZTCCwR39ytHuiOC\\n36ts94Tv4LPXWxlRloyKlOqxTdFWFvC0zIt2SizoqxhbtNkQrRVi33j/VclAFQSO\\n0J0OlLvRc0XwQ2lA+25B1gTr/L4mS7xP2mpBCr9I4vq4+CyfrMlDjW7+XYCEbSLk\\nbFgPzoLnAgMBAAECggEAGolMO9WmsrzAd9T1Pt5k2uPGoIwTmJ+9L3hsXU515vII\\nsELl4zy7MSB4LwYOhIOylgSPAthZZ1qCb9Q+u91slHYtHywvg2zAAPhV3M2lUeiI\\nJuEmtDILEdsYaVZODOT22F9je05D5WtCDAVbFi1oNqRvq8grKS1E6jiAzjMd3yBY\\n5AgUUP8sbS7BDdPus2t3mCAXqdtFkxn8wo/4WdMV6vG4p9p+a8dIoiRYHNBIw4sU\\nCYPiE7q52tVqVl10ShrJQlvoyDvmJam4inbl8ZWtbQLUsQxfzoEUfuuC0mYc2pES\\n1kp1pWfJNc5JtAXXZ/k9F4jLvjMp9KJximOG+E5tmQKBgQD5RwG8xu7s2JcocogD\\nuJjWfLz/4ab7Zs2NReCX6jmb522d1TqyiU9ilc0XBz8gDNeMwnOfWIyGR3Vtuub+\\nU56Y9/q+IMZ0ewdrUTR2NADZdqLH1ViVGacnnmJXJ+30Z8eUO0GCU7++DSQ1u8xh\\ng0mmrj6++xHsYJM3bCxBstNzGQKBgQDYIfOUIK+JKvYZ5idUrFWZnnVSdLu8IJHi\\nA0osaMB9VNtAdqrPVIA3L9AbIR1/la3gb3ILP0hIM7glt4i78WTwVrh0qkaz/8dw\\nX3EL5u9OAMecVFn4+gZon4RqbzjdtHso87v2GrIZ88eVOWjXRq93gWAe17G+noyS\\n44PgB7Zl/wKBgC3bOh6YGevIDEaMiyjkFHmgiMQppqYoyzdp218W33ImqKuYRiwB\\nxnDETe4mjx4+PojOXKa7i15IVvnQoB25FDvfomjHbrqOx1aeoZ/9AQsAIAHS5XDI\\nP0+yezS9S7DiRnymSe7HqUY09KxN19M4a5wWAcTwOuPZADv50kpjszJBAoGAS0io\\nO7SW8ESSrLrKgGf2+SeE3k/jBMijiAJ1V7q1MfLY3D95h/Z7Ir340zpZuBM/Gao4\\nI0rLtrqtLhYb/rs62aybW6fkMNarda0JB4hNWvJSlVWccWlFyjOmQBy1xiQTslQT\\n6Mmrt/Z+UrBIoJPyksHx5UxkkW1QsemmCecl1akCgYAys4PNRhdToxS5WDF7Tddg\\n8ghhobuwUCP2pxNYakX9HeHZsAjwwQhx6vGJqZPs5hh8HRinGWSvKyVXaUGVN+5b\\nFjMz/rhO9vTCZJS3aJSXxr0PFTbpP/AZSXwCBxjp+uEFTD8GJRX7/6+wlTk7+uTj\\nl7klp31noFtzz+onGSmqvA==\\n-----END PRIVATE KEY-----\"}}', '');
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platform', '平台后台', '{\"token_config\": {\"is_ip\": 1, \"is_unique\": 1, \"sign_type\": \"ES256\", \"public_key\": \"-----BEGIN PUBLIC KEY-----\\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42\\nhF5PnYenbWprk1MQIzT2+V4rRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\\n-----END PUBLIC KEY-----\", \"token_type\": 0, \"active_time\": 3600, \"expire_time\": 604800, \"private_key\": \"-----BEGIN EC PRIVATE KEY-----\\nMHcCAQEEIKvYPRtCqy9MI/yhx4L4+Sog/5lntHbuwxPg/JI0qW6LoAoGCCqGSM49\\nAwEHoUQDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42hF5PnYenbWprk1MQIzT2+V4r\\nRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\\n-----END EC PRIVATE KEY-----\"}}', '');

-- ----------------------------
-- Table structure for org
-- ----------------------------
DROP TABLE IF EXISTS `org`;
CREATE TABLE `org`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `org_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '机构ID',
  `org_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '机构名称',
  PRIMARY KEY (`org_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '机构表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of org
-- ----------------------------

-- ----------------------------
-- Table structure for org_admin
-- ----------------------------
DROP TABLE IF EXISTS `org_admin`;
CREATE TABLE `org_admin`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `admin_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `org_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `is_super` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '超管：0否 1是',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  PRIMARY KEY (`admin_id`) USING BTREE,
  UNIQUE INDEX `org_id_2`(`org_id` ASC, `phone` ASC) USING BTREE,
  UNIQUE INDEX `org_id_3`(`org_id` ASC, `email` ASC) USING BTREE,
  UNIQUE INDEX `org_id_4`(`org_id` ASC, `account` ASC) USING BTREE,
  INDEX `org_id`(`org_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '机构管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of org_admin
-- ----------------------------

-- ----------------------------
-- Table structure for org_config
-- ----------------------------
DROP TABLE IF EXISTS `org_config`;
CREATE TABLE `org_config`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `org_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `config_key` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置键',
  `config_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值',
  PRIMARY KEY (`org_id`, `config_key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '机构配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of org_config
-- ----------------------------

-- ----------------------------
-- Table structure for pay
-- ----------------------------
DROP TABLE IF EXISTS `pay`;
CREATE TABLE `pay`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `pay_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付ID',
  `pay_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `pay_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0支付宝 1微信',
  `pay_config` json NOT NULL COMMENT '配置。JSON格式，根据类型设置',
  `pay_rate` decimal(4, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '费率',
  `total_amount` decimal(14, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '总额',
  `balance` decimal(18, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '余额',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`pay_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '支付表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `channel_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '通道ID',
  `channel_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `channel_icon` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标',
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `pay_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付ID',
  `pay_method` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 100 COMMENT '排序值。从大到小排序',
  `total_amount` decimal(14, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '总额',
  PRIMARY KEY (`channel_id`) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE,
  INDEX `pay_id`(`pay_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '支付通道表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `order_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `rel_order_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联订单类型：0默认',
  `rel_order_user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联订单用户ID',
  `pay_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付ID',
  `channel_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道ID',
  `pay_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0支付宝 1微信',
  `amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '实付金额',
  `pay_status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0未付款 1已付款',
  `pay_time` datetime NULL DEFAULT NULL COMMENT '支付时间',
  `pay_rate` decimal(4, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '费率',
  `third_order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '第三方订单号',
  PRIMARY KEY (`order_id`) USING BTREE,
  UNIQUE INDEX `order_no`(`order_no` ASC) USING BTREE,
  INDEX `pay_id`(`pay_id` ASC) USING BTREE,
  INDEX `channel_id`(`channel_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '支付订单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_order
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order_rel
-- ----------------------------
DROP TABLE IF EXISTS `pay_order_rel`;
CREATE TABLE `pay_order_rel`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `order_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单ID',
  `rel_order_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联订单类型：0默认',
  `rel_order_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联订单ID',
  `rel_order_no` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '关联订单号',
  `rel_order_user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '关联订单用户ID',
  `rel_order_amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '关联订单实付金额',
  INDEX `order_id`(`order_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '支付订单关联表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_order_rel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_scene
-- ----------------------------
DROP TABLE IF EXISTS `pay_scene`;
CREATE TABLE `pay_scene`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `scene_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '场景ID',
  `scene_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`scene_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '支付场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_scene
-- ----------------------------

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS `platform_admin`;
CREATE TABLE `platform_admin`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `admin_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员ID',
  `is_super` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '超管：0否 1是',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  PRIMARY KEY (`admin_id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台管理员表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO `platform_admin` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, 1, '超级管理员', '', NULL, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS `platform_config`;
CREATE TABLE `platform_config`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `config_key` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '配置键',
  `config_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '配置值',
  PRIMARY KEY (`config_key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台配置表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_config
-- ----------------------------
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'email_code', '{\"subject\":\"您的验证码\",\"template\":\"验证码：{code}\\n说明：\\n1. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。\\n2. 出于安全考虑，请不要将此验证码分享给任何人。\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'email_of_common', '{\"from_email\":\"xxxxxxxx@qq.com\",\"password\":\"xxxxxxxx\",\"smtp_host\":\"smtp.qq.com\",\"smtp_port\":\"465\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'email_type', 'email_of_common');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'id_card_of_aliyun', '{\"appcode\":\"appcode\",\"url\":\"http://idcard.market.alicloudapi.com/lianzhuo/idcard\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'id_card_type', 'id_card_of_aliyun');

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS `platform_server`;
CREATE TABLE `platform_server`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `server_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '服务器ID',
  `network_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '公网IP',
  `local_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '内网IP',
  PRIMARY KEY (`server_id`) USING BTREE,
  UNIQUE INDEX `network_ip`(`network_ip` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '平台服务器表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for upload
-- ----------------------------
DROP TABLE IF EXISTS `upload`;
CREATE TABLE `upload`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `upload_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '上传ID',
  `upload_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0本地 1阿里云OSS',
  `upload_config` json NOT NULL COMMENT '配置。JSON格式，根据类型设置',
  `is_default` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '默认：0否 1是',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`upload_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '上传表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of upload
-- ----------------------------
INSERT INTO `upload` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 0, '{\"sign_key\": \"signKey\", \"is_cluster\": 1, \"server_list\": [], \"is_same_server\": 0}', 1, '此项目自带文件上传下载功能，可直接部署成文件服务器使用');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `user_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '性别：0未设置 1男 2女',
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '手机',
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `account` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '账号',
  `wx_openid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信openid',
  `wx_unionid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '微信unionid',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `phone`(`phone` ASC) USING BTREE,
  UNIQUE INDEX `email`(`email` ASC) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `wx_openid`(`wx_openid` ASC) USING BTREE,
  UNIQUE INDEX `wx_unionid`(`wx_unionid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表（postgresql中user是关键字，使用需要加双引号。程序中考虑与mysql通用，故命名成users）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Table structure for users_privacy
-- ----------------------------
DROP TABLE IF EXISTS `users_privacy`;
CREATE TABLE `users_privacy`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `user_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码。md5保存',
  `salt` char(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `id_card_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证号码',
  `id_card_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证姓名',
  `id_card_gender` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '身份证性别：0未设置 1男 2女',
  `id_card_birthday` date NULL DEFAULT NULL COMMENT '身份证生日',
  `id_card_address` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '身份证地址',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户隐私表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of users_privacy
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
