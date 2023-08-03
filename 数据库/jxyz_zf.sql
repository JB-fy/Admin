/*
 Navicat Premium Data Transfer

 Source Server         : 本地-8.0.32
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 192.168.0.16:3306
 Source Schema         : jxyz_zf

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 29/04/2023 09:56:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ad
-- ----------------------------
DROP TABLE IF EXISTS `ad`;
CREATE TABLE `ad`  (
  `adId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道ID',
  `zfbId` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付宝id',
  `zfbOrderId` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付宝订单id',
  `price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '金额',
  `remark` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `adStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '广告状态：0未发送，1已发送， 2未收红包',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`adId`) USING BTREE,
  INDEX `zfbId`(`zfbId` ASC) USING BTREE,
  INDEX `adStatus`(`adStatus` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE
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
  `googleAuthenticatorSecret` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '谷歌身份验证器密钥',
  `allowIpList` json NULL COMMENT '允许登录的ip列表',
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
INSERT INTO `admin` VALUES (1, 0, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'ASCJM2AC7FSY4JY3', NULL, 0, '2021-04-07 14:42:08', '2019-08-28 15:50:40');

-- ----------------------------
-- Table structure for auth_group
-- ----------------------------
DROP TABLE IF EXISTS `auth_group`;
CREATE TABLE `auth_group`  (
  `authGroupId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `scene` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '适用场景：platform平台，其他待扩展',
  `tableId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '表id（0表示平台创建，其他值根据authType对应不同表id）',
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
  `menuUrl` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '菜单链接地址(和authType组成唯一索引，程序内与authType，actionCode<表auth_action>同时使用，用于判断权限)',
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
INSERT INTO `auth_menu` VALUES (1, 'platform', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-03-15 22:46:44', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (2, 'platform', 0, '权限管理', 'el-icon-lock', '/menuAuth', 100, 0, '2021-03-10 14:45:36', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (3, 'platform', 2, '管理员', '', '/admin/index', 50, 0, '2021-03-04 11:21:25', '2021-03-04 11:19:59');
INSERT INTO `auth_menu` VALUES (4, 'platform', 2, '权限菜单', '', '/authMenu/index', 100, 0, '2021-03-04 11:26:13', '2021-03-04 11:24:18');
INSERT INTO `auth_menu` VALUES (5, 'platform', 2, '权限组', '', '/authGroup/index', 50, 0, '2021-03-04 11:26:47', '2021-03-04 11:25:13');
INSERT INTO `auth_menu` VALUES (6, 'platform', 0, '系统管理', 'el-icon-setting', '/menuSystem', 99, 0, '2021-03-10 14:45:27', '2021-03-04 11:25:47');
INSERT INTO `auth_menu` VALUES (7, 'platform', 6, '系统设置', '', '/config/index', 50, 0, '2021-03-04 11:28:43', '2021-03-04 11:27:54');
INSERT INTO `auth_menu` VALUES (8, 'platform', 0, '日志管理', 'el-icon-view', '/menuLog', 95, 0, '2021-03-10 14:45:20', '2021-03-04 11:28:35');
INSERT INTO `auth_menu` VALUES (9, 'platform', 8, '请求日志', '', '/logRequest/index', 50, 0, '2021-03-06 11:12:03', '2021-03-04 11:28:55');
INSERT INTO `auth_menu` VALUES (10, 'platform', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 50, 0, '2021-03-10 14:45:04', '2021-03-10 14:45:04');
INSERT INTO `auth_menu` VALUES (11, 'platform', 10, '订单列表', '', '/order/index', 50, 0, '2021-03-10 15:03:09', '2021-03-10 15:00:38');
INSERT INTO `auth_menu` VALUES (12, 'platform', 10, '业务员抽成', '', '/orderRelSalesman/index', 50, 0, '2021-03-10 15:02:14', '2021-03-10 15:02:14');
INSERT INTO `auth_menu` VALUES (13, 'platform', 10, '通道失败回调', '', '/payChannelNotify/index', 50, 0, '2021-04-06 13:32:01', '2021-03-10 15:04:28');
INSERT INTO `auth_menu` VALUES (14, 'platform', 0, '通道管理', 'el-icon-coin', '/menuPayChannel', 50, 0, '2021-03-15 22:11:11', '2021-03-15 22:02:07');
INSERT INTO `auth_menu` VALUES (15, 'platform', 14, '通道列表', '', '/payChannel/index', 50, 0, '2021-03-15 22:06:05', '2021-03-15 22:06:05');
INSERT INTO `auth_menu` VALUES (16, 'platform', 14, '通道组', '', '/payChannelGroup/index', 50, 0, '2021-03-15 22:07:13', '2021-03-15 22:07:13');
INSERT INTO `auth_menu` VALUES (17, 'platform', 0, '机构管理', 'el-icon-s-tools', '/menuOrganization', 50, 0, '2021-03-15 22:12:11', '2021-03-15 22:12:11');
INSERT INTO `auth_menu` VALUES (18, 'platform', 17, '机构列表', '', '/organization/index', 50, 0, '2021-03-15 22:25:05', '2021-03-15 22:13:26');
INSERT INTO `auth_menu` VALUES (19, 'platform', 17, '机构管理员', '', '/organizationAdmin/index', 50, 0, '2021-03-15 22:14:53', '2021-03-15 22:14:53');
INSERT INTO `auth_menu` VALUES (20, 'platform', 17, '提款列表', '', '/withdraw/index', 50, 0, '2021-03-15 22:15:52', '2021-03-15 22:15:52');
INSERT INTO `auth_menu` VALUES (21, 'platform', 0, '业务管理', 'el-icon-s-data', '/menuSalesman', 50, 0, '2021-03-15 22:23:06', '2021-03-15 22:23:06');
INSERT INTO `auth_menu` VALUES (22, 'platform', 21, '业务员', '', '/salesman/index', 50, 0, '2021-03-15 22:24:36', '2021-03-15 22:23:40');
INSERT INTO `auth_menu` VALUES (23, 'salesman', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-03-15 22:51:57', '2021-03-15 22:47:58');
INSERT INTO `auth_menu` VALUES (24, 'salesman', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 50, 0, '2021-03-15 22:48:55', '2021-03-15 22:48:55');
INSERT INTO `auth_menu` VALUES (25, 'salesman', 24, '订单列表', '', '/order/index', 50, 0, '2021-03-15 22:49:38', '2021-03-15 22:49:38');
INSERT INTO `auth_menu` VALUES (26, 'merchant', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-03-15 22:51:53', '2021-03-15 22:51:53');
INSERT INTO `auth_menu` VALUES (27, 'merchant', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 50, 0, '2021-03-15 22:53:40', '2021-03-15 22:53:40');
INSERT INTO `auth_menu` VALUES (28, 'merchant', 27, '订单列表', '', '/order/index', 50, 0, '2021-03-15 22:54:13', '2021-03-15 22:54:13');
INSERT INTO `auth_menu` VALUES (29, 'merchant', 0, '机构管理', 'el-icon-s-tools', '/menuOrganization', 50, 0, '2021-03-15 22:54:51', '2021-03-15 22:54:51');
INSERT INTO `auth_menu` VALUES (30, 'merchant', 29, '提款列表', '', '/withdraw/index', 50, 0, '2021-03-15 22:56:08', '2021-03-15 22:56:08');
INSERT INTO `auth_menu` VALUES (31, 'payService', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-03-15 23:03:32', '2021-03-15 23:03:32');
INSERT INTO `auth_menu` VALUES (32, 'payService', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 50, 0, '2021-03-15 23:03:59', '2021-03-15 23:03:59');
INSERT INTO `auth_menu` VALUES (33, 'payService', 32, '订单列表', '', '/order/index', 50, 0, '2021-03-15 23:04:35', '2021-03-15 23:04:35');
INSERT INTO `auth_menu` VALUES (34, 'payService', 0, '通道管理', 'el-icon-coin', '/menuPayChannel', 50, 0, '2021-03-18 16:58:04', '2021-03-15 23:05:13');
INSERT INTO `auth_menu` VALUES (35, 'payService', 34, '通道列表', '', '/payChannel/index', 50, 0, '2021-12-02 15:29:25', '2021-03-15 23:06:34');
INSERT INTO `auth_menu` VALUES (36, 'platform', 6, '广告', '', '/ad/index', 50, 0, '2022-02-09 10:04:08', '2022-02-09 10:04:08');

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
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `auth_menu_action` VALUES (17, 11, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (19, 12, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (20, 13, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (21, 15, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (22, 15, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (23, 15, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (24, 15, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (25, 16, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (26, 16, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (27, 16, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (28, 16, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (29, 18, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (30, 18, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (31, 18, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (32, 18, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (33, 19, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (34, 19, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (35, 19, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (36, 19, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (37, 20, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (38, 20, 'audit', '审核');
INSERT INTO `auth_menu_action` VALUES (39, 22, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (40, 22, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (41, 22, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (42, 22, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (43, 23, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (44, 25, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (45, 26, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (46, 28, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (47, 30, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (48, 30, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (49, 30, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (50, 31, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (51, 33, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (52, 35, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (53, 35, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (54, 35, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (55, 35, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (56, 11, 'notify', '回调');
INSERT INTO `auth_menu_action` VALUES (57, 11, 'addSupplement', '补单');
INSERT INTO `auth_menu_action` VALUES (58, 11, 'updateSupplement', '补单反馈');
INSERT INTO `auth_menu_action` VALUES (59, 28, 'addSupplement', '补单');
INSERT INTO `auth_menu_action` VALUES (60, 20, 'sendPayChannel', '转发通道');
INSERT INTO `auth_menu_action` VALUES (61, 15, 'config', '配置');
INSERT INTO `auth_menu_action` VALUES (62, 36, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (63, 36, 'send', '发送');

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
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order`  (
  `orderId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organizationId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构id',
  `payment` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付方式',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用的支付通道id',
  `payChannelOrderCode` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '通道返回的订单Code（即上游订单号）',
  `organizationOrderCode` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '机构订单Code（即下游订单号）',
  `orderCode` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '订单Code(唯一索引，平台订单号)',
  `myPayCode` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '我的支付Code(支付通道类型是MyAlipay和MyQqpay时有)',
  `orderPayStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单付款状态：0未付款，1已付款',
  `payTime` datetime NULL DEFAULT NULL COMMENT '付款时间',
  `orderName` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '订单名称',
  `orderPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '订单价格',
  `extendParams` json NULL COMMENT '扩展参数(回调时原样返回)',
  `notifyUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '异步回调地址',
  `redirectUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '同步回调地址',
  `organizationRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '机构费率',
  `payChannelRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '通道费率',
  `payChannelReturnData` json NULL COMMENT '通道返回数据',
  `returnData` json NULL COMMENT '返回给机构的数据',
  `payUserId` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '付款用户支付宝id',
  `supplementStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '补单状态：0不补单，1需补单，2已处理',
  `supplementImage` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '补单图片',
  `supplementMsg` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '补单反馈',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`orderId`, `addTime`) USING BTREE,
  INDEX `orderCode`(`orderCode` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE,
  INDEX `payChannelOrderCode`(`payChannelOrderCode` ASC) USING BTREE,
  INDEX `orderPayStatus`(`orderPayStatus` ASC) USING BTREE,
  INDEX `myPayCode`(`myPayCode` ASC) USING BTREE,
  INDEX `payment`(`payment` ASC) USING BTREE,
  INDEX `organizationId`(`organizationId` ASC) USING BTREE,
  INDEX `organizationOrderCode`(`organizationOrderCode` ASC) USING BTREE,
  INDEX `addTime`(`addTime` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC PARTITION BY RANGE (to_days(`addTime`))
PARTITIONS 140
(PARTITION `p20210421` VALUES LESS THAN (738267) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210422` VALUES LESS THAN (738268) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210423` VALUES LESS THAN (738269) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210424` VALUES LESS THAN (738270) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210425` VALUES LESS THAN (738271) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210426` VALUES LESS THAN (738272) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210427` VALUES LESS THAN (738273) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210428` VALUES LESS THAN (738274) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210506` VALUES LESS THAN (738282) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210507` VALUES LESS THAN (738283) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210508` VALUES LESS THAN (738284) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210509` VALUES LESS THAN (738285) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210510` VALUES LESS THAN (738286) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210511` VALUES LESS THAN (738287) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210512` VALUES LESS THAN (738288) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210608` VALUES LESS THAN (738315) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210609` VALUES LESS THAN (738316) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210610` VALUES LESS THAN (738317) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210611` VALUES LESS THAN (738318) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210612` VALUES LESS THAN (738319) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210613` VALUES LESS THAN (738320) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210614` VALUES LESS THAN (738321) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210623` VALUES LESS THAN (738330) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210624` VALUES LESS THAN (738331) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210625` VALUES LESS THAN (738332) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210626` VALUES LESS THAN (738333) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210627` VALUES LESS THAN (738334) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210628` VALUES LESS THAN (738335) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210629` VALUES LESS THAN (738336) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210706` VALUES LESS THAN (738343) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210707` VALUES LESS THAN (738344) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210708` VALUES LESS THAN (738345) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210709` VALUES LESS THAN (738346) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210710` VALUES LESS THAN (738347) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210711` VALUES LESS THAN (738348) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210712` VALUES LESS THAN (738349) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210716` VALUES LESS THAN (738353) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210717` VALUES LESS THAN (738354) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210718` VALUES LESS THAN (738355) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210719` VALUES LESS THAN (738356) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210720` VALUES LESS THAN (738357) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210721` VALUES LESS THAN (738358) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210722` VALUES LESS THAN (738359) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210723` VALUES LESS THAN (738360) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210724` VALUES LESS THAN (738361) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210725` VALUES LESS THAN (738362) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210726` VALUES LESS THAN (738363) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210727` VALUES LESS THAN (738364) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210728` VALUES LESS THAN (738365) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210729` VALUES LESS THAN (738366) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210730` VALUES LESS THAN (738367) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210731` VALUES LESS THAN (738368) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210801` VALUES LESS THAN (738369) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210802` VALUES LESS THAN (738370) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210804` VALUES LESS THAN (738372) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210805` VALUES LESS THAN (738373) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210806` VALUES LESS THAN (738374) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210807` VALUES LESS THAN (738375) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210808` VALUES LESS THAN (738376) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210809` VALUES LESS THAN (738377) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210810` VALUES LESS THAN (738378) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210811` VALUES LESS THAN (738379) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210812` VALUES LESS THAN (738380) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210813` VALUES LESS THAN (738381) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210814` VALUES LESS THAN (738382) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210815` VALUES LESS THAN (738383) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20210816` VALUES LESS THAN (738384) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211030` VALUES LESS THAN (738459) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211031` VALUES LESS THAN (738460) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211101` VALUES LESS THAN (738461) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211102` VALUES LESS THAN (738462) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211103` VALUES LESS THAN (738463) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211104` VALUES LESS THAN (738464) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211105` VALUES LESS THAN (738465) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211127` VALUES LESS THAN (738487) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211128` VALUES LESS THAN (738488) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211129` VALUES LESS THAN (738489) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211130` VALUES LESS THAN (738490) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211201` VALUES LESS THAN (738491) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211202` VALUES LESS THAN (738492) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211203` VALUES LESS THAN (738493) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211204` VALUES LESS THAN (738494) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211205` VALUES LESS THAN (738495) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211206` VALUES LESS THAN (738496) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211207` VALUES LESS THAN (738497) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211208` VALUES LESS THAN (738498) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211209` VALUES LESS THAN (738499) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211210` VALUES LESS THAN (738500) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211211` VALUES LESS THAN (738501) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211212` VALUES LESS THAN (738502) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211213` VALUES LESS THAN (738503) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211218` VALUES LESS THAN (738508) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211219` VALUES LESS THAN (738509) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211220` VALUES LESS THAN (738510) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211221` VALUES LESS THAN (738511) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211222` VALUES LESS THAN (738512) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211223` VALUES LESS THAN (738513) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20211224` VALUES LESS THAN (738514) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
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
PARTITION `p20220105` VALUES LESS THAN (738526) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220106` VALUES LESS THAN (738527) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220107` VALUES LESS THAN (738528) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220108` VALUES LESS THAN (738529) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220109` VALUES LESS THAN (738530) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220110` VALUES LESS THAN (738531) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220111` VALUES LESS THAN (738532) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
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
PARTITION `p20220127` VALUES LESS THAN (738548) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220128` VALUES LESS THAN (738549) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220129` VALUES LESS THAN (738550) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220130` VALUES LESS THAN (738551) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220131` VALUES LESS THAN (738552) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220201` VALUES LESS THAN (738553) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220202` VALUES LESS THAN (738554) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220203` VALUES LESS THAN (738555) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220210` VALUES LESS THAN (738562) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220211` VALUES LESS THAN (738563) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220212` VALUES LESS THAN (738564) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220213` VALUES LESS THAN (738565) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220214` VALUES LESS THAN (738566) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220215` VALUES LESS THAN (738567) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
PARTITION `p20220216` VALUES LESS THAN (738568) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 )
;

-- ----------------------------
-- Records of order
-- ----------------------------

-- ----------------------------
-- Table structure for order_notify
-- ----------------------------
DROP TABLE IF EXISTS `order_notify`;
CREATE TABLE `order_notify`  (
  `orderNotifyId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `orderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
  `isNotify` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否已通知：0否  1是',
  `notifyTime` datetime NULL DEFAULT NULL COMMENT '通知时间',
  `notifyNumber` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通知次数',
  `notifyUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '异步回调地址',
  `notifyData` json NULL COMMENT '发送数据（json格式保存）',
  `notifyResultData` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL COMMENT '通知返回结果',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`orderNotifyId`) USING BTREE,
  INDEX `orderId`(`orderId` ASC) USING BTREE,
  INDEX `notifyTime`(`notifyTime` ASC) USING BTREE,
  INDEX `isNotify`(`isNotify` ASC) USING BTREE,
  INDEX `notifyNumber`(`notifyNumber` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order_notify
-- ----------------------------

-- ----------------------------
-- Table structure for order_rel_salesman
-- ----------------------------
DROP TABLE IF EXISTS `order_rel_salesman`;
CREATE TABLE `order_rel_salesman`  (
  `orderRelSalesmanId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `orderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
  `salesmanId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '业务员id',
  `salesmanRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '业务员费率',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`orderRelSalesmanId`) USING BTREE,
  INDEX `orderId`(`orderId` ASC) USING BTREE,
  INDEX `salesmanId`(`salesmanId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of order_rel_salesman
-- ----------------------------

-- ----------------------------
-- Table structure for organization
-- ----------------------------
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization`  (
  `organizationId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organizationType` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '机构类型(与权限类型对应)：merchant商户，payService支付服务商',
  `payChannelGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付通道分组id',
  `appid` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '为兼容旧系统设置（有值为旧系统用户,不是旧系统用户不要填写该值,否则回调将出错）',
  `token` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT 'token(随机字符串的md5保存)',
  `organizationName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '机构名称',
  `organizationAdminNamePrefix` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '管理员前缀(该值创建机构管理员时作用)',
  `organizationWallet` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '机构钱包',
  `organizationRate` json NULL COMMENT '机构费率（与交易金额相乘结果为平台收益，各个支付方式费率不一样）',
  `phoneNumber` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `legalPerson` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '法人',
  `legalPersonIdentityCard` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '法人身份证',
  `organizationBankName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '机构开户银行名称',
  `organizationBankNumber` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '机构开户银行账号',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `googleAuthenticatorSecret` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '谷歌身份验证器密钥',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`organizationId`) USING BTREE,
  UNIQUE INDEX `organizationAdminNamePrefix`(`organizationAdminNamePrefix` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE,
  INDEX `payChannelGroupId`(`payChannelGroupId` ASC) USING BTREE,
  INDEX `appid`(`appid` ASC) USING BTREE,
  INDEX `organizationType`(`organizationType` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of organization
-- ----------------------------

-- ----------------------------
-- Table structure for organization_admin
-- ----------------------------
DROP TABLE IF EXISTS `organization_admin`;
CREATE TABLE `organization_admin`  (
  `organizationAdminId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organizationId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构id',
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id',
  `organizationAdminName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '机构管理员名称',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `allowIpList` json NULL COMMENT '允许登录的ip列表',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`organizationAdminId`) USING BTREE,
  UNIQUE INDEX `organizationAdminName`(`organizationAdminName` ASC) USING BTREE,
  INDEX `organizationId`(`organizationId` ASC) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of organization_admin
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel`  (
  `payChannelId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organizationId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构id（支付服务商类型机构可自建通道）',
  `payment` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付方式（客户请求生成订单接口时选择的支付方式）',
  `payChannelType` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通道类型：各个支付通道封装好且对外服务的类名',
  `payChannelName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通道名称（唯一索引，类型为MyAlipay时，为登录账号，密码保存在配置参数中）',
  `payChannelConfig` json NULL COMMENT '支付通道配置参数（json格式保存）',
  `payChannelRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '通道费率',
  `limitedAmount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '限制额度（0不限制；大于0时，值大于订单金额，订单才能使用该通道）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `serverIp` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '服务器IP',
  `webSocketId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT 'webSocket链接Id',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`payChannelId`) USING BTREE,
  UNIQUE INDEX `payChannelName`(`payChannelName` ASC) USING BTREE,
  INDEX `payChannelType`(`payChannelType` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE,
  INDEX `organizationId`(`organizationId` ASC) USING BTREE,
  INDEX `payment`(`payment` ASC) USING BTREE,
  INDEX `serverIp`(`serverIp` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel_group
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel_group`;
CREATE TABLE `pay_channel_group`  (
  `payChannelGroupId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `payChannelGroupName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通道分组名称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`payChannelGroupId`) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel_group
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel_group_rel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel_group_rel`;
CREATE TABLE `pay_channel_group_rel`  (
  `payChannelGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付通道分组id',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付通道id',
  `useNumber` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '使用次数（轮询时：与权重相乘的结果来决定）',
  `useWeight` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '使用权重（0-100，默认：50）',
  `useTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '使用时间',
  PRIMARY KEY (`payChannelGroupId`, `payChannelId`) USING BTREE,
  INDEX `payChannelGroupId`(`payChannelGroupId` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel_group_rel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel_notify
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel_notify`;
CREATE TABLE `pay_channel_notify`  (
  `payChannelNotifyId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `notifyData` json NULL COMMENT '通知数据',
  `returnData` json NULL COMMENT '返回数据',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`payChannelNotifyId`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel_notify
-- ----------------------------

-- ----------------------------
-- Table structure for salesman
-- ----------------------------
DROP TABLE IF EXISTS `salesman`;
CREATE TABLE `salesman`  (
  `salesmanId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id',
  `salesmanName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '业务员名称',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `salesmanWallet` decimal(16, 7) UNSIGNED NOT NULL DEFAULT 0.0000000 COMMENT '钱包金额',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `allowIpList` json NULL COMMENT '允许登录的ip列表',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`salesmanId`) USING BTREE,
  UNIQUE INDEX `salesmanName`(`salesmanName` ASC) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of salesman
-- ----------------------------

-- ----------------------------
-- Table structure for salesman_rel_organization
-- ----------------------------
DROP TABLE IF EXISTS `salesman_rel_organization`;
CREATE TABLE `salesman_rel_organization`  (
  `salesmanId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '业务员id',
  `organizationId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构id',
  `salesmanRate` json NULL COMMENT '费率（与交易金额相乘结果即为收益，各个支付方式费率不一样）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  INDEX `salesmanId`(`salesmanId` ASC) USING BTREE,
  INDEX `organizationId`(`organizationId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of salesman_rel_organization
-- ----------------------------

-- ----------------------------
-- Table structure for withdraw
-- ----------------------------
DROP TABLE IF EXISTS `withdraw`;
CREATE TABLE `withdraw`  (
  `withdrawId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `organizationId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `payeeName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收款人姓名',
  `bankName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '银行名称',
  `bankNumber` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '银行卡号',
  `withdrawAmount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '提款金额',
  `serviceFee` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '手续费',
  `withdrawStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '下发状态：0未下发，1已下发， 2下发中',
  `withdrawImageList` json NULL COMMENT '提款图片列表',
  `notifyUrl` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '通知回调地址',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道ID',
  `isPayChannelNotify` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '通道是否回调：0否  1是',
  `withdrawShowData` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '下发显示数据（商户看）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`withdrawId`) USING BTREE,
  INDEX `organizationId`(`organizationId` ASC) USING BTREE,
  INDEX `withdrawStatus`(`withdrawStatus` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE,
  INDEX `isPayChannelNotify`(`isPayChannelNotify` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of withdraw
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
