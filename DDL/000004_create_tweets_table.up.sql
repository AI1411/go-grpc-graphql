DROP TABLE IF EXISTS tweets;
CREATE TABLE IF NOT EXISTS tweets
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(), -- ID
    user_id    UUID                     NOT NULL,                           -- ユーザID
    body       VARCHAR(255)             NOT NULL,                           -- 本文
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),             -- 作成日時
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()              -- 更新日時
);

COMMENT ON TABLE tweets IS 'つぶやきテーブル';
COMMENT ON COLUMN tweets.id IS 'ID';
COMMENT ON COLUMN tweets.user_id IS 'ユーザID';
COMMENT ON COLUMN tweets.body IS '本文';
COMMENT ON COLUMN tweets.created_at IS '作成日時';
COMMENT ON COLUMN tweets.updated_at IS '更新日時';

CREATE INDEX IF NOT EXISTS user_id_idx on tweets (user_id);
CREATE INDEX IF NOT EXISTS created_at_idx on chats (created_at);