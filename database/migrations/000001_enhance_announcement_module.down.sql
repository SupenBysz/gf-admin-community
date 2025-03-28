-- 删除公告确认表
DROP TABLE IF EXISTS sys_announcement_confirm;

-- 删除公告分类表
DROP TABLE IF EXISTS sys_announcement_category;

-- 删除公告表新增的字段
ALTER TABLE sys_announcement 
    DROP COLUMN IF EXISTS category_id,
    DROP COLUMN IF EXISTS priority,
    DROP COLUMN IF EXISTS is_pinned,
    DROP COLUMN IF EXISTS is_force_read,
    DROP COLUMN IF EXISTS tags,
    DROP COLUMN IF EXISTS read_count,
    DROP COLUMN IF EXISTS confirm_required; 