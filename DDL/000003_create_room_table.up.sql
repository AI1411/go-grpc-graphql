DROP TABLE IF EXISTS rooms CASCADE;
CREATE TABLE rooms
(
    id         UUID PRIMARY KEY   DEFAULT gen_random_uuid(), -- ルームID
    user_id    UUID      NOT NULL,                           -- ユーザID
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),             -- 作成日時
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()              -- 更新日時
);

COMMENT ON TABLE rooms IS 'ルーム';
COMMENT ON COLUMN rooms.id IS 'ルームID';
COMMENT ON COLUMN rooms.user_id IS 'ユーザID';
COMMENT ON COLUMN rooms.created_at IS '作成日時';
COMMENT ON COLUMN rooms.updated_at IS '更新日時';

ALTER TABLE chats
    ADD CONSTRAINT chats_room_id_fkey FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE;