/* sys_user_detail IP字段数据长度从 varchar(64) 改成 varchar(128) */
ALTER TABLE "public"."sys_user_detail"
ALTER COLUMN "last_login_ip" TYPE varchar(128) COLLATE "pg_catalog"."default";