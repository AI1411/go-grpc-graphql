DROP TABLE IF EXISTS user_room;
CREATE TABLE user_room
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(), -- ID
    user_id    UUID      NOT NULL,                           -- ユーザID
    room_id    UUID      NOT NULL,                           -- ルームID
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),             -- 作成日時
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()              -- 更新日時
);

COMMENT ON TABLE user_room IS 'ユーザルーム';
COMMENT ON COLUMN user_room.id IS 'ID';
COMMENT ON COLUMN user_room.user_id IS 'ユーザID';
COMMENT ON COLUMN user_room.room_id IS 'ルームID';
COMMENT ON COLUMN user_room.created_at IS '作成日時';
COMMENT ON COLUMN user_room.updated_at IS '更新日時';

CREATE INDEX IF NOT EXISTS user_id_idx on user_room (user_id);
CREATE INDEX IF NOT EXISTS room_id_idx on user_room (room_id);
CREATE INDEX IF NOT EXISTS created_at_idx on user_room (created_at);