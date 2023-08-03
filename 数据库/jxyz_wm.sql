/*
 Navicat Premium Data Transfer

 Source Server         : 本地-Mysql8
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.0.200:3306
 Source Schema         : jxyz_wm

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 28/07/2023 20:35:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
INSERT INTO `admin` VALUES (1, 0, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 0, '2021-04-26 19:31:07', '2019-08-28 15:50:40');

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
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of auth_group
-- ----------------------------
INSERT INTO `auth_group` VALUES (1, 'platform', 0, '平台超级管理员', 0, '2021-04-26 14:43:23', '2021-04-26 14:43:23');
INSERT INTO `auth_group` VALUES (2, 'takeaway', 0, '外卖商家超级管理员', 0, '2021-04-28 15:26:20', '2021-04-28 15:26:20');

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
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 1, 1);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 6, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 7, 14);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 7, 15);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 8, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 9, 16);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 10, 17);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 10, 18);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 10, 19);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 10, 20);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 11, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 12, 21);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 12, 22);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 12, 23);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 12, 24);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 13, 25);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 13, 26);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 14, 27);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 14, 28);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 15, 29);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 15, 30);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 16, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 17, 31);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 17, 32);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 18, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 19, 33);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 19, 34);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 20, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 21, 35);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 21, 36);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 22, 37);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 23, 38);
INSERT INTO `auth_group_rel_menu_action` VALUES (1, 23, 39);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 24, 40);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 25, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 26, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 27, 41);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 27, 42);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 28, 43);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 28, 44);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 28, 45);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 28, 46);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 29, 47);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 29, 48);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 29, 49);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 29, 50);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 30, 51);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 30, 52);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 30, 53);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 30, 54);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 31, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 32, 55);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 32, 56);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 33, 0);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 34, 57);
INSERT INTO `auth_group_rel_menu_action` VALUES (2, 34, 58);

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
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `auth_menu` VALUES (10, 'platform', 6, '支付通道', '', '/payChannel/index', 0, 0, '2021-04-27 09:26:10', '2021-04-26 18:41:40');
INSERT INTO `auth_menu` VALUES (11, 'platform', 0, '外卖管理', 'el-icon-food', '/menuTakeaway', 50, 0, '2021-04-27 09:54:27', '2021-04-27 09:27:23');
INSERT INTO `auth_menu` VALUES (12, 'platform', 11, '店铺分类', '', '/tyStoreCategory/index', 50, 0, '2021-04-27 09:30:55', '2021-04-27 09:28:31');
INSERT INTO `auth_menu` VALUES (13, 'platform', 11, '店铺列表', '', '/tyStore/index', 50, 0, '2021-04-27 09:51:58', '2021-04-27 09:28:31');
INSERT INTO `auth_menu` VALUES (14, 'platform', 11, '商品列表', '', '/tyGoods/index', 50, 0, '2021-04-27 09:52:56', '2021-04-27 09:52:56');
INSERT INTO `auth_menu` VALUES (15, 'platform', 11, '外卖店铺管理员', '', '/tyAdmin/index', 100, 0, '2021-04-27 09:53:55', '2021-04-27 09:53:55');
INSERT INTO `auth_menu` VALUES (16, 'platform', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 20, 0, '2021-04-27 09:55:36', '2021-04-27 09:55:36');
INSERT INTO `auth_menu` VALUES (17, 'platform', 16, '外卖订单', '', '/tyOrder/index', 50, 0, '2021-04-27 09:56:50', '2021-04-27 09:56:50');
INSERT INTO `auth_menu` VALUES (18, 'platform', 0, '外包管理', 'el-icon-toilet-paper', '/menuOutsource', 50, 0, '2021-04-27 09:57:48', '2021-04-27 09:57:48');
INSERT INTO `auth_menu` VALUES (19, 'platform', 18, '外卖小哥', '', '/tyDeliveryWorker/index', 50, 0, '2021-04-27 09:58:44', '2021-04-27 09:58:44');
INSERT INTO `auth_menu` VALUES (20, 'platform', 0, '用户管理', 'el-icon-user-solid', '/menuUser', 80, 0, '2021-04-27 10:00:00', '2021-04-27 10:00:00');
INSERT INTO `auth_menu` VALUES (21, 'platform', 20, '用户列表', '', '/user/index', 50, 0, '2021-04-27 10:17:13', '2021-04-27 10:00:46');
INSERT INTO `auth_menu` VALUES (22, 'platform', 11, '评论列表', '', '/tyComment/index', 50, 0, '2021-04-27 10:01:37', '2021-04-27 10:01:37');
INSERT INTO `auth_menu` VALUES (23, 'platform', 16, '支付订单', '', '/payOrder/index', 50, 0, '2021-04-27 10:02:39', '2021-04-27 10:02:39');
INSERT INTO `auth_menu` VALUES (24, 'takeaway', 0, '主页', 'el-icon-s-home', '/index/index', 0, 0, '2021-04-27 10:07:01', '2021-04-27 10:07:01');
INSERT INTO `auth_menu` VALUES (25, 'takeaway', 0, '订单管理', 'el-icon-s-order', '/menuOrder', 20, 0, '2021-04-27 10:08:14', '2021-04-27 10:08:14');
INSERT INTO `auth_menu` VALUES (26, 'takeaway', 0, '商品管理', 'el-icon-food', '/menuGoods', 50, 0, '2021-04-27 10:08:58', '2021-04-27 10:08:58');
INSERT INTO `auth_menu` VALUES (27, 'takeaway', 25, '订单列表', '', '/tyOrder/index', 50, 0, '2021-04-27 10:10:22', '2021-04-27 10:10:22');
INSERT INTO `auth_menu` VALUES (28, 'takeaway', 26, '商品分类', '', '/tyGoodsCategory/index', 50, 0, '2021-04-27 10:11:08', '2021-04-27 10:11:08');
INSERT INTO `auth_menu` VALUES (29, 'takeaway', 26, '可选项列表', '', '/tyGoodsOption/index', 50, 0, '2021-04-27 10:11:59', '2021-04-27 10:11:59');
INSERT INTO `auth_menu` VALUES (30, 'takeaway', 26, '商品列表', '', '/tyGoods/index', 50, 0, '2021-04-27 10:12:28', '2021-04-27 10:12:28');
INSERT INTO `auth_menu` VALUES (31, 'takeaway', 0, '店铺管理', 'el-icon-setting', '/menuStore', 99, 0, '2021-04-27 10:13:31', '2021-04-27 10:13:31');
INSERT INTO `auth_menu` VALUES (32, 'takeaway', 31, '店铺设置', '', '/config/index', 50, 0, '2021-04-29 20:25:33', '2021-04-27 10:14:12');
INSERT INTO `auth_menu` VALUES (33, 'takeaway', 0, '评论管理', 'el-icon-s-comment', '/menuComment', 50, 0, '2021-04-27 10:14:53', '2021-04-27 10:14:53');
INSERT INTO `auth_menu` VALUES (34, 'takeaway', 33, '评论列表', '', '/tyComment/index', 50, 0, '2021-04-28 13:00:45', '2021-04-27 10:15:24');

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
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `auth_menu_action` VALUES (17, 10, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (18, 10, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (19, 10, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (20, 10, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (21, 12, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (22, 12, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (23, 12, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (24, 12, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (25, 13, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (26, 13, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (27, 14, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (28, 14, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (29, 15, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (30, 15, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (31, 17, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (32, 17, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (33, 19, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (34, 19, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (35, 21, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (36, 21, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (37, 22, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (38, 23, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (39, 23, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (40, 24, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (41, 27, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (42, 27, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (43, 28, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (44, 28, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (45, 28, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (46, 28, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (47, 29, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (48, 29, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (49, 29, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (50, 29, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (51, 30, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (52, 30, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (53, 30, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (54, 30, 'del', '删除');
INSERT INTO `auth_menu_action` VALUES (55, 32, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (56, 32, 'save', '保存');
INSERT INTO `auth_menu_action` VALUES (57, 34, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (58, 34, 'edit', '编辑');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of log_request
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS `pay_channel`;
CREATE TABLE `pay_channel`  (
  `payChannelId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付通道id',
  `payChannelType` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型：0付款，1打款',
  `payChannelName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '通道名称',
  `payChannelConfig` json NOT NULL COMMENT '配置参数',
  `payChannelRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '通道费率',
  `balance` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '通道余额',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`payChannelId`) USING BTREE,
  INDEX `payChannelType`(`payChannelType` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_channel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS `pay_order`;
CREATE TABLE `pay_order`  (
  `payOrderId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '支付订单id',
  `payOrderSn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付订单编号',
  `payChannelId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付通道id',
  `payOrderType` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '付款订单类型（takeaway外卖订单）',
  `orderIdList` json NOT NULL COMMENT '订单id列表（该订单id列表根据payOrderType不同指向不同表）',
  `amount` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '金额',
  `payChannelRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '通道费率',
  `payStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态：0未付款，1已付款',
  `payChannelSn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '支付通道交易流水号',
  `requestData` json NULL COMMENT '通道请求数据',
  `responseData` json NULL COMMENT '通道响应数据',
  `payTime` datetime NULL DEFAULT NULL COMMENT '付款时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`payOrderId`) USING BTREE,
  UNIQUE INDEX `payOrderSn`(`payOrderSn` ASC) USING BTREE,
  INDEX `payChannelId`(`payChannelId` ASC) USING BTREE,
  INDEX `payOrderType`(`payOrderType` ASC) USING BTREE,
  INDEX `payStatus`(`payStatus` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pay_order
-- ----------------------------

-- ----------------------------
-- Table structure for ty_admin
-- ----------------------------
DROP TABLE IF EXISTS `ty_admin`;
CREATE TABLE `ty_admin`  (
  `adminId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `authGroupId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限组id（平台超级管理员为0无权限限制，其他管理员必须设置权限组id）',
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '管理员名称',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`adminId`) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE,
  INDEX `authGroupId`(`authGroupId` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_admin
-- ----------------------------
INSERT INTO `ty_admin` VALUES (1, 1, 2, 'testTyAdmin', 'e10adc3949ba59abbe56e057f20f883e', '外卖商家超级管理员', 0, '2023-07-28 20:26:59', '2023-07-28 20:26:37');

-- ----------------------------
-- Table structure for ty_comment
-- ----------------------------
DROP TABLE IF EXISTS `ty_comment`;
CREATE TABLE `ty_comment`  (
  `commentId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `orderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id（冗余字段，用于优化查询）',
  `deliveryWorkerId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送人员id（冗余字段，用于优化查询）',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称（冗余字段，用于优化查询）',
  `goodsNameStr` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品名称（冗余字段，用于优化查询。格式：商品名称1,商品名称2(规格及选中项),... ）',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '评论内容',
  `image` json NULL COMMENT '评论图片',
  `deliveryRank` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送评价等级(范围0-5)',
  `goodsRank` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品评价等级(范围0-5)',
  `serviceRank` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商家服务态度评价等级(范围0-5)',
  `isAnonymous` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '匿名评价：0否，1是',
  `storeReplyStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商家回复：0否，1是',
  `deliveryWorkerReplyStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送员回复：0否，1是',
  `userReplyStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户追评：0否，1是',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`commentId`) USING BTREE,
  INDEX `orderId`(`orderId` ASC) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `deliveryWorkerId`(`deliveryWorkerId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_comment
-- ----------------------------

-- ----------------------------
-- Table structure for ty_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `ty_comment_reply`;
CREATE TABLE `ty_comment_reply`  (
  `replyId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '回复id',
  `commentId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论id（指向评论表，用于取出评论下面的所有回复）',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级id（0表示回复评论表，大于0表示回复上级回复）',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `deliveryWorkerId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送人员id',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id（用户id大于0时，当pid为0时表示为用户追加评论，pid大于0表示为用户回复其他评论）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称（冗余字段，用于优化查询）',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '评论内容',
  `image` json NULL COMMENT '评论图片',
  `isAnonymous` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否匿名评价:0否，1是',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`replyId`) USING BTREE,
  INDEX `commentId`(`commentId` ASC) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE,
  INDEX `deliveryWorkerId`(`deliveryWorkerId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_comment_reply
-- ----------------------------

-- ----------------------------
-- Table structure for ty_delivery_worker
-- ----------------------------
DROP TABLE IF EXISTS `ty_delivery_worker`;
CREATE TABLE `ty_delivery_worker`  (
  `deliveryWorkerId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '配送人员id',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`deliveryWorkerId`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_delivery_worker
-- ----------------------------

-- ----------------------------
-- Table structure for ty_goods
-- ----------------------------
DROP TABLE IF EXISTS `ty_goods`;
CREATE TABLE `ty_goods`  (
  `goodsId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `categoryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品分类id',
  `goodsSn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品编号',
  `option` json NULL COMMENT '可选项（由所有不可以影响价格和库存的optionName和optionValue组成。json格式：{\"选项名1\":[\"选项值1\",...],...}）',
  `goodsName` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品名称',
  `sellPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '售卖价',
  `marketPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '市场价',
  `costPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '成本价',
  `packPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '打包费',
  `stockNum` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存数',
  `clickNum` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '点击量',
  `saleNum` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '销量',
  `mainImage` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '主图',
  `contentImage` json NULL COMMENT '详情图片（json格式：[\"图片1\",...]）',
  `attribute` json NULL COMMENT '商品属性（json格式：{\"属性名1\":\"属性值\",...}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`goodsId`) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE,
  INDEX `categoryId`(`categoryId` ASC) USING BTREE,
  INDEX `goodsSn`(`goodsSn` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_goods
-- ----------------------------

-- ----------------------------
-- Table structure for ty_goods_category
-- ----------------------------
DROP TABLE IF EXISTS `ty_goods_category`;
CREATE TABLE `ty_goods_category`  (
  `categoryId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品分类id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `categoryName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `categoryIcon` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`categoryId`) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_goods_category
-- ----------------------------

-- ----------------------------
-- Table structure for ty_goods_option
-- ----------------------------
DROP TABLE IF EXISTS `ty_goods_option`;
CREATE TABLE `ty_goods_option`  (
  `optionId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '选项id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `optionName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '选项名',
  `optionValueList` json NOT NULL COMMENT '选项值列表（json格式：[\"值1\",...]）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`optionId`) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_goods_option
-- ----------------------------

-- ----------------------------
-- Table structure for ty_goods_spec
-- ----------------------------
DROP TABLE IF EXISTS `ty_goods_spec`;
CREATE TABLE `ty_goods_spec`  (
  `specId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '规格id',
  `goodsId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
  `specName` json NOT NULL COMMENT '规格名称（由所有可以影响价格和库存的optionName和optionValue组成，内部所有字段按ascii码排序。json格式：{\"选项名1\":\"选项值\",...}）',
  `sellPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '售卖价',
  `marketPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '市场价',
  `costPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '成本价',
  `packPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '打包费',
  `stockNum` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存数',
  `mainImage` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '主图',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`specId`) USING BTREE,
  INDEX `goodsId`(`goodsId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_goods_spec
-- ----------------------------

-- ----------------------------
-- Table structure for ty_order
-- ----------------------------
DROP TABLE IF EXISTS `ty_order`;
CREATE TABLE `ty_order`  (
  `orderId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `orderSn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '订单编号',
  `payOrderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付订单id',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `deliveryWorkerId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送人员id',
  `isStoreDelivery` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '商家配送：0否，1是',
  `totalGoodsFinalPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '合计最终价（该订单下所有商品价格总和）',
  `totalGoodsPackPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '合计打包费（该订单下所有商品打包费总和）',
  `deliveryPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '配送费',
  `totalPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '订单总价（最终价格+打包费+配送费）',
  `orderPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '实际付款金额（订单总价-各种优惠）',
  `storeRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '店铺费率',
  `orderStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单状态：0未完成，1已完成',
  `payStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态：0未付款，1已付款',
  `acceptStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '接单状态：0未接单，1店铺接单，2配送员接单，3店铺拒单',
  `deliveryStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流状态：0未发货，1已收货，2已发货',
  `commentStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '评价状态：0未评价，1已评价',
  `consignee` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收货人',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `longitude` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '经度',
  `latitude` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '纬度',
  `userRemark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户备注',
  `payTime` datetime NULL DEFAULT NULL COMMENT '付款时间',
  `acceptTime` datetime NULL DEFAULT NULL COMMENT '接单时间',
  `sendTime` datetime NULL DEFAULT NULL COMMENT '发货时间',
  `receiveTime` datetime NULL DEFAULT NULL COMMENT '收货时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`orderId`) USING BTREE,
  UNIQUE INDEX `orderSn`(`orderSn` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `deliveryWorkerId`(`deliveryWorkerId` ASC) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE,
  INDEX `payOrderId`(`payOrderId` ASC) USING BTREE,
  INDEX `isStoreDelivery`(`isStoreDelivery` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_order
-- ----------------------------

-- ----------------------------
-- Table structure for ty_order_action
-- ----------------------------
DROP TABLE IF EXISTS `ty_order_action`;
CREATE TABLE `ty_order_action`  (
  `orderActionId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单操作id',
  `orderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
  `adminId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '总后台管理员id',
  `tyAdminId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '外卖管理员id',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `deliveryWorkerId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '配送人员id',
  `newOrderStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单状态(改动后)',
  `newDeliveryStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流状态(改动后)',
  `newAcceptStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '接单状态(改动后)',
  `newPayStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态(改动后)',
  `oldOrderStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单状态(改动前)',
  `oldDeliveryStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '物流状态(改动前)',
  `oldAcceptStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '接单状态(改动前)',
  `oldPayStatus` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '支付状态(改动前)',
  `remark` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`orderActionId`) USING BTREE,
  INDEX `adminId`(`adminId` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `deliveryWorkerId`(`deliveryWorkerId` ASC) USING BTREE,
  INDEX `orderId`(`orderId` ASC) USING BTREE,
  INDEX `tyAdminId`(`tyAdminId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_order_action
-- ----------------------------

-- ----------------------------
-- Table structure for ty_order_goods
-- ----------------------------
DROP TABLE IF EXISTS `ty_order_goods`;
CREATE TABLE `ty_order_goods`  (
  `orderGoodsId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '订单商品id',
  `orderId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '订单id',
  `goodsId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品id',
  `goodsName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品名称',
  `specId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品规格id',
  `specName` json NULL COMMENT '规格名称（json格式：{\"选项名1\":\"可选值\",...}）',
  `optionSelected` json NULL COMMENT '选中项（json格式：{\"选项名1\":\"可选值\",...}）',
  `mainImage` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '主图',
  `goodsSn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '商品编号',
  `finalPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '最终价（单价。打折等因素影响后的价格）',
  `sellPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '售卖价（单价）',
  `marketPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '市场价（单价）',
  `costPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '成本价（单价）',
  `packPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '打包费（单价）',
  `goodsNum` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品数量',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`orderGoodsId`) USING BTREE,
  INDEX `orderId`(`orderId` ASC) USING BTREE,
  INDEX `goodsId`(`goodsId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_order_goods
-- ----------------------------

-- ----------------------------
-- Table structure for ty_store
-- ----------------------------
DROP TABLE IF EXISTS `ty_store`;
CREATE TABLE `ty_store`  (
  `storeId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺id',
  `storeName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '店铺名称',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `storeRate` decimal(5, 4) UNSIGNED NOT NULL DEFAULT 0.0000 COMMENT '店铺费率',
  `balance` decimal(15, 6) UNSIGNED NOT NULL DEFAULT 0.000000 COMMENT '店铺余额',
  `payeeName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收款人姓名',
  `payeeMobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收款人手机',
  `payeeAccountType` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收款账户类型：待定（比如bc中国银行，cbc建设银行，zfb支付宝）',
  `payeeAccount` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收款账户',
  `longitude` decimal(10, 7) NOT NULL DEFAULT 0.0000000 COMMENT '经度',
  `latitude` decimal(10, 7) NOT NULL DEFAULT 0.0000000 COMMENT '纬度',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`storeId`) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE,
  INDEX `longitude`(`longitude` ASC) USING BTREE,
  INDEX `latitude`(`latitude` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_store
-- ----------------------------
INSERT INTO `ty_store` VALUES (1, '测试店铺', '', 0.0000, 0.000000, '', '', '', '', 118.5700000, 24.8300000, 0, '2023-07-28 20:32:47', '2023-07-28 20:24:27');

-- ----------------------------
-- Table structure for ty_store_category
-- ----------------------------
DROP TABLE IF EXISTS `ty_store_category`;
CREATE TABLE `ty_store_category`  (
  `categoryId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '店铺分类id',
  `pid` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级id（只支持2级，够用）',
  `categoryName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `categoryIcon` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '分类图标',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`categoryId`) USING BTREE,
  INDEX `pid`(`pid` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_store_category
-- ----------------------------

-- ----------------------------
-- Table structure for ty_store_config
-- ----------------------------
DROP TABLE IF EXISTS `ty_store_config`;
CREATE TABLE `ty_store_config`  (
  `configId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  `configKey` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '配置项Key',
  `configValue` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '配置项值',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`configId`) USING BTREE,
  UNIQUE INDEX `storeId`(`storeId` ASC, `configKey` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_store_config
-- ----------------------------

-- ----------------------------
-- Table structure for ty_store_rel_category
-- ----------------------------
DROP TABLE IF EXISTS `ty_store_rel_category`;
CREATE TABLE `ty_store_rel_category`  (
  `categoryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺分类id',
  `storeId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '店铺id',
  PRIMARY KEY (`categoryId`, `storeId`) USING BTREE,
  INDEX `categoryId`(`categoryId` ASC) USING BTREE,
  INDEX `storeId`(`storeId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ty_store_rel_category
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `userName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE INDEX `userName`(`userName` ASC) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for user_address
-- ----------------------------
DROP TABLE IF EXISTS `user_address`;
CREATE TABLE `user_address`  (
  `addressId` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '地址id',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `consignee` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '收货人',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `address` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地址',
  `addressMap` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地图定位地址（前端需要）',
  `addressMapS` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '地图定位地址简称（前端需要）',
  `longitude` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '经度',
  `latitude` decimal(10, 6) NOT NULL DEFAULT 0.000000 COMMENT '纬度',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`addressId`) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_address
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
