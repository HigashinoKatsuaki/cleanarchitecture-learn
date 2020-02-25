
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE todos (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'プライマリーキー',
  status VARCHAR(16) NOT NULL COMMENT 'ステータス',
  task VARCHAR(32) NOT NULL COMMENT 'タスク',
  limit_date DATETIME NOT NULL COMMENT '期限'
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE todos;
