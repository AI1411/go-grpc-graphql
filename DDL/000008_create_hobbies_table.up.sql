DROP TABLE IF EXISTS hobbies CASCADE;
CREATE TABLE hobbies
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- ID
    "name"      VARCHAR(100) NOT NULL UNIQUE,               -- 趣味名
    description TEXT         NULL,                          -- 説明
    category_id UUID                                        -- カテゴリーID
);

COMMENT ON TABLE hobbies IS '趣味テーブル';
COMMENT ON COLUMN hobbies.id IS 'ID';
COMMENT ON COLUMN hobbies.name IS '趣味名';
COMMENT ON COLUMN hobbies.description IS '説明';
COMMENT ON COLUMN hobbies.category_id IS 'カテゴリーID';

CREATE INDEX hobbies_id_idx ON hobbies (id);