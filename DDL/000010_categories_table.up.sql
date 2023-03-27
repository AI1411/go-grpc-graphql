DROP TABLE IF EXISTS categories CASCADE;
CREATE TABLE categories
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- ID
    "name"      VARCHAR(100) NOT NULL UNIQUE,               -- カテゴリ名
    description TEXT         NULL                           -- カテゴリ名
);

COMMENT ON TABLE categories IS '趣味カテゴリーテーブル';
COMMENT ON COLUMN categories.id IS 'ID';
COMMENT ON COLUMN categories.name IS 'カテゴリ名';
COMMENT ON COLUMN categories.description IS '説明';


ALTER TABLE hobbies
    DROP CONSTRAINT IF EXISTS fk_hobbies_category_id;
ALTER TABLE hobbies
    ADD CONSTRAINT fk_hobbies_category_id
        FOREIGN KEY (category_id)
            REFERENCES categories (id)
            ON DELETE CASCADE;

CREATE INDEX categories_id_idx ON categories (id);