
-- ----------------------------
-- Table structure for sys_message
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_message";
CREATE TABLE "public"."sys_message" (
                                        "id" int8 NOT NULL,
                                        "title" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                        "summary" varchar(128) COLLATE "pg_catalog"."default",
                                        "content" text COLLATE "pg_catalog"."default" NOT NULL,
                                        "type" int2 NOT NULL,
                                        "link" text COLLATE "pg_catalog"."default",
                                        "to_user_ids" text COLLATE "pg_catalog"."default" NOT NULL,
                                        "to_user_type" int2,
                                        "from_user_id" int8 NOT NULL,
                                        "from_user_type" int4,
                                        "send_at" timestamptz(6),
                                        "ext_json" json,
                                        "read_user_ids" text COLLATE "pg_catalog"."default" NOT NULL,
                                        "data_identifier" varchar(255) COLLATE "pg_catalog"."default",
                                        "created_at" timestamptz(6),
                                        "updated_at" timestamptz(6),
                                        "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."sys_message" OWNER TO "mianlajie";
COMMENT ON COLUMN "public"."sys_message"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_message"."title" IS '标题';
COMMENT ON COLUMN "public"."sys_message"."summary" IS '摘要';
COMMENT ON COLUMN "public"."sys_message"."content" IS '内容';
COMMENT ON COLUMN "public"."sys_message"."type" IS '消息类型';
COMMENT ON COLUMN "public"."sys_message"."link" IS '跳转链接';
COMMENT ON COLUMN "public"."sys_message"."to_user_ids" IS '接收者UserIds，允许有多个接收者';
COMMENT ON COLUMN "public"."sys_message"."to_user_type" IS '接收者类型用户类型，和UserType保持一致';
COMMENT ON COLUMN "public"."sys_message"."from_user_id" IS '发送者ID，为-1代表系统消息';
COMMENT ON COLUMN "public"."sys_message"."from_user_type" IS '发送者类型';
COMMENT ON COLUMN "public"."sys_message"."send_at" IS '发送时间';
COMMENT ON COLUMN "public"."sys_message"."ext_json" IS '拓展数据Json';
COMMENT ON COLUMN "public"."sys_message"."read_user_ids" IS '已读用户UserIds';
COMMENT ON COLUMN "public"."sys_message"."data_identifier" IS '关联的数据标识';

-- ----------------------------
-- Records of sys_message
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Primary Key structure for table sys_message
-- ----------------------------
ALTER TABLE "public"."sys_message" ADD CONSTRAINT "sys_message_pkey" PRIMARY KEY ("id");


