/*
 Navicat Premium Data Transfer

 Source Server         : 本地-Postgresql16
 Source Server Type    : PostgreSQL
 Source Server Version : 160002 (160002)
 Source Host           : 192.168.2.200:5432
 Source Catalog        : admin
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 160002 (160002)
 File Encoding         : 65001

 Date: 07/05/2024 17:33:12
*/


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
-- Sequence structure for auth_scene_scene_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_scene_scene_id_seq";
CREATE SEQUENCE "public"."auth_scene_scene_id_seq" 
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
-- Sequence structure for user_user_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_user_user_id_seq";
CREATE SEQUENCE "public"."user_user_user_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action";
CREATE TABLE "public"."auth_action" (
  "action_id" int4 NOT NULL DEFAULT nextval('auth_action_action_id_seq'::regclass),
  "action_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "action_code" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_action"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action"."action_name" IS '名称';
COMMENT ON COLUMN "public"."auth_action"."action_code" IS '标识';
COMMENT ON COLUMN "public"."auth_action"."remark" IS '备注';
COMMENT ON COLUMN "public"."auth_action"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_action"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO "public"."auth_action" VALUES (1, '权限管理-场景-查看', 'authSceneLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (2, '权限管理-场景-新增', 'authSceneCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (3, '权限管理-场景-编辑', 'authSceneUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (4, '权限管理-场景-删除', 'authSceneDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (5, '权限操作-查看', 'authActionLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (6, '权限操作-新增', 'authActionCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (7, '权限操作-编辑', 'authActionUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (8, '权限操作-删除', 'authActionDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (9, '权限菜单-查看', 'authMenuLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (10, '权限菜单-新增', 'authMenuCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (11, '权限菜单-编辑', 'authMenuUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (12, '权限菜单-删除', 'authMenuDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (13, '权限角色-查看', 'authRoleLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (14, '权限角色-新增', 'authRoleCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (15, '权限角色-编辑', 'authRoleUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (16, '权限角色-删除', 'authRoleDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (17, '平台管理员-查看', 'platformAdminLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (18, '平台管理员-新增', 'platformAdminCreate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (19, '平台管理员-编辑', 'platformAdminUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (20, '平台管理员-删除', 'platformAdminDelete', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (21, '平台配置-查看', 'platformConfigLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (22, '平台配置-保存', 'platformConfigSave', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (23, '用户-查看', 'userUserLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (24, '用户-编辑', 'userUserUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action_rel_to_scene";
CREATE TABLE "public"."auth_action_rel_to_scene" (
  "action_id" int4 NOT NULL DEFAULT 0,
  "scene_id" int4 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_action_rel_to_scene
-- ----------------------------
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (1, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (2, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (3, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (4, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (5, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (6, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (7, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (8, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (9, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (10, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (11, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (12, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (13, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (14, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (15, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (16, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (17, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (18, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (19, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (20, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (21, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (22, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (23, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action_rel_to_scene" VALUES (24, 1, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_menu";
CREATE TABLE "public"."auth_menu" (
  "menu_id" int4 NOT NULL DEFAULT nextval('auth_menu_menu_id_seq'::regclass),
  "menu_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_id" int4 NOT NULL DEFAULT 0,
  "pid" int4 NOT NULL DEFAULT 0,
  "level" int2 NOT NULL DEFAULT 0,
  "id_path" text COLLATE "pg_catalog"."default",
  "menu_icon" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "menu_url" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "extra_data" json,
  "sort" int2 NOT NULL DEFAULT 50,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_menu"."menu_id" IS '菜单ID';
COMMENT ON COLUMN "public"."auth_menu"."menu_name" IS '名称';
COMMENT ON COLUMN "public"."auth_menu"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_menu"."pid" IS '父ID';
COMMENT ON COLUMN "public"."auth_menu"."level" IS '层级';
COMMENT ON COLUMN "public"."auth_menu"."id_path" IS '层级路径';
COMMENT ON COLUMN "public"."auth_menu"."menu_icon" IS '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}';
COMMENT ON COLUMN "public"."auth_menu"."menu_url" IS '链接';
COMMENT ON COLUMN "public"."auth_menu"."extra_data" IS '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}';
COMMENT ON COLUMN "public"."auth_menu"."sort" IS '排序值。从小到大排序，默认50，范围0-100';
COMMENT ON COLUMN "public"."auth_menu"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_menu"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_menu"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO "public"."auth_menu" VALUES (1, '主页', 1, 0, 1, '0-1', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 0, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (2, '权限管理', 1, 0, 1, '0-2', 'autoicon-ep-lock', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 90, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (3, '场景', 1, 2, 2, '0-2-3', 'autoicon-ep-flag', '/auth/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "场景"}}}', 100, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (4, '操作', 1, 2, 2, '0-2-4', 'autoicon-ep-coordinate', '/auth/action', '{"i18n": {"title": {"en": "Action", "zh-cn": "操作"}}}', 90, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (5, '菜单', 1, 2, 2, '0-2-5', 'autoicon-ep-menu', '/auth/menu', '{"i18n": {"title": {"en": "Menu", "zh-cn": "菜单"}}}', 80, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (6, '角色', 1, 2, 2, '0-2-6', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 70, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (7, '平台管理员', 1, 2, 2, '0-2-7', 'vant-manager-o', '/platform/admin', '{"i18n": {"title": {"en": "Platform Admin", "zh-cn": "平台管理员"}}}', 60, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (8, '系统管理', 1, 0, 1, '0-8', 'autoicon-ep-platform', '', '{"i18n": {"title": {"en": "System Manage", "zh-cn": "系统管理"}}}', 85, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (9, '配置中心', 1, 8, 2, '0-8-9', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 100, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (10, '平台配置', 1, 9, 3, '0-8-9-10', '', '/platform/config/platform', '{"i18n": {"title": {"en": "Platform Config", "zh-cn": "平台配置"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (11, '插件配置', 1, 9, 3, '0-8-9-11', '', '/platform/config/plugin', '{"i18n": {"title": {"en": "Plugin Config", "zh-cn": "插件配置"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (12, '用户管理', 1, 0, 1, '0-12', 'vant-friends', '', '{"i18n": {"title": {"en": "User Manage", "zh-cn": "用户管理"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_menu" VALUES (13, '用户', 1, 12, 2, '0-12-13', 'vant-user-o', '/user/user', '{"i18n": {"title": {"en": "User", "zh-cn": "用户"}}}', 50, 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role";
CREATE TABLE "public"."auth_role" (
  "role_id" int4 NOT NULL DEFAULT nextval('auth_role_role_id_seq'::regclass),
  "role_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_id" int4 NOT NULL DEFAULT 0,
  "table_id" int4 NOT NULL DEFAULT 0,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role"."role_name" IS '名称';
COMMENT ON COLUMN "public"."auth_role"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_role"."table_id" IS '关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建';
COMMENT ON COLUMN "public"."auth_role"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_role"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_of_platform_admin";
CREATE TABLE "public"."auth_role_rel_of_platform_admin" (
  "role_id" int4 NOT NULL DEFAULT 0,
  "admin_id" int4 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_action";
CREATE TABLE "public"."auth_role_rel_to_action" (
  "role_id" int4 NOT NULL DEFAULT 0,
  "action_id" int4 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_menu";
CREATE TABLE "public"."auth_role_rel_to_menu" (
  "role_id" int4 NOT NULL DEFAULT 0,
  "menu_id" int4 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."menu_id" IS '菜单ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_scene";
CREATE TABLE "public"."auth_scene" (
  "scene_id" int4 NOT NULL DEFAULT nextval('auth_scene_scene_id_seq'::regclass),
  "scene_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_code" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_config" json NOT NULL,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_scene"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_scene"."scene_name" IS '名称';
COMMENT ON COLUMN "public"."auth_scene"."scene_code" IS '标识';
COMMENT ON COLUMN "public"."auth_scene"."scene_config" IS '配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}';
COMMENT ON COLUMN "public"."auth_scene"."remark" IS '备注';
COMMENT ON COLUMN "public"."auth_scene"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_scene"."created_at" IS '创建时间';

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO "public"."auth_scene" VALUES (1, '平台后台', 'platform', '{"signKey": "www.admin.com_platform", "signType": "HS256", "expireTime": 14400}', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_scene" VALUES (2, 'APP', 'app', '{"signKey": "www.admin.com_app", "signType": "HS256", "expireTime": 604800}', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_admin";
CREATE TABLE "public"."platform_admin" (
  "admin_id" int4 NOT NULL DEFAULT nextval('platform_admin_admin_id_seq'::regclass),
  "phone" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_admin"."admin_id" IS '管理员ID';
COMMENT ON COLUMN "public"."platform_admin"."phone" IS '手机';
COMMENT ON COLUMN "public"."platform_admin"."account" IS '账号';
COMMENT ON COLUMN "public"."platform_admin"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."platform_admin"."salt" IS '密码盐';
COMMENT ON COLUMN "public"."platform_admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."platform_admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."platform_admin"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."platform_admin"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_admin"."created_at" IS '创建时间';

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO "public"."platform_admin" VALUES (1, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB', '超级管理员', 'http://JB.Admin.com/common/20240106/1704522339892_31917913.png', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_config";
CREATE TABLE "public"."platform_config" (
  "config_key" varchar(60) COLLATE "pg_catalog"."default" NOT NULL,
  "config_value" text COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::text,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_config"."config_key" IS '配置Key';
COMMENT ON COLUMN "public"."platform_config"."config_value" IS '配置值';
COMMENT ON COLUMN "public"."platform_config"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_config"."created_at" IS '创建时间';

-- ----------------------------
-- Records of platform_config
-- ----------------------------
INSERT INTO "public"."platform_config" VALUES ('idCardOfAliyunAppcode', 'appcode', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('idCardOfAliyunHost', 'http://idcard.market.alicloudapi.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('idCardOfAliyunPath', '/lianzhuo/idcard', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('idCardType', 'idCardOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfAliAppId', 'appId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfAliNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfAli', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfAliOpAppId', 'opAppId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfAliPrivateKey', '****************', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfAliPublicKey', '****************', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxApiV3Key', '********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxAppId', 'appId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxMchid', 'mchId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfWx', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxPrivateKey', '-----BEGIN RSA PRIVATE KEY-----
****************************************************************
-----END RSA PRIVATE KEY-----', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('payOfWxSerialNo', '********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxAndroidAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxAndroidSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxHost', 'https://api.tpns.tencent.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxIosAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxIosSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxMacOSAccessID', 'accessID', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushOfTxMacOSSecretKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('pushType', 'pushOfTx', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsOfAliyunAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsOfAliyunAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsOfAliyunEndpoint', 'dysmsapi.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsOfAliyunSignName', 'JB Admin', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsOfAliyunTemplateCode', 'SMS_********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('smsType', 'smsOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssBucket', 'bucket', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssCallbackUrl', 'http://JB.Admin.com/upload/notify/uploadOfAliyunOss', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssEndpoint', 'sts.cn-hangzhou.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssHost', 'https://oss-cn-hangzhou.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfAliyunOssRoleArn', 'acs:ram::********:role/********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfLocalFileSaveDir', '../public/', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfLocalFileUrlPrefix', 'http://JB.Admin.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfLocalSignKey', 'secretKey', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadOfLocalUrl', 'http://JB.Admin.com/upload/upload', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('uploadType', 'uploadOfLocal', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('vodOfAliyunAccessKeyId', 'accessKeyId', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('vodOfAliyunAccessKeySecret', 'accessKeySecret', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('vodOfAliyunEndpoint', 'sts.cn-shanghai.aliyuncs.com', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('vodOfAliyunRoleArn', 'acs:ram::********:role/********', '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."platform_config" VALUES ('vodType', 'vodOfAliyun', '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_server
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_server";
CREATE TABLE "public"."platform_server" (
  "server_id" int4 NOT NULL DEFAULT nextval('platform_server_server_id_seq'::regclass),
  "network_ip" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "local_ip" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_server"."server_id" IS '服务器ID';
COMMENT ON COLUMN "public"."platform_server"."network_ip" IS '公网IP';
COMMENT ON COLUMN "public"."platform_server"."local_ip" IS '内网IP';
COMMENT ON COLUMN "public"."platform_server"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."platform_server"."created_at" IS '创建时间';

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for user_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_user";
CREATE TABLE "public"."user_user" (
  "user_id" int4 NOT NULL DEFAULT nextval('user_user_user_id_seq'::regclass),
  "phone" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "gender" int2 NOT NULL DEFAULT 0,
  "birthday" date,
  "address" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "id_card_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "id_card_no" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "open_id_of_wx" varchar(128) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "union_id_of_wx" varchar(64) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying
)
;
COMMENT ON COLUMN "public"."user_user"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."user_user"."phone" IS '手机';
COMMENT ON COLUMN "public"."user_user"."account" IS '账号';
COMMENT ON COLUMN "public"."user_user"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."user_user"."salt" IS '密码盐';
COMMENT ON COLUMN "public"."user_user"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."user_user"."avatar" IS '头像';
COMMENT ON COLUMN "public"."user_user"."gender" IS '性别：0未设置 1男 2女';
COMMENT ON COLUMN "public"."user_user"."birthday" IS '生日';
COMMENT ON COLUMN "public"."user_user"."address" IS '详细地址';
COMMENT ON COLUMN "public"."user_user"."id_card_name" IS '身份证姓名';
COMMENT ON COLUMN "public"."user_user"."id_card_no" IS '身份证号码';
COMMENT ON COLUMN "public"."user_user"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."user_user"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."user_user"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."user_user"."open_id_of_wx" IS '微信openId';
COMMENT ON COLUMN "public"."user_user"."union_id_of_wx" IS '微信unionId';

-- ----------------------------
-- Records of user_user
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_action_action_id_seq"
OWNED BY "public"."auth_action"."action_id";
SELECT setval('"public"."auth_action_action_id_seq"', 24, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_menu_menu_id_seq"
OWNED BY "public"."auth_menu"."menu_id";
SELECT setval('"public"."auth_menu_menu_id_seq"', 13, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_role_role_id_seq"
OWNED BY "public"."auth_role"."role_id";
SELECT setval('"public"."auth_role_role_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_scene_scene_id_seq"
OWNED BY "public"."auth_scene"."scene_id";
SELECT setval('"public"."auth_scene_scene_id_seq"', 2, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_admin_admin_id_seq"
OWNED BY "public"."platform_admin"."admin_id";
SELECT setval('"public"."platform_admin_admin_id_seq"', 1, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_server_server_id_seq"
OWNED BY "public"."platform_server"."server_id";
SELECT setval('"public"."platform_server_server_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_user_user_id_seq"
OWNED BY "public"."user_user"."user_id";
SELECT setval('"public"."user_user_user_id_seq"', 1, false);

-- ----------------------------
-- Indexes structure for table auth_action
-- ----------------------------
CREATE UNIQUE INDEX "auth_action_action_code_idx" ON "public"."auth_action" USING btree (
  "action_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_action
-- ----------------------------
ALTER TABLE "public"."auth_action" ADD CONSTRAINT "auth_action_pkey" PRIMARY KEY ("action_id");

-- ----------------------------
-- Indexes structure for table auth_action_rel_to_scene
-- ----------------------------
CREATE INDEX "auth_action_rel_to_scene_action_id_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "action_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_action_rel_to_scene_scene_id_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "scene_id" "pg_catalog"."int4_ops" ASC NULLS LAST
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
  "scene_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_menu
-- ----------------------------
ALTER TABLE "public"."auth_menu" ADD CONSTRAINT "auth_menu_pkey" PRIMARY KEY ("menu_id");

-- ----------------------------
-- Indexes structure for table auth_role
-- ----------------------------
CREATE INDEX "auth_role_scene_id_idx" ON "public"."auth_role" USING btree (
  "scene_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_table_id_idx" ON "public"."auth_role" USING btree (
  "table_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role
-- ----------------------------
ALTER TABLE "public"."auth_role" ADD CONSTRAINT "auth_role_pkey" PRIMARY KEY ("role_id");

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
ALTER TABLE "public"."auth_role_rel_of_platform_admin" ADD CONSTRAINT "auth_role_rel_of_platform_admin_pkey" PRIMARY KEY ("role_id", "admin_id");

-- ----------------------------
-- Indexes structure for table auth_role_rel_to_action
-- ----------------------------
CREATE INDEX "auth_role_rel_to_action_action_id_idx" ON "public"."auth_role_rel_to_action" USING btree (
  "action_id" "pg_catalog"."int4_ops" ASC NULLS LAST
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
-- Indexes structure for table auth_scene
-- ----------------------------
CREATE UNIQUE INDEX "auth_scene_scene_code_idx" ON "public"."auth_scene" USING btree (
  "scene_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_scene
-- ----------------------------
ALTER TABLE "public"."auth_scene" ADD CONSTRAINT "auth_scene_pkey" PRIMARY KEY ("scene_id");

-- ----------------------------
-- Indexes structure for table platform_admin
-- ----------------------------
CREATE UNIQUE INDEX "platform_admin_account_idx" ON "public"."platform_admin" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
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
-- Indexes structure for table user_user
-- ----------------------------
CREATE UNIQUE INDEX "user_user_account_idx" ON "public"."user_user" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "user_user_open_id_of_wx_idx" ON "public"."user_user" USING btree (
  "open_id_of_wx" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "user_user_phone_idx" ON "public"."user_user" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "user_user_union_id_of_wx_idx" ON "public"."user_user" USING btree (
  "union_id_of_wx" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user_user
-- ----------------------------
ALTER TABLE "public"."user_user" ADD CONSTRAINT "user_user_pkey" PRIMARY KEY ("user_id");
