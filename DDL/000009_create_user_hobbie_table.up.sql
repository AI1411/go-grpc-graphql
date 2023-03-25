DROP TABLE IF EXISTS user_hobbies CASCADE;
CREATE TABLE user_hobbies
(
    id       UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- ID
    user_id  UUID NOT NULL,                              -- ユーザID
    hobby_id UUID NOT NULL                               -- 趣味ID
);

ALTER TABLE user_hobbies
    ADD CONSTRAINT fk_user_hobbies_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE;

ALTER TABLE user_hobbies
    ADD CONSTRAINT fk_user_hobbies_hobby_id
        FOREIGN KEY (hobby_id)
            REFERENCES hobbies (id)
            ON DELETE CASCADE;

COMMENT ON TABLE user_hobbies IS 'ユーザと趣味の関連テーブル';
COMMENT ON COLUMN user_hobbies.id IS 'ID';
COMMENT ON COLUMN user_hobbies.user_id IS 'ユーザID';
COMMENT ON COLUMN user_hobbies.hobby_id IS '趣味ID';

CREATE INDEX user_hobbies_user_id_idx ON user_hobbies (user_id);
CREATE INDEX user_hobbies_hobby_id_idx ON user_hobbies (hobby_id);