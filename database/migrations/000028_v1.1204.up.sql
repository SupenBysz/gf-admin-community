CREATE TABLE "public"."sys_comment" (
                                        "id" int8 NOT NULL,
                                        "user_id" int8 NOT NULL DEFAULT 0,
                                        "union_main_id" int8 NOT NULL DEFAULT 0,
                                        "union_main_type" int4 NOT NULL,
                                        "body" text,
                                        "media_ids" text,
                                        "reply_id" int8 NOT NULL DEFAULT 0,
                                        "score" int4 NOT NULL DEFAULT 5,
                                        "created_at" timestamp,
                                        PRIMARY KEY ("id", "reply_id")
)
;

ALTER TABLE "public"."sys_comment"
    OWNER TO "kuaimk";

COMMENT ON COLUMN "public"."sys_comment"."id" IS 'ID';

COMMENT ON COLUMN "public"."sys_comment"."user_id" IS '用户ID';

COMMENT ON COLUMN "public"."sys_comment"."union_main_id" IS '关联主体ID';

COMMENT ON COLUMN "public"."sys_comment"."union_main_type" IS '关联主体类型';

COMMENT ON COLUMN "public"."sys_comment"."body" IS '图文评论';

COMMENT ON COLUMN "public"."sys_comment"."media_ids" IS '媒体资源：图文、视频等';

COMMENT ON COLUMN "public"."sys_comment"."reply_id" IS '评论回复信息ID，即关联父级评论ID';

COMMENT ON COLUMN "public"."sys_comment"."score" IS '评分0-5，间隔0.1';

COMMENT ON COLUMN "public"."sys_comment"."created_at" IS '评论发表时间';