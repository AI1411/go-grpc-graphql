DROP TABLE IF EXISTS hobby_categories CASCADE;
CREATE TABLE hobby_categories
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- ID
    "name"      VARCHAR(100) NOT NULL UNIQUE,               -- カテゴリ名
    description TEXT         NULL                           -- カテゴリ名
);

COMMENT ON TABLE hobby_categories IS '趣味カテゴリーテーブル';
COMMENT ON COLUMN hobby_categories.id IS 'ID';
COMMENT ON COLUMN hobby_categories.name IS 'カテゴリ名';
COMMENT ON COLUMN hobby_categories.description IS '説明';

ALTER TABLE hobbies
    ADD CONSTRAINT fk_hobbies_category_id
        FOREIGN KEY (category_id)
            REFERENCES hobby_categories (id)
            ON DELETE SET NULL;

CREATE INDEX hobby_categories_id_idx ON hobby_categories (id);