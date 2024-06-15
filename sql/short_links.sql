use shortlink;
-- 短链表
CREATE TABLE short_links (
                             id INT AUTO_INCREMENT PRIMARY KEY, # 主键
        original_url VARCHAR(255) NOT NULL,   # 原始链接
                                 short_code VARCHAR(255) NOT NULL UNIQUE, # 短链码
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP # 创建时间
);