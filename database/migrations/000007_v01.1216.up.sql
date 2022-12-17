ALTER TABLE "public"."sys_permission"
DROP COLUMN "created_at",
  DROP COLUMN "updated_at",
  ADD COLUMN "is_show" int4 DEFAULT 1,
  ADD COLUMN "created_at" timestamp(6),
  ADD COLUMN "updated_at" timestamp(6);

COMMENT ON COLUMN "public"."sys_permission"."is_show" IS '是否显示：0不显示 1显示';