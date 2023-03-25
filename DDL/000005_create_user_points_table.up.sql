DROP TABLE IF EXISTS user_points CASCADE;
CREATE TABLE IF NOT EXISTS user_points
(
    id         UUID PRIMARY KEY                  DEFAULT gen_random_uuid(), -- ID
    user_id    UUID                     NOT NULL UNIQUE,                    -- ユーザID
    point      INTEGER                           DEFAULT 0,                 -- ポイント
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),             -- 作成日時
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()              -- 更新日時
);

COMMENT ON TABLE user_points IS 'ユーザポイント';
COMMENT ON COLUMN user_points.id IS 'ID';
COMMENT ON COLUMN user_points.user_id IS 'ユーザID';
COMMENT ON COLUMN user_points.point IS 'ポイント';
COMMENT ON COLUMN user_points.created_at IS '作成日時';
COMMENT ON COLUMN user_points.updated_at IS '更新日時';

CREATE INDEX IF NOT EXISTS user_id_idx on user_points (user_id);
CREATE INDEX IF NOT EXISTS created_at_idx on user_points (created_at);

ALTER TABLE user_points
    ADD CONSTRAINT user_points_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;