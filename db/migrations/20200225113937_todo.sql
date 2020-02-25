
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE todos (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'プライマリーキー',
  status VARCHAR(16) NOT NULL COMMENT 'ステータス',
  task VARCHAR(32) NOT NULL COMMENT 'タスク',
  limit_date DATETIME NOT NULL DEFAULT '9999-12-31 23:59:59' COMMENT '期限'
);

INSERT INTO todos(status, task, limit_date) VALUES('completed', 'hoge', '2018-01-08 08:00:00');
INSERT INTO todos(status, task) VALUES('in progress', 'foo');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE todos;
