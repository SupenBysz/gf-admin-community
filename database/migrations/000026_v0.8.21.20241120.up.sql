/*
 Navicat Premium Dump SQL

 Source Server         : 10.68.74.250
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 127.0.0.1:5432
 Source Catalog        : kuaimkdb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 20/11/2024 15:47:38
*/


-- ----------------------------
-- Table structure for sys_category
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_category";
CREATE TABLE "public"."sys_category" (
  "id" int8 NOT NULL,
  "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
  "parent_id" int8,
  "picture_path" varchar(255) COLLATE "pg_catalog"."default",
  "hidden" int2,
  "sort" int2,
  "union_main_id" int8
)
;
ALTER TABLE "public"."sys_category" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_category"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_category"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."sys_category"."picture_path" IS '分类图片';
COMMENT ON COLUMN "public"."sys_category"."hidden" IS '是否隐藏';
COMMENT ON COLUMN "public"."sys_category"."sort" IS '顺序';
COMMENT ON COLUMN "public"."sys_category"."union_main_id" IS '关联主体ID（保留字段）';

-- ----------------------------
-- Indexes structure for table sys_category
-- ----------------------------
CREATE INDEX "sys_category_name_idx" ON "public"."sys_category" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Uniques structure for table sys_category
-- ----------------------------
ALTER TABLE "public"."sys_category" ADD CONSTRAINT "sys_category_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table sys_category
-- ----------------------------
ALTER TABLE "public"."sys_category" ADD CONSTRAINT "sys_category_pkey" PRIMARY KEY ("id");
