ALTER TABLE "public"."sys_role"
    ADD COLUMN "union_main_id" int8,
  ADD COLUMN "is_sys" varchar(255);

COMMENT ON COLUMN "public"."sys_role"."union_main_id" IS '主体id';

COMMENT ON COLUMN "public"."sys_role"."is_sys" IS '0允许删除和修改 1禁止修改删除，并拥有默认权限 ';