ALTER TABLE "public"."sys_permission"
    ADD COLUMN "sort" int4;

COMMENT ON COLUMN "public"."sys_permission"."sort" IS '排序字段';