DROP TABLE IF EXISTS chats;
CREATE TABLE IF NOT EXISTS chats
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(), -- ID
    from_user_id UUID                     NOT NULL,                           -- 送り元ユーザID
    to_user_id   UUID                     NOT NULL,                           -- 送り先ユーザID
    body         VARCHAR(255)             NOT NULL,                           -- 本文
    is_read      BOOLEAN                  NOT NULL DEFAULT FALSE,             -- 既読フラグ
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),             -- 作成日時
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()              -- 更新日時
);

COMMENT ON TABLE chats IS 'チャットテーブル';
COMMENT ON COLUMN chats.id IS 'ID';
COMMENT ON COLUMN chats.from_user_id IS '送り元ユーザID';
COMMENT ON COLUMN chats.to_user_id IS '送り先ユーザID';
COMMENT ON COLUMN chats.body IS '本文';
COMMENT ON COLUMN chats.IS_read IS '既読フラグ';
COMMENT ON COLUMN chats.created_at IS '作成日時';
COMMENT ON COLUMN chats.updated_at IS '更新日時';

CREATE INDEX IF NOT EXISTS from_user_id_idx on chats (from_user_id);
CREATE INDEX IF NOT EXISTS to_user_id_idx on chats (to_user_id);
CREATE INDEX IF NOT EXISTS is_read_idx on chats (is_read);
CREATE INDEX IF NOT EXISTS created_at_idx on chats (created_at);