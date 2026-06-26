/*
 Navicat Premium Dump SQL

 Source Server         : PostgreSQL
 Source Server Type    : PostgreSQL
 Source Server Version : 180003 (180003)
 Source Host           : 192.168.0.200:5432
 Source Catalog        : admin_bak
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 180003 (180003)
 File Encoding         : 65001

 Date: 26/06/2026 18:05:46
*/


-- ----------------------------
-- Sequence structure for admin_admin_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."admin_admin_id_seq";
CREATE SEQUENCE "public"."admin_admin_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

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
-- Table structure for admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin";
CREATE TABLE "public"."admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "admin_id" int4 NOT NULL DEFAULT nextval('admin_admin_id_seq'::regclass),
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "rel_id" int4 NOT NULL DEFAULT 0,
  "admin_type" int2 NOT NULL,
  "account" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "email" varchar(60) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_super" int2 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."admin"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."admin"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."admin"."rel_id" IS '关联ID。根据scene_id对应不同表';
COMMENT ON COLUMN "public"."admin"."admin_type" IS '类型：0平台 10机构';
COMMENT ON COLUMN "public"."admin"."account" IS '账号';
COMMENT ON COLUMN "public"."admin"."phone" IS '手机';
COMMENT ON COLUMN "public"."admin"."email" IS '邮箱';
COMMENT ON COLUMN "public"."admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."admin"."is_super" IS '超管：0否 1是';
COMMENT ON TABLE "public"."admin" IS '管理员';

-- ----------------------------
-- Records of admin
-- ----------------------------
INSERT INTO "public"."admin" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, 'platform', 0, 0, 'admin', NULL, NULL, '超级管理员', '', 1);

-- ----------------------------
-- Table structure for admin_privacy
-- ----------------------------
DROP TABLE IF EXISTS "public"."admin_privacy";
CREATE TABLE "public"."admin_privacy" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "admin_id" int4 NOT NULL DEFAULT 0,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar
)
;
COMMENT ON COLUMN "public"."admin_privacy"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."admin_privacy"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."admin_privacy"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."admin_privacy"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."admin_privacy"."salt" IS '密码盐';
COMMENT ON TABLE "public"."admin_privacy" IS '管理员隐私';

-- ----------------------------
-- Records of admin_privacy
-- ----------------------------
INSERT INTO "public"."admin_privacy" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB');

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
COMMENT ON TABLE "public"."app" IS 'APP';

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
COMMENT ON TABLE "public"."app_pkg" IS 'APP安装包';

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
COMMENT ON TABLE "public"."auth_action" IS '权限操作';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'adminCreate', '权限管理-管理员-新增', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'adminDelete', '权限管理-管理员-删除', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'adminRead', '权限管理-管理员-查看', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'adminUpdate', '权限管理-管理员-编辑', '');
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
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configCommonRead', '应用配置-常用-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configCommonSave', '应用配置-常用-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-应用配置-常用');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configEmailRead', '插件配置-邮箱-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configEmailSave', '插件配置-邮箱-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configIdCardRead', '插件配置-实名认证-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configIdCardSave', '插件配置-实名认证-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configOneClickRead', '插件配置-一键登录-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configOneClickSave', '插件配置-一键登录-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configOrgRead', '应用配置-机构-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-应用配置-机构');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configOrgSave', '应用配置-机构-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-应用配置-机构');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configPlatformRead', '应用配置-平台-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-应用配置-平台');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configPlatformSave', '应用配置-平台-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-应用配置-平台');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configPushRead', '插件配置-推送-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configPushSave', '插件配置-推送-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configRead', '配置-查看', '可任意读取配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configSave', '配置-保存', '可任意保存配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configSmsRead', '插件配置-短信-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configSmsSave', '插件配置-短信-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configVodRead', '插件配置-视频点播-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configVodSave', '插件配置-视频点播-保存', '只能保存配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configWxRead', '插件配置-微信-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'configWxSave', '插件配置-微信-查看', '只能读取配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
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
COMMENT ON TABLE "public"."auth_action_rel_to_scene" IS '权限操作，权限场景关联（操作可用在哪些场景）';

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminCreate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminDelete', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminUpdate', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'adminUpdate', 'platform');
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
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneCreate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneDelete', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'authSceneUpdate', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configCommonRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configCommonRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configCommonSave', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configCommonSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configEmailRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configEmailSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configIdCardRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configIdCardSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configOneClickRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configOneClickSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configOrgRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configOrgSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configPlatformRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configPlatformSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configPushRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configPushSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configRead', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configSave', 'org');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configSmsRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configSmsSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configVodRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configVodSave', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configWxRead', 'platform');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'configWxSave', 'platform');
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
  "menu_icon" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
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
COMMENT ON COLUMN "public"."auth_menu"."menu_icon" IS '图标';
COMMENT ON COLUMN "public"."auth_menu"."menu_url" IS '链接';
COMMENT ON COLUMN "public"."auth_menu"."extra_data" IS '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}';
COMMENT ON COLUMN "public"."auth_menu"."sort" IS '排序值。从大到小排序';
COMMENT ON TABLE "public"."auth_menu" IS '权限菜单';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '主页', 'platform', 0, 1, 1, '0-1', '-主页', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理', 'platform', 0, 0, 1, '0-2', '-权限管理', 'autoicon-ep-lock', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '场景', 'platform', 2, 1, 2, '0-2-3', '-权限管理-场景', 'autoicon-ep-flag', '/auth/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "场景"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '操作', 'platform', 2, 1, 2, '0-2-4', '-权限管理-操作', 'autoicon-ep-coordinate', '/auth/action', '{"i18n": {"title": {"en": "Action", "zh-cn": "操作"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '菜单', 'platform', 2, 1, 2, '0-2-5', '-权限管理-菜单', 'autoicon-ep-menu', '/auth/menu', '{"i18n": {"title": {"en": "Menu", "zh-cn": "菜单"}}}', 30);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '角色', 'platform', 2, 1, 2, '0-2-6', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '管理员', 'platform', 2, 1, 2, '0-2-7', '-权限管理-管理员', 'autoicon-ep-avatar', '/admin/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "管理员"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '系统管理', 'platform', 0, 0, 1, '0-8', '-系统管理', 'autoicon-ep-platform', '', '{"i18n": {"title": {"en": "System Manage", "zh-cn": "系统管理"}}}', 20);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '配置中心', 'platform', 8, 0, 2, '0-8-9', '-系统管理-配置中心', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '上传配置', 'platform', 9, 1, 3, '0-8-9-10', '-系统管理-配置中心-上传配置', 'autoicon-ep-upload', '/upload/upload', '{"i18n": {"title": {"en": "Upload", "zh-cn": "上传配置"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '支付管理', 'platform', 9, 0, 3, '0-8-9-11', '-系统管理-配置中心-支付管理', 'autoicon-ep-coin', '', '{"i18n": {"title": {"en": "Pay Manage", "zh-cn": "支付管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '支付配置', 'platform', 11, 1, 4, '0-8-9-11-12', '-系统管理-配置中心-支付管理-支付配置', 'autoicon-ep-money', '/pay/pay', '{"i18n": {"title": {"en": "Pay", "zh-cn": "支付配置"}}}', 50);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '支付场景', 'platform', 11, 1, 4, '0-8-9-11-13', '-系统管理-配置中心-支付管理-支付场景', 'autoicon-ep-guide', '/pay/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "支付场景"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '支付通道', 'platform', 11, 1, 4, '0-8-9-11-14', '-系统管理-配置中心-支付管理-支付通道', 'autoicon-ep-connection', '/pay/channel', '{"i18n": {"title": {"en": "Channel", "zh-cn": "支付通道"}}}', 150);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '插件配置', 'platform', 9, 1, 3, '0-8-9-15', '-系统管理-配置中心-插件配置', 'autoicon-ep-ticket', '/config/config/plugin', '{"i18n": {"title": {"en": "Plugin Config", "zh-cn": "插件配置"}}}', 150);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '应用配置', 'platform', 9, 1, 3, '0-8-9-16', '-系统管理-配置中心-应用配置', 'autoicon-ep-set-up', '/config/config/app', '{"i18n": {"title": {"en": "APP Config", "zh-cn": "应用配置"}}}', 200);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, 'APP管理', 'platform', 8, 0, 2, '0-8-17', '-系统管理-APP管理', 'autoicon-ep-suitcase-line', '', '{"i18n": {"title": {"en": "APP Manage", "zh-cn": "APP管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, 'APP', 'platform', 17, 1, 3, '0-8-17-18', '-系统管理-APP管理-APP', 'autoicon-ep-apple', '/app/app', '{"i18n": {"title": {"en": "App", "zh-cn": "APP"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '安装包', 'platform', 17, 1, 3, '0-8-17-19', '-系统管理-APP管理-安装包', 'autoicon-ep-box', '/app/pkg', '{"i18n": {"title": {"en": "Pkg", "zh-cn": "安装包"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '用户管理', 'platform', 0, 0, 1, '0-20', '-用户管理', 'autoicon-ep-user-filled', '', '{"i18n": {"title": {"en": "User Manage", "zh-cn": "用户管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '用户', 'platform', 20, 1, 2, '0-20-21', '-用户管理-用户', 'autoicon-ep-user', '/users/users', '{"i18n": {"title": {"en": "Users", "zh-cn": "用户"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '机构管理', 'platform', 0, 0, 1, '0-22', '-机构管理', 'autoicon-ep-office-building', '', '{"i18n": {"title": {"en": "Org Manage", "zh-cn": "机构管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '机构', 'platform', 22, 1, 2, '0-22-23', '-机构管理-机构', 'autoicon-ep-school', '/org/org', '{"i18n": {"title": {"en": "Org", "zh-cn": "机构"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 101, '主页', 'org', 0, 1, 1, '0-101', '-主页', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 102, '权限管理', 'org', 0, 0, 1, '0-102', '-权限管理', 'autoicon-ep-menu', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 103, '角色', 'org', 26, 1, 2, '0-102-103', '-权限管理-角色', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 104, '管理员', 'org', 26, 1, 2, '0-102-104', '-权限管理-管理员', 'autoicon-ep-avatar', '/admin/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "管理员"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 105, '配置中心', 'org', 0, 0, 1, '0-105', '-配置中心', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 20);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 106, '应用配置', 'org', 29, 1, 2, '0-105-106', '-配置中心-应用配置', 'autoicon-ep-set-up', '/config/config/app', '{"i18n": {"title": {"en": "APP Config", "zh-cn": "应用配置"}}}', 200);

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
COMMENT ON TABLE "public"."auth_role" IS '权限角色';

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_of_admin";
CREATE TABLE "public"."auth_role_rel_of_admin" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "admin_id" int4 NOT NULL,
  "role_id" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."auth_role_rel_of_admin"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_admin"."role_id" IS '角色ID';
COMMENT ON TABLE "public"."auth_role_rel_of_admin" IS '管理员，权限角色关联（管理员包含哪些角色）';

-- ----------------------------
-- Records of auth_role_rel_of_admin
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
COMMENT ON TABLE "public"."auth_role_rel_to_action" IS '权限角色，权限操作关联（角色包含哪些操作）';

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
COMMENT ON TABLE "public"."auth_role_rel_to_menu" IS '权限角色，权限菜单关联（角色包含哪些菜单）';

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
COMMENT ON TABLE "public"."auth_scene" IS '权限场景';

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'app', 'APP', '{"token_config": {"is_ip": 0, "is_unique": 0, "sign_type": "HS256", "token_type": 0, "active_time": 0, "expire_time": 604800, "private_key": "任意字符串"}}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'org', '机构后台', '{"token_config": {"is_ip": 1, "is_unique": 1, "sign_type": "RS256", "public_key": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0nT4zSS2O+2yib+gdBm0\nW0kU2AL6felZFp5U6B/ySeJnM8UE+fZwytgPuND5Z07khxtHR/YYH7huPGZ/fAgh\nZHmNE5K5phMI5eETwPjk3RDeyYAyOosrKr3SAjjEQxJBISvBEillH4bKjaa4WF5/\n+nGcp7f4e49caW/CfwuC2ZVrvySCPf1lR8o7/4Zz/hWUwgsEd/crR7ojgt+rbPeE\n7+Cz11sZUZaMipTqsU3RVhbwtMyLdkos6KsYW7TZEK0VYt94/1XJQBUEjtCdDpS7\n0XNF8ENpQPtuQdYE6/y+Jku8T9pqQQq/SOL6uPgsn6zJQ41u/l2AhG0i5GxYD86C\n5wIDAQAB\n-----END PUBLIC KEY-----", "token_type": 0, "active_time": 3600, "expire_time": 14400, "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDSdPjNJLY77bKJ\nv6B0GbRbSRTYAvp96VkWnlToH/JJ4mczxQT59nDK2A+40PlnTuSHG0dH9hgfuG48\nZn98CCFkeY0TkrmmEwjl4RPA+OTdEN7JgDI6iysqvdICOMRDEkEhK8ESKWUfhsqN\nprhYXn/6cZynt/h7j1xpb8J/C4LZlWu/JII9/WVHyjv/hnP+FZTCCwR39ytHuiOC\n36ts94Tv4LPXWxlRloyKlOqxTdFWFvC0zIt2SizoqxhbtNkQrRVi33j/VclAFQSO\n0J0OlLvRc0XwQ2lA+25B1gTr/L4mS7xP2mpBCr9I4vq4+CyfrMlDjW7+XYCEbSLk\nbFgPzoLnAgMBAAECggEAGolMO9WmsrzAd9T1Pt5k2uPGoIwTmJ+9L3hsXU515vII\nsELl4zy7MSB4LwYOhIOylgSPAthZZ1qCb9Q+u91slHYtHywvg2zAAPhV3M2lUeiI\nJuEmtDILEdsYaVZODOT22F9je05D5WtCDAVbFi1oNqRvq8grKS1E6jiAzjMd3yBY\n5AgUUP8sbS7BDdPus2t3mCAXqdtFkxn8wo/4WdMV6vG4p9p+a8dIoiRYHNBIw4sU\nCYPiE7q52tVqVl10ShrJQlvoyDvmJam4inbl8ZWtbQLUsQxfzoEUfuuC0mYc2pES\n1kp1pWfJNc5JtAXXZ/k9F4jLvjMp9KJximOG+E5tmQKBgQD5RwG8xu7s2JcocogD\nuJjWfLz/4ab7Zs2NReCX6jmb522d1TqyiU9ilc0XBz8gDNeMwnOfWIyGR3Vtuub+\nU56Y9/q+IMZ0ewdrUTR2NADZdqLH1ViVGacnnmJXJ+30Z8eUO0GCU7++DSQ1u8xh\ng0mmrj6++xHsYJM3bCxBstNzGQKBgQDYIfOUIK+JKvYZ5idUrFWZnnVSdLu8IJHi\nA0osaMB9VNtAdqrPVIA3L9AbIR1/la3gb3ILP0hIM7glt4i78WTwVrh0qkaz/8dw\nX3EL5u9OAMecVFn4+gZon4RqbzjdtHso87v2GrIZ88eVOWjXRq93gWAe17G+noyS\n44PgB7Zl/wKBgC3bOh6YGevIDEaMiyjkFHmgiMQppqYoyzdp218W33ImqKuYRiwB\nxnDETe4mjx4+PojOXKa7i15IVvnQoB25FDvfomjHbrqOx1aeoZ/9AQsAIAHS5XDI\nP0+yezS9S7DiRnymSe7HqUY09KxN19M4a5wWAcTwOuPZADv50kpjszJBAoGAS0io\nO7SW8ESSrLrKgGf2+SeE3k/jBMijiAJ1V7q1MfLY3D95h/Z7Ir340zpZuBM/Gao4\nI0rLtrqtLhYb/rs62aybW6fkMNarda0JB4hNWvJSlVWccWlFyjOmQBy1xiQTslQT\n6Mmrt/Z+UrBIoJPyksHx5UxkkW1QsemmCecl1akCgYAys4PNRhdToxS5WDF7Tddg\n8ghhobuwUCP2pxNYakX9HeHZsAjwwQhx6vGJqZPs5hh8HRinGWSvKyVXaUGVN+5b\nFjMz/rhO9vTCZJS3aJSXxr0PFTbpP/AZSXwCBxjp+uEFTD8GJRX7/6+wlTk7+uTj\nl7klp31noFtzz+onGSmqvA==\n-----END PRIVATE KEY-----"}}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 'platform', '平台后台', '{"token_config": {"is_ip": 1, "is_unique": 1, "sign_type": "ES256", "public_key": "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42\nhF5PnYenbWprk1MQIzT2+V4rRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\n-----END PUBLIC KEY-----", "token_type": 0, "active_time": 3600, "expire_time": 604800, "private_key": "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIKvYPRtCqy9MI/yhx4L4+Sog/5lntHbuwxPg/JI0qW6LoAoGCCqGSM49\nAwEHoUQDQgAEvqHRsI0W+SABR4hYOXrbXR4EiC42hF5PnYenbWprk1MQIzT2+V4r\nRJc7nyXQ/ntRK7B/rN9mpc3Vot02bwp02w==\n-----END EC PRIVATE KEY-----"}}', '');

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS "public"."config";
CREATE TABLE "public"."config" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "scene_id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "rel_id" int4 NOT NULL DEFAULT 0,
  "config_key" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "config_value" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text
)
;
COMMENT ON COLUMN "public"."config"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."config"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."config"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."config"."rel_id" IS '关联ID。根据scene_id对应不同表';
COMMENT ON COLUMN "public"."config"."config_key" IS '配置键';
COMMENT ON COLUMN "public"."config"."config_value" IS '配置值';
COMMENT ON TABLE "public"."config" IS '配置';

-- ----------------------------
-- Records of config
-- ----------------------------
INSERT INTO "public"."config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platform', 0, 'email_code', '{"subject":"您的验证码","template":"验证码：{code}\n说明：\n1. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。\n2. 出于安全考虑，请不要将此验证码分享给任何人。"}');
INSERT INTO "public"."config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platform', 0, 'email_of_common', '{"from_email":"xxxxxxxx@qq.com","password":"xxxxxxxx","smtp_host":"smtp.qq.com","smtp_port":"465"}');
INSERT INTO "public"."config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platform', 0, 'email_type', 'email_of_common');
INSERT INTO "public"."config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platform', 0, 'id_card_of_aliyun', '{"appcode":"appcode","url":"http://idcard.market.alicloudapi.com/lianzhuo/idcard"}');
INSERT INTO "public"."config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'platform', 0, 'id_card_type', 'id_card_of_aliyun');

-- ----------------------------
-- Table structure for org
-- ----------------------------
DROP TABLE IF EXISTS "public"."org";
CREATE TABLE "public"."org" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "org_id" int4 NOT NULL DEFAULT nextval('org_org_id_seq'::regclass),
  "org_name" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "org_type" int2 NOT NULL DEFAULT 10
)
;
COMMENT ON COLUMN "public"."org"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."org"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."org"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."org"."org_id" IS '机构ID';
COMMENT ON COLUMN "public"."org"."org_name" IS '机构名称';
COMMENT ON COLUMN "public"."org"."org_type" IS '类型：10默认';
COMMENT ON TABLE "public"."org" IS '机构';

-- ----------------------------
-- Records of org
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
COMMENT ON COLUMN "public"."pay"."pay_config" IS '配置。JSON格式，根据类型设置';
COMMENT ON COLUMN "public"."pay"."pay_rate" IS '费率';
COMMENT ON COLUMN "public"."pay"."total_amount" IS '总额';
COMMENT ON COLUMN "public"."pay"."balance" IS '余额';
COMMENT ON COLUMN "public"."pay"."remark" IS '备注';
COMMENT ON TABLE "public"."pay" IS '支付';

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
COMMENT ON TABLE "public"."pay_channel" IS '支付通道';

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
  "order_type" int2 NOT NULL DEFAULT 0,
  "rel_id" int4 NOT NULL DEFAULT 0,
  "pay_id" int4 NOT NULL DEFAULT 0,
  "channel_id" int4 NOT NULL DEFAULT 0,
  "pay_type" int2 NOT NULL DEFAULT 0,
  "amount" numeric(10,2) NOT NULL DEFAULT 0.00,
  "pay_status" int2 NOT NULL DEFAULT 0,
  "pay_at" timestamp(6),
  "pay_rate" numeric(4,4) NOT NULL DEFAULT 0.0000,
  "third_order_no" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "ext_data" varchar(120) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "order_ip" varchar(15) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."pay_order"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."pay_order"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."pay_order"."order_id" IS '订单ID';
COMMENT ON COLUMN "public"."pay_order"."order_no" IS '订单号';
COMMENT ON COLUMN "public"."pay_order"."order_type" IS '订单类型：0默认';
COMMENT ON COLUMN "public"."pay_order"."rel_id" IS '关联ID。根据order_type对应不同表';
COMMENT ON COLUMN "public"."pay_order"."pay_id" IS '支付ID';
COMMENT ON COLUMN "public"."pay_order"."channel_id" IS '通道ID';
COMMENT ON COLUMN "public"."pay_order"."pay_type" IS '类型：0支付宝 1微信';
COMMENT ON COLUMN "public"."pay_order"."amount" IS '实付金额';
COMMENT ON COLUMN "public"."pay_order"."pay_status" IS '状态：0未付款 1已付款';
COMMENT ON COLUMN "public"."pay_order"."pay_at" IS '支付时间';
COMMENT ON COLUMN "public"."pay_order"."pay_rate" IS '费率';
COMMENT ON COLUMN "public"."pay_order"."third_order_no" IS '第三方订单号';
COMMENT ON COLUMN "public"."pay_order"."ext_data" IS '扩展数据';
COMMENT ON COLUMN "public"."pay_order"."order_ip" IS '订单IP';
COMMENT ON TABLE "public"."pay_order" IS '支付订单';

-- ----------------------------
-- Records of pay_order
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
COMMENT ON TABLE "public"."pay_scene" IS '支付场景';

-- ----------------------------
-- Records of pay_scene
-- ----------------------------

-- ----------------------------
-- Table structure for server
-- ----------------------------
DROP TABLE IF EXISTS "public"."server";
CREATE TABLE "public"."server" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "server_id" int4 NOT NULL DEFAULT nextval('platform_server_server_id_seq'::regclass),
  "network_ip" varchar(15) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "local_ip" varchar(15) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."server"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."server"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."server"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."server"."server_id" IS '服务器ID';
COMMENT ON COLUMN "public"."server"."network_ip" IS '公网IP';
COMMENT ON COLUMN "public"."server"."local_ip" IS '内网IP';
COMMENT ON TABLE "public"."server" IS '服务器';

-- ----------------------------
-- Records of server
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
COMMENT ON TABLE "public"."upload" IS '上传';

-- ----------------------------
-- Records of upload
-- ----------------------------
INSERT INTO "public"."upload" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 0, '{"sign_key": "signKey", "is_cluster": 1, "server_list": [], "is_same_server": 0}', 1, '此项目自带文件上传下载功能，可直接部署成文件服务器使用');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "user_id" int4 NOT NULL DEFAULT nextval('users_user_id_seq'::regclass),
  "account" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "phone" varchar(20) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "email" varchar(60) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "gender" int2 NOT NULL DEFAULT 0,
  "birthday" date,
  "address" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "wx_openid" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "wx_unionid" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
COMMENT ON COLUMN "public"."users"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."users"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."users"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."users"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."users"."account" IS '账号';
COMMENT ON COLUMN "public"."users"."phone" IS '手机';
COMMENT ON COLUMN "public"."users"."email" IS '邮箱';
COMMENT ON COLUMN "public"."users"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."users"."avatar" IS '头像';
COMMENT ON COLUMN "public"."users"."gender" IS '性别：0未设置 1男 2女';
COMMENT ON COLUMN "public"."users"."birthday" IS '生日';
COMMENT ON COLUMN "public"."users"."address" IS '地址';
COMMENT ON COLUMN "public"."users"."wx_openid" IS '微信openid';
COMMENT ON COLUMN "public"."users"."wx_unionid" IS '微信unionid';
COMMENT ON TABLE "public"."users" IS '用户（postgresql中user是关键字，使用需要加双引号。程序中考虑与mysql通用，故命名成users）';

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
COMMENT ON TABLE "public"."users_privacy" IS '用户隐私';

-- ----------------------------
-- Records of users_privacy
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."admin_admin_id_seq"
OWNED BY "public"."admin"."admin_id";
SELECT setval('"public"."admin_admin_id_seq"', 1, false);

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
ALTER SEQUENCE "public"."platform_server_server_id_seq"
OWNED BY "public"."server"."server_id";
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
-- Indexes structure for table admin
-- ----------------------------
CREATE UNIQUE INDEX "admin_account_admin_type_idx" ON "public"."admin" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "admin_type" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE INDEX "admin_admin_type_idx" ON "public"."admin" USING btree (
  "admin_type" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "admin_email_admin_type_idx" ON "public"."admin" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "admin_type" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "admin_phone_admin_type_idx" ON "public"."admin" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
  "admin_type" "pg_catalog"."int2_ops" ASC NULLS LAST
);
CREATE INDEX "admin_rel_id_idx" ON "public"."admin" USING btree (
  "rel_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "admin_scene_id_idx" ON "public"."admin" USING btree (
  "scene_id" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table admin
-- ----------------------------
ALTER TABLE "public"."admin" ADD CONSTRAINT "admin_pkey" PRIMARY KEY ("admin_id");

-- ----------------------------
-- Primary Key structure for table admin_privacy
-- ----------------------------
ALTER TABLE "public"."admin_privacy" ADD CONSTRAINT "users_privacy_copy1_pkey" PRIMARY KEY ("admin_id");

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
-- Indexes structure for table auth_role_rel_of_admin
-- ----------------------------
CREATE INDEX "auth_role_rel_of_admin_admin_id_idx" ON "public"."auth_role_rel_of_admin" USING btree (
  "admin_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_of_admin_role_id_idx" ON "public"."auth_role_rel_of_admin" USING btree (
  "role_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_of_admin
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_of_admin" ADD CONSTRAINT "auth_role_rel_of_org_admin_pkey" PRIMARY KEY ("admin_id", "role_id");

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
-- Primary Key structure for table config
-- ----------------------------
ALTER TABLE "public"."config" ADD CONSTRAINT "config_pkey" PRIMARY KEY ("scene_id", "rel_id", "config_key");

-- ----------------------------
-- Primary Key structure for table org
-- ----------------------------
ALTER TABLE "public"."org" ADD CONSTRAINT "org_pkey" PRIMARY KEY ("org_id");

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
CREATE INDEX "pay_order_rel_id_idx" ON "public"."pay_order" USING btree (
  "rel_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "pay_order_third_order_no_idx" ON "public"."pay_order" USING btree (
  "third_order_no" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table pay_order
-- ----------------------------
ALTER TABLE "public"."pay_order" ADD CONSTRAINT "pay_order_pkey" PRIMARY KEY ("order_id");

-- ----------------------------
-- Primary Key structure for table pay_scene
-- ----------------------------
ALTER TABLE "public"."pay_scene" ADD CONSTRAINT "pay_scene_pkey" PRIMARY KEY ("scene_id");

-- ----------------------------
-- Indexes structure for table server
-- ----------------------------
CREATE UNIQUE INDEX "platform_server_network_ip_idx" ON "public"."server" USING btree (
  "network_ip" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table server
-- ----------------------------
ALTER TABLE "public"."server" ADD CONSTRAINT "platform_server_pkey" PRIMARY KEY ("server_id");

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
