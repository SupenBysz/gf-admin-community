/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:16:56
*/


-- ----------------------------
-- Table structure for sms_app_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_app_config";
CREATE TABLE "public"."sms_app_config" (
                                           "id" int8 NOT NULL,
                                           "app_no" varchar(255) COLLATE "pg_catalog"."default",
                                           "app_name" varchar(255) COLLATE "pg_catalog"."default",
                                           "available_number" int4,
                                           "current_limiting" int4,
                                           "use_number" int4,
                                           "remark" varchar(255) COLLATE "pg_catalog"."default",
                                           "status" int4,
                                           "created_at" timestamp(6),
                                           "updated_at" timestamp(6),
                                           "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_app_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_app_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_app_config"."app_no" IS '应用编号';
COMMENT ON COLUMN "public"."sms_app_config"."app_name" IS '应用名称';
COMMENT ON COLUMN "public"."sms_app_config"."available_number" IS '可用数量';
COMMENT ON COLUMN "public"."sms_app_config"."current_limiting" IS '限流数量';
COMMENT ON COLUMN "public"."sms_app_config"."use_number" IS '已用数量';
COMMENT ON COLUMN "public"."sms_app_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_app_config"."status" IS '状态: 0禁用 1正常';

-- ----------------------------
-- Records of sms_app_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_app_config
-- ----------------------------
ALTER TABLE "public"."sms_app_config" ADD CONSTRAINT "sms_app_config_pkey" PRIMARY KEY ("id");


/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:17:48
*/


-- ----------------------------
-- Table structure for sms_business_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_business_config";
CREATE TABLE "public"."sms_business_config" (
                                                "id" int8 NOT NULL,
                                                "app_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "business_name" varchar(255) COLLATE "pg_catalog"."default",
                                                "business_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "template_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "business_desc" varchar(255) COLLATE "pg_catalog"."default",
                                                "remark" varchar(255) COLLATE "pg_catalog"."default",
                                                "status" int4,
                                                "created_at" timestamp(6),
                                                "updated_at" timestamp(6),
                                                "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_business_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_business_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_business_config"."app_no" IS '应用ID';
COMMENT ON COLUMN "public"."sms_business_config"."business_name" IS '业务名称';
COMMENT ON COLUMN "public"."sms_business_config"."business_no" IS '业务编号';
COMMENT ON COLUMN "public"."sms_business_config"."template_no" IS '模版编号';
COMMENT ON COLUMN "public"."sms_business_config"."business_desc" IS '业务说明';
COMMENT ON COLUMN "public"."sms_business_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_business_config"."status" IS '状态: 0禁用 1正常';

-- ----------------------------
-- Records of sms_business_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_business_config
-- ----------------------------
ALTER TABLE "public"."sms_business_config" ADD CONSTRAINT "sms_business_config_pkey1" PRIMARY KEY ("id");

/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:17:55
*/


-- ----------------------------
-- Table structure for sms_send_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_send_log";
CREATE TABLE "public"."sms_send_log" (
                                         "id" int8 NOT NULL,
                                         "app_no" varchar(255) COLLATE "pg_catalog"."default",
                                         "business_no" varchar(255) COLLATE "pg_catalog"."default",
                                         "status" int4,
                                         "fee" varchar(255) COLLATE "pg_catalog"."default",
                                         "phone_number" varchar(255) COLLATE "pg_catalog"."default",
                                         "message" varchar(255) COLLATE "pg_catalog"."default",
                                         "code" varchar(255) COLLATE "pg_catalog"."default",
                                         "content" varchar(255) COLLATE "pg_catalog"."default",
                                         "remark" varchar(255) COLLATE "pg_catalog"."default",
                                         "created_at" timestamp(6),
                                         "updated_at" timestamp(6),
                                         "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_send_log" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_send_log"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_send_log"."app_no" IS '应用ID';
COMMENT ON COLUMN "public"."sms_send_log"."business_no" IS '业务编号';
COMMENT ON COLUMN "public"."sms_send_log"."status" IS '状态';
COMMENT ON COLUMN "public"."sms_send_log"."fee" IS '计价条数';
COMMENT ON COLUMN "public"."sms_send_log"."phone_number" IS '发送手机号';
COMMENT ON COLUMN "public"."sms_send_log"."message" IS '接口响应消息';
COMMENT ON COLUMN "public"."sms_send_log"."code" IS '接口响应状态码';
COMMENT ON COLUMN "public"."sms_send_log"."content" IS '发送内容';
COMMENT ON COLUMN "public"."sms_send_log"."remark" IS '备注';

-- ----------------------------
-- Records of sms_send_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_send_log
-- ----------------------------
ALTER TABLE "public"."sms_send_log" ADD CONSTRAINT "sms_send_log_pkey" PRIMARY KEY ("id");

/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:18:01
*/


-- ----------------------------
-- Table structure for sms_service_provider_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_service_provider_config";
CREATE TABLE "public"."sms_service_provider_config" (
                                                        "id" int8 NOT NULL,
                                                        "provider_no" varchar(255) COLLATE "pg_catalog"."default",
                                                        "provider_name" varchar(255) COLLATE "pg_catalog"."default",
                                                        "access_key_id" varchar(255) COLLATE "pg_catalog"."default",
                                                        "access_key_secret" varchar(255) COLLATE "pg_catalog"."default",
                                                        "endpoint" varchar(255) COLLATE "pg_catalog"."default",
                                                        "sdk_app_id" varchar(255) COLLATE "pg_catalog"."default",
                                                        "region" varchar(255) COLLATE "pg_catalog"."default",
                                                        "remark" varchar(255) COLLATE "pg_catalog"."default",
                                                        "status" int4,
                                                        "created_at" timestamp(6),
                                                        "updated_at" timestamp(6),
                                                        "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_service_provider_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_service_provider_config"."id" IS '渠道商id';
COMMENT ON COLUMN "public"."sms_service_provider_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_service_provider_config"."provider_name" IS '渠道商名字';
COMMENT ON COLUMN "public"."sms_service_provider_config"."access_key_id" IS '身份标识';
COMMENT ON COLUMN "public"."sms_service_provider_config"."access_key_secret" IS '身份认证密钥';
COMMENT ON COLUMN "public"."sms_service_provider_config"."endpoint" IS '域名调用';
COMMENT ON COLUMN "public"."sms_service_provider_config"."sdk_app_id" IS '应用id';
COMMENT ON COLUMN "public"."sms_service_provider_config"."region" IS '地域';
COMMENT ON COLUMN "public"."sms_service_provider_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_service_provider_config"."status" IS '状态: 0禁用 1正常';

-- ----------------------------
-- Records of sms_service_provider_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_service_provider_config
-- ----------------------------
ALTER TABLE "public"."sms_service_provider_config" ADD CONSTRAINT "sms_service_provider_config_pkey" PRIMARY KEY ("id");

/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:18:09
*/


-- ----------------------------
-- Table structure for sms_sign_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_sign_config";
CREATE TABLE "public"."sms_sign_config" (
                                            "id" int8 NOT NULL,
                                            "sign_name" varchar(255) COLLATE "pg_catalog"."default",
                                            "provider_no" varchar(255) COLLATE "pg_catalog"."default",
                                            "provider_name" varchar(255) COLLATE "pg_catalog"."default",
                                            "remark" varchar(255) COLLATE "pg_catalog"."default",
                                            "status" int4,
                                            "audit_user_id" int8,
                                            "audit_reply_msg" varchar(255) COLLATE "pg_catalog"."default",
                                            "audit_at" timestamp(6),
                                            "created_at" timestamp(6),
                                            "updated_at" timestamp(6),
                                            "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_sign_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_sign_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_sign_config"."sign_name" IS '短信签名名称';
COMMENT ON COLUMN "public"."sms_sign_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_sign_config"."provider_name" IS '渠道商名字';
COMMENT ON COLUMN "public"."sms_sign_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_sign_config"."status" IS '状态: -1不通过 0待审核 1正常';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_user_id" IS '审核者UserID';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_reply_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."sms_sign_config"."audit_at" IS '审核时间';

-- ----------------------------
-- Records of sms_sign_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_sign_config
-- ----------------------------
ALTER TABLE "public"."sms_sign_config" ADD CONSTRAINT "sms_business_config_pkey" PRIMARY KEY ("id");

/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.105
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 192.168.1.105:5432
 Source Catalog        : kuaimk_test
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 14/02/2023 19:18:16
*/


-- ----------------------------
-- Table structure for sms_template_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sms_template_config";
CREATE TABLE "public"."sms_template_config" (
                                                "id" int8 NOT NULL,
                                                "sign_name" varchar(255) COLLATE "pg_catalog"."default",
                                                "template_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "template_name" varchar(255) COLLATE "pg_catalog"."default",
                                                "template_content" varchar(255) COLLATE "pg_catalog"."default",
                                                "third_party_template_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "provider_no" varchar(255) COLLATE "pg_catalog"."default",
                                                "remark" varchar(255) COLLATE "pg_catalog"."default",
                                                "status" int4,
                                                "audit_user_id" int8,
                                                "audit_reply_msg" varchar(255) COLLATE "pg_catalog"."default",
                                                "audit_at" varchar(255) COLLATE "pg_catalog"."default",
                                                "created_at" timestamp(6),
                                                "updated_at" timestamp(6),
                                                "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sms_template_config" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sms_template_config"."id" IS 'ID';
COMMENT ON COLUMN "public"."sms_template_config"."sign_name" IS '签名名称';
COMMENT ON COLUMN "public"."sms_template_config"."template_no" IS '模版编号';
COMMENT ON COLUMN "public"."sms_template_config"."template_name" IS '模版名称';
COMMENT ON COLUMN "public"."sms_template_config"."template_content" IS '模版内容';
COMMENT ON COLUMN "public"."sms_template_config"."third_party_template_no" IS '第三方模版编号';
COMMENT ON COLUMN "public"."sms_template_config"."provider_no" IS '渠道商编号';
COMMENT ON COLUMN "public"."sms_template_config"."remark" IS '备注';
COMMENT ON COLUMN "public"."sms_template_config"."status" IS '状态: 0禁用 1正常';
COMMENT ON COLUMN "public"."sms_template_config"."audit_user_id" IS '审核者UserID
审核者UserID
';
COMMENT ON COLUMN "public"."sms_template_config"."audit_reply_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."sms_template_config"."audit_at" IS '审核时间';

-- ----------------------------
-- Records of sms_template_config
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sms_template_config
-- ----------------------------
ALTER TABLE "public"."sms_template_config" ADD CONSTRAINT "sms_template_config_pkey" PRIMARY KEY ("id");
