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

 Date: 18/06/2024 15:18:35
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
-- Table structure for auth_action
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action";
CREATE TABLE "public"."auth_action" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "is_stop" int2 NOT NULL DEFAULT 0,
  "action_id" int4 NOT NULL DEFAULT nextval('auth_action_action_id_seq'::regclass),
  "action_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "action_code" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."auth_action"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_action"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_action"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_action"."action_id" IS '操作ID';
COMMENT ON COLUMN "public"."auth_action"."action_name" IS '名称';
COMMENT ON COLUMN "public"."auth_action"."action_code" IS '标识';
COMMENT ON COLUMN "public"."auth_action"."remark" IS '备注';
COMMENT ON TABLE "public"."auth_action" IS '权限操作表';

-- ----------------------------
-- Records of auth_action
-- ----------------------------
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '权限管理-场景-查看', 'authSceneRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理-场景-新增', 'authSceneCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '权限管理-场景-编辑', 'authSceneUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '权限管理-场景-删除', 'authSceneDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '权限管理-操作-查看', 'authActionRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '权限管理-操作-新增', 'authActionCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '权限管理-操作-编辑', 'authActionUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '权限管理-操作-删除', 'authActionDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '权限管理-菜单-查看', 'authMenuRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '权限管理-菜单-新增', 'authMenuCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '权限管理-菜单-编辑', 'authMenuUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '权限管理-菜单-删除', 'authMenuDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '权限管理-角色-查看', 'authRoleRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '权限管理-角色-新增', 'authRoleCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '权限管理-角色-编辑', 'authRoleUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '权限管理-角色-删除', 'authRoleDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, '权限管理-平台管理员-查看', 'platformAdminRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, '权限管理-平台管理员-新增', 'platformAdminCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '权限管理-平台管理员-编辑', 'platformAdminUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '权限管理-平台管理员-删除', 'platformAdminDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 21, '平台配置-查看', 'platformConfigRead', '可任意读取平台配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 22, '平台配置-保存', 'platformConfigSave', '可任意保存平台配置表');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 23, '应用配置-网站-查看', 'platformConfigWebsiteRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-应用配置-网站');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 24, '应用配置-网站-保存', 'platformConfigWebsiteSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-应用配置-网站');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 25, '应用配置-APP-查看', 'platformConfigAppRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-应用配置-APP');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 26, '应用配置-APP-保存', 'platformConfigAppSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-应用配置-APP');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 27, '插件配置-上传-查看', 'platformConfigUploadRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-上传');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 28, '插件配置-上传-保存', 'platformConfigUploadSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-上传');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 29, '插件配置-支付-查看', 'platformConfigPayRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-支付');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 30, '插件配置-支付-查看', 'platformConfigPaySave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-支付');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 31, '插件配置-短信-查看', 'platformConfigSmsRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 32, '插件配置-短信-保存', 'platformConfigSmsSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-短信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 33, '插件配置-实名认证-查看', 'platformConfigIdCardRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 34, '插件配置-实名认证-查看', 'platformConfigIdCardSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-实名认证');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 35, '插件配置-一键登录-查看', 'platformConfigOneClickRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 36, '插件配置-一键登录-保存', 'platformConfigOneClickSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-一键登录');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 37, '插件配置-推送-查看', 'platformConfigPushRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 38, '插件配置-推送-查看', 'platformConfigPushSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-推送');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 39, '插件配置-视频点播-查看', 'platformConfigVodRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 40, '插件配置-视频点播-保存', 'platformConfigVodSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-视频点播');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 41, '插件配置-微信-查看', 'platformConfigWxRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 42, '插件配置-微信-查看', 'platformConfigWxSave', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-微信');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 43, '插件配置-邮箱-查看', 'platformConfigEmailRead', '只能读取平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 44, '插件配置-邮箱-保存', 'platformConfigEmailSave', '只能保存平台配置表中的某些配置。对应前端页面：系统管理-插件配置-邮箱');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 45, '用户管理-用户-查看', 'usersRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 46, '用户管理-用户-编辑', 'usersUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 47, '机构管理-机构-查看', 'orgRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 48, '机构管理-机构-新增', 'orgCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 49, '机构管理-机构-编辑', 'orgUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 50, '机构管理-机构-删除', 'orgDelete', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 51, '权限管理-机构管理员-查看', 'orgAdminRead', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 52, '权限管理-机构管理员-新增', 'orgAdminCreate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 53, '权限管理-机构管理员-编辑', 'orgAdminUpdate', '');
INSERT INTO "public"."auth_action" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 54, '权限管理-机构管理员-删除', 'orgAdminDelete', '');

-- ----------------------------
-- Table structure for auth_action_rel_to_scene
-- ----------------------------
DROP TABLE IF EXISTS "public"."auth_action_rel_to_scene";
CREATE TABLE "public"."auth_action_rel_to_scene" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "action_id" int4 NOT NULL DEFAULT 0,
  "scene_id" int4 NOT NULL DEFAULT 0
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
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 1, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 2, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 3, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 4, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 5, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 6, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 7, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 8, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 9, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 10, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 11, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 12, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 13, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 13, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 14, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 14, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 15, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 15, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 16, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 16, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 17, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 18, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 19, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 20, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 21, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 22, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 23, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 24, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 25, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 26, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 27, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 28, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 29, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 30, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 31, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 32, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 33, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 34, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 35, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 36, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 37, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 38, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 39, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 40, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 41, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 42, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 43, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 44, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 45, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 46, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 47, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 48, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 49, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 50, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 51, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 51, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 52, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 52, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 53, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 53, 2);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 54, 1);
INSERT INTO "public"."auth_action_rel_to_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 54, 2);

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
  "scene_id" int4 NOT NULL DEFAULT 0,
  "pid" int4 NOT NULL DEFAULT 0,
  "level" int2 NOT NULL DEFAULT 0,
  "id_path" text COLLATE "pg_catalog"."default" DEFAULT ''::text,
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
COMMENT ON COLUMN "public"."auth_menu"."level" IS '层级';
COMMENT ON COLUMN "public"."auth_menu"."id_path" IS '层级路径';
COMMENT ON COLUMN "public"."auth_menu"."menu_icon" IS '图标。常用格式：autoicon-{集合}-{标识}；vant格式：vant-{标识}';
COMMENT ON COLUMN "public"."auth_menu"."menu_url" IS '链接';
COMMENT ON COLUMN "public"."auth_menu"."extra_data" IS '额外数据。JSON格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}';
COMMENT ON COLUMN "public"."auth_menu"."sort" IS '排序值。从大到小排序';
COMMENT ON TABLE "public"."auth_menu" IS '权限菜单表';

