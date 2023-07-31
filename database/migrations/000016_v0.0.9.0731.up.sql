CREATE TABLE "public"."sys_invite" (
                                   "id" int8 NOT NULL,
                                   "user_id" int8 NOT NULL,
                                   "value" json,
                                   "expire_at" timestamptz,
                                   "activate_number" int2,
                                   "state" int2 NOT NULL DEFAULT 1,
                                   "type" int2 NOT NULL,
                                   "created_at" timestamptz,
                                   PRIMARY KEY ("id")
)
;

COMMENT ON COLUMN "public"."sys_invite"."id" IS 'ID';

COMMENT ON COLUMN "public"."sys_invite"."user_id" IS '用户ID, 也就是邀约人ID';

COMMENT ON COLUMN "public"."sys_invite"."value" IS '邀约Json数据';

COMMENT ON COLUMN "public"."sys_invite"."expire_at" IS '邀约码背后的关联业务数据,';

COMMENT ON COLUMN "public"."sys_invite"."activate_number" IS '邀约码的激活次数限制';

COMMENT ON COLUMN "public"."sys_invite"."state" IS '状态： 0失效、1正常';

COMMENT ON COLUMN "public"."sys_invite"."type" IS '类型： 1注册、2加入团队、4加入角色 (复合类型)';

    COMMENT ON COLUMN "public"."sys_invite"."value" IS '邀约码背后的关联业务Json数据,';

COMMENT ON COLUMN "public"."sys_invite"."expire_at" IS '邀约码的过期失效';