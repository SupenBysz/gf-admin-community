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

 Date: 05/11/2024 18:04:44
*/


-- ----------------------------
-- Table structure for sys_announcement
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_announcement";
CREATE TABLE "public"."sys_announcement" (
                                             "id" int8 NOT NULL,
                                             "title" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
                                             "union_main_id" int8,
                                             "public_at" timestamptz(6),
                                             "body" text COLLATE "pg_catalog"."default",
                                             "user_type_scope" int2 NOT NULL DEFAULT 0,
                                             "expire_at" timestamptz(6),
                                             "state" int2 NOT NULL,
                                             "created_at" timestamptz(6),
                                             "updated_at" timestamptz(6),
                                             "created_by" int8,
                                             "updated_by" int8,
                                             "deleted_at" timestamptz(6),
                                             "deleted_by" int8,
                                             "ext_data_json" json
)
;
ALTER TABLE "public"."sys_announcement" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_announcement"."title" IS '公告标题';
COMMENT ON COLUMN "public"."sys_announcement"."union_main_id" IS '发布主体，0值则代表平台发布的公告';
COMMENT ON COLUMN "public"."sys_announcement"."public_at" IS '公示时间，只有到了公示时间用户才可见';
COMMENT ON COLUMN "public"."sys_announcement"."body" IS '公告正文';
COMMENT ON COLUMN "public"."sys_announcement"."user_type_scope" IS '受众用户类型：0则所有，复合类型';
COMMENT ON COLUMN "public"."sys_announcement"."expire_at" IS '过期时间，过期后前端用户不可见';
COMMENT ON COLUMN "public"."sys_announcement"."state" IS '状态：1草稿、2待发布、4已发布、8已过期、16已撤销';
COMMENT ON COLUMN "public"."sys_announcement"."created_by" IS '创建用户';
COMMENT ON COLUMN "public"."sys_announcement"."updated_by" IS '最后修改用户';
COMMENT ON COLUMN "public"."sys_announcement"."ext_data_json" IS '扩展json数据';

-- ----------------------------
-- Table structure for sys_announcement_read_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_announcement_read_user";
CREATE TABLE "public"."sys_announcement_read_user" (
                                                       "id" int8 NOT NULL,
                                                       "user_id" int8 NOT NULL,
                                                       "read_announcement_id" text COLLATE "pg_catalog"."default",
                                                       "read_at" timestamptz(6),
                                                       "ext_data_json" json,
                                                       "flag_read" int2
)
;
ALTER TABLE "public"."sys_announcement_read_user" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_announcement_read_user"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_announcement_read_user"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_announcement_read_user"."read_announcement_id" IS '用户已阅读的公告id';
COMMENT ON COLUMN "public"."sys_announcement_read_user"."read_at" IS '用户阅读时间';
COMMENT ON COLUMN "public"."sys_announcement_read_user"."ext_data_json" IS '扩展数据Json，由业务端决定用途';
COMMENT ON COLUMN "public"."sys_announcement_read_user"."flag_read" IS '标记已读，0未读，1已读，用户首次打开公告即标记已读，可手动标记未读，但read_at 数据不变，下次点开时更新阅读时间，并标记已读';

-- ----------------------------
-- Indexes structure for table sys_announcement
-- ----------------------------
CREATE INDEX "body_index" ON "public"."sys_announcement" USING btree (
    "body" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "title_index" ON "public"."sys_announcement" USING btree (
    "title" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "user_type_scope_index" ON "public"."sys_announcement" USING btree (
    "user_type_scope" "pg_catalog"."int2_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table sys_announcement
-- ----------------------------
ALTER TABLE "public"."sys_announcement" ADD CONSTRAINT "sys_announcement_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_announcement_read_user
-- ----------------------------
ALTER TABLE "public"."sys_announcement_read_user" ADD CONSTRAINT "sys_announcement_read_user_pkey" PRIMARY KEY ("id");
