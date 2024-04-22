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

 Date: 22/04/2024 12:18:08
*/


-- ----------------------------
-- Sequence structure for auth_action_actionId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_action_actionId_seq";
CREATE SEQUENCE "public"."auth_action_actionId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_menu_menuId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_menu_menuId_seq";
CREATE SEQUENCE "public"."auth_menu_menuId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_role_roleId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_role_roleId_seq";
CREATE SEQUENCE "public"."auth_role_roleId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for auth_scene_sceneId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."auth_scene_sceneId_seq";
CREATE SEQUENCE "public"."auth_scene_sceneId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for platform_admin_adminId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."platform_admin_adminId_seq";
CREATE SEQUENCE "public"."platform_admin_adminId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for platform_server_serverId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."platform_server_serverId_seq";
CREATE SEQUENCE "public"."platform_server_serverId_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for user_userId_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_userId_seq";
CREATE SEQUENCE "public"."user_userId_seq" 
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
  "actionId" int4 NOT NULL DEFAULT nextval('"auth_action_actionId_seq"'::regclass),
  "actionName" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "actionCode" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_action"."actionId" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action"."actionName" IS '名称';
COMMENT ON COLUMN "public"."auth_action"."actionCode" IS '标识';
COMMENT ON COLUMN "public"."auth_action"."remark" IS '备注';
COMMENT ON COLUMN "public"."auth_action"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_action"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action"."createdAt" IS '创建时间';

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
INSERT INTO "public"."auth_action" VALUES (23, '用户-查看', 'userLook', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');
INSERT INTO "public"."auth_action" VALUES (24, '用户-编辑', 'userUpdate', '', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action_rel_to_scene";
CREATE TABLE "public"."auth_action_rel_to_scene" (
  "actionId" int4 NOT NULL DEFAULT 0,
  "sceneId" int4 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."actionId" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."sceneId" IS '场景ID';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action_rel_to_scene"."createdAt" IS '创建时间';

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
  "menuId" int4 NOT NULL DEFAULT nextval('"auth_menu_menuId_seq"'::regclass),
  "menuName" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sceneId" int4 NOT NULL DEFAULT 0,
  "pid" int4 NOT NULL DEFAULT 0,
  "level" int2 NOT NULL DEFAULT 0,
  "idPath" text COLLATE "pg_catalog"."default",
  "menuIcon" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "menuUrl" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "extraData" json,
  "sort" int2 NOT NULL DEFAULT 0,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_menu"."menuId" IS '菜单ID';
COMMENT ON COLUMN "public"."auth_menu"."menuName" IS '名称';
COMMENT ON COLUMN "public"."auth_menu"."sceneId" IS '场景ID';
COMMENT ON COLUMN "public"."auth_menu"."pid" IS '父ID';
COMMENT ON COLUMN "public"."auth_menu"."level" IS '层级';
COMMENT ON COLUMN "public"."auth_menu"."idPath" IS '层级路径';
COMMENT ON COLUMN "public"."auth_menu"."menuIcon" IS '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}';
COMMENT ON COLUMN "public"."auth_menu"."menuUrl" IS '链接';
COMMENT ON COLUMN "public"."auth_menu"."extraData" IS '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}';
COMMENT ON COLUMN "public"."auth_menu"."sort" IS '排序值。从小到大排序，默认50，范围0-100';
COMMENT ON COLUMN "public"."auth_menu"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_menu"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_menu"."createdAt" IS '创建时间';

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
  "roleId" int4 NOT NULL DEFAULT nextval('"auth_role_roleId_seq"'::regclass),
  "roleName" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sceneId" int4 NOT NULL DEFAULT 0,
  "tableId" int4 NOT NULL DEFAULT 0,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role"."roleId" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role"."roleName" IS '名称';
COMMENT ON COLUMN "public"."auth_role"."sceneId" IS '场景ID';
COMMENT ON COLUMN "public"."auth_role"."tableId" IS '关联表ID。0表示平台创建，其它值根据sceneId对应不同表，表示由哪个机构或个人创建';
COMMENT ON COLUMN "public"."auth_role"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_role"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of auth_role
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_of_platform_admin
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_of_platform_admin";
CREATE TABLE "public"."auth_role_rel_of_platform_admin" (
  "roleId" int4 NOT NULL DEFAULT 0,
  "adminId" int4 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."roleId" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."adminId" IS '管理员ID';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_of_platform_admin"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_of_platform_admin
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_action";
CREATE TABLE "public"."auth_role_rel_to_action" (
  "roleId" int4 NOT NULL DEFAULT 0,
  "actionId" int4 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."roleId" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."actionId" IS '操作ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_action"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_to_action
-- ----------------------------

-- ----------------------------
-- Table structure for auth_role_rel_to_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_role_rel_to_menu";
CREATE TABLE "public"."auth_role_rel_to_menu" (
  "roleId" int4 NOT NULL DEFAULT 0,
  "menuId" int4 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."roleId" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."menuId" IS '菜单ID';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role_rel_to_menu"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of auth_role_rel_to_menu
-- ----------------------------

-- ----------------------------
-- Table structure for auth_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_scene";
CREATE TABLE "public"."auth_scene" (
  "sceneId" int4 NOT NULL DEFAULT nextval('"auth_scene_sceneId_seq"'::regclass),
  "sceneName" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sceneCode" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "sceneConfig" json NOT NULL,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."auth_scene"."sceneId" IS '场景ID';
COMMENT ON COLUMN "public"."auth_scene"."sceneName" IS '名称';
COMMENT ON COLUMN "public"."auth_scene"."sceneCode" IS '标识';
COMMENT ON COLUMN "public"."auth_scene"."sceneConfig" IS '配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}';
COMMENT ON COLUMN "public"."auth_scene"."remark" IS '备注';
COMMENT ON COLUMN "public"."auth_scene"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_scene"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."auth_scene"."createdAt" IS '创建时间';

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
  "adminId" int4 NOT NULL DEFAULT nextval('"platform_admin_adminId_seq"'::regclass),
  "phone" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_admin"."adminId" IS '管理员ID';
COMMENT ON COLUMN "public"."platform_admin"."phone" IS '手机号';
COMMENT ON COLUMN "public"."platform_admin"."account" IS '账号';
COMMENT ON COLUMN "public"."platform_admin"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."platform_admin"."salt" IS '密码盐';
COMMENT ON COLUMN "public"."platform_admin"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."platform_admin"."avatar" IS '头像';
COMMENT ON COLUMN "public"."platform_admin"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."platform_admin"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."platform_admin"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of platform_admin
-- ----------------------------
INSERT INTO "public"."platform_admin" VALUES (1, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB', '超级管理员', 'http://JB.Admin.com/common/20240106/1704522339892_31917913.png', 0, '2024-01-01 00:00:00', '2024-01-01 00:00:00');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_config";
CREATE TABLE "public"."platform_config" (
  "configKey" varchar(60) COLLATE "pg_catalog"."default" NOT NULL,
  "configValue" text COLLATE "pg_catalog"."default" NOT NULL,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_config"."configKey" IS '配置Key';
COMMENT ON COLUMN "public"."platform_config"."configValue" IS '配置值';
COMMENT ON COLUMN "public"."platform_config"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."platform_config"."createdAt" IS '创建时间';

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
  "serverId" int4 NOT NULL DEFAULT nextval('"platform_server_serverId_seq"'::regclass),
  "networkIp" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "localIp" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."platform_server"."serverId" IS '服务器ID';
COMMENT ON COLUMN "public"."platform_server"."networkIp" IS '公网IP';
COMMENT ON COLUMN "public"."platform_server"."localIp" IS '内网IP';
COMMENT ON COLUMN "public"."platform_server"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."platform_server"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of platform_server
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "userId" int4 NOT NULL DEFAULT nextval('"user_userId_seq"'::regclass),
  "phone" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "account" varchar(30) COLLATE "pg_catalog"."default" DEFAULT NULL::character varying,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "salt" char(8) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::bpchar,
  "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "avatar" varchar(200) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "gender" int2 NOT NULL DEFAULT 0,
  "birthday" date,
  "address" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "idCardName" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "idCardNo" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "isStop" int2 NOT NULL DEFAULT 0,
  "updatedAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "createdAt" timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;
COMMENT ON COLUMN "public"."user"."userId" IS '用户ID';
COMMENT ON COLUMN "public"."user"."phone" IS '手机号';
COMMENT ON COLUMN "public"."user"."account" IS '账号';
COMMENT ON COLUMN "public"."user"."password" IS '密码。md5保存';
COMMENT ON COLUMN "public"."user"."salt" IS '密码盐';
COMMENT ON COLUMN "public"."user"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."user"."avatar" IS '头像';
COMMENT ON COLUMN "public"."user"."gender" IS '性别：0未设置 1男 2女';
COMMENT ON COLUMN "public"."user"."birthday" IS '生日';
COMMENT ON COLUMN "public"."user"."address" IS '详细地址';
COMMENT ON COLUMN "public"."user"."idCardName" IS '身份证姓名';
COMMENT ON COLUMN "public"."user"."idCardNo" IS '身份证号码';
COMMENT ON COLUMN "public"."user"."isStop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."user"."updatedAt" IS '更新时间';
COMMENT ON COLUMN "public"."user"."createdAt" IS '创建时间';

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_action_actionId_seq"
OWNED BY "public"."auth_action"."actionId";
SELECT setval('"public"."auth_action_actionId_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_menu_menuId_seq"
OWNED BY "public"."auth_menu"."menuId";
SELECT setval('"public"."auth_menu_menuId_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_role_roleId_seq"
OWNED BY "public"."auth_role"."roleId";
SELECT setval('"public"."auth_role_roleId_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_scene_sceneId_seq"
OWNED BY "public"."auth_scene"."sceneId";
SELECT setval('"public"."auth_scene_sceneId_seq"', 4, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_admin_adminId_seq"
OWNED BY "public"."platform_admin"."adminId";
SELECT setval('"public"."platform_admin_adminId_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."platform_server_serverId_seq"
OWNED BY "public"."platform_server"."serverId";
SELECT setval('"public"."platform_server_serverId_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_userId_seq"
OWNED BY "public"."user"."userId";
SELECT setval('"public"."user_userId_seq"', 1, false);

-- ----------------------------
-- Indexes structure for table auth_action
-- ----------------------------
CREATE UNIQUE INDEX "auth_action_actionCode_idx" ON "public"."auth_action" USING btree (
  "actionCode" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_action
-- ----------------------------
ALTER TABLE "public"."auth_action" ADD CONSTRAINT "auth_action_pkey" PRIMARY KEY ("actionId");

-- ----------------------------
-- Indexes structure for table auth_action_rel_to_scene
-- ----------------------------
CREATE INDEX "auth_action_rel_to_scene_actionId_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "actionId" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_action_rel_to_scene_sceneId_idx" ON "public"."auth_action_rel_to_scene" USING btree (
  "sceneId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_action_rel_to_scene
-- ----------------------------
ALTER TABLE "public"."auth_action_rel_to_scene" ADD CONSTRAINT "auth_action_rel_to_scene_pkey" PRIMARY KEY ("actionId", "sceneId");

-- ----------------------------
-- Indexes structure for table auth_menu
-- ----------------------------
CREATE INDEX "auth_menu_pid_idx" ON "public"."auth_menu" USING btree (
  "pid" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_menu_sceneId_idx" ON "public"."auth_menu" USING btree (
  "sceneId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_menu
-- ----------------------------
ALTER TABLE "public"."auth_menu" ADD CONSTRAINT "auth_menu_pkey" PRIMARY KEY ("menuId");

-- ----------------------------
-- Indexes structure for table auth_role
-- ----------------------------
CREATE INDEX "auth_role_sceneId_idx" ON "public"."auth_role" USING btree (
  "sceneId" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_tableId_idx" ON "public"."auth_role" USING btree (
  "tableId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role
-- ----------------------------
ALTER TABLE "public"."auth_role" ADD CONSTRAINT "auth_role_pkey" PRIMARY KEY ("roleId");

-- ----------------------------
-- Indexes structure for table auth_role_rel_of_platform_admin
-- ----------------------------
CREATE INDEX "auth_role_rel_of_platform_admin_adminId_idx" ON "public"."auth_role_rel_of_platform_admin" USING btree (
  "adminId" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_of_platform_admin_roleId_idx" ON "public"."auth_role_rel_of_platform_admin" USING btree (
  "roleId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_of_platform_admin
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_of_platform_admin" ADD CONSTRAINT "auth_role_rel_of_platform_admin_pkey" PRIMARY KEY ("roleId", "adminId");

-- ----------------------------
-- Indexes structure for table auth_role_rel_to_action
-- ----------------------------
CREATE INDEX "auth_role_rel_to_action_actionId_idx" ON "public"."auth_role_rel_to_action" USING btree (
  "actionId" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_to_action_roleId_idx" ON "public"."auth_role_rel_to_action" USING btree (
  "roleId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_to_action
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_to_action" ADD CONSTRAINT "auth_role_rel_to_action_pkey" PRIMARY KEY ("roleId", "actionId");

-- ----------------------------
-- Indexes structure for table auth_role_rel_to_menu
-- ----------------------------
CREATE INDEX "auth_role_rel_to_menu_menuId_idx" ON "public"."auth_role_rel_to_menu" USING btree (
  "menuId" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_rel_to_menu_roleId_idx" ON "public"."auth_role_rel_to_menu" USING btree (
  "roleId" "pg_catalog"."int4_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_role_rel_to_menu
-- ----------------------------
ALTER TABLE "public"."auth_role_rel_to_menu" ADD CONSTRAINT "auth_role_rel_to_menu_pkey" PRIMARY KEY ("roleId", "menuId");

-- ----------------------------
-- Indexes structure for table auth_scene
-- ----------------------------
CREATE UNIQUE INDEX "auth_scene_sceneCode_idx" ON "public"."auth_scene" USING btree (
  "sceneCode" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table auth_scene
-- ----------------------------
ALTER TABLE "public"."auth_scene" ADD CONSTRAINT "auth_scene_pkey" PRIMARY KEY ("sceneId");

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
ALTER TABLE "public"."platform_admin" ADD CONSTRAINT "platform_admin_pkey" PRIMARY KEY ("adminId");

-- ----------------------------
-- Primary Key structure for table platform_config
-- ----------------------------
ALTER TABLE "public"."platform_config" ADD CONSTRAINT "platform_config_pkey" PRIMARY KEY ("configKey");

-- ----------------------------
-- Indexes structure for table platform_server
-- ----------------------------
CREATE UNIQUE INDEX "platform_server_networkIp_idx" ON "public"."platform_server" USING btree (
  "networkIp" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table platform_server
-- ----------------------------
ALTER TABLE "public"."platform_server" ADD CONSTRAINT "platform_server_pkey" PRIMARY KEY ("serverId");

-- ----------------------------
-- Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX "user_account_idx" ON "public"."user" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "user_phone_idx" ON "public"."user" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table user
-- ----------------------------
ALTER TABLE "public"."user" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("userId");
