-- 为sys_user表添加salt字段以支持随机盐值密码加密
ALTER TABLE sys_user 
    ADD COLUMN IF NOT EXISTS salt VARCHAR(32) DEFAULT '' COMMENT '密码盐值';

-- 为salt字段添加索引（可选，用于查询优化）
CREATE INDEX IF NOT EXISTS idx_sys_user_salt ON sys_user(salt);