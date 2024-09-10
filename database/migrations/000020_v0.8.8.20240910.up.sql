/* 个人资质表增加关联的用户ID */
ALTER TABLE "public"."sys_person_license"
    ADD COLUMN "user_id" int8 NOT NULL;

COMMENT ON COLUMN "public"."sys_person_license"."user_id" IS '关联的用户ID';