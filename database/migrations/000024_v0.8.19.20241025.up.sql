/*
 Navicat Premium Data Transfer

 Source Server         : 筷满客-Dev
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 10.68.74.250:5432
 Source Catalog        : kuaimkdb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 25/10/2024 16:39:21
*/


-- ----------------------------
-- Table structure for sys_member_level
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_member_level";
CREATE TABLE "public"."sys_member_level" (
                                             "id" int8 NOT NULL,
                                             "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
                                             "desc" varchar(255) COLLATE "pg_catalog"."default",
                                             "identifier" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
                                             "created_at" timestamp(6),
                                             "updated_at" timestamp(6),
                                             "created_by" int8,
                                             "union_main_id" int4 NOT NULL
)
;
ALTER TABLE "public"."sys_member_level" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_member_level"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_member_level"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_member_level"."desc" IS '描述';
COMMENT ON COLUMN "public"."sys_member_level"."identifier" IS '级别标识符';

-- ----------------------------
-- Primary Key structure for table sys_member_level
-- ----------------------------
ALTER TABLE "public"."sys_member_level" ADD CONSTRAINT "sys_member_level_pkey" PRIMARY KEY ("id");


/*
 Navicat Premium Data Transfer

 Source Server         : 筷满客-Dev
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 10.68.74.250:5432
 Source Catalog        : kuaimkdb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 25/10/2024 16:39:34
*/


-- ----------------------------
-- Table structure for sys_member_level_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_member_level_user";
CREATE TABLE "public"."sys_member_level_user" (
                                                  "id" int8 NOT NULL,
                                                  "user_id" int8 NOT NULL,
                                                  "ext_member_level_id" int8 NOT NULL,
                                                  "union_main_id" int8 NOT NULL
)
;
ALTER TABLE "public"."sys_member_level_user" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_member_level_user"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_member_level_user"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_member_level_user"."ext_member_level_id" IS '会员级别';
COMMENT ON COLUMN "public"."sys_member_level_user"."union_main_id" IS '保留字段';

-- ----------------------------
-- Primary Key structure for table sys_member_level_user
-- ----------------------------
ALTER TABLE "public"."sys_member_level_user" ADD CONSTRAINT "sys_member_level_user_pkey" PRIMARY KEY ("id");



