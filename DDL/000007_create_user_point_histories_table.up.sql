DROP TABLE IF EXISTS user_point_histories;
CREATE TABLE user_point_histories
(
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id        UUID         NOT NULL,
    point          INTEGER      NOT NULL,
    operation_type VARCHAR(255) NOT NULL, -- 'add' or 'subtract'など
    description    TEXT,                  -- ポイント操作の説明（任意）
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NOT NULL
);

-- 外部キー制約を追加
ALTER TABLE user_point_histories
    ADD CONSTRAINT fk_user_point_histories_user_id
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- インデックスを追加
CREATE INDEX idx_user_point_histories_user_id_created_at
    ON user_point_histories (user_id, created_at);

-- 日本語のコメント
COMMENT ON TABLE user_point_histories IS 'ユーザーのポイント履歴を追跡するためのテーブル';
COMMENT ON COLUMN user_point_histories.id IS 'user_point_historiesテーブルの主キー';
COMMENT ON COLUMN user_point_histories.user_id IS 'usersテーブルを参照する外部キー';
COMMENT ON COLUMN user_point_histories.point IS 'ポイント操作の金額（正または負）';
COMMENT ON COLUMN user_point_histories.operation_type IS 'ポイント操作の種類（例：追加、減算）';
COMMENT ON COLUMN user_point_histories.description IS 'ポイント操作の説明（任意）';
COMMENT ON COLUMN user_point_histories.created_at IS 'レコードが作成されたタイムスタンプ';
COMMENT ON COLUMN user_point_histories.updated_at IS 'レコードが最後に更新されたタイムスタンプ';
