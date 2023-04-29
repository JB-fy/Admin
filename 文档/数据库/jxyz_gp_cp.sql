/*
 Navicat Premium Data Transfer

 Source Server         : 本地-8.0.32
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : 192.168.0.16:3306
 Source Schema         : jxyz_gp_cp

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 29/04/2023 09:56:36
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
INSERT INTO `admin` VALUES (1, 0, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 0, '2021-03-04 15:19:32', '2019-08-28 15:50:40');

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
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `auth_menu` VALUES (10, 'platform', 0, '用户管理', 'el-icon-user-solid', '/menuUser', 90, 0, '2021-07-28 15:33:40', '2021-07-28 15:28:25');
INSERT INTO `auth_menu` VALUES (11, 'platform', 10, '用户列表', '', '/user/index', 50, 0, '2021-07-28 15:31:26', '2021-07-28 15:31:26');
INSERT INTO `auth_menu` VALUES (12, 'platform', 0, '彩票管理', 'el-icon-grape', '/menuLottery', 50, 0, '2021-10-28 09:33:40', '2021-07-28 15:32:25');
INSERT INTO `auth_menu` VALUES (13, 'platform', 12, '订单列表', '', '/lotteryOrder/index', 50, 0, '2022-01-17 09:56:43', '2021-07-28 15:36:53');
INSERT INTO `auth_menu` VALUES (14, 'platform', 12, '彩票列表', '', '/lottery/index', 50, 0, '2022-01-17 17:11:20', '2022-01-17 09:54:34');
INSERT INTO `auth_menu` VALUES (15, 'platform', 12, '期号列表', '', '/lotteryIssue/index', 50, 0, '2022-01-17 17:11:04', '2021-07-28 15:38:36');
INSERT INTO `auth_menu` VALUES (16, 'platform', 12, '玩法列表', '', '/lotteryPlay/index', 50, 0, '2022-01-17 09:56:55', '2021-07-30 17:24:24');

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
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

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
INSERT INTO `auth_menu_action` VALUES (18, 13, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (19, 15, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (20, 11, 'audit', '审核');
INSERT INTO `auth_menu_action` VALUES (21, 16, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (22, 16, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (23, 16, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (25, 15, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (26, 15, 'edit', '编辑');
INSERT INTO `auth_menu_action` VALUES (27, 15, 'settle', '结算');
INSERT INTO `auth_menu_action` VALUES (28, 14, 'sel', '查看');
INSERT INTO `auth_menu_action` VALUES (29, 14, 'add', '新增');
INSERT INTO `auth_menu_action` VALUES (30, 14, 'edit', '编辑');

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
-- Table structure for lottery
-- ----------------------------
DROP TABLE IF EXISTS `lottery`;
CREATE TABLE `lottery`  (
  `lotteryId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `lotteryCategory` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '彩票分类',
  `lotteryName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '彩票名称',
  `lotteryDes` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '彩票说明',
  `isManual` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否手动开奖：0否 1是',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`lotteryId`) USING BTREE,
  INDEX `lotteryCategory`(`lotteryCategory` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of lottery
-- ----------------------------
INSERT INTO `lottery` VALUES (1, 'lhc', '极速六合彩', '本站自有彩种，全天24小时，1分钟1期，每期从01-49中随机选出7个号码，前6个为正码，最后1个为特别码', 0, 50, 0, '2022-02-09 15:58:01', '2022-01-25 17:28:48');
INSERT INTO `lottery` VALUES (2, 'ssc', '本站1分彩', '本站自有彩种，全天24小时，1分钟1期，每期5个号码，每个号码都从0-9中随机生成', 0, 50, 0, '2022-02-09 15:55:35', '2022-01-17 10:15:06');
INSERT INTO `lottery` VALUES (3, 'ssc', '本站5分彩', '本站自有彩种，全天24小时，5分钟1期，每期5个号码，每个号码都从0-9中随机生成', 0, 50, 0, '2022-02-09 15:55:31', '2022-01-17 10:25:02');
INSERT INTO `lottery` VALUES (4, 'ssc', '本站10分彩', '本站自有彩种，全天24小时，10分钟1期，每期5个号码，每个号码都从0-9中随机生成', 0, 50, 0, '2022-02-09 15:55:29', '2022-01-17 10:25:32');
INSERT INTO `lottery` VALUES (5, 'lhc', '香港六合彩', '与香港六合彩官方保持一致，一般于每周二，四，六的21:35开奖，每期从01-49中随机选出7个号码，前6个为正码，最后1个为特别码', 0, 50, 0, '2022-02-09 15:55:27', '2022-01-17 10:25:46');
INSERT INTO `lottery` VALUES (6, 'lhc', '澳门六合彩', '与澳门六合彩官方保持一致，每天21:32开奖，每期从01-49中随机选出7个号码，前6个为正码，最后1个为特别码', 0, 50, 0, '2022-02-09 15:55:25', '2022-01-20 13:02:21');
INSERT INTO `lottery` VALUES (7, 'ssc', '天津时时彩', '与天津时时彩官方保持一致，每天9:00到23:00，20分钟1期，每期5个号码，每个号码都从0-9中随机生成', 0, 50, 0, '2022-02-09 15:55:22', '2022-01-20 13:02:38');
INSERT INTO `lottery` VALUES (8, 'ssc', '新疆时时彩', '与新疆时时彩官方保持一致，每天10:00到次日2:00，20分钟1期，每期5个号码，每个号码都从0-9中随机生成', 0, 50, 0, '2022-02-09 15:55:20', '2022-01-20 13:02:53');

-- ----------------------------
-- Table structure for lottery_issue
-- ----------------------------
DROP TABLE IF EXISTS `lottery_issue`;
CREATE TABLE `lottery_issue`  (
  `issueId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `lotteryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '彩票id',
  `lotterySn` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '彩票期号（表示哪一期）',
  `result` json NULL COMMENT '开奖结果',
  `startTime` datetime NOT NULL COMMENT '开始时间',
  `endTime` datetime NOT NULL COMMENT '结束时间',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`issueId`) USING BTREE,
  UNIQUE INDEX `lotteryId`(`lotteryId` ASC, `lotterySn` ASC) USING BTREE,
  INDEX `lotteryId_2`(`lotteryId` ASC) USING BTREE,
  INDEX `lotterySn`(`lotterySn` ASC) USING BTREE,
  INDEX `startTime`(`startTime` ASC) USING BTREE,
  INDEX `endTime`(`endTime` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of lottery_issue
-- ----------------------------

-- ----------------------------
-- Table structure for lottery_order
-- ----------------------------
DROP TABLE IF EXISTS `lottery_order`;
CREATE TABLE `lottery_order`  (
  `orderId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `lotteryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '彩票id（冗余字段，用来优化查询）',
  `userId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `issueId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '期号id',
  `playId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '玩法id',
  `orderCode` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '单号',
  `count` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '注数',
  `onePrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '单注金额',
  `orderPrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '订单金额（单注金额*注数）',
  `settlePrice` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '结算金额（由onePrice，winCount及oddsConfig计算得来）',
  `orderResult` json NULL COMMENT '下注结果',
  `oddsConfig` json NULL COMMENT '下注赔率（格式：{\"中奖等级1\":\"赔率\",\"中奖等级2\":\"赔率\",...},示例：{\"赢\":\"5.00\",\"平局\":\"1.00\"}）',
  `winCount` json NULL COMMENT '中奖注数（格式：{\"中奖等级1\":\"中奖注数\",\"中奖等级2\":\"中奖注数\",...},示例：{\"赢\":\"3\",\"平局\":\"5\"}）',
  `isSettle` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否结算：0否，1是',
  `thirdCode` varchar(60) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '三方单号',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`orderId`) USING BTREE,
  UNIQUE INDEX `orderCode`(`orderCode` ASC) USING BTREE,
  INDEX `userId`(`userId` ASC) USING BTREE,
  INDEX `lotteryId`(`lotteryId` ASC) USING BTREE,
  INDEX `playId`(`playId` ASC) USING BTREE,
  INDEX `isSettle`(`isSettle` ASC) USING BTREE,
  INDEX `thirdCode`(`thirdCode` ASC) USING BTREE,
  INDEX `issueId`(`issueId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of lottery_order
-- ----------------------------

-- ----------------------------
-- Table structure for lottery_play
-- ----------------------------
DROP TABLE IF EXISTS `lottery_play`;
CREATE TABLE `lottery_play`  (
  `playId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `lotteryCategory` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '彩票分类',
  `playName` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '玩法名称',
  `playConfig` json NULL COMMENT '玩法配置参数（格式：{\"className\":\"后端该玩法对应的类文件名\",\"remark\":\"玩法说明（前端显示用）\"},示例：{\"className\":\"zx5bz\",\"remark\":\"玩法说明（前端显示用）\"}）',
  `playIns` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '玩法说明',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`playId`) USING BTREE,
  INDEX `lotteryCategory`(`lotteryCategory` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 58 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of lottery_play
-- ----------------------------
INSERT INTO `lottery_play` VALUES (1, 'lhc', '特别码', '{\"className\": \"Lhc\\\\Tbm\"}', '所选号码与开奖号码的特别码一致，即为中奖。（1个号码1注，超过则为复式票）', 0, '2022-01-25 09:59:09', '2021-09-08 10:57:54');
INSERT INTO `lottery_play` VALUES (2, 'lhc', '色波-特', '{\"className\": \"Lhc\\\\TSb\"}', '所选颜色与开奖号码中特别码颜色一致，即为中奖。（1个颜色1注，超过则为复式票）', 0, '2022-01-25 09:59:48', '2021-09-11 16:51:07');
INSERT INTO `lottery_play` VALUES (3, 'lhc', '生肖-特', '{\"className\": \"Lhc\\\\TSx\"}', '所选生肖与开奖号码中特别码生肖一致，即为中奖。（1个生肖1注，超过则为复式票）', 0, '2022-01-25 09:59:59', '2021-09-14 14:39:38');
INSERT INTO `lottery_play` VALUES (4, 'lhc', '五行-特', '{\"className\": \"Lhc\\\\TWx\"}', '所选五行与开奖号码中特别码五行一致，即为中奖。（1个五行1注，超过则为复式票）', 0, '2022-01-25 10:00:10', '2021-09-14 14:41:41');
INSERT INTO `lottery_play` VALUES (5, 'lhc', '头尾-特', '{\"className\": \"Lhc\\\\TTw\"}', '所选头尾与开奖号码中特别码头尾一致，即为中奖。（1个头尾1注，超过则为复式票）', 0, '2022-01-25 10:00:23', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (6, 'lhc', '色波-合', '{\"className\": \"Lhc\\\\HSb\"}', '开奖号码中6个正码各以1个颜色计算，特别码以1.5个颜色计算。所选颜色占比最大，且是唯一最大，即为中奖。如果同时有两个颜色占比一样且都最大，占比都是3，则为中奖双色。（1个选项1注，超过则为复式票）', 0, '2022-01-25 10:00:46', '2021-09-08 10:03:43');
INSERT INTO `lottery_play` VALUES (7, 'lhc', '生肖-合', '{\"className\": \"Lhc\\\\HSx\"}', '所选生肖总数与开奖号码中的不同生肖总数一致，即为中奖。如果生肖总数为单数，则为中奖肖单；为双数，则为中奖肖双。（1个选项1注，超过则为复式票）', 0, '2022-01-25 10:00:58', '2021-09-08 10:03:43');
INSERT INTO `lottery_play` VALUES (8, 'lhc', '2中-全', '{\"className\": \"Lhc\\\\Q2z\"}', '所选号码包含在开奖号码中，即为中奖。如所选号码包含特别码，即中一等奖，不包含特别码，即中二等奖。（2个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:01:10', '2021-11-04 09:40:51');
INSERT INTO `lottery_play` VALUES (9, 'lhc', '一肖-全', '{\"className\": \"Lhc\\\\Q1x\"}', '所选生肖包含在开奖号码生肖中，即为中奖。（1个生肖1注，超过则为复式票）', 0, '2022-01-25 10:01:21', '2021-09-14 14:39:38');
INSERT INTO `lottery_play` VALUES (10, 'lhc', '二肖-全', '{\"className\": \"Lhc\\\\Q2x\"}', '所选生肖包含在开奖号码生肖中，即为中奖。（2个生肖组成1注，超过则为复式票）', 0, '2022-01-25 10:01:31', '2021-11-04 09:40:51');
INSERT INTO `lottery_play` VALUES (11, 'lhc', '三肖-全', '{\"className\": \"Lhc\\\\Q3x\"}', '所选生肖包含在开奖号码生肖中，即为中奖。（3个生肖组成1注，超过则为复式票）', 0, '2022-01-25 10:01:44', '2021-11-04 09:40:51');
INSERT INTO `lottery_play` VALUES (12, 'lhc', '四肖-全', '{\"className\": \"Lhc\\\\Q4x\"}', '所选生肖包含在开奖号码生肖中，即为中奖。（4个生肖组成1注，超过则为复式票）', 0, '2022-01-25 10:01:54', '2021-11-04 09:40:51');
INSERT INTO `lottery_play` VALUES (13, 'lhc', '五肖-全', '{\"className\": \"Lhc\\\\Q5x\"}', '所选生肖包含在开奖号码生肖中，即为中奖。（5个生肖组成1注，超过则为复式票）', 0, '2022-01-25 10:02:03', '2021-11-04 09:40:51');
INSERT INTO `lottery_play` VALUES (14, 'lhc', '一尾-全', '{\"className\": \"Lhc\\\\Q1w\"}', '所选尾包含在开奖号码尾中，即为中奖。（1个尾1注，超过则为复式票）', 0, '2022-01-25 10:02:12', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (15, 'lhc', '二尾-全', '{\"className\": \"Lhc\\\\Q2w\"}', '所选尾包含在开奖号码尾中，即为中奖。（2个尾1注，超过则为复式票）', 0, '2022-01-25 10:02:20', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (16, 'lhc', '三尾-全', '{\"className\": \"Lhc\\\\Q3w\"}', '所选尾包含在开奖号码尾中，即为中奖。（3个尾1注，超过则为复式票）', 0, '2022-01-25 10:02:28', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (17, 'lhc', '四尾-全', '{\"className\": \"Lhc\\\\Q4w\"}', '所选尾包含在开奖号码尾中，即为中奖。（4个尾1注，超过则为复式票）', 0, '2022-01-25 10:02:36', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (18, 'lhc', '五尾-全', '{\"className\": \"Lhc\\\\Q5w\"}', '所选尾包含在开奖号码尾中，即为中奖。（5个尾1注，超过则为复式票）', 0, '2022-01-25 10:02:46', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (19, 'lhc', '1中-正', '{\"className\": \"Lhc\\\\Z1z\"}', '所选号码包含在开奖号码6个正码中，即为中奖。（1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:01', '2021-10-28 11:10:20');
INSERT INTO `lottery_play` VALUES (20, 'lhc', '2中-正', '{\"className\": \"Lhc\\\\Z2z\"}', '所选号码包含在开奖号码6个正码中，即为中奖。（2个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:09', '2021-11-02 11:14:53');
INSERT INTO `lottery_play` VALUES (21, 'lhc', '3中-正', '{\"className\": \"Lhc\\\\Z3z\"}', '所选号码包含在开奖号码6个正码中，即为中奖。（3个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:18', '2021-11-02 11:16:35');
INSERT INTO `lottery_play` VALUES (22, 'lhc', '4中-正', '{\"className\": \"Lhc\\\\Z4z\"}', '所选号码包含在开奖号码6个正码中，即为中奖。（1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:25', '2021-11-02 11:17:42');
INSERT INTO `lottery_play` VALUES (23, 'lhc', '5不中-全', '{\"className\": \"Lhc\\\\Q5bz\"}', '所选号码不包含在开奖号码中，即为中奖。（5个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:32', '2021-07-30 17:56:55');
INSERT INTO `lottery_play` VALUES (24, 'lhc', '6不中-全', '{\"className\": \"Lhc\\\\Q6bz\"}', '所选号码不包含在开奖号码中，即为中奖。（6个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:40', '2021-08-21 17:16:27');
INSERT INTO `lottery_play` VALUES (25, 'lhc', '7不中-全', '{\"className\": \"Lhc\\\\Q7bz\"}', '所选号码不包含在开奖号码中，即为中奖。（7个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:49', '2021-08-21 18:22:40');
INSERT INTO `lottery_play` VALUES (26, 'lhc', '8不中-全', '{\"className\": \"Lhc\\\\Q8bz\"}', '所选号码不包含在开奖号码中，即为中奖。（8个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:03:57', '2021-09-06 10:17:24');
INSERT INTO `lottery_play` VALUES (27, 'lhc', '9不中-全', '{\"className\": \"Lhc\\\\Q9bz\"}', '所选号码不包含在开奖号码中，即为中奖。（9个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:04:05', '2021-09-06 10:18:15');
INSERT INTO `lottery_play` VALUES (28, 'lhc', '10不中-全', '{\"className\": \"Lhc\\\\Q10bz\"}', '所选号码不包含在开奖号码中，即为中奖。（10个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:04:13', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (29, 'lhc', '11不中-全', '{\"className\": \"Lhc\\\\Q11bz\"}', '所选号码不包含在开奖号码中，即为中奖。（11个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:04:19', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (30, 'lhc', '12不中-全', '{\"className\": \"Lhc\\\\Q12bz\"}', '所选号码不包含在开奖号码中，即为中奖。（12个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:04:26', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (31, 'ssc', '比大小', '{\"className\": \"Ssc\\\\Bdx\"}', '比较结果与开奖号码中所选2个位置比较结果一致，即为中奖。（先两边各选1个位置再选比较结果组成1注，超过则为复式票）', 0, '2022-01-25 10:04:34', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (32, 'ssc', '诈金花', '{\"className\": \"Ssc\\\\Zjh\"}', '所选结果与开奖号码中所选3个位置结果一致，即为中奖。3个位置只会确定一种结果，如确定为豹子，则对子不算中奖，优先级：豹子>顺子>对子>半顺>其他。豹子：3个位置号码相同；顺子：3个位置号码相连，顺序不限；对子：3个位置任意2个位置号码相同；半顺：3个位置任意2个位置号码相连，顺序不限；其他：3个位置号码不能确定为豹子，顺子，对子，半顺的情况；（先选3个位置再选1个结果组成1注，超过则为复式票）', 0, '2022-01-25 10:04:46', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (33, 'ssc', '斗牛', '{\"className\": \"Ssc\\\\Dn\"}', '所选斗牛与开奖号码斗牛一致，即为中奖。开奖号码中任意三个号码之和（0当作10计算）不是10的倍数则无牛，是则有牛，且用其余两个位置号码之和的个位数来决定是牛几，若其余两个位置号码之和也是10的倍数则是牛牛。（1个斗牛1注，超过则为复式票）', 0, '2022-01-25 10:04:54', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (34, 'ssc', '1直中-定位', '{\"className\": \"Ssc\\\\Dw1Zz\"}', '所选号码与开奖号码中所在位置号码相同，且位置一致，即为中奖。（在任意1个位置选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:02', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (35, 'ssc', '2直中-定位', '{\"className\": \"Ssc\\\\Dw2Zz\"}', '所选号码与开奖号码中所在位置号码相同，且位置一致，即为中奖。（在任意2个位置各选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:10', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (36, 'ssc', '3直中-定位', '{\"className\": \"Ssc\\\\Dw3Zz\"}', '所选号码与开奖号码中所在位置号码相同，且位置一致，即为中奖。（在任意3个位置各选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:18', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (37, 'ssc', '4直中-定位', '{\"className\": \"Ssc\\\\Dw4Zz\"}', '所选号码与开奖号码中所在位置号码相同，且位置一致，即为中奖。（在任意4个位置各选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:26', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (38, 'ssc', '5直中-定位', '{\"className\": \"Ssc\\\\Dw5Zz\"}', '所选号码与开奖号码中所在位置号码相同，且位置一致，即为中奖。（在5个位置各选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:38', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (39, 'ssc', '5选中-定位', '{\"className\": \"Ssc\\\\Dw5Xz\"}', '所选号码与开奖号码全部相同，且位置一致，即为一等奖；与开奖号码中前3位或后3位相同，且位置一致，即中二等奖；与开奖号码前2位和后2位相同，且位置一致，即中三等奖；与开奖号码前2位或后2位相同，且位置一致，即中四等奖。（在5个位置各选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:05:45', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (40, 'ssc', '2任中-不定位', '{\"className\": \"Ssc\\\\Bdw2Rz\"}', '所选号码与开奖号码中所选位置号码相同，顺序不限，即为中奖。（先选2个位置再选2个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:05:52', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (41, 'ssc', '3任中-不定位', '{\"className\": \"Ssc\\\\Bdw3Rz\"}', '所选号码与开奖号码中所选位置号码相同，顺序不限，即为中奖。（先选3个位置再选3个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:06:00', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (42, 'ssc', '4任中-不定位', '{\"className\": \"Ssc\\\\Bdw4Rz\"}', '所选号码与开奖号码中所选位置号码相同，顺序不限，即为中奖。（先选4个位置再选4个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:06:07', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (43, 'ssc', '5任中-不定位', '{\"className\": \"Ssc\\\\Bdw5Rz\"}', '所选号码与开奖号码相同，顺序不限，即为中奖。（选5个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:06:15', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (44, 'ssc', '1码3中-不定位', '{\"className\": \"Ssc\\\\Bdw1m3z\"}', '所选号码包含在开奖号码所选位置中，即为中奖。（先选3个位置再选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:06:22', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (45, 'ssc', '1码4中-不定位', '{\"className\": \"Ssc\\\\Bdw1m4z\"}', '所选号码包含在开奖号码所选位置中，即为中奖。（先选4个位置再选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:06:31', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (46, 'ssc', '1码5中-不定位', '{\"className\": \"Ssc\\\\Bdw1m5z\"}', '所选号码包含在开奖号码中，即为中奖。（选1个号码组成1注，超过则为复式票）', 0, '2022-01-25 10:06:38', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (47, 'ssc', '2码3中-不定位', '{\"className\": \"Ssc\\\\Bdw2m3z\"}', '所选号码包含在开奖号码所选位置中，即为中奖。（先选3个位置再选2个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:06:47', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (48, 'ssc', '2码4中-不定位', '{\"className\": \"Ssc\\\\Bdw2m4z\"}', '所选号码包含在开奖号码所选位置中，即为中奖。（先选4个位置再选2个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:06:54', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (49, 'ssc', '2码5中-不定位', '{\"className\": \"Ssc\\\\Bdw2m5z\"}', '所选号码包含在开奖号码中，即为中奖。（选2个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:07:01', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (50, 'ssc', '3码4中-不定位', '{\"className\": \"Ssc\\\\Bdw3m4z\"}', '所选号码包含在开奖号码所选位置中，即为中奖。（先选4个位置再选3个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:07:09', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (51, 'ssc', '3码5中-不定位', '{\"className\": \"Ssc\\\\Bdw3m5z\"}', '所选号码包含在开奖号码中，即为中奖。（选3个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:07:18', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (52, 'ssc', '4码5中-不定位', '{\"className\": \"Ssc\\\\Bdw4m5z\"}', '所选号码包含在开奖号码中，即为中奖。（选4个号码组成1注，超过则为复式票，但有号码多次选中时，会对号码组合做去重）', 0, '2022-01-25 10:07:25', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (53, 'ssc', '1中-大小单双', '{\"className\": \"Ssc\\\\Dxds1z\"}', '所选大小单双与开奖号码中所在位置大小单双相同，且位置一致，即为中奖。（在任意1个位置选1个大小单双组成1注，超过则为复式票。小：0~4；大：5~9；单：奇数；双：偶数；同一位置大和小只能选一个，单和双也是）', 0, '2022-01-25 10:07:32', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (54, 'ssc', '2中-大小单双', '{\"className\": \"Ssc\\\\Dxds2z\"}', '所选大小单双与开奖号码中所在位置大小单双相同，且位置一致，即为中奖。（在任意2个位置各选1个大小单双组成1注，超过则为复式票。小：0~4；大：5~9；单：奇数；双：偶数；同一位置大和小只能选一个，单和双也是）', 0, '2022-01-25 10:07:39', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (55, 'ssc', '3中-大小单双', '{\"className\": \"Ssc\\\\Dxds3z\"}', '所选大小单双与开奖号码中所在位置大小单双相同，且位置一致，即为中奖。（在任意3个位置各选1个大小单双组成1注，超过则为复式票。小：0~4；大：5~9；单：奇数；双：偶数；同一位置大和小只能选一个，单和双也是）', 0, '2022-01-25 10:07:46', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (56, 'ssc', '4中-大小单双', '{\"className\": \"Ssc\\\\Dxds4z\"}', '所选大小单双与开奖号码中所在位置大小单双相同，且位置一致，即为中奖。（在任意4个位置各选1个大小单双组成1注，超过则为复式票。小：0~4；大：5~9；单：奇数；双：偶数；同一位置大和小只能选一个，单和双也是）', 0, '2022-01-25 10:07:53', '2021-09-06 10:30:39');
INSERT INTO `lottery_play` VALUES (57, 'ssc', '5中-大小单双', '{\"className\": \"Ssc\\\\Dxds5z\"}', '所选大小单双与开奖号码中所在位置大小单双相同，且位置一致，即为中奖。（在5个位置各选1个大小单双组成1注，超过则为复式票。小：0~4；大：5~9；单：奇数；双：偶数；同一位置大和小只能选一个，单和双也是）', 0, '2022-01-25 10:10:12', '2021-09-06 10:30:39');

-- ----------------------------
-- Table structure for lottery_rel_play
-- ----------------------------
DROP TABLE IF EXISTS `lottery_rel_play`;
CREATE TABLE `lottery_rel_play`  (
  `lotteryId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '彩票id',
  `playId` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '玩法id',
  `oddsConfig` json NULL COMMENT '下注赔率（格式：{\"中奖等级1\":\"赔率\",\"中奖等级2\":\"赔率\",...},示例：{\"赢\":\"5.00\",\"平局\":\"1.00\"}）',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 50 COMMENT '排序值（从小到大排序，默认50，范围0-100）',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`lotteryId`, `playId`) USING BTREE,
  INDEX `lotteryId`(`lotteryId` ASC) USING BTREE,
  INDEX `playId`(`playId` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of lottery_rel_play
-- ----------------------------
INSERT INTO `lottery_rel_play` VALUES (1, 1, '{\"中奖\": \"48.80\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 2, '{\"其他颜色\": \"2.70\", \"最多颜色\": \"2.60\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 3, '{\"其他生肖\": \"12.00\", \"最多生肖\": \"9.50\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 4, '{\"其他五行\": \"4.41\", \"最多五行\": \"4.01\", \"最少五行\": \"5.51\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 5, '{\"其他头\": \"4.55\", \"其他尾\": \"9.50\", \"最少头\": \"5.10\", \"最少尾\": \"12.00\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 6, '{\"双色\": \"20.90\", \"其他颜色\": \"2.55\", \"最多颜色\": \"2.50\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 7, '{\"2肖\": \"12.73\", \"3肖\": \"12.73\", \"4肖\": \"12.73\", \"5肖\": \"2.77\", \"6肖\": \"1.77\", \"7肖\": \"4.82\", \"肖单\": \"1.77\", \"肖双\": \"1.68\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 8, '{\"一等奖\": \"110.00\", \"二等奖\": \"50.00\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 9, '{\"其他生肖\": \"1.90\", \"最多生肖\": \"1.62\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 10, '{\"中奖\": \"3.68\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 11, '{\"中奖\": \"9.95\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 12, '{\"中奖\": \"28.78\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 13, '{\"中奖\": \"88.25\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 14, '{\"其他尾\": \"1.62\", \"最少尾\": \"1.90\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 15, '{\"中奖\": \"2.88\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 16, '{\"中奖\": \"7.20\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 17, '{\"中奖\": \"17.00\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 18, '{\"中奖\": \"36.36\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 19, '{\"中奖\": \"8.02\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 20, '{\"中奖\": \"70.00\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 21, '{\"中奖\": \"650.00\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 22, '{\"中奖\": \"12729.84\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 23, '{\"中奖\": \"2.18\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 24, '{\"中奖\": \"2.60\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 25, '{\"中奖\": \"3.10\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 26, '{\"中奖\": \"3.72\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 27, '{\"中奖\": \"4.30\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 28, '{\"中奖\": \"5.20\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 29, '{\"中奖\": \"6.35\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (1, 30, '{\"中奖\": \"8.35\"}', 50, '2022-02-09 15:58:33');
INSERT INTO `lottery_rel_play` VALUES (2, 31, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 32, '{\"其他\": \"2.00\", \"半顺\": \"1.50\", \"对子\": \"4.00\", \"豹子\": \"15.00\", \"顺子\": \"8.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 33, '{\"无牛\": \"2.18\", \"其他牛\": \"10.89\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 34, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 35, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 36, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 37, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 38, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 39, '{\"一等奖\": \"1.00\", \"三等奖\": \"1.00\", \"二等奖\": \"1.00\", \"四等奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 40, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 41, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 42, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 43, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 44, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 45, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 46, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 47, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 48, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 49, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 50, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 51, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 52, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 53, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 54, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 55, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 56, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (2, 57, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:11');
INSERT INTO `lottery_rel_play` VALUES (3, 31, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 32, '{\"其他\": \"2.00\", \"半顺\": \"1.50\", \"对子\": \"4.00\", \"豹子\": \"15.00\", \"顺子\": \"8.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 33, '{\"无牛\": \"2.18\", \"其他牛\": \"10.89\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 34, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 35, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 36, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 37, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 38, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 39, '{\"一等奖\": \"1.00\", \"三等奖\": \"1.00\", \"二等奖\": \"1.00\", \"四等奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 40, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 41, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 42, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 43, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 44, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 45, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 46, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 47, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 48, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 49, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 50, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 51, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 52, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 53, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 54, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 55, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 56, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (3, 57, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:57:02');
INSERT INTO `lottery_rel_play` VALUES (4, 31, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 32, '{\"其他\": \"2.00\", \"半顺\": \"1.50\", \"对子\": \"4.00\", \"豹子\": \"15.00\", \"顺子\": \"8.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 33, '{\"无牛\": \"2.18\", \"其他牛\": \"10.89\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 34, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 35, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 36, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 37, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 38, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 39, '{\"一等奖\": \"1.00\", \"三等奖\": \"1.00\", \"二等奖\": \"1.00\", \"四等奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 40, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 41, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 42, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 43, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 44, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 45, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 46, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 47, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 48, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 49, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 50, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 51, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 52, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 53, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 54, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 55, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 56, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (4, 57, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:55');
INSERT INTO `lottery_rel_play` VALUES (5, 1, '{\"中奖\": \"48.80\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 2, '{\"其他颜色\": \"2.70\", \"最多颜色\": \"2.60\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 3, '{\"其他生肖\": \"12.00\", \"最多生肖\": \"9.50\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 4, '{\"其他五行\": \"4.41\", \"最多五行\": \"4.01\", \"最少五行\": \"5.51\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 5, '{\"其他头\": \"4.55\", \"其他尾\": \"9.50\", \"最少头\": \"5.10\", \"最少尾\": \"12.00\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 6, '{\"双色\": \"20.90\", \"其他颜色\": \"2.55\", \"最多颜色\": \"2.50\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 7, '{\"2肖\": \"12.73\", \"3肖\": \"12.73\", \"4肖\": \"12.73\", \"5肖\": \"2.77\", \"6肖\": \"1.77\", \"7肖\": \"4.82\", \"肖单\": \"1.77\", \"肖双\": \"1.68\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 8, '{\"一等奖\": \"110.00\", \"二等奖\": \"50.00\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 9, '{\"其他生肖\": \"1.90\", \"最多生肖\": \"1.62\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 10, '{\"中奖\": \"3.68\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 11, '{\"中奖\": \"9.95\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 12, '{\"中奖\": \"28.78\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 13, '{\"中奖\": \"88.25\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 14, '{\"其他尾\": \"1.62\", \"最少尾\": \"1.90\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 15, '{\"中奖\": \"2.88\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 16, '{\"中奖\": \"7.20\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 17, '{\"中奖\": \"17.00\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 18, '{\"中奖\": \"36.36\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 19, '{\"中奖\": \"8.02\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 20, '{\"中奖\": \"70.00\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 21, '{\"中奖\": \"650.00\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 22, '{\"中奖\": \"12729.84\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 23, '{\"中奖\": \"2.18\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 24, '{\"中奖\": \"2.60\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 25, '{\"中奖\": \"3.10\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 26, '{\"中奖\": \"3.72\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 27, '{\"中奖\": \"4.30\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 28, '{\"中奖\": \"5.20\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 29, '{\"中奖\": \"6.35\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (5, 30, '{\"中奖\": \"8.35\"}', 50, '2022-02-09 15:56:45');
INSERT INTO `lottery_rel_play` VALUES (6, 1, '{\"中奖\": \"48.80\"}', 50, '2022-02-09 15:56:36');
INSERT INTO `lottery_rel_play` VALUES (6, 2, '{\"其他颜色\": \"2.70\", \"最多颜色\": \"2.60\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 3, '{\"其他生肖\": \"12.00\", \"最多生肖\": \"9.50\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 4, '{\"其他五行\": \"4.41\", \"最多五行\": \"4.01\", \"最少五行\": \"5.51\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 5, '{\"其他头\": \"4.55\", \"其他尾\": \"9.50\", \"最少头\": \"5.10\", \"最少尾\": \"12.00\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 6, '{\"双色\": \"20.90\", \"其他颜色\": \"2.55\", \"最多颜色\": \"2.50\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 7, '{\"2肖\": \"12.73\", \"3肖\": \"12.73\", \"4肖\": \"12.73\", \"5肖\": \"2.77\", \"6肖\": \"1.77\", \"7肖\": \"4.82\", \"肖单\": \"1.77\", \"肖双\": \"1.68\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 8, '{\"一等奖\": \"110.00\", \"二等奖\": \"50.00\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 9, '{\"其他生肖\": \"1.90\", \"最多生肖\": \"1.62\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 10, '{\"中奖\": \"3.68\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 11, '{\"中奖\": \"9.95\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 12, '{\"中奖\": \"28.78\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 13, '{\"中奖\": \"88.25\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 14, '{\"其他尾\": \"1.62\", \"最少尾\": \"1.90\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 15, '{\"中奖\": \"2.88\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 16, '{\"中奖\": \"7.20\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 17, '{\"中奖\": \"17.00\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 18, '{\"中奖\": \"36.36\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 19, '{\"中奖\": \"8.02\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 20, '{\"中奖\": \"70.00\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 21, '{\"中奖\": \"650.00\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 22, '{\"中奖\": \"12729.84\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 23, '{\"中奖\": \"2.18\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 24, '{\"中奖\": \"2.60\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 25, '{\"中奖\": \"3.10\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 26, '{\"中奖\": \"3.72\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 27, '{\"中奖\": \"4.30\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 28, '{\"中奖\": \"5.20\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 29, '{\"中奖\": \"6.35\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (6, 30, '{\"中奖\": \"8.35\"}', 50, '2022-02-09 15:56:35');
INSERT INTO `lottery_rel_play` VALUES (7, 31, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 32, '{\"其他\": \"2.00\", \"半顺\": \"1.50\", \"对子\": \"4.00\", \"豹子\": \"15.00\", \"顺子\": \"8.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 33, '{\"无牛\": \"2.18\", \"其他牛\": \"10.89\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 34, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 35, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 36, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 37, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 38, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 39, '{\"一等奖\": \"1.00\", \"三等奖\": \"1.00\", \"二等奖\": \"1.00\", \"四等奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 40, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 41, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 42, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 43, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 44, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 45, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 46, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 47, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 48, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 49, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 50, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 51, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 52, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 53, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 54, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 55, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 56, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (7, 57, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:21');
INSERT INTO `lottery_rel_play` VALUES (8, 31, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 32, '{\"其他\": \"2.00\", \"半顺\": \"1.50\", \"对子\": \"4.00\", \"豹子\": \"15.00\", \"顺子\": \"8.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 33, '{\"无牛\": \"2.18\", \"其他牛\": \"10.89\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 34, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 35, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 36, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 37, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 38, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 39, '{\"一等奖\": \"1.00\", \"三等奖\": \"1.00\", \"二等奖\": \"1.00\", \"四等奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 40, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 41, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 42, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 43, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 44, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 45, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 46, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 47, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 48, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 49, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 50, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 51, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 52, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 53, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 54, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 55, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 56, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');
INSERT INTO `lottery_rel_play` VALUES (8, 57, '{\"中奖\": \"1.00\"}', 50, '2022-02-09 15:56:11');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `userId` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `account` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '账号',
  `mobile` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '手机',
  `password` char(32) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '密码（md5保存）',
  `nickname` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `userWallet` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '用户钱包',
  `isStop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否停用：0否 1是',
  `uid` char(12) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '用户uid',
  `thirdUid` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '三方用户uid',
  `updateTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `addTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`userId`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile` ASC) USING BTREE,
  UNIQUE INDEX `uid`(`uid` ASC) USING BTREE,
  UNIQUE INDEX `thirdUid`(`thirdUid` ASC) USING BTREE,
  UNIQUE INDEX `account`(`account` ASC) USING BTREE,
  INDEX `isStop`(`isStop` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
