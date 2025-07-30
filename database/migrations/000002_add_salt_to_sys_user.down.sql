-- 回滚：删除sys_user表的salt字段
ALTER TABLE sys_user 
    DROP INDEX IF EXISTS idx_sys_user_salt,
    DROP COLUMN IF EXISTS salt;