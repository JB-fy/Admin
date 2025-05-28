/*
 Navicat Premium Dump SQL

 Source Server         : Postgresql
 Source Server Type    : PostgreSQL
 Source Server Version : 170005 (170005)
 Source Host           : 192.168.0.200:5432
 Source Catalog        : admin
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 170005 (170005)
 File Encoding         : 65001

 Date: 29/05/2025 01:20:11
*/


-- ----------------------------
-- Sequence structure for app_pkg_pkg_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."app_pkg_pkg_id_seq";
CREATE SEQUENCE "public"."app_pkg_pkg_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_action_action_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_action_action_id_seq";
CREATE SEQUENCE "public"."auth_action_action_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_menu_menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_menu_menu_id_seq";
CREATE SEQUENCE "public"."auth_menu_menu_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_role_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_role_role_id_seq";
CREATE SEQUENCE "public"."auth_role_role_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for org_admin_admin_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."org_admin_admin_id_seq";
CREATE SEQUENCE "public"."org_admin_admin_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for org_org_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."org_org_id_seq";
CREATE SEQUENCE "public"."org_org_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pay_channel_channel_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."pay_channel_channel_id_seq";
CREATE SEQUENCE "public"."pay_channel_channel_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pay_order_order_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."pay_order_order_id_seq";
CREATE SEQUENCE "public"."pay_order_order_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pay_pay_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."pay_pay_id_seq";
CREATE SEQUENCE "public"."pay_pay_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for pay_scene_scene_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."pay_scene_scene_id_seq";
CREATE SEQUENCE "public"."pay_scene_scene_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for platform_admin_admin_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."platform_admin_admin_id_seq";
CREATE SEQUENCE "public"."platform_admin_admin_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for platform_server_server_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."platform_server_server_id_seq";
CREATE SEQUENCE "public"."platform_server_server_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for upload_upload_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."upload_upload_id_seq";
CREATE SEQUENCE "public"."upload_upload_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for users_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."users_user_id_seq";
CREATE SEQUENCE "public"."users_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS "public"."app";
CREATE TABLE "public"."app" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "app_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "app_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "app_config" json,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."app"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."app"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."app"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."app"."app_id" IS 'APPID';
COMMENT ON COLUMN "public"."app"."app_name" IS '名称';
COMMENT ON COLUMN "public"."app"."app_config" IS '配置。JSON格式，需要时设置';
COMMENT ON COLUMN "public"."app"."remark" IS '备注';

-- ----------------------------
-- Records of app
-- ----------------------------

-- ----------------------------
-- Table structure for app_pkg
-- ----------------------------
DROP TABLE IF EXISTS "public"."app_pkg";
CREATE TABLE "public"."app_pkg" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "pkg_id" int4 NOT NULL DEFAULT nextval('app_pkg_pkg_id_seq'::regclass),
  "app_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "pkg_type" int2 NOT NULL DEFAULT 0,
  "pkg_name" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "pkg_file" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ver_no" int4 NOT NULL DEFAULT 0,
  "ver_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ver_intro" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "extra_config" json,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_force_prev" int2 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."app_pkg"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."app_pkg"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."app_pkg"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."app_pkg"."pkg_id" IS '安装包ID';
COMMENT ON COLUMN "public"."app_pkg"."app_id" IS 'APPID';
COMMENT ON COLUMN "public"."app_pkg"."pkg_type" IS '类型：0安卓 1苹果 2PC';
COMMENT ON COLUMN "public"."app_pkg"."pkg_name" IS '包名';
COMMENT ON COLUMN "public"."app_pkg"."pkg_file" IS '安装包';
COMMENT ON COLUMN "public"."app_pkg"."ver_no" IS '版本号';
COMMENT ON COLUMN "public"."app_pkg"."ver_name" IS '版本名称';
COMMENT ON COLUMN "public"."app_pkg"."ver_intro" IS '版本介绍';
COMMENT ON COLUMN "public"."app_pkg"."extra_config" IS '额外配置。JSON格式，需要时设置';
COMMENT ON COLUMN "public"."app_pkg"."remark" IS '备注';
COMMENT ON COLUMN "public"."app_pkg"."is_force_prev" IS '强制更新：0否 1是。注意：只根据前一个版本来设置，与更早之前的版本无关';

-- ----------------------------
-- Records of app_pkg
-- ----------------------------

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action";
CREATE TABLE "public"."auth_action" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "action_id" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "action_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."auth_action"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_action"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_action"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action"."action_name" IS '名称';
COMMENT ON COLUMN "public"."auth_action"."remark" IS '备注';
COMMENT ON TABLE "public"."auth_action" IS '权限操作表';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appCreate', '系统管理-APP-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appDelete', '系统管理-APP-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgCreate', '系统管理-APP管理-安装包-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgDelete', '系统管理-APP管理-安装包-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgRead', '系统管理-APP管理-安装包-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appPkgUpdate', '系统管理-APP管理-安装包-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appRead', '系统管理-APP-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'appUpdate', '系统管理-APP-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionCreate', '权限管理-操作-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionDelete', '权限管理-操作-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionRead', '权限管理-操作-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authActionUpdate', '权限管理-操作-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuCreate', '权限管理-菜单-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuDelete', '权限管理-菜单-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuRead', '权限管理-菜单-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authMenuUpdate', '权限管理-菜单-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleCreate', '权限管理-角色-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleDelete', '权限管理-角色-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleRead', '权限管理-角色-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authRoleUpdate', '权限管理-角色-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneCreate', '权限管理-场景-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneDelete', '权限管理-场景-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneRead', '权限管理-场景-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'authSceneUpdate', '权限管理-场景-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminCreate', '权限管理-机构管理员-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminDelete', '权限管理-机构管理员-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminRead', '权限管理-机构管理员-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgAdminUpdate', '权限管理-机构管理员-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigCommonRead', '应用配置-常用-查看', '只能读取机构配置表中的某些配置。对应前端页面：配置中心-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigCommonSave', '应用配置-常用-保存', '只能保存机构配置表中的某些配置。对应前端页面：配置中心-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigRead', '配置中心-查看', '可任意读取机构配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgConfigSave', '配置中心-保存', '可任意保存机构配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgCreate', '机构管理-机构-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgDelete', '机构管理-机构-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgRead', '机构管理-机构-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'orgUpdate', '机构管理-机构-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelCreate', '系统管理-配置中心-支付管理-支付通道-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelDelete', '系统管理-配置中心-支付管理-支付通道-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelRead', '系统管理-配置中心-支付管理-支付通道-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payChannelUpdate', '系统管理-配置中心-支付管理-支付通道-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payCreate', '系统管理-配置中心-支付管理-支付配置-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payDelete', '系统管理-配置中心-支付管理-支付配置-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payRead', '系统管理-配置中心-支付管理-支付配置-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneCreate', '系统管理-配置中心-支付管理-支付场景-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneDelete', '系统管理-配置中心-支付管理-支付场景-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneRead', '系统管理-配置中心-支付管理-支付场景-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'paySceneUpdate', '系统管理-配置中心-支付管理-支付场景-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'payUpdate', '系统管理-配置中心-支付管理-支付配置-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminCreate', '权限管理-平台管理员-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminDelete', '权限管理-平台管理员-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminRead', '权限管理-平台管理员-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformAdminUpdate', '权限管理-平台管理员-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigCommonRead', '应用配置-常用-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigCommonSave', '应用配置-常用-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigEmailRead', '插件配置-邮箱-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigEmailSave', '插件配置-邮箱-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigIdCardRead', '插件配置-实名认证-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigIdCardSave', '插件配置-实名认证-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigOneClickRead', '插件配置-一键登录-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigOneClickSave', '插件配置-一键登录-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigPushRead', '插件配置-推送-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigPushSave', '插件配置-推送-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigRead', '平台配置-查看', '可任意读取平台配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSave', '平台配置-保存', '可任意保存平台配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSmsRead', '插件配置-短信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigSmsSave', '插件配置-短信-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigVodRead', '插件配置-视频点播-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigVodSave', '插件配置-视频点播-保存', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigWxRead', '插件配置-微信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platformConfigWxSave', '插件配置-微信-查看', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadCreate', '系统管理-配置中心-上传配置-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadDelete', '系统管理-配置中心-上传配置-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadRead', '系统管理-配置中心-上传配置-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'uploadUpdate', '系统管理-配置中心-上传配置-编辑', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'usersRead', '用户管理-用户-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'usersUpdate', '用户管理-用户-编辑', '');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action_rel_to_scene";
CREATE TABLE "public"."auth_action_rel_to_scene" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "action_id" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."scene_id" IS '场景ID';
COMMENT ON TABLE "public"."auth_action_rel_to_scene" IS '权限操作，权限场景关联表（操作可用在哪些场景）';

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appPkgCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appPkgDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appPkgRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appPkgUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'appUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authActionUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authMenuUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleCreate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleDelete', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleUpdate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authRoleUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneCreate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneDelete', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneUpdate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminCreate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminDelete', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminUpdate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgAdminUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'orgUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payChannelUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'paySceneUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformAdminUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigCommonRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigCommonSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigEmailRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigEmailSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigIdCardRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigIdCardSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigOneClickRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigOneClickSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigPushRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigPushSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSmsRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigSmsSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigVodRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigVodSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigWxRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platformConfigWxSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'usersRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'usersUpdate', 'platform');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_menu";
CREATE TABLE "public"."auth_menu" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "menu_id" int4 NOT NULL DEFAULT nextval('auth_menu_menu_id_seq'::regclass),
  "menu_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "pid" int4 NOT NULL DEFAULT 0,
  "is_leaf" int2 NOT NULL DEFAULT 1,
  "level" int2 NOT NULL DEFAULT 0,
  "id_path" text COLLATE "pg_catalog"."default" DEFAULT ''::text,
  "name_path" text COLLATE "pg_catalog"."default" DEFAULT ''::text,
  "menu_icon" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "menu_url" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "extra_data" json,
  "sort" int2 NOT NULL DEFAULT 100
)
;
COMMENT ON COLUMN "public"."auth_menu"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_menu"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_menu"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_menu"."menu_id" IS '菜单ID';
COMMENT ON COLUMN "public"."auth_menu"."menu_name" IS '名称';
COMMENT ON COLUMN "public"."auth_menu"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_menu"."pid" IS '父ID';
COMMENT ON COLUMN "public"."auth_menu"."is_leaf" IS '叶子：0否 1是';
COMMENT ON COLUMN "public"."auth_menu"."level" IS '层级';
COMMENT ON COLUMN "public"."auth_menu"."id_path" IS 'ID路径';
COMMENT ON COLUMN "public"."auth_menu"."name_path" IS '名称路径';
COMMENT ON COLUMN "public"."auth_menu"."menu_icon" IS '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}';
COMMENT ON COLUMN "public"."auth_menu"."menu_url" IS '链接';
COMMENT ON COLUMN "public"."auth_menu"."extra_data" IS '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}';
COMMENT ON COLUMN "public"."auth_menu"."sort" IS '排序值。从大到小排序';
COMMENT ON TABLE "public"."auth_menu" IS '权限菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '主页', 'platform', 0, 1, 1, '0-1', '-主页', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理', 'platform', 0, 0, 1, '0-2', '-权限管理', 'autoicon-ep-lock', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '场景', 'platform', 2, 1, 2, '0-2-3', '-权限管理-场景', 'autoicon-ep-flag', '/auth/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "场景"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '操作', 'platform', 2, 1, 2, '0-2-4', '-权限管理-操作', 'autoicon-ep-coordinate', '/auth/action', '{"i18n": {"title": {"en": "Action", "zh-cn": "操作"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '菜单', 'platform', 2, 1, 2, '0-2-5', '-权限管理-菜单', 'autoicon-ep-menu', '/auth/menu', '{"i18n": {"title": {"en": "Menu", "zh-cn": "菜单"}}}', 30);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '角色', 'platform', 2, 1, 2, '0-2-6', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '平台管理员', 'platform', 2, 1, 2, '0-2-7', '-权限管理-平台管理员', 'autoicon-ep-avatar', '/platform/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "平台管理员"}}}', 50);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '系统管理', 'platform', 0, 0, 1, '0-8', '-系统管理', 'autoicon-ep-platform', '', '{"i18n": {"title": {"en": "System Manage", "zh-cn": "系统管理"}}}', 20);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '配置中心', 'platform', 8, 0, 2, '0-8-9', '-系统管理-配置中心', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '上传配置', 'platform', 9, 1, 3, '0-8-9-10', '-系统管理-配置中心-上传配置', 'autoicon-ep-upload', '/upload/upload', '{"i18n": {"title": {"en": "Upload", "zh-cn": "上传配置"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '支付管理', 'platform', 9, 0, 3, '0-8-9-11', '-系统管理-配置中心-支付管理', 'autoicon-ep-coin', '', '{"i18n": {"title": {"en": "Pay Manage", "zh-cn": "支付管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '支付配置', 'platform', 11, 1, 4, '0-8-9-11-12', '-系统管理-配置中心-支付管理-支付配置', 'autoicon-ep-money', '/pay/pay', '{"i18n": {"title": {"en": "Pay", "zh-cn": "支付配置"}}}', 50);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '支付场景', 'platform', 11, 1, 4, '0-8-9-11-13', '-系统管理-配置中心-支付管理-支付场景', 'autoicon-ep-guide', '/pay/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "支付场景"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '支付通道', 'platform', 11, 1, 4, '0-8-9-11-14', '-系统管理-配置中心-支付管理-支付通道', 'autoicon-ep-connection', '/pay/channel', '{"i18n": {"title": {"en": "Channel", "zh-cn": "支付通道"}}}', 150);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '插件配置', 'platform', 9, 1, 3, '0-8-9-15', '-系统管理-配置中心-插件配置', 'autoicon-ep-ticket', '/platform/config/plugin', '{"i18n": {"title": {"en": "Plugin Config", "zh-cn": "插件配置"}}}', 150);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '应用配置', 'platform', 9, 1, 3, '0-8-9-16', '-系统管理-配置中心-应用配置', 'autoicon-ep-set-up', '/platform/config/app', '{"i18n": {"title": {"en": "APP Config", "zh-cn": "应用配置"}}}', 200);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, 'APP管理', 'platform', 8, 0, 2, '0-8-17', '-系统管理-APP管理', 'autoicon-ep-suitcase-line', '', '{"i18n": {"title": {"en": "APP Manage", "zh-cn": "APP管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, 'APP', 'platform', 17, 1, 3, '0-8-17-18', '-系统管理-APP管理-APP', 'autoicon-ep-apple', '/app/app', '{"i18n": {"title": {"en": "App", "zh-cn": "APP"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '安装包', 'platform', 17, 1, 3, '0-8-17-19', '-系统管理-APP管理-安装包', 'autoicon-ep-box', '/app/pkg', '{"i18n": {"title": {"en": "Pkg", "zh-cn": "安装包"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '用户管理', 'platform', 0, 0, 1, '0-20', '-用户管理', 'autoicon-ep-user-filled', '', '{"i18n": {"title": {"en": "User Manage", "zh-cn": "用户管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '用户', 'platform', 20, 1, 2, '0-20-21', '-用户管理-用户', 'autoicon-ep-user', '/users/users', '{"i18n": {"title": {"en": "Users", "zh-cn": "用户"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '机构管理', 'platform', 0, 0, 1, '0-22', '-机构管理', 'autoicon-ep-office-building', '', '{"i18n": {"title": {"en": "Org Manage", "zh-cn": "机构管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '机构', 'platform', 22, 1, 2, '0-22-23', '-机构管理-机构', 'autoicon-ep-school', '/org/org', '{"i18n": {"title": {"en": "Org", "zh-cn": "机构"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 24, '机构管理员', 'platform', 2, 1, 2, '0-2-24', '-权限管理-机构管理员', 'autoicon-ep-avatar', '/org/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "机构管理员"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 25, '主页', 'org', 0, 1, 1, '0-25', '-主页', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 26, '权限管理', 'org', 0, 0, 1, '0-26', '-权限管理', 'autoicon-ep-menu', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 27, '角色', 'org', 26, 1, 2, '0-26-27', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 28, '管理员', 'org', 26, 1, 2, '0-26-28', '-权限管理-管理员', 'autoicon-ep-avatar', '/org/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "管理员"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 29, '配置中心', 'org', 0, 0, 1, '0-29', '-配置中心', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 20);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 30, '应用配置', 'org', 29, 1, 2, '0-29-30', '-配置中心-应用配置', 'autoicon-ep-set-up', '/org/config/app', '{"i18n": {"title": {"en": "APP Config", "zh-cn": "应用配置"}}}', 200);

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role";
CREATE TABLE "public"."auth_role" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "role_id" int4 NOT NULL DEFAULT nextval('auth_role_role_id_seq'::regclass),
  "role_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "rel_id" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."auth_role"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role"."role_name" IS '名称';
COMMENT ON COLUMN "public"."auth_role"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_role"."rel_id" IS '关联ID。0表示平台创建，其它值根据scene_id对应不同表';
COMMENT ON TABLE "public"."auth_role" IS '权限角色表';

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_org_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_of_org_admin";
CREATE TABLE "public"."auth_role_rel_of_org_admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "admin_id" int4 NOT NULL,
  "role_id" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."auth_role_rel_of_org_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_org_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_org_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_org_admin"."role_id" IS '角色ID';
COMMENT ON TABLE "public"."auth_role_rel_of_org_admin" IS '机构管理员，权限角色关联表（机构管理员包含哪些角色）';

-- ----------------------------
-- Records of auth_role_rel_of_org_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_of_platform_admin";
CREATE TABLE "public"."auth_role_rel_of_platform_admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "admin_id" int4 NOT NULL DEFAULT 0,
  "role_id" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."role_id" IS '角色ID';
COMMENT ON TABLE "public"."auth_role_rel_of_platform_admin" IS '平台管理员，权限角色关联表（平台管理员包含哪些角色）';

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_action";
CREATE TABLE "public"."auth_role_rel_to_action" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "role_id" int4 NOT NULL DEFAULT 0,
  "action_id" varchar(30) COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."action_id" IS '操作ID';
COMMENT ON TABLE "public"."auth_role_rel_to_action" IS '权限角色，权限操作关联表（角色包含哪些操作）';

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_menu";
CREATE TABLE "public"."auth_role_rel_to_menu" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "role_id" int4 NOT NULL DEFAULT 0,
  "menu_id" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."menu_id" IS '菜单ID';
COMMENT ON TABLE "public"."auth_role_rel_to_menu" IS '权限角色，权限菜单关联表（角色包含哪些菜单）';

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_scene";
CREATE TABLE "public"."auth_scene" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "scene_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_config" json,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."auth_scene"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_scene"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_scene"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_scene"."scene_name" IS '名称';
COMMENT ON COLUMN "public"."auth_scene"."scene_config" IS '配置。JSON格式，根据场景设置';
COMMENT ON COLUMN "public"."auth_scene"."remark" IS '备注';
COMMENT ON TABLE "public"."auth_scene" IS '权限场景表';

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'app', 'APP', '{"token_config": {"is_ip": 0, "is_unique": 0, "sign_type": "HS256", "token_type": 0, "active_time": 0, "expire_time": 604800, "private_key": "任意字符串"}}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'org', '机构后台', '{"token_config": {"is_ip": 1, "is_unique": 1, "sign_type": "RS256", "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0nT4zSS2O+2yib+gdBm0\nW0kU2AL6felZFp5U6B/ySeJnM8UE+fZwytgPuND5Z07khxtHR/YYH7huPGZ/fAgh\nZHmNE5K5phMI5eETwPjk3RDeyYAyOosrKr3SAjjEQxJBISvBEillH4bKjaa4WF5/\n+nGcp7f4e49caW/CfwuC2ZVrvySCPf1lR8o7/4Zz/hWUwgsEd/crR7ojgt+rbPeE\n7+Cz11sZUZaMipTqsU3RVhbwtMyLdkos6KsYW7TZEK0VYt94/1XJQBUEjtCdDpS7\n0XNF8ENpQPtuQdYE6/y+Jku8T9pqQQq/SOL6uPgsn6zJQ41u/l2AhG0i5GxYD86C\n5wIDAQAB\n-----END PUBLIC KEY-----", "token_type": 0, "active_time": 3600, "expire_time": 14400, "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDSdPjNJLY77bKJ\nv6B0GbRbSRTYAvp96VkWnlToH/JJ4mczxQT59nDK2A+40PlnTuSHG0dH9hgfuG48\nZn98CCFkeY0TkrmmEwjl4RPA+OTdEN7JgDI6iysqvdICOMRDEkEhK8ESKWUfhsqN\nprhYXn/6cZynt/h7j1xpb8J/C4LZlWu/JII9/WVHyjv/hnP+FZTCCwR39ytHuiOC\n36ts94Tv4LPXWxlRloyKlOqxTdFWFvC0zIt2SizoqxhbtNkQrRVi33j/VclAFQSO\n0J0OlLvRc0XwQ2lA+25B1gTr/L4mS7xP2mpBCr9I4vq4+CyfrMlDjW7+XYCEbSLk\nbFgPzoLnAgMBAAECggEAGolMO9WmsrzAd9T1Pt5k2uPGoIwTmJ+9L3hsXU515vII\nsELl4zy7MSB4LwYOhIOylgSPAthZZ1qCb9Q+u91slHYtHywvg2zAAPhV3M2lUeiI\nJuEmtDILEdsYaVZODOT22F9je05D5WtCDAVbFi1oNqRvq8grKS1E6jiAzjMd3yBY\n5AgUUP8sbS7BDdPus2t3mCAXqdtFkxn8wo/4WdMV6vG4p9p+a8dIoiRYHNBIw4sU\nCYPiE7q52tVqVl10ShrJQlvoyDvmJam4inbl8ZWtbQLUsQxfzoEUfuuC0mYc2pES\n1kp1pWfJNc5JtAXXZ/k9F4jLvjMp9KJximOG+E5tmQKBgQD5RwG8xu7s2JcocogD\nuJjWfLz/4ab7Zs2NReCX6jmb522d1TqyiU9ilc0XBz8gDNeMwnOfWIyGR3Vtuub+\nU56Y9/q+IMZ0ewdrUTR2NADZdqLH1ViVGacnnmJXJ+30Z8eUO0GCU7++DSQ1u8xh\ng0mmrj6++xHsYJM3bCxBstNzGQKBgQDYIfOUIK+JKvYZ5idUrFWZnnVSdLu8IJHi\nA0osaMB9VNtAdqrPVIA3L9AbIR1/la3gb3ILP0hIM7glt4i78WTwVrh0qkaz/8dw\nX3EL5u9OAMecVFn4+gZon4RqbzjdtHso87v2GrIZ88eVOWjXRq93gWAe17G+noyS\n44PgB7Zl/wKBgC3bOh6YGevIDEaMiyjkFHmgiMQppqYoyzdp218W33ImqKuYRiwB\nxnDETe4mjx4+PojOXKa7i15IVvnQoB25FDvfomjHbrqOx1aeoZ/9AQsAIAHS5XDI\nP0+yezS9S7DiRnymSe7HqUY09KxN19M4a5wWAcTwOuPZADv50kpjszJBAoGAS0io\nO7SW8ESSrLrKgGf2+SeE3k/jBMijiAJ1V7q1MfLY3D95h/Z7Ir340zpZuBM/Gao4\nI0rLtrqtLhYb/rs62aybW6fkMNarda0JB4hNWvJSlVWccWlFyjOmQBy1xiQTslQT\n6Mmrt/Z+UrBIoJPyksHx5UxkkW1QsemmCecl1akCgYAys4PNRhdToxS5WDF7Tddg\n8ghhobuwUCP2pxNYakX9HeHZsAjwwQhx6vGJqZPs5hh8HRinGWSvKyVXaUGVN+5b\nFjMz/rhO9vTCZJS3aJSXxr0PFTbpP/AZSXwCBxjp+uEFTD8GJRX7/6+wlTk7+uTj\nl7klp31noFtzz+onGSmqvA==\n-----END PRIVATE KEY-----"}}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platform', '平台后台', '{"token_config": {"is_ip": 1, "is_unique": 1, "sign_type": "ES256", "public_key": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42\nhF5PnYenbWprk1MQIzT2+V4rRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\n-----END PUBLIC KEY-----", "token_type": 0, "active_time": 3600, "expire_time": 604800, "private_key": "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIKvYPRtCqy9MI/yhx4L4+Sog/5lntHbuwxPg/JI0qW6LoAoGCCqGSM49\nAwEHoUQDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42hF5PnYenbWprk1MQIzT2+V4r\nRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\n-----END EC PRIVATE KEY-----"}}', '');

-- ----------------------------
-- Table structure for org
-- ----------------------------
DROP TABLE IF EXISTS "public"."org";
CREATE TABLE "public"."org" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "org_id" int4 NOT NULL DEFAULT nextval('org_org_id_seq'::regclass),
  "org_name" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."org"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."org"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."org"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."org"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."org"."org_name" IS '机构名称';
COMMENT ON TABLE "public"."org" IS '机构表';

-- ----------------------------
-- Records of org
-- ----------------------------

-- ----------------------------
-- Table structure for org_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."org_admin";
CREATE TABLE "public"."org_admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "admin_id" int4 NOT NULL DEFAULT nextval('org_admin_admin_id_seq'::regclass),
  "org_id" int4 NOT NULL DEFAULT 0,
  "is_super" int2 NOT NULL DEFAULT 0,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "email" varchar(60) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar
)
;
COMMENT ON COLUMN "public"."org_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."org_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."org_admin"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."org_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."org_admin"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."org_admin"."is_super" IS '超管：0否 1是';
COMMENT ON COLUMN "public"."org_admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."org_admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."org_admin"."phone" IS '手机';
COMMENT ON COLUMN "public"."org_admin"."email" IS '邮箱';
COMMENT ON COLUMN "public"."org_admin"."account" IS '账号';
COMMENT ON COLUMN "public"."org_admin"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."org_admin"."salt" IS '密码盐';
COMMENT ON TABLE "public"."org_admin" IS '机构管理员表';

-- ----------------------------
-- Records of org_admin
-- ----------------------------

-- ----------------------------
-- Table structure for org_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."org_config";
CREATE TABLE "public"."org_config" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "org_id" int4 NOT NULL DEFAULT 0,
  "config_key" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "config_value" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text
)
;
COMMENT ON COLUMN "public"."org_config"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."org_config"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."org_config"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."org_config"."config_key" IS '配置键';
COMMENT ON COLUMN "public"."org_config"."config_value" IS '配置值';
COMMENT ON TABLE "public"."org_config" IS '机构配置表';

-- ----------------------------
-- Records of org_config
-- ----------------------------

-- ----------------------------
-- Table structure for pay
-- ----------------------------
DROP TABLE IF EXISTS "public"."pay";
CREATE TABLE "public"."pay" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "pay_id" int4 NOT NULL DEFAULT nextval('pay_pay_id_seq'::regclass),
  "pay_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "pay_type" int2 NOT NULL DEFAULT 0,
  "pay_config" json NOT NULL,
  "pay_rate" numeric(4,4) NOT NULL DEFAULT 0.0000,
  "total_amount" numeric(14,2) NOT NULL DEFAULT 0.00,
  "balance" numeric(18,6) NOT NULL DEFAULT 0.000000,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."pay"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay"."pay_id" IS '支付ID';
COMMENT ON COLUMN "public"."pay"."pay_name" IS '名称';
COMMENT ON COLUMN "public"."pay"."pay_type" IS '类型：0支付宝 1微信';
COMMENT ON COLUMN "public"."pay"."pay_config" IS '配置。根据pay_type类型设置';
COMMENT ON COLUMN "public"."pay"."pay_rate" IS '费率';
COMMENT ON COLUMN "public"."pay"."total_amount" IS '总额';
COMMENT ON COLUMN "public"."pay"."balance" IS '余额';
COMMENT ON COLUMN "public"."pay"."remark" IS '备注';
COMMENT ON TABLE "public"."pay" IS '支付表';

-- ----------------------------
-- Records of pay
-- ----------------------------

-- ----------------------------
-- Table structure for pay_channel
-- ----------------------------
DROP TABLE IF EXISTS "public"."pay_channel";
CREATE TABLE "public"."pay_channel" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "channel_id" int4 NOT NULL DEFAULT nextval('pay_channel_channel_id_seq'::regclass),
  "channel_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_id" int4 NOT NULL DEFAULT 0,
  "pay_id" int4 NOT NULL DEFAULT 0,
  "pay_method" int2 NOT NULL DEFAULT 0,
  "channel_icon" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sort" int2 NOT NULL DEFAULT 100,
  "total_amount" numeric(14,2) NOT NULL DEFAULT 0.00
)
;
COMMENT ON COLUMN "public"."pay_channel"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay_channel"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay_channel"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."pay_channel"."channel_id" IS '通道ID';
COMMENT ON COLUMN "public"."pay_channel"."channel_name" IS '名称';
COMMENT ON COLUMN "public"."pay_channel"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."pay_channel"."pay_id" IS '支付ID';
COMMENT ON COLUMN "public"."pay_channel"."pay_method" IS '支付方法：0APP支付 1H5支付 2扫码支付 3小程序支付';
COMMENT ON COLUMN "public"."pay_channel"."channel_icon" IS '图标';
COMMENT ON COLUMN "public"."pay_channel"."sort" IS '排序值。从大到小排序';
COMMENT ON COLUMN "public"."pay_channel"."total_amount" IS '总额';
COMMENT ON TABLE "public"."pay_channel" IS '支付通道表';

-- ----------------------------
-- Records of pay_channel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order
-- ----------------------------
DROP TABLE IF EXISTS "public"."pay_order";
CREATE TABLE "public"."pay_order" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "order_id" int4 NOT NULL DEFAULT nextval('pay_order_order_id_seq'::regclass),
  "order_no" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "rel_order_type" int2 NOT NULL DEFAULT 0,
  "rel_order_user_id" int4 NOT NULL DEFAULT 0,
  "pay_id" int4 NOT NULL DEFAULT 0,
  "channel_id" int4 NOT NULL DEFAULT 0,
  "pay_type" int2 NOT NULL DEFAULT 0,
  "amount" numeric(10,2) NOT NULL DEFAULT 0.00,
  "pay_status" int2 NOT NULL DEFAULT 0,
  "pay_time" timestamp(6),
  "pay_rate" numeric(4,4) NOT NULL DEFAULT 0.0000,
  "third_order_no" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."pay_order"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay_order"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay_order"."order_id" IS '订单ID';
COMMENT ON COLUMN "public"."pay_order"."order_no" IS '订单号';
COMMENT ON COLUMN "public"."pay_order"."rel_order_type" IS '关联订单类型：0默认';
COMMENT ON COLUMN "public"."pay_order"."rel_order_user_id" IS '关联订单用户ID';
COMMENT ON COLUMN "public"."pay_order"."pay_id" IS '支付ID';
COMMENT ON COLUMN "public"."pay_order"."channel_id" IS '通道ID';
COMMENT ON COLUMN "public"."pay_order"."pay_type" IS '类型：0支付宝 1微信';
COMMENT ON COLUMN "public"."pay_order"."amount" IS '实付金额';
COMMENT ON COLUMN "public"."pay_order"."pay_status" IS '状态：0未付款 1已付款';
COMMENT ON COLUMN "public"."pay_order"."pay_time" IS '支付时间';
COMMENT ON COLUMN "public"."pay_order"."pay_rate" IS '费率';
COMMENT ON COLUMN "public"."pay_order"."third_order_no" IS '第三方订单号';
COMMENT ON TABLE "public"."pay_order" IS '支付订单表';

-- ----------------------------
-- Records of pay_order
-- ----------------------------

-- ----------------------------
-- Table structure for pay_order_rel
-- ----------------------------
DROP TABLE IF EXISTS "public"."pay_order_rel";
CREATE TABLE "public"."pay_order_rel" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "order_id" int4 NOT NULL DEFAULT 0,
  "rel_order_type" int2 NOT NULL DEFAULT 0,
  "rel_order_id" int4 NOT NULL DEFAULT 0,
  "rel_order_no" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "rel_order_user_id" int4 NOT NULL DEFAULT 0,
  "rel_order_amount" numeric(10,2) NOT NULL DEFAULT 0.00
)
;
COMMENT ON COLUMN "public"."pay_order_rel"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay_order_rel"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay_order_rel"."order_id" IS '订单ID';
COMMENT ON COLUMN "public"."pay_order_rel"."rel_order_type" IS '关联订单类型：0默认';
COMMENT ON COLUMN "public"."pay_order_rel"."rel_order_id" IS '关联订单ID';
COMMENT ON COLUMN "public"."pay_order_rel"."rel_order_no" IS '关联订单号';
COMMENT ON COLUMN "public"."pay_order_rel"."rel_order_user_id" IS '关联订单用户ID';
COMMENT ON COLUMN "public"."pay_order_rel"."rel_order_amount" IS '关联订单实付金额';
COMMENT ON TABLE "public"."pay_order_rel" IS '支付订单关联表';

-- ----------------------------
-- Records of pay_order_rel
-- ----------------------------

-- ----------------------------
-- Table structure for pay_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."pay_scene";
CREATE TABLE "public"."pay_scene" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "scene_id" int4 NOT NULL DEFAULT nextval('pay_scene_scene_id_seq'::regclass),
  "scene_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."pay_scene"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay_scene"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."pay_scene"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."pay_scene"."scene_name" IS '名称';
COMMENT ON COLUMN "public"."pay_scene"."remark" IS '备注';
COMMENT ON TABLE "public"."pay_scene" IS '支付场景表';

-- ----------------------------
-- Records of pay_scene
-- ----------------------------

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_admin";
CREATE TABLE "public"."platform_admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "admin_id" int4 NOT NULL DEFAULT nextval('platform_admin_admin_id_seq'::regclass),
  "is_super" int2 NOT NULL DEFAULT 0,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "email" varchar(60) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar
)
;
COMMENT ON COLUMN "public"."platform_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."platform_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_admin"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."platform_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."platform_admin"."is_super" IS '超管：0否 1是';
COMMENT ON COLUMN "public"."platform_admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."platform_admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."platform_admin"."phone" IS '手机';
COMMENT ON COLUMN "public"."platform_admin"."email" IS '邮箱';
COMMENT ON COLUMN "public"."platform_admin"."account" IS '账号';
COMMENT ON COLUMN "public"."platform_admin"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."platform_admin"."salt" IS '密码盐';
COMMENT ON TABLE "public"."platform_admin" IS '平台管理员表';

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO "public"."platform_admin" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, 1, '超级管理员', '', NULL, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_config";
CREATE TABLE "public"."platform_config" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "config_key" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "config_value" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text
)
;
COMMENT ON COLUMN "public"."platform_config"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."platform_config"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_config"."config_key" IS '配置键';
COMMENT ON COLUMN "public"."platform_config"."config_value" IS '配置值';
COMMENT ON TABLE "public"."platform_config" IS '平台配置表';

-- ----------------------------
-- Records of platform_config
-- ----------------------------
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailCode', '{"subject":"您的验证码","template":"验证码：{code}\n说明：\n1. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。\n2. 出于安全考虑，请不要将此验证码分享给任何人。"}');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommon', '{"fromEmail":"xxxxxxxx@qq.com","password":"xxxxxxxx","smtpHost":"smtp.qq.com","smtpPort":"465"}');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailType', 'emailOfCommon');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardOfAliyun', '{"appcode":"appcode","url":"http://idcard.market.alicloudapi.com/lianzhuo/idcard"}');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardType', 'idCardOfAliyun');

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_server";
CREATE TABLE "public"."platform_server" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "server_id" int4 NOT NULL DEFAULT nextval('platform_server_server_id_seq'::regclass),
  "network_ip" varchar(15) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "local_ip" varchar(15) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."platform_server"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."platform_server"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_server"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."platform_server"."server_id" IS '服务器ID';
COMMENT ON COLUMN "public"."platform_server"."network_ip" IS '公网IP';
COMMENT ON COLUMN "public"."platform_server"."local_ip" IS '内网IP';
COMMENT ON TABLE "public"."platform_server" IS '平台服务器表';

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for upload
-- ----------------------------
DROP TABLE IF EXISTS "public"."upload";
CREATE TABLE "public"."upload" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "upload_id" int4 NOT NULL DEFAULT nextval('upload_upload_id_seq'::regclass),
  "upload_type" int2 NOT NULL DEFAULT 0,
  "upload_config" json NOT NULL,
  "is_default" int2 NOT NULL DEFAULT 0,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."upload"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."upload"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."upload"."upload_id" IS '上传ID';
COMMENT ON COLUMN "public"."upload"."upload_type" IS '类型：0本地 1阿里云OSS';
COMMENT ON COLUMN "public"."upload"."upload_config" IS '配置。JSON格式，根据类型设置';
COMMENT ON COLUMN "public"."upload"."is_default" IS '默认：0否 1是';
COMMENT ON COLUMN "public"."upload"."remark" IS '备注';
COMMENT ON TABLE "public"."upload" IS '上传表';

-- ----------------------------
-- Records of upload
-- ----------------------------
INSERT INTO "public"."upload" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 0, '{"sign_key": "secretKey", "is_cluster": 1, "server_list": [], "is_same_server": 0}', 1, '此项目自带文件上传下载功能，可直接部署成文件服务器使用');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "user_id" int4 NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "gender" int2 NOT NULL DEFAULT 0,
  "birthday" date,
  "address" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "email" varchar(60) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "wx_openid" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "wx_unionid" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
COMMENT ON COLUMN "public"."users"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."users"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."users"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."users"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."users"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."users"."avatar" IS '头像';
COMMENT ON COLUMN "public"."users"."gender" IS '性别：0未设置 1男 2女';
COMMENT ON COLUMN "public"."users"."birthday" IS '生日';
COMMENT ON COLUMN "public"."users"."address" IS '地址';
COMMENT ON COLUMN "public"."users"."phone" IS '手机';
COMMENT ON COLUMN "public"."users"."email" IS '邮箱';
COMMENT ON COLUMN "public"."users"."account" IS '账号';
COMMENT ON COLUMN "public"."users"."wx_openid" IS '微信openid';
COMMENT ON COLUMN "public"."users"."wx_unionid" IS '微信unionid';
COMMENT ON TABLE "public"."users" IS '用户表（postgresql中user是关键字，使用需要加双引号。程序中考虑与mysql通用，故命名成users）';

-- ----------------------------
-- Records of users
-- ----------------------------

-- ----------------------------
-- Table structure for users_privacy
-- ----------------------------
DROP TABLE IF EXISTS "public"."users_privacy";
CREATE TABLE "public"."users_privacy" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "user_id" int4 NOT NULL DEFAULT 0,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "id_card_no" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "id_card_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "id_card_gender" int2 NOT NULL DEFAULT 0,
  "id_card_birthday" date,
  "id_card_address" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."users_privacy"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."users_privacy"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."users_privacy"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."users_privacy"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."users_privacy"."salt" IS '密码盐';
COMMENT ON COLUMN "public"."users_privacy"."id_card_no" IS '身份证号码';
COMMENT ON COLUMN "public"."users_privacy"."id_card_name" IS '身份证姓名';
COMMENT ON COLUMN "public"."users_privacy"."id_card_gender" IS '身份证性别：0未设置 1男 2女';
COMMENT ON COLUMN "public"."users_privacy"."id_card_birthday" IS '身份证生日';
COMMENT ON COLUMN "public"."users_privacy"."id_card_address" IS '身份证地址';
COMMENT ON TABLE "public"."users_privacy" IS '用户隐私表';

-- ----------------------------
-- Records of users_privacy
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."app_pkg_pkg_id_seq"
OWNED BY "public"."app_pkg"."pkg_id";
SELECT setval('"public"."app_pkg_pkg_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_action_action_id_seq"
OWNED BY "public"."auth_action"."action_id";
SELECT setval('"public"."auth_action_action_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_menu_menu_id_seq"
OWNED BY "public"."auth_menu"."menu_id";
SELECT setval('"public"."auth_menu_menu_id_seq"', 28, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_role_role_id_seq"
OWNED BY "public"."auth_role"."role_id";
SELECT setval('"public"."auth_role_role_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."org_admin_admin_id_seq"
OWNED BY "public"."org_admin"."admin_id";
SELECT setval('"public"."org_admin_admin_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."org_org_id_seq"
OWNED BY "public"."org"."org_id";
SELECT setval('"public"."org_org_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."pay_channel_channel_id_seq"
OWNED BY "public"."pay_channel"."channel_id";
SELECT setval('"public"."pay_channel_channel_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."pay_order_order_id_seq"
OWNED BY "public"."pay_order"."order_id";
SELECT setval('"public"."pay_order_order_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."pay_pay_id_seq"
OWNED BY "public"."pay"."pay_id";
SELECT setval('"public"."pay_pay_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."pay_scene_scene_id_seq"
OWNED BY "public"."pay_scene"."scene_id";
SELECT setval('"public"."pay_scene_scene_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_admin_admin_id_seq"
OWNED BY "public"."platform_admin"."admin_id";
SELECT setval('"public"."platform_admin_admin_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_server_server_id_seq"
OWNED BY "public"."platform_server"."server_id";
SELECT setval('"public"."platform_server_server_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."upload_upload_id_seq"
OWNED BY "public"."upload"."upload_id";
SELECT setval('"public"."upload_upload_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."users_user_id_seq"
OWNED BY "public"."users"."user_id";
SELECT setval('"public"."users_user_id_seq"', 1, false);

-- ----------------------------
-- Primary Key structure for table app
-- ----------------------------
ALTER TABLE "public"."app" ADD CONSTRAINT "app_pkey" PRIMARY KEY ("app_id");

-- ----------------------------
-- Indexes structure for table app_pkg
-- ----------------------------
CREATE INDEX "app_pkg_app_id_idx" ON "public"."app_pkg" USING btree (
  "app_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table app_pkg
-- ----------------------------
ALTER TABLE "public"."app_pkg" ADD CONSTRAINT "app_pkg_pkey" PRIMARY KEY ("pkg_id");

-- ----------------------------
-- Primary Key structure for table auth_action
-- ----------------------------
ALTER TABLE "public"."auth_action" ADD CONSTRAINT "auth_action_pkey" PRIMARY KEY ("action_id");

-- ----------------------------
-- Indexes structure for table auth_action_rel_to_scene
-- ----------------------------
CREATE INDEX "auth_action_rel_to_scene_action_id_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "action_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "auth_action_rel_to_scene_scene_id_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "scene_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_action_rel_to_scene
-- ----------------------------
ALTER TABLE "public"."auth_action_rel_to_scene" ADD CONSTRAINT "auth_action_rel_to_scene_pkey" PRIMARY KEY ("action_id", "scene_id");

-- ----------------------------
-- Indexes structure for table auth_menu
-- ----------------------------
CREATE INDEX "auth_menu_pid_idx" ON "public"."auth_menu" USING btree (
  "pid" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_menu_scene_id_idx" ON "public"."auth_menu" USING btree (
  "scene_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_menu
-- ----------------------------
ALTER TABLE "public"."auth_menu" ADD CONSTRAINT "auth_menu_pkey" PRIMARY KEY ("menu_id");

-- ----------------------------
-- Indexes structure for table auth_role
-- ----------------------------
CREATE INDEX "auth_role_rel_id_idx" ON "public"."auth_role" USING btree (
  "rel_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_scene_id_idx" ON "public"."auth_role" USING btree (
  "scene_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role
-- ----------------------------
ALTER TABLE "public"."auth_role" ADD CONSTRAINT "auth_role_pkey" PRIMARY KEY ("role_id");

-- ----------------------------
-- Indexes structure for table auth_role_rel_of_org_admin
-- ----------------------------
CREATE INDEX "auth_role_rel_of_org_admin_admin_id_idx" ON "public"."auth_role_rel_of_org_admin" USING btree (
  "admin_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_of_org_admin_role_id_idx" ON "public"."auth_role_rel_of_org_admin" USING btree (
  "role_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_of_org_admin
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_of_org_admin" ADD CONSTRAINT "auth_role_rel_of_org_admin_pkey" PRIMARY KEY ("admin_id", "role_id");

-- ----------------------------
-- Indexes structure for table auth_role_rel_of_platform_admin
-- ----------------------------
CREATE INDEX "auth_role_rel_of_platform_admin_admin_id_idx" ON "public"."auth_role_rel_of_platform_admin" USING btree (
  "admin_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_of_platform_admin_role_id_idx" ON "public"."auth_role_rel_of_platform_admin" USING btree (
  "role_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_of_platform_admin
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_of_platform_admin" ADD CONSTRAINT "auth_role_rel_of_platform_admin_pkey" PRIMARY KEY ("admin_id", "role_id");

-- ----------------------------
-- Indexes structure for table auth_role_rel_to_action
-- ----------------------------
CREATE INDEX "auth_role_rel_to_action_action_id_idx" ON "public"."auth_role_rel_to_action" USING btree (
  "action_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_to_action_role_id_idx" ON "public"."auth_role_rel_to_action" USING btree (
  "role_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_to_action
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_to_action" ADD CONSTRAINT "auth_role_rel_to_action_pkey" PRIMARY KEY ("role_id", "action_id");

-- ----------------------------
-- Indexes structure for table auth_role_rel_to_menu
-- ----------------------------
CREATE INDEX "auth_role_rel_to_menu_menu_id_idx" ON "public"."auth_role_rel_to_menu" USING btree (
  "menu_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_to_menu_role_id_idx" ON "public"."auth_role_rel_to_menu" USING btree (
  "role_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_to_menu
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_to_menu" ADD CONSTRAINT "auth_role_rel_to_menu_pkey" PRIMARY KEY ("role_id", "menu_id");

-- ----------------------------
-- Primary Key structure for table auth_scene
-- ----------------------------
ALTER TABLE "public"."auth_scene" ADD CONSTRAINT "auth_scene_pkey" PRIMARY KEY ("scene_id");

-- ----------------------------
-- Primary Key structure for table org
-- ----------------------------
ALTER TABLE "public"."org" ADD CONSTRAINT "org_pkey" PRIMARY KEY ("org_id");

-- ----------------------------
-- Indexes structure for table org_admin
-- ----------------------------
CREATE UNIQUE INDEX "org_admin_org_id_account_idx" ON "public"."org_admin" USING btree (
  "org_id" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "org_admin_org_id_email_idx" ON "public"."org_admin" USING btree (
  "org_id" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "org_admin_org_id_idx" ON "public"."org_admin" USING btree (
  "org_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "org_admin_org_id_phone_idx" ON "public"."org_admin" USING btree (
  "org_id" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table org_admin
-- ----------------------------
ALTER TABLE "public"."org_admin" ADD CONSTRAINT "org_admin_pkey" PRIMARY KEY ("admin_id");

-- ----------------------------
-- Primary Key structure for table org_config
-- ----------------------------
ALTER TABLE "public"."org_config" ADD CONSTRAINT "org_config_pkey" PRIMARY KEY ("org_id", "config_key");

-- ----------------------------
-- Primary Key structure for table pay
-- ----------------------------
ALTER TABLE "public"."pay" ADD CONSTRAINT "pay_pkey" PRIMARY KEY ("pay_id");

-- ----------------------------
-- Indexes structure for table pay_channel
-- ----------------------------
CREATE INDEX "pay_channel_pay_id_idx" ON "public"."pay_channel" USING btree (
  "pay_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "pay_channel_scene_id_idx" ON "public"."pay_channel" USING btree (
  "scene_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table pay_channel
-- ----------------------------
ALTER TABLE "public"."pay_channel" ADD CONSTRAINT "pay_channel_pkey" PRIMARY KEY ("channel_id");

-- ----------------------------
-- Indexes structure for table pay_order
-- ----------------------------
CREATE INDEX "pay_order_channel_id_idx" ON "public"."pay_order" USING btree (
  "channel_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "pay_order_order_no_idx" ON "public"."pay_order" USING btree (
  "order_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "pay_order_pay_id_idx" ON "public"."pay_order" USING btree (
  "pay_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "pay_order_rel_order_user_id_idx" ON "public"."pay_order" USING btree (
  "rel_order_user_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "pay_order_third_order_no_idx" ON "public"."pay_order" USING btree (
  "third_order_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table pay_order
-- ----------------------------
ALTER TABLE "public"."pay_order" ADD CONSTRAINT "pay_order_pkey" PRIMARY KEY ("order_id");

-- ----------------------------
-- Indexes structure for table pay_order_rel
-- ----------------------------
CREATE INDEX "pay_order_rel_order_id_idx" ON "public"."pay_order_rel" USING btree (
  "order_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table pay_scene
-- ----------------------------
ALTER TABLE "public"."pay_scene" ADD CONSTRAINT "pay_scene_pkey" PRIMARY KEY ("scene_id");

-- ----------------------------
-- Indexes structure for table platform_admin
-- ----------------------------
CREATE UNIQUE INDEX "platform_admin_account_idx" ON "public"."platform_admin" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "platform_admin_email_idx" ON "public"."platform_admin" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "platform_admin_phone_idx" ON "public"."platform_admin" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table platform_admin
-- ----------------------------
ALTER TABLE "public"."platform_admin" ADD CONSTRAINT "platform_admin_pkey" PRIMARY KEY ("admin_id");

-- ----------------------------
-- Primary Key structure for table platform_config
-- ----------------------------
ALTER TABLE "public"."platform_config" ADD CONSTRAINT "platform_config_pkey" PRIMARY KEY ("config_key");

-- ----------------------------
-- Indexes structure for table platform_server
-- ----------------------------
CREATE UNIQUE INDEX "platform_server_network_ip_idx" ON "public"."platform_server" USING btree (
  "network_ip" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table platform_server
-- ----------------------------
ALTER TABLE "public"."platform_server" ADD CONSTRAINT "platform_server_pkey" PRIMARY KEY ("server_id");

-- ----------------------------
-- Primary Key structure for table upload
-- ----------------------------
ALTER TABLE "public"."upload" ADD CONSTRAINT "upload_pkey" PRIMARY KEY ("upload_id");

-- ----------------------------
-- Indexes structure for table users
-- ----------------------------
CREATE UNIQUE INDEX "users_account_idx" ON "public"."users" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "users_email_idx" ON "public"."users" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "users_phone_idx" ON "public"."users" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "users_wx_openid_idx" ON "public"."users" USING btree (
  "wx_openid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "users_wx_unionid_idx" ON "public"."users" USING btree (
  "wx_unionid" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "users_pkey" PRIMARY KEY ("user_id");

-- ----------------------------
-- Primary Key structure for table users_privacy
-- ----------------------------
ALTER TABLE "public"."users_privacy" ADD CONSTRAINT "users_privacy_pkey" PRIMARY KEY ("user_id");
