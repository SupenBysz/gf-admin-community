ALTER TABLE "public"."sys_role"
DROP COLUMN "is_sys";

COMMENT ON COLUMN "public"."sys_role"."is_system" IS '是否默认角色，true仅能修改名称，不允许删除和修改';