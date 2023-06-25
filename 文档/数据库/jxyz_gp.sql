/*
 Navicat Premium Data Transfer

 Source Server         : 本地-8.0.32
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 192.168.0.16:3306
 Source Schema         : jxyz_gp

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 29/04/2023 09:56:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ad
-- ----------------------------
DROP TABLE IF EXISTS `ad`;
CREATE TABLE `ad`  (
  `adId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `adType` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：1活动，2公告',
  `adName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `adImage` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '图片',
  `adUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `startTime` datetime NOT NULL COMMENT '开始时间',
  `endTime` datetime NOT NULL COMMENT '结束时间',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ad
-- ----------------------------

-- ----------------------------
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS `admin`;
CREATE TABLE `admin`  (
  `adminId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id（平台超级管理员为0无权限限制，其他管理员必须设置权限组id）',
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO `admin` VALUES (1, 0, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 0, '2021-03-04 15:19:32', '2019-08-28 15:50:40');

-- ----------------------------
-- Table structure for agent
-- ----------------------------
DROP TABLE IF EXISTS `agent`;
CREATE TABLE `agent`  (
  `agentId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id',
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '账号',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `agentWallet` decimal(18, 8) NOT NULL DEFAULT 0.00000000 COMMENT '代理钱包',
  `agentRate` json NULL COMMENT '费率（格式：{\"计算方式1\":\"不同方式不同\",...}）',
  `inviteCode` char(4) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '邀请码',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`agentId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  UNIQUE INDEX `inviteCode`(`inviteCode` ASC) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of agent
-- ----------------------------

-- ----------------------------
-- Table structure for agent_rel_user
-- ----------------------------
DROP TABLE IF EXISTS `agent_rel_user`;
CREATE TABLE `agent_rel_user`  (
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  UNIQUE INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of agent_rel_user
-- ----------------------------

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app`  (
  `appId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `appCode` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT 'app标识',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`appId`) USING BTREE,
  UNIQUE INDEX `appCode`(`appCode` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app
-- ----------------------------

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group`  (
  `authGroupId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `scene` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '适用场景：platform平台，其他待扩展',
  `tableId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '表id（0表示平台创建，其他值根据scene对应不同表id）',
  `groupName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '组名称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`authGroupId`) USING BTREE,
  INDEX `scene`(`scene` ASC) USING BTREE,
  INDEX `tableId`(`tableId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_group
-- ----------------------------

-- ----------------------------
-- Table structure for auth_group_rel_menu_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_group_rel_menu_action`;
CREATE TABLE `auth_group_rel_menu_action`  (
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id',
  `authMenuId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限菜单id',
  `authMenuActionId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限菜单操作id',
  PRIMARY KEY (`authGroupId`, `authMenuId`, `authMenuActionId`) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `authMenuId`(`authMenuId` ASC) USING BTREE,
  INDEX `authMenuActionId`(`authMenuActionId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_group_rel_menu_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu`;
CREATE TABLE `auth_menu`  (
  `authMenuId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `scene` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '权限类型：platform平台，其他待扩展',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级id（只支持2级，后台侧边菜单两层够用了）',
  `menuName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `menuIcon` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '后台显示权限菜单时的图标（根据后台模板设置）',
  `menuUrl` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '菜单链接地址(和scene组成唯一索引，程序内与scene，actionCode<表auth_action>同时使用，用于判断权限)',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`authMenuId`) USING BTREE,
  UNIQUE INDEX `scene`(`scene` ASC, `menuUrl` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO `auth_menu` VALUES (1, 'platform', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-03-15 22:47:10', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (2, 'platform', 0, '权限管理', 'el-icon-lock', '/menuAuth', 100, 0, '2021-03-10 14:45:36', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (3, 'platform', 2, '管理员', '', '/admin/index', 50, 0, '2021-03-04 11:21:25', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (4, 'platform', 2, '权限菜单', '', '/authMenu/index', 100, 0, '2021-03-04 11:26:13', '2021-03-04 11:24:18');
INSERT INTO `auth_menu` VALUES (5, 'platform', 2, '权限组', '', '/authGroup/index', 50, 0, '2021-03-04 11:26:47', '2021-03-04 11:25:13');
INSERT INTO `auth_menu` VALUES (6, 'platform', 0, '系统管理', 'el-icon-setting', '/menuSystem', 99, 0, '2021-03-10 14:45:27', '2021-03-04 11:25:47');
INSERT INTO `auth_menu` VALUES (7, 'platform', 6, '系统设置', '', '/config/index', 50, 0, '2021-03-04 11:28:43', '2021-03-04 11:27:54');
INSERT INTO `auth_menu` VALUES (8, 'platform', 0, '日志管理', 'el-icon-view', '/menuLog', 95, 0, '2021-03-10 14:45:20', '2021-03-04 11:28:35');
INSERT INTO `auth_menu` VALUES (9, 'platform', 8, '请求日志', '', '/logRequest/index', 50, 0, '2021-03-06 11:12:03', '2021-03-04 11:28:55');
INSERT INTO `auth_menu` VALUES (10, 'platform', 0, '游戏管理', 'el-icon-s-grid', '/menuGame', 50, 0, '2021-05-06 16:05:28', '2021-05-06 16:05:28');
INSERT INTO `auth_menu` VALUES (11, 'platform', 0, '用户管理', 'el-icon-user-solid', '/menuUser', 50, 0, '2021-05-06 16:06:02', '2021-05-06 16:06:02');
INSERT INTO `auth_menu` VALUES (12, 'platform', 11, '用户列表', '', '/user/index', 50, 0, '2021-05-06 16:06:58', '2021-05-06 16:06:58');
INSERT INTO `auth_menu` VALUES (13, 'platform', 11, '充值记录', '', '/userRecharge/index', 50, 0, '2021-12-22 15:29:50', '2021-05-06 16:07:37');
INSERT INTO `auth_menu` VALUES (14, 'platform', 11, '提现记录', '', '/userWithdraw/index', 50, 0, '2021-12-22 15:29:52', '2021-05-06 16:08:04');
INSERT INTO `auth_menu` VALUES (15, 'platform', 22, '资金流水', '', '/userWalletLog/index', 50, 0, '2021-12-23 14:55:46', '2021-05-06 16:10:47');
INSERT INTO `auth_menu` VALUES (16, 'platform', 6, '广告', '', '/ad/index', 50, 0, '2021-05-06 16:11:24', '2021-05-06 16:11:24');
INSERT INTO `auth_menu` VALUES (17, 'platform', 6, '支付通道', '', '/payChannel/index', 50, 0, '2021-05-06 16:12:41', '2021-05-06 16:12:23');
INSERT INTO `auth_menu` VALUES (18, 'platform', 10, '游戏分类', '', '/gameCategory/index', 100, 0, '2021-05-06 16:13:29', '2021-05-06 16:13:29');
INSERT INTO `auth_menu` VALUES (19, 'platform', 10, '游戏', '', '/game/index', 50, 0, '2021-05-06 16:14:10', '2021-05-06 16:14:10');
INSERT INTO `auth_menu` VALUES (20, 'platform', 0, '代理管理', 'el-icon-s-check', '/menuAgent', 50, 0, '2021-12-23 17:20:55', '2021-12-18 09:42:28');
INSERT INTO `auth_menu` VALUES (21, 'platform', 20, '代理列表', '', '/agent/index', 50, 0, '2021-12-18 09:43:21', '2021-12-18 09:43:12');
INSERT INTO `auth_menu` VALUES (22, 'platform', 0, '数据统计', 'el-icon-data-analysis', '/menuStatistic', 50, 0, '2021-12-20 16:52:43', '2021-12-20 16:51:58');
INSERT INTO `auth_menu` VALUES (23, 'platform', 6, 'APP列表', '', '/app/index', 50, 0, '2021-12-22 13:35:47', '2021-12-22 13:35:47');
INSERT INTO `auth_menu` VALUES (24, 'platform', 11, '游戏记录', '', '/userGameRecord/index', 50, 0, '2021-12-22 15:30:02', '2021-12-22 15:29:24');
INSERT INTO `auth_menu` VALUES (25, 'platform', 22, '日统计-游戏', '', '/statisticDayGame/index', 50, 0, '2021-12-23 17:24:26', '2021-12-23 17:24:26');
INSERT INTO `auth_menu` VALUES (26, 'platform', 22, '日统计-用户', '', '/statisticDayUser/index', 50, 0, '2021-12-23 17:25:58', '2021-12-23 17:24:38');
INSERT INTO `auth_menu` VALUES (27, 'platform', 22, '日统计-代理', '', '/statisticDayAgent/index', 50, 0, '2021-12-23 17:25:36', '2021-12-23 17:25:36');
INSERT INTO `auth_menu` VALUES (28, 'platform', 22, '日统计-通道', '', '/statisticDayPayChannel/index', 50, 0, '2021-12-23 17:26:37', '2021-12-23 17:26:37');
INSERT INTO `auth_menu` VALUES (29, 'agent', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-12-25 10:58:03', '2021-12-25 10:58:03');
INSERT INTO `auth_menu` VALUES (30, 'agent', 0, '用户管理', 'el-icon-user-solid', '/menuUser', 50, 0, '2021-12-25 10:59:39', '2021-12-25 10:59:39');
INSERT INTO `auth_menu` VALUES (31, 'agent', 30, '用户列表', '', '/user/index', 50, 0, '2021-12-25 11:00:11', '2021-12-25 11:00:11');
INSERT INTO `auth_menu` VALUES (32, 'agent', 0, '数据统计', 'el-icon-data-analysis', '/menuStatistic', 50, 0, '2021-12-25 11:00:53', '2021-12-25 11:00:53');
INSERT INTO `auth_menu` VALUES (33, 'agent', 32, '资金流水', '', '/userWalletLog/index', 50, 0, '2021-12-25 11:02:03', '2021-12-25 11:02:03');
INSERT INTO `auth_menu` VALUES (34, 'agent', 32, '日统计-代理', '', '/statisticDayAgent/index', 50, 0, '2021-12-25 11:02:44', '2021-12-25 11:02:44');
INSERT INTO `auth_menu` VALUES (35, 'platform', 22, '日统计-平台', '', '/statisticDay/index', 50, 0, '2021-12-25 17:40:42', '2021-12-25 17:40:42');
INSERT INTO `auth_menu` VALUES (36, 'platform', 8, '用户日志', '', '/logUser/index', 50, 0, '2022-02-14 10:56:47', '2022-02-14 10:56:47');

-- ----------------------------
-- Table structure for auth_menu_action
-- ----------------------------
DROP TABLE IF EXISTS `auth_menu_action`;
CREATE TABLE `auth_menu_action`  (
  `authMenuActionId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `authMenuId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限菜单id',
  `actionCode` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '操作Code（程序内与authType，menuUrl<表auth_menu>同时使用，用于判断权限)',
  `actionName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '操作名称',
  PRIMARY KEY (`authMenuActionId`) USING BTREE,
  UNIQUE INDEX `authMenuId_2`(`authMenuId` ASC, `actionCode` ASC) USING BTREE,
  INDEX `authMenuId`(`authMenuId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 60 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_menu_action
-- ----------------------------
INSERT INTO `auth_menu_action` VALUES (1, 1, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (2, 3, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (3, 3, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (4, 3, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (5, 3, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (6, 4, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (7, 4, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (8, 4, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (9, 4, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (10, 5, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (11, 5, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (12, 5, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (13, 5, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (14, 7, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (15, 7, 'save', '保存');
INSERT INTO `auth_menu_action` VALUES (16, 9, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (17, 12, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (18, 12, 'audit', '审核');
INSERT INTO `auth_menu_action` VALUES (19, 13, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (20, 13, 'audit', '审核');
INSERT INTO `auth_menu_action` VALUES (21, 14, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (22, 14, 'audit', '审核');
INSERT INTO `auth_menu_action` VALUES (23, 15, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (24, 16, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (25, 16, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (26, 16, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (27, 16, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (28, 17, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (29, 17, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (30, 17, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (31, 17, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (32, 18, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (33, 18, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (34, 18, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (35, 18, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (36, 19, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (37, 19, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (38, 19, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (39, 19, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (40, 21, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (41, 21, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (42, 21, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (43, 21, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (44, 23, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (45, 23, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (46, 23, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (47, 23, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (48, 24, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (49, 25, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (50, 27, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (51, 26, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (52, 28, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (53, 29, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (54, 31, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (55, 31, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (56, 33, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (57, 34, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (58, 35, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (59, 36, 'sel', '查看');

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config`  (
  `configId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `configKey` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '配置项Key',
  `configValue` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '配置项值',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `configKey`(`configKey` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of config
-- ----------------------------

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS `game`;
CREATE TABLE `game`  (
  `gameId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `gameCategoryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏分类id',
  `gameType` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏类型：1即时，2过程',
  `mid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '母id（设置母id，签名将用母id游戏的token签名）',
  `gameName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏名称',
  `gameUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏链接',
  `gameIcon` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '游戏图标',
  `gameWallet` decimal(15, 5) NOT NULL DEFAULT 0.00000 COMMENT '游戏钱包',
  `gameRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '费率',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isHot` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否热门：0否 1是',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `tokenId` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'tokenId（接口调用身份证明）',
  `token` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'token（接口调用签名验证使用）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`gameId`) USING BTREE,
  UNIQUE INDEX `tokenId`(`tokenId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE,
  INDEX `gameCategoryId`(`gameCategoryId` ASC) USING BTREE,
  INDEX `isHot`(`isHot` ASC) USING BTREE,
  INDEX `gameType`(`gameType` ASC) USING BTREE,
  INDEX `mid`(`mid` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of game
-- ----------------------------

-- ----------------------------
-- Table structure for game_category
-- ----------------------------
DROP TABLE IF EXISTS `game_category`;
CREATE TABLE `game_category`  (
  `gameCategoryId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级id',
  `idPath` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '层级路径',
  `level` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '层级',
  `gameCategoryName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `gameCategoryIcon` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`gameCategoryId`) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE,
  INDEX `level`(`level` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of game_category
-- ----------------------------

-- ----------------------------
-- Table structure for log_request
-- ----------------------------
DROP TABLE IF EXISTS `log_request`;
CREATE TABLE `log_request`  (
  `logRequestId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `requestUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求路径',
  `requestData` json NULL COMMENT '请求参数',
  `requestHeaders` json NULL COMMENT '请求头',
  `responseData` json NULL COMMENT '响应数据',
  `runTime` decimal(8, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '响应时间（单位：毫秒）',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`logRequestId`) USING BTREE,
  INDEX `requestUrl`(`requestUrl` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_request
-- ----------------------------

-- ----------------------------
-- Table structure for log_user
-- ----------------------------
DROP TABLE IF EXISTS `log_user`;
CREATE TABLE `log_user`  (
  `logUserId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `logUserType` varchar(10) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '类型：register注册，login登录',
  `userId` int UNSIGNED NOT NULL DEFAULT 0,
  `requestUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '请求链接',
  `clientIp` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '客户端IP',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`logUserId`) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `clientIp`(`clientIp` ASC) USING BTREE,
  INDEX `logUserType`(`logUserType` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_user
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel`  (
  `payChannelId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `payment` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付方式：aliPay（支付宝），wxPay（微信），ysfPay（云闪付）',
  `payChannelName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付名称',
  `payChannelConfig` json NOT NULL COMMENT '配置参数',
  `payChannelWallet` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '通道金额',
  `payChannelRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '费率',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`payChannelId`) USING BTREE,
  INDEX `payment`(`payment` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel
-- ----------------------------

-- ----------------------------
-- Table structure for statistic_day
-- ----------------------------
DROP TABLE IF EXISTS `statistic_day`;
CREATE TABLE `statistic_day`  (
  `statisticId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `day` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '统计日期（格式：年月日）',
  `recharge` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值',
  `rechargeReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值奖励',
  `payChannelPercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '通道提成（每笔充值*通道费率的总和）',
  `withdraw` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提现',
  `gameGain` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '游戏收益',
  `gameGainPercentage` decimal(15, 5) NOT NULL DEFAULT 0.00000 COMMENT '游戏收益提成（游戏每笔交易*游戏费率的总和）',
  `peopleReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '代理人头费',
  `rechargePercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '代理充值提成',
  `profit` decimal(15, 5) NOT NULL COMMENT '利润',
  `profitPercentage` decimal(18, 8) UNSIGNED NOT NULL DEFAULT 0.00000000 COMMENT '代理利润提成',
  PRIMARY KEY (`statisticId`) USING BTREE,
  UNIQUE INDEX `day_2`(`day` ASC) USING BTREE,
  INDEX `day`(`day` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of statistic_day
-- ----------------------------

-- ----------------------------
-- Table structure for statistic_day_agent
-- ----------------------------
DROP TABLE IF EXISTS `statistic_day_agent`;
CREATE TABLE `statistic_day_agent`  (
  `statisticId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `day` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '统计日期（格式：年月日）',
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `peopleReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '人头费',
  `rechargePercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '充值提成',
  `profit` decimal(15, 5) NOT NULL COMMENT '利润',
  `profitRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '利润费率',
  PRIMARY KEY (`statisticId`) USING BTREE,
  UNIQUE INDEX `day_2`(`day` ASC, `agentId` ASC) USING BTREE,
  INDEX `day`(`day` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of statistic_day_agent
-- ----------------------------

-- ----------------------------
-- Table structure for statistic_day_game
-- ----------------------------
DROP TABLE IF EXISTS `statistic_day_game`;
CREATE TABLE `statistic_day_game`  (
  `statisticId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `day` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '统计日期（格式：年月日）',
  `gameId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `gameGain` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '游戏收益',
  `gameGainPercentage` decimal(15, 5) NOT NULL DEFAULT 0.00000 COMMENT '游戏收益提成（游戏每笔交易*游戏费率的总和）',
  PRIMARY KEY (`statisticId`) USING BTREE,
  UNIQUE INDEX `day_2`(`day` ASC, `gameId` ASC) USING BTREE,
  INDEX `day`(`day` ASC) USING BTREE,
  INDEX `gameId`(`gameId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of statistic_day_game
-- ----------------------------

-- ----------------------------
-- Table structure for statistic_day_payChannel
-- ----------------------------
DROP TABLE IF EXISTS `statistic_day_payChannel`;
CREATE TABLE `statistic_day_payChannel`  (
  `statisticId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `day` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '统计日期（格式：年月日）',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道ID',
  `payChannelAmount` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '通道当天的充值金额',
  `payChannelPercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '通道提成（每笔充值*通道费率的总和）',
  PRIMARY KEY (`statisticId`) USING BTREE,
  UNIQUE INDEX `day_2`(`day` ASC, `payChannelId` ASC) USING BTREE,
  INDEX `day`(`day` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of statistic_day_payChannel
-- ----------------------------

-- ----------------------------
-- Table structure for statistic_day_user
-- ----------------------------
DROP TABLE IF EXISTS `statistic_day_user`;
CREATE TABLE `statistic_day_user`  (
  `statisticId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `day` char(8) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '统计日期（格式：年月日）',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `recharge` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值',
  `rechargeReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值奖励',
  `payChannelPercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '通道提成（每笔充值*通道费率的总和）',
  `withdraw` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提现',
  `gameGain` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '游戏收益',
  `gameGainPercentage` decimal(15, 5) NOT NULL DEFAULT 0.00000 COMMENT '游戏收益提成（游戏每笔交易*游戏费率的总和）',
  `peopleReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '代理人头费',
  `rechargePercentage` decimal(15, 5) UNSIGNED NOT NULL DEFAULT 0.00000 COMMENT '代理充值提成',
  `profit` decimal(15, 5) NOT NULL COMMENT '利润',
  `profitRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '代理利润费率',
  PRIMARY KEY (`statisticId`) USING BTREE,
  UNIQUE INDEX `day_2`(`day` ASC, `userId` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `day`(`day` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of statistic_day_user
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `userName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '手机',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `userWallet` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '用户钱包',
  `gameId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '当isLock为1时表示正在玩的游戏id',
  `isLock` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否锁定：0否 1是（正在玩游戏时需锁定，不能同时玩另一个游戏，否则用户钱包计算上会出问题）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `uid` char(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '用户uid（给第三方的用户标识，用数字容易被猜到）',
  `loginIp` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '登录IP',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE INDEX `uid`(`uid` ASC) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  UNIQUE INDEX `userName`(`userName` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE,
  INDEX `loginIp`(`loginIp` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for user_game_record
-- ----------------------------
DROP TABLE IF EXISTS `user_game_record`;
CREATE TABLE `user_game_record`  (
  `userGameRecordId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `appId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'appid',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `gameId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `recordCode` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '单号',
  `betAmount` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '下注金额',
  `settleAmount` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '结算金额',
  `betGameRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '下注游戏费率',
  `settleGameRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '结算游戏费率',
  `startGameRequestData` json NULL COMMENT '开始游戏数据',
  `endGameRequestData` json NULL COMMENT '结束游戏数据',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userGameRecordId`) USING BTREE,
  UNIQUE INDEX `recordCode`(`recordCode` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `gameId`(`gameId` ASC) USING BTREE,
  INDEX `appId`(`appId` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_game_record
-- ----------------------------

-- ----------------------------
-- Table structure for user_recharge
-- ----------------------------
DROP TABLE IF EXISTS `user_recharge`;
CREATE TABLE `user_recharge`  (
  `userRechargeId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `appId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'APPID',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付通道ID',
  `orderCode` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '订单号',
  `payChannelOrderCode` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通道的订单号',
  `rechargeStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值状态：0未充值，1已充值',
  `amount` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值金额',
  `peopleReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '代理人头费',
  `rechargeRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '代理充值费率',
  `rechargeReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值奖励',
  `payChannelRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '通道费率',
  `payChannelResponseData` json NULL COMMENT '通道响应数据',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userRechargeId`) USING BTREE,
  UNIQUE INDEX `orderCode`(`orderCode` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE,
  INDEX `payChannelOrderCode`(`payChannelOrderCode` ASC) USING BTREE,
  INDEX `rechargeStatus`(`rechargeStatus` ASC) USING BTREE,
  INDEX `appId`(`appId` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_recharge
-- ----------------------------

-- ----------------------------
-- Table structure for user_wallet_log
-- ----------------------------
DROP TABLE IF EXISTS `user_wallet_log`;
CREATE TABLE `user_wallet_log`  (
  `userWalletLogId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `userWalletLogType` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '来源：1充值，2提现，3游戏',
  `appId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'APPID',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户ID',
  `agentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '代理ID',
  `userRechargeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '充值ID',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道ID',
  `userWithdrawId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现ID',
  `userGameRecordId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏记录ID',
  `gameId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '游戏ID',
  `amount` decimal(12, 2) NOT NULL DEFAULT 0.00 COMMENT '金额',
  `peopleReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '代理人头费',
  `rechargeRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '代理充值费率',
  `rechargeReward` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值奖励',
  `payChannelRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '通道费率',
  `gameRate` decimal(4, 3) UNSIGNED NOT NULL DEFAULT 0.000 COMMENT '游戏费率',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userWalletLogId`, `addTime`) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `userRechargeId`(`userRechargeId` ASC) USING BTREE,
  INDEX `userWithdrawId`(`userWithdrawId` ASC) USING BTREE,
  INDEX `gameId`(`gameId` ASC) USING BTREE,
  INDEX `appId`(`appId` ASC) USING BTREE,
  INDEX `userWalletLogType`(`userWalletLogType` ASC) USING BTREE,
  INDEX `agentId`(`agentId` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC PARTITION BY RANGE (to_days(`addTime`))
PARTITIONS 45
(PARTITION `p20211224` VALUES LESS THAN (738514) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211225` VALUES LESS THAN (738515) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211226` VALUES LESS THAN (738516) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211227` VALUES LESS THAN (738517) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211228` VALUES LESS THAN (738518) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211229` VALUES LESS THAN (738519) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211230` VALUES LESS THAN (738520) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211231` VALUES LESS THAN (738521) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220101` VALUES LESS THAN (738522) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220102` VALUES LESS THAN (738523) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220103` VALUES LESS THAN (738524) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220112` VALUES LESS THAN (738533) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220113` VALUES LESS THAN (738534) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220114` VALUES LESS THAN (738535) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220115` VALUES LESS THAN (738536) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220116` VALUES LESS THAN (738537) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220117` VALUES LESS THAN (738538) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220118` VALUES LESS THAN (738539) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220119` VALUES LESS THAN (738540) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220120` VALUES LESS THAN (738541) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220121` VALUES LESS THAN (738542) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220122` VALUES LESS THAN (738543) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220123` VALUES LESS THAN (738544) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220124` VALUES LESS THAN (738545) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220125` VALUES LESS THAN (738546) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220127` VALUES LESS THAN (738548) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220128` VALUES LESS THAN (738549) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220129` VALUES LESS THAN (738550) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220130` VALUES LESS THAN (738551) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220131` VALUES LESS THAN (738552) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220201` VALUES LESS THAN (738553) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220202` VALUES LESS THAN (738554) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220210` VALUES LESS THAN (738562) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220211` VALUES LESS THAN (738563) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220212` VALUES LESS THAN (738564) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220213` VALUES LESS THAN (738565) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220214` VALUES LESS THAN (738566) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220215` VALUES LESS THAN (738567) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220216` VALUES LESS THAN (738568) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220217` VALUES LESS THAN (738569) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220218` VALUES LESS THAN (738570) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220219` VALUES LESS THAN (738571) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220220` VALUES LESS THAN (738572) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220221` VALUES LESS THAN (738573) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220222` VALUES LESS THAN (738574) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 )
;

-- ----------------------------
-- Records of user_wallet_log
-- ----------------------------

-- ----------------------------
-- Table structure for user_withdraw
-- ----------------------------
DROP TABLE IF EXISTS `user_withdraw`;
CREATE TABLE `user_withdraw`  (
  `userWithdrawId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `appId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'appid',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `bankName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '银行名称',
  `bankAccountName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '户名',
  `bankNumber` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '卡号',
  `amount` decimal(12, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '充值金额',
  `withdrawStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '提现状态：0未提现，1已提现，2打款中，3禁止提现',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userWithdrawId`) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `withdrawStatus`(`withdrawStatus` ASC) USING BTREE,
  INDEX `appId`(`appId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_withdraw
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
