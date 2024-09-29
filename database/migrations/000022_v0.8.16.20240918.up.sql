/* 增加前端的配置表 */
CREATE TABLE "public"."sys_front_settings" (
                                   "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                   "values" json,
                                   "desc" varchar(128) COLLATE "pg_catalog"."default",
                                   "union_main_id" int8,
                                   "user_id" int8,
                                   "created_at" timestamptz(6),
                                   "updated_at" timestamptz(6)
)
;

COMMENT ON COLUMN "public"."sys_front_settings"."name" IS '配置名称';

COMMENT ON COLUMN "public"."sys_front_settings"."values" IS '配置信息JSON格式';

COMMENT ON COLUMN "public"."sys_front_settings"."desc" IS '描述';

COMMENT ON COLUMN "public"."sys_front_settings"."union_main_id" IS '关联的主体id，为0代表是平台配置';

COMMENT ON COLUMN "public"."sys_front_settings"."user_id" IS '关联的用户id，为0代表平台配置';


ALTER TABLE "public"."sys_front_settings"
    OWNER TO "当前数据库的所属用户";