/*
 Navicat Premium Dump SQL

 Source Server         : 本地-Mysql8
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.1.200:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 26/10/2024 17:34:34
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
  `app_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'APPID',
  `app_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0安卓 1苹果',
  `package_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '包名',
  `package_file` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '安装包',
  `ver_no` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '版本号',
  `ver_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '版本名称',
  `ver_intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '版本介绍',
  `extra_config` json NULL COMMENT '额外配置',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_force_prev` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关',
  PRIMARY KEY (`app_id`) USING BTREE,
  UNIQUE INDEX `app_type`(`app_type` ASC, `ver_no` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = 'APP表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app
-- ----------------------------

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_action`;
CREATE TABLE `auth_action`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `action_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作ID',
  `action_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `action_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`action_id`) USING BTREE,
  UNIQUE INDEX `action_code`(`action_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 69 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '权限管理-场景-查看', 'authSceneRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理-场景-新增', 'authSceneCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '权限管理-场景-编辑', 'authSceneUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '权限管理-场景-删除', 'authSceneDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '权限管理-操作-查看', 'authActionRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '权限管理-操作-新增', 'authActionCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '权限管理-操作-编辑', 'authActionUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '权限管理-操作-删除', 'authActionDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '权限管理-菜单-查看', 'authMenuRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '权限管理-菜单-新增', 'authMenuCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '权限管理-菜单-编辑', 'authMenuUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '权限管理-菜单-删除', 'authMenuDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '权限管理-角色-查看', 'authRoleRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '权限管理-角色-新增', 'authRoleCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '权限管理-角色-编辑', 'authRoleUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '权限管理-角色-删除', 'authRoleDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, '权限管理-平台管理员-查看', 'platformAdminRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, '权限管理-平台管理员-新增', 'platformAdminCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '权限管理-平台管理员-编辑', 'platformAdminUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '权限管理-平台管理员-删除', 'platformAdminDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '系统管理-配置中心-上传配置-查看', 'uploadRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '系统管理-配置中心-上传配置-新增', 'uploadCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '系统管理-配置中心-上传配置-编辑', 'uploadUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 24, '系统管理-配置中心-上传配置-删除', 'uploadDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 25, '系统管理-配置中心-支付管理-支付配置-查看', 'payRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 26, '系统管理-配置中心-支付管理-支付配置-新增', 'payCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 27, '系统管理-配置中心-支付管理-支付配置-编辑', 'payUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 28, '系统管理-配置中心-支付管理-支付配置-删除', 'payDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 29, '系统管理-配置中心-支付管理-支付场景-查看', 'paySceneRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 30, '系统管理-配置中心-支付管理-支付场景-新增', 'paySceneCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 31, '系统管理-配置中心-支付管理-支付场景-编辑', 'paySceneUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 32, '系统管理-配置中心-支付管理-支付场景-删除', 'paySceneDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 33, '系统管理-配置中心-支付管理-支付通道-查看', 'payChannelRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 34, '系统管理-配置中心-支付管理-支付通道-新增', 'payChannelCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 35, '系统管理-配置中心-支付管理-支付通道-编辑', 'payChannelUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 36, '系统管理-配置中心-支付管理-支付通道-删除', 'payChannelDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 37, '平台配置-查看', 'platformConfigRead', '可任意读取平台配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 38, '平台配置-保存', 'platformConfigSave', '可任意保存平台配置表');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 39, '应用配置-常用-查看', 'platformConfigCommonRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 40, '应用配置-常用-保存', 'platformConfigCommonSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 41, '插件配置-短信-查看', 'platformConfigSmsRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 42, '插件配置-短信-保存', 'platformConfigSmsSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 43, '插件配置-实名认证-查看', 'platformConfigIdCardRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 44, '插件配置-实名认证-查看', 'platformConfigIdCardSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 45, '插件配置-一键登录-查看', 'platformConfigOneClickRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 46, '插件配置-一键登录-保存', 'platformConfigOneClickSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 47, '插件配置-推送-查看', 'platformConfigPushRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 48, '插件配置-推送-查看', 'platformConfigPushSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 49, '插件配置-视频点播-查看', 'platformConfigVodRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 50, '插件配置-视频点播-保存', 'platformConfigVodSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 51, '插件配置-微信-查看', 'platformConfigWxRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 52, '插件配置-微信-查看', 'platformConfigWxSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 53, '插件配置-邮箱-查看', 'platformConfigEmailRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 54, '插件配置-邮箱-保存', 'platformConfigEmailSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 55, '系统管理-APP-查看', 'appRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 56, '系统管理-APP-新增', 'appCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 57, '系统管理-APP-编辑', 'appUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 58, '系统管理-APP-删除', 'appDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 59, '用户管理-用户-查看', 'usersRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 60, '用户管理-用户-编辑', 'usersUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 61, '机构管理-机构-查看', 'orgRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 62, '机构管理-机构-新增', 'orgCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 63, '机构管理-机构-编辑', 'orgUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 64, '机构管理-机构-删除', 'orgDelete', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 65, '权限管理-机构管理员-查看', 'orgAdminRead', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 66, '权限管理-机构管理员-新增', 'orgAdminCreate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 67, '权限管理-机构管理员-编辑', 'orgAdminUpdate', '');
INSERT INTO `auth_action` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 68, '权限管理-机构管理员-删除', 'orgAdminDelete', '');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS `auth_action_rel_to_scene`;
CREATE TABLE `auth_action_rel_to_scene`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `action_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  PRIMARY KEY (`action_id`, `scene_id`) USING BTREE,
  INDEX `action_id`(`action_id` ASC) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限操作，权限场景关联表（操作可用在哪些场景）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 2, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 3, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 4, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 5, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 6, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 7, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 8, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 9, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 10, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 11, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 12, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 13, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 13, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 14, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 14, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 15, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 15, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 16, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 16, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 17, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 18, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 19, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 20, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 21, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 22, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 23, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 24, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 25, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 26, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 27, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 28, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 29, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 30, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 31, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 32, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 33, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 34, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 35, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 36, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 37, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 38, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 39, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 40, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 41, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 42, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 43, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 44, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 45, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 46, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 47, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 48, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 49, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 50, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 51, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 52, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 53, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 54, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 55, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 56, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 57, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 58, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 59, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 60, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 61, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 62, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 63, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 64, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 65, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 65, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 66, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 66, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 67, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 67, 2);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 68, 1);
INSERT INTO `auth_action_rel_to_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 68, 2);

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
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `id_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '层级路径',
  `menu_icon` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}',
  `menu_url` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '链接',
  `extra_data` json NULL COMMENT '额外数据。JSON格式：{\"i18n（国际化设置）\": {\"title\": {\"语言标识\":\"标题\",...}}',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 100 COMMENT '排序值。从大到小排序',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `scene_id`(`scene_id` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 27 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限菜单表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '主页', 1, 0, 1, '0-1', 'autoicon-ep-home-filled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 255);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理', 1, 0, 1, '0-2', 'autoicon-ep-lock', '', '{\"i18n\": {\"title\": {\"en\": \"Auth Manage\", \"zh-cn\": \"权限管理\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '场景', 1, 2, 2, '0-2-3', 'autoicon-ep-flag', '/auth/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"场景\"}}}', 0);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '操作', 1, 2, 2, '0-2-4', 'autoicon-ep-coordinate', '/auth/action', '{\"i18n\": {\"title\": {\"en\": \"Action\", \"zh-cn\": \"操作\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '菜单', 1, 2, 2, '0-2-5', 'autoicon-ep-menu', '/auth/menu', '{\"i18n\": {\"title\": {\"en\": \"Menu\", \"zh-cn\": \"菜单\"}}}', 30);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '角色', 1, 2, 2, '0-2-6', 'autoicon-ep-view', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 40);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '平台管理员', 1, 2, 2, '0-2-7', 'vant-manager-o', '/platform/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"平台管理员\"}}}', 50);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '系统管理', 1, 0, 1, '0-8', 'autoicon-ep-platform', '', '{\"i18n\": {\"title\": {\"en\": \"System Manage\", \"zh-cn\": \"系统管理\"}}}', 20);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '配置中心', 1, 8, 2, '0-8-9', 'autoicon-ep-setting', '', '{\"i18n\": {\"title\": {\"en\": \"Config Center\", \"zh-cn\": \"配置中心\"}}}', 0);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '上传配置', 1, 9, 3, '0-8-9-10', 'autoicon-ep-upload', '/upload/upload', '{\"i18n\": {\"title\": {\"en\": \"Upload\", \"zh-cn\": \"上传配置\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '支付管理', 1, 9, 3, '0-8-9-11', 'autoicon-ep-coin', '', '{\"i18n\": {\"title\": {\"en\": \"\", \"zh-cn\": \"支付管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '支付配置', 1, 11, 4, '0-8-9-11-12', 'autoicon-ep-money', '/pay/pay', '{\"i18n\": {\"title\": {\"en\": \"Pay\", \"zh-cn\": \"支付配置\"}}}', 50);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '支付场景', 1, 11, 4, '0-8-9-11-13', 'autoicon-ep-guide', '/pay/scene', '{\"i18n\": {\"title\": {\"en\": \"Scene\", \"zh-cn\": \"支付场景\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '支付通道', 1, 11, 4, '0-8-9-11-14', 'autoicon-ep-connection', '/pay/channel', '{\"i18n\": {\"title\": {\"en\": \"Channel\", \"zh-cn\": \"支付通道\"}}}', 150);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '插件配置', 1, 9, 3, '0-8-9-15', 'autoicon-ep-ticket', '/platform/config/plugin', '{\"i18n\": {\"title\": {\"en\": \"Plugin Config\", \"zh-cn\": \"插件配置\"}}}', 150);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '应用配置', 1, 9, 3, '0-8-9-16', 'autoicon-ep-set-up', '/platform/config/app', '{\"i18n\": {\"title\": {\"en\": \"APP Config\", \"zh-cn\": \"应用配置\"}}}', 200);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, 'APP', 1, 8, 2, '0-8-17', 'vant-apps-o', '/app/app', '{\"i18n\": {\"title\": {\"en\": \"App\", \"zh-cn\": \"APP\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, '用户管理', 1, 0, 1, '0-18', 'vant-friends', '', '{\"i18n\": {\"title\": {\"en\": \"User Manage\", \"zh-cn\": \"用户管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '用户', 1, 18, 2, '0-18-19', 'vant-user-o', '/users/users', '{\"i18n\": {\"title\": {\"en\": \"Users\", \"zh-cn\": \"用户\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '机构管理', 1, 0, 1, '0-20', 'autoicon-ep-office-building', '', '{\"i18n\": {\"title\": {\"en\": \"\", \"zh-cn\": \"机构管理\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '机构', 1, 20, 2, '0-20-21', 'autoicon-ep-school', '/org/org', '{\"i18n\": {\"title\": {\"en\": \"Org\", \"zh-cn\": \"机构\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '机构管理员', 1, 2, 2, '0-2-22', 'vant-manager-o', '/org/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"机构管理员\"}}}', 100);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '主页', 2, 0, 1, '0-23', 'autoicon-ep-home-filled', '/', '{\"i18n\": {\"title\": {\"en\": \"Homepage\", \"zh-cn\": \"主页\"}}}', 255);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 24, '权限管理', 2, 0, 1, '0-24', 'autoicon-ep-menu', '', '{\"i18n\": {\"title\": {\"en\": \"\", \"zh-cn\": \"权限管理\"}}}', 10);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 25, '角色', 2, 24, 2, '0-24-25', 'autoicon-ep-view', '/auth/role', '{\"i18n\": {\"title\": {\"en\": \"Role\", \"zh-cn\": \"角色\"}}}', 40);
INSERT INTO `auth_menu` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 26, '管理员', 2, 24, 2, '0-24-26', 'vant-manager-o', '/org/admin', '{\"i18n\": {\"title\": {\"en\": \"Admin\", \"zh-cn\": \"管理员\"}}}', 100);

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
  `scene_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '场景ID',
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
  `action_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作ID',
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
  `scene_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '场景ID',
  `scene_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `scene_code` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标识',
  `scene_config` json NULL COMMENT '配置。JSON格式，根据场景设置',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`scene_id`) USING BTREE,
  UNIQUE INDEX `scene_code`(`scene_code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限场景表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '平台后台', 'platform', '{\"token_config\": {\"is_ip\": 1, \"is_unique\": 1, \"sign_type\": \"ES256\", \"public_key\": \"-----BEGIN PUBLIC KEY-----\\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42\\nhF5PnYenbWprk1MQIzT2+V4rRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\\n-----END PUBLIC KEY-----\", \"token_type\": 0, \"active_time\": 3600, \"expire_time\": 604800, \"private_key\": \"-----BEGIN EC PRIVATE KEY-----\\nMHcCAQEEIKvYPRtCqy9MI/yhx4L4+Sog/5lntHbuwxPg/JI0qW6LoAoGCCqGSM49\\nAwEHoUQDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42hF5PnYenbWprk1MQIzT2+V4r\\nRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\\n-----END EC PRIVATE KEY-----\"}}', '');
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '机构后台', 'org', '{\"token_config\": {\"is_ip\": 1, \"is_unique\": 1, \"sign_type\": \"RS256\", \"public_key\": \"-----BEGIN PUBLIC KEY-----\\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0nT4zSS2O+2yib+gdBm0\\nW0kU2AL6felZFp5U6B/ySeJnM8UE+fZwytgPuND5Z07khxtHR/YYH7huPGZ/fAgh\\nZHmNE5K5phMI5eETwPjk3RDeyYAyOosrKr3SAjjEQxJBISvBEillH4bKjaa4WF5/\\n+nGcp7f4e49caW/CfwuC2ZVrvySCPf1lR8o7/4Zz/hWUwgsEd/crR7ojgt+rbPeE\\n7+Cz11sZUZaMipTqsU3RVhbwtMyLdkos6KsYW7TZEK0VYt94/1XJQBUEjtCdDpS7\\n0XNF8ENpQPtuQdYE6/y+Jku8T9pqQQq/SOL6uPgsn6zJQ41u/l2AhG0i5GxYD86C\\n5wIDAQAB\\n-----END PUBLIC KEY-----\", \"token_type\": 0, \"active_time\": 3600, \"expire_time\": 14400, \"private_key\": \"-----BEGIN PRIVATE KEY-----\\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDSdPjNJLY77bKJ\\nv6B0GbRbSRTYAvp96VkWnlToH/JJ4mczxQT59nDK2A+40PlnTuSHG0dH9hgfuG48\\nZn98CCFkeY0TkrmmEwjl4RPA+OTdEN7JgDI6iysqvdICOMRDEkEhK8ESKWUfhsqN\\nprhYXn/6cZynt/h7j1xpb8J/C4LZlWu/JII9/WVHyjv/hnP+FZTCCwR39ytHuiOC\\n36ts94Tv4LPXWxlRloyKlOqxTdFWFvC0zIt2SizoqxhbtNkQrRVi33j/VclAFQSO\\n0J0OlLvRc0XwQ2lA+25B1gTr/L4mS7xP2mpBCr9I4vq4+CyfrMlDjW7+XYCEbSLk\\nbFgPzoLnAgMBAAECggEAGolMO9WmsrzAd9T1Pt5k2uPGoIwTmJ+9L3hsXU515vII\\nsELl4zy7MSB4LwYOhIOylgSPAthZZ1qCb9Q+u91slHYtHywvg2zAAPhV3M2lUeiI\\nJuEmtDILEdsYaVZODOT22F9je05D5WtCDAVbFi1oNqRvq8grKS1E6jiAzjMd3yBY\\n5AgUUP8sbS7BDdPus2t3mCAXqdtFkxn8wo/4WdMV6vG4p9p+a8dIoiRYHNBIw4sU\\nCYPiE7q52tVqVl10ShrJQlvoyDvmJam4inbl8ZWtbQLUsQxfzoEUfuuC0mYc2pES\\n1kp1pWfJNc5JtAXXZ/k9F4jLvjMp9KJximOG+E5tmQKBgQD5RwG8xu7s2JcocogD\\nuJjWfLz/4ab7Zs2NReCX6jmb522d1TqyiU9ilc0XBz8gDNeMwnOfWIyGR3Vtuub+\\nU56Y9/q+IMZ0ewdrUTR2NADZdqLH1ViVGacnnmJXJ+30Z8eUO0GCU7++DSQ1u8xh\\ng0mmrj6++xHsYJM3bCxBstNzGQKBgQDYIfOUIK+JKvYZ5idUrFWZnnVSdLu8IJHi\\nA0osaMB9VNtAdqrPVIA3L9AbIR1/la3gb3ILP0hIM7glt4i78WTwVrh0qkaz/8dw\\nX3EL5u9OAMecVFn4+gZon4RqbzjdtHso87v2GrIZ88eVOWjXRq93gWAe17G+noyS\\n44PgB7Zl/wKBgC3bOh6YGevIDEaMiyjkFHmgiMQppqYoyzdp218W33ImqKuYRiwB\\nxnDETe4mjx4+PojOXKa7i15IVvnQoB25FDvfomjHbrqOx1aeoZ/9AQsAIAHS5XDI\\nP0+yezS9S7DiRnymSe7HqUY09KxN19M4a5wWAcTwOuPZADv50kpjszJBAoGAS0io\\nO7SW8ESSrLrKgGf2+SeE3k/jBMijiAJ1V7q1MfLY3D95h/Z7Ir340zpZuBM/Gao4\\nI0rLtrqtLhYb/rs62aybW6fkMNarda0JB4hNWvJSlVWccWlFyjOmQBy1xiQTslQT\\n6Mmrt/Z+UrBIoJPyksHx5UxkkW1QsemmCecl1akCgYAys4PNRhdToxS5WDF7Tddg\\n8ghhobuwUCP2pxNYakX9HeHZsAjwwQhx6vGJqZPs5hh8HRinGWSvKyVXaUGVN+5b\\nFjMz/rhO9vTCZJS3aJSXxr0PFTbpP/AZSXwCBxjp+uEFTD8GJRX7/6+wlTk7+uTj\\nl7klp31noFtzz+onGSmqvA==\\n-----END PRIVATE KEY-----\"}}', '');
INSERT INTO `auth_scene` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, 'APP', 'app', '{\"token_config\": {\"is_ip\": 0, \"is_unique\": 0, \"sign_type\": \"HS256\", \"token_type\": 0, \"active_time\": 0, \"expire_time\": 604800, \"private_key\": \"任意字符串\"}}', '');

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
-- Table structure for pay
-- ----------------------------
DROP TABLE IF EXISTS `pay`;
CREATE TABLE `pay`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `pay_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付ID',
  `pay_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `pay_type` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0支付宝 1微信',
  `pay_config` json NOT NULL COMMENT '配置。根据pay_type类型设置',
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
  PRIMARY KEY (`order_id`, `rel_order_type`, `rel_order_id`) USING BTREE,
  INDEX `order_id`(`order_id` ASC) USING BTREE,
  INDEX `rel_order_id`(`rel_order_id` ASC) USING BTREE,
  INDEX `rel_order_no`(`rel_order_no` ASC) USING BTREE,
  INDEX `rel_order_user_id`(`rel_order_user_id` ASC) USING BTREE
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
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailCode', '{\"subject\":\"您的验证码\",\"template\":\"验证码：{code}\\n说明：\\n1. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。\\n2. 出于安全考虑，请不要将此验证码分享给任何人。\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommon', '{\"fromEmail\":\"xxxxxxxx@qq.com\",\"password\":\"xxxxxxxx\",\"smtpHost\":\"smtp.qq.com\",\"smtpPort\":\"465\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailType', 'emailOfCommon');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardOfAliyun', '{\"appcode\":\"appcode\",\"url\":\"http://idcard.market.alicloudapi.com/lianzhuo/idcard\"}');
INSERT INTO `platform_config` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardType', 'idCardOfAliyun');

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
  `upload_config` json NOT NULL COMMENT '配置。根据upload_type类型设置',
  `is_default` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '默认：0否 1是',
  `remark` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`upload_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '上传表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of upload
-- ----------------------------
INSERT INTO `upload` VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 0, '{\"url\": \"/upload/upload\", \"signKey\": \"secretKey\", \"fileSaveDir\": \"../public/\", \"fileUrlPrefix\": \"\"}', 1, '此项目自带简易文件上传接口，故可将此项目部署到服务器，对外提供文件上传下载服务');

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