-- ----------------------------
-- Records of auth_menu
-- ----------------------------
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '主页', 1, 0, 1, '0-1', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '权限管理', 1, 0, 1, '0-2', 'autoicon-ep-lock', '', '{"i18n": {"title": {"en": "Auth Manage", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, '场景', 1, 2, 2, '0-2-3', 'autoicon-ep-flag', '/auth/scene', '{"i18n": {"title": {"en": "Scene", "zh-cn": "场景"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 4, '操作', 1, 2, 2, '0-2-4', 'autoicon-ep-coordinate', '/auth/action', '{"i18n": {"title": {"en": "Action", "zh-cn": "操作"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 5, '菜单', 1, 2, 2, '0-2-5', 'autoicon-ep-menu', '/auth/menu', '{"i18n": {"title": {"en": "Menu", "zh-cn": "菜单"}}}', 30);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 6, '角色', 1, 2, 2, '0-2-6', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 7, '平台管理员', 1, 2, 2, '0-2-7', 'vant-manager-o', '/platform/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "平台管理员"}}}', 50);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 8, '系统管理', 1, 0, 1, '0-8', 'autoicon-ep-platform', '', '{"i18n": {"title": {"en": "System Manage", "zh-cn": "系统管理"}}}', 20);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 9, '配置中心', 1, 8, 2, '0-8-9', 'autoicon-ep-setting', '', '{"i18n": {"title": {"en": "Config Center", "zh-cn": "配置中心"}}}', 0);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 10, '应用配置', 1, 9, 3, '0-8-9-10', '', '/platform/config/app', '{"i18n": {"title": {"en": "APP Config", "zh-cn": "应用配置"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 11, '插件配置', 1, 9, 3, '0-8-9-11', '', '/platform/config/plugin', '{"i18n": {"title": {"en": "Plugin Config", "zh-cn": "插件配置"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 12, '用户管理', 1, 0, 1, '0-12', 'vant-friends', '', '{"i18n": {"title": {"en": "User Manage", "zh-cn": "用户管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 13, '用户', 1, 12, 2, '0-12-13', 'vant-user-o', '/users/users', '{"i18n": {"title": {"en": "Users", "zh-cn": "用户"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 14, '机构管理', 1, 0, 1, '0-14', 'autoicon-ep-office-building', '', '{"i18n": {"title": {"en": "", "zh-cn": "机构管理"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 15, '机构', 1, 14, 2, '0-14-15', 'autoicon-ep-school', '/org/org', '{"i18n": {"title": {"en": "Org", "zh-cn": "机构"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 16, '机构管理员', 1, 2, 2, '0-2-16', 'vant-manager-o', '/org/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "机构管理员"}}}', 100);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 17, '主页', 2, 0, 1, '0-17', 'autoicon-ep-home-filled', '/', '{"i18n": {"title": {"en": "Homepage", "zh-cn": "主页"}}}', 255);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 18, '权限管理', 2, 0, 1, '0-18', 'autoicon-ep-menu', '', '{"i18n": {"title": {"en": "", "zh-cn": "权限管理"}}}', 10);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 19, '角色', 2, 18, 2, '0-18-19', 'autoicon-ep-view', '/auth/role', '{"i18n": {"title": {"en": "Role", "zh-cn": "角色"}}}', 40);
INSERT INTO "public"."auth_menu" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 20, '管理员', 2, 18, 2, '0-18-20', 'vant-manager-o', '/org/admin', '{"i18n": {"title": {"en": "Admin", "zh-cn": "管理员"}}}', 100);

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
  "scene_id" int4 NOT NULL DEFAULT 0,
  "rel_id" int4 NOT NULL DEFAULT 0
)
;
COMMENT ON COLUMN "public"."auth_role"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_role"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_role"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."auth_role"."role_name" IS '名称';
COMMENT ON COLUMN "public"."auth_role"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_role"."rel_id" IS '关联ID。0表示平台创建，其它值根据sceneId对应不同表';
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
  "action_id" int4 NOT NULL DEFAULT 0
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
  "scene_id" int4 NOT NULL DEFAULT nextval('auth_scene_scene_id_seq'::regclass),
  "scene_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_code" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "scene_config" json NOT NULL,
  "remark" varchar(120) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying
)
;
COMMENT ON COLUMN "public"."auth_scene"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."auth_scene"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."auth_scene"."is_stop" IS '停用：0否 1是';
COMMENT ON COLUMN "public"."auth_scene"."scene_id" IS '场景ID';
COMMENT ON COLUMN "public"."auth_scene"."scene_name" IS '名称';
COMMENT ON COLUMN "public"."auth_scene"."scene_code" IS '标识';
COMMENT ON COLUMN "public"."auth_scene"."scene_config" IS '配置。JSON格式，字段根据场景自定义。如下为场景使用JWT的示例：{"signType": "算法","signKey": "密钥","expireTime": 过期时间,...}';
COMMENT ON COLUMN "public"."auth_scene"."remark" IS '备注';
COMMENT ON TABLE "public"."auth_scene" IS '权限场景表';

-- ----------------------------
-- Records of auth_scene
-- ----------------------------
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, '平台后台', 'platform', '{"signKey": "www.admin.com_platform", "signType": "HS256", "expireTime": 14400}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 2, '机构后台', 'org', '{"signKey": "www.admin.com_org", "signType": "HS256", "expireTime": 14400}', '');
INSERT INTO "public"."auth_scene" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 3, 'APP', 'app', '{"signKey": "www.admin.com_app", "signType": "HS256", "expireTime": 604800}', '');

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
INSERT INTO "public"."platform_admin" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 0, 1, 1, '超级管理员', 'http://JB.Admin.com/common/20240106/1704522339892_31917913.png', NULL, NULL, 'admin', '0930b03ed8d217f1c5756b1a2e898e50', 'u74XLJAB');

-- ----------------------------
-- Table structure for platform_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."platform_config";
CREATE TABLE "public"."platform_config" (
  "created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "config_key" varchar(60) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
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
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailCodeSubject', '您的验证码');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailCodeTemplate', '验证码：{code}
说明：
1. 验证码在发送后的5分钟内有效。如果验证码过期，请重新请求一个新的验证码。
2. 出于安全考虑，请不要将此验证码分享给任何人。');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommonFromEmail', 'xxxx@qq.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommonPassword', 'xxxxxx');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommonSmtpHost', 'smtp.qq.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailOfCommonSmtpPort', '465');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'emailType', 'emailOfCommon');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardOfAliyunAppcode', 'appcode');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardOfAliyunHost', 'http://idcard.market.alicloudapi.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardOfAliyunPath', '/lianzhuo/idcard');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'idCardType', 'idCardOfAliyun');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfAliAppId', 'appId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfAliNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfAli');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfAliOpAppId', 'opAppId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfAliPrivateKey', '****************');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfAliPublicKey', '****************');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxApiV3Key', '********');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxAppId', 'appId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxMchid', 'mchId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxNotifyUrl', 'http://JB.Admin.com/pay/notify/payOfWx');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxPrivateKey', '-----BEGIN RSA PRIVATE KEY-----
****************************************************************
-----END RSA PRIVATE KEY-----');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'payOfWxSerialNo', '********');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxAndroidAccessID', 'accessID');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxAndroidSecretKey', 'secretKey');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxHost', 'https://api.tpns.tencent.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxIosAccessID', 'accessID');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxIosSecretKey', 'secretKey');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxMacOSAccessID', 'accessID');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushOfTxMacOSSecretKey', 'secretKey');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'pushType', 'pushOfTx');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsOfAliyunAccessKeyId', 'accessKeyId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsOfAliyunAccessKeySecret', 'accessKeySecret');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsOfAliyunEndpoint', 'dysmsapi.aliyuncs.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsOfAliyunSignName', 'JB Admin');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsOfAliyunTemplateCode', 'SMS_********');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'smsType', 'smsOfAliyun');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssAccessKeyId', 'accessKeyId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssAccessKeySecret', 'accessKeySecret');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssBucket', 'bucket');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssCallbackUrl', 'http://JB.Admin.com/upload/notify/uploadOfAliyunOss');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssEndpoint', 'sts.cn-hangzhou.aliyuncs.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssHost', 'https://oss-cn-hangzhou.aliyuncs.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfAliyunOssRoleArn', 'acs:ram::********:role/********');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfLocalFileSaveDir', '../public/');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfLocalFileUrlPrefix', 'http://JB.Admin.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfLocalSignKey', 'secretKey');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadOfLocalUrl', 'http://JB.Admin.com/upload/upload');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'uploadType', 'uploadOfLocal');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'vodOfAliyunAccessKeyId', 'accessKeyId');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'vodOfAliyunAccessKeySecret', 'accessKeySecret');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'vodOfAliyunEndpoint', 'sts.cn-shanghai.aliyuncs.com');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'vodOfAliyunRoleArn', 'acs:ram::********:role/********');
INSERT INTO "public"."platform_config" VALUES ('2024-01-01 00:00:00', '2024-01-01 00:00:00', 'vodType', 'vodOfAliyun');

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
ALTER SEQUENCE "public"."auth_action_action_id_seq"
OWNED BY "public"."auth_action"."action_id";
SELECT setval('"public"."auth_action_action_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."auth_menu_menu_id_seq"
OWNED BY "public"."auth_menu"."menu_id";
SELECT setval('"public"."auth_menu_menu_id_seq"', 1, false);

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
SELECT setval('"public"."auth_scene_scene_id_seq"', 1, false);

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
ALTER SEQUENCE "public"."users_user_id_seq"
OWNED BY "public"."users"."user_id";
SELECT setval('"public"."users_user_id_seq"', 1, false);

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
CREATE INDEX "auth_role_rel_id_idx" ON "public"."auth_role" USING btree (
  "rel_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE INDEX "auth_role_scene_id_idx" ON "public"."auth_role" USING btree (
  "scene_id" "pg_catalog"."int4_ops" ASC NULLS LAST
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
-- Primary Key structure for table org
-- ----------------------------
ALTER TABLE "public"."org" ADD CONSTRAINT "org_pkey" PRIMARY KEY ("org_id");

-- ----------------------------
-- Indexes structure for table org_admin
-- ----------------------------
CREATE UNIQUE INDEX "org_admin_account_idx" ON "public"."org_admin" USING btree (
  "account" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "org_admin_email_idx" ON "public"."org_admin" USING btree (
  "email" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "org_admin_org_id_idx" ON "public"."org_admin" USING btree (
  "org_id" "pg_catalog"."int4_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "org_admin_phone_idx" ON "public"."org_admin" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table org_admin
-- ----------------------------
ALTER TABLE "public"."org_admin" ADD CONSTRAINT "org_admin_pkey" PRIMARY KEY ("admin_id");

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
