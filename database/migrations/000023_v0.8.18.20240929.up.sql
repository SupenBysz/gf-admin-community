/* 优化前端配置表 */
ALTER TABLE "public"."sys_front_settings"
    ADD COLUMN "version" varchar(16),
  ADD COLUMN "sys" int2 DEFAULT 0;

COMMENT ON COLUMN "public"."sys_front_settings"."version" IS '版本';

COMMENT ON COLUMN "public"."sys_front_settings"."sys" IS '1除主体管理员外，主体下的其他用户仅有只读权限，（默认0）';