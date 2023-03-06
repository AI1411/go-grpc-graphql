DROP TABLE IF EXISTS rooms;
CREATE TABLE rooms
(
    id         UUID      NOT NULL DEFAULT gen_random_uuid(), -- ルームID
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),             -- 作成日時
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()              -- 更新日時
);

COMMENT ON TABLE rooms IS 'ルーム';
COMMENT ON COLUMN rooms.id IS 'ルームID';
COMMENT ON COLUMN rooms.created_at IS '作成日時';
COMMENT ON COLUMN rooms.updated_at IS '更新日時';