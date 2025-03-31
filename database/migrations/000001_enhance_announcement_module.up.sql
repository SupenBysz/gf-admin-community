-- 扩展现有公告表
ALTER TABLE sys_announcement 
    ADD COLUMN IF NOT EXISTS category_id BIGINT DEFAULT 0 COMMENT '公告分类ID',
    ADD COLUMN IF NOT EXISTS priority INT DEFAULT 1 COMMENT '优先级：1普通、2重要、3紧急',
    ADD COLUMN IF NOT EXISTS is_pinned INT DEFAULT 0 COMMENT '是否置顶：0否、1是',
    ADD COLUMN IF NOT EXISTS is_force_read INT DEFAULT 0 COMMENT '是否强制阅读：0否、1是',
    ADD COLUMN IF NOT EXISTS tags VARCHAR(255) DEFAULT '' COMMENT '公告标签，多个用逗号分隔',
    ADD COLUMN IF NOT EXISTS read_count BIGINT DEFAULT 0 COMMENT '阅读次数',
    ADD COLUMN IF NOT EXISTS confirm_required INT DEFAULT 0 COMMENT '是否需要确认：0否、1是';

-- 为新字段添加索引
CREATE INDEX IF NOT EXISTS idx_sys_announcement_category_id ON sys_announcement(category_id);
CREATE INDEX IF NOT EXISTS idx_sys_announcement_priority ON sys_announcement(priority);
CREATE INDEX IF NOT EXISTS idx_sys_announcement_is_pinned ON sys_announcement(is_pinned);

-- 创建公告分类表
CREATE TABLE IF NOT EXISTS sys_announcement_category (
    id BIGINT NOT NULL COMMENT '主键ID',
    name VARCHAR(50) NOT NULL COMMENT '分类名称',
    code VARCHAR(50) NOT NULL COMMENT '分类编码',
    union_main_id BIGINT DEFAULT 0 COMMENT '所属主体ID',
    sort INT DEFAULT 0 COMMENT '排序',
    description VARCHAR(255) DEFAULT '' COMMENT '描述',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    created_by BIGINT DEFAULT 0 COMMENT '创建者ID',
    updated_by BIGINT DEFAULT 0 COMMENT '更新者ID',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    deleted_by BIGINT DEFAULT 0 COMMENT '删除者ID',
    PRIMARY KEY (id),
    INDEX idx_sys_announcement_category_code (code),
    INDEX idx_sys_announcement_category_union_main_id (union_main_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告分类表';

-- 创建公告确认表
CREATE TABLE IF NOT EXISTS sys_announcement_confirm (
    id BIGINT NOT NULL COMMENT '主键ID',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    announcement_id BIGINT NOT NULL COMMENT '公告ID',
    confirm_at TIMESTAMP NULL COMMENT '确认时间',
    confirm_comment VARCHAR(255) DEFAULT '' COMMENT '确认备注',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    created_by BIGINT DEFAULT 0 COMMENT '创建者ID',
    updated_by BIGINT DEFAULT 0 COMMENT '更新者ID',
    deleted_at TIMESTAMP NULL COMMENT '删除时间',
    deleted_by BIGINT DEFAULT 0 COMMENT '删除者ID',
    PRIMARY KEY (id),
    UNIQUE KEY uk_user_announcement (user_id, announcement_id),
    INDEX idx_sys_announcement_confirm_user_id (user_id),
    INDEX idx_sys_announcement_confirm_announcement_id (announcement_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='公告确认表';

-- 初始化公告分类数据
INSERT INTO sys_announcement_category (id, name, code, union_main_id, sort, description, created_at, created_by)
VALUES
    (1, '系统通知', 'system', 0, 10, '系统相关的通知公告', NOW(), 1),
    (2, '功能更新', 'feature', 0, 20, '新功能或功能更新通知', NOW(), 1),
    (3, '用户公告', 'user', 0, 30, '针对用户的重要公告', NOW(), 1),
    (4, '活动通知', 'activity', 0, 40, '各类活动的通知公告', NOW(), 1);

-- 更新现有公告的分类
UPDATE sys_announcement SET category_id = 1 WHERE category_id = 0; 