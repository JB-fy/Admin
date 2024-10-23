/*
 Navicat Premium Dump SQL

 Source Server         : 本地-Mysql-8
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : 192.168.1.200:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 23/10/2024 11:31:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for goods
-- ----------------------------
DROP TABLE IF EXISTS `goods`;
CREATE TABLE `goods`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_stop` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '停用：0否 1是',
  `goods_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `goods_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `org_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `goods_no` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '编号',
  `image` json NOT NULL COMMENT '图片',
  `attr_show` json NULL COMMENT '展示属性。JSON格式：[{\"name\":\"属性名\",\"val\":\"属性值\"},...]',
  `attr_opt` json NULL COMMENT '可选属性。通常由不会影响价格和库存的属性组成。JSON格式：[{\"name\":\"属性名\",\"val_arr\":[\"属性值1\",\"属性值2\",...]},...]',
  `status` tinyint UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态：0上架 1下架',
  `sort` tinyint UNSIGNED NOT NULL DEFAULT 100 COMMENT '排序值。从大到小排序',
  PRIMARY KEY (`goods_id`) USING BTREE,
  INDEX `org_id`(`org_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for goods_attr
-- ----------------------------
DROP TABLE IF EXISTS `goods_attr`;
CREATE TABLE `goods_attr`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `attr_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '属性ID',
  `attr_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '名称',
  `org_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '机构ID',
  `attr_val_arr` json NOT NULL COMMENT '值。JSON格式：[\"值1\",\"值2\",...]',
  PRIMARY KEY (`attr_id`) USING BTREE,
  INDEX `org_id`(`org_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '属性表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for goods_spec
-- ----------------------------
DROP TABLE IF EXISTS `goods_spec`;
CREATE TABLE `goods_spec`  (
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `spec_id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '规格ID',
  `goods_id` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '商品ID',
  `attr_spec` json NULL COMMENT '规格属性。通常由会影响价格和库存的属性组成。JSON格式：[{\"name\":\"属性名\",\"val\":\"属性值\"},...]',
  `spec_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '图片',
  `price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '价格',
  `cost_price` decimal(10, 2) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT '成本价',
  `stock_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '库存',
  `sale_num` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '销量',
  PRIMARY KEY (`spec_id`) USING BTREE,
  INDEX `goods_id`(`goods_id` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '商品规格表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
