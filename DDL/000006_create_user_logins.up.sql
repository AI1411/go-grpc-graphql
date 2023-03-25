DROP TABLE IF EXISTS user_logins CASCADE;
CREATE TABLE user_logins
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY, -- UUIDを主キーに設定し、デフォルトでuuidv4を生成
    user_id    UUID      NOT NULL,                         -- UUIDを使用
    login_date DATE      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- 外部キー制約を追加
ALTER TABLE user_logins
    ADD CONSTRAINT fk_user_logins_user_id
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- UNIQUE制約を追加
ALTER TABLE user_logins
    ADD CONSTRAINT uq_user_logins_user_id_login_date
        UNIQUE (user_id, login_date);

CREATE INDEX idx_user_logins_user_id_login_date
    ON user_logins (user_id, login_date);

-- 日本語コメント
COMMENT ON TABLE user_logins IS 'ユーザーの毎日のログインポイントを追跡するためのテーブル';
COMMENT ON COLUMN user_logins.id IS 'user_loginsテーブルの主キー';
COMMENT ON COLUMN user_logins.user_id IS 'usersテーブルを参照する外部キー';
COMMENT ON COLUMN user_logins.login_date IS 'ユーザーがログインした日付。毎日のログインを追跡するために使用されます';
COMMENT ON COLUMN user_logins.created_at IS 'レコードが作成されたタイムスタンプ';
COMMENT ON COLUMN user_logins.updated_at IS 'レコードが最後に更新されたタイムスタンプ';
