DROP TABLE IF EXISTS users CASCADE;
DROP TYPE IF EXISTS user_status;
DROP TYPE IF EXISTS prefecture;
DROP TYPE IF EXISTS blood_type;
-- user_status型: ユーザのステータスを表すenum型
CREATE TYPE user_status AS ENUM ('通常会員', '退会済', 'アカウント停止', 'プレミアム');
-- prefecture型: 都道府県名を表すenum型
CREATE TYPE prefecture AS ENUM (
    'ひみつにする',
    '北海道',
    '青森県',
    '岩手県',
    '宮城県',
    '秋田県',
    '山形県',
    '福島県',
    '茨城県',
    '栃木県',
    '群馬県',
    '埼玉県',
    '千葉県',
    '東京都',
    '神奈川県',
    '新潟県',
    '富山県',
    '石川県',
    '福井県',
    '山梨県',
    '長野県',
    '岐阜県',
    '静岡県',
    '愛知県',
    '三重県',
    '滋賀県',
    '京都府',
    '大阪府',
    '兵庫県',
    '奈良県',
    '和歌山県',
    '鳥取県',
    '島根県',
    '岡山県',
    '広島県',
    '山口県',
    '徳島県',
    '香川県',
    '愛媛県',
    '高知県',
    '福岡県',
    '佐賀県',
    '長崎県',
    '熊本県',
    '大分県',
    '宮崎県',
    '鹿児島県',
    '沖縄県',
    '海外'
    );
-- blood_type型: 血液型を表すenum型
CREATE TYPE blood_type AS ENUM (
    'ひみつにする',
    'A型',
    'B型',
    'O型',
    'AB型'
    );
CREATE TABLE users
(
    id                     UUID PRIMARY KEY                  DEFAULT gen_random_uuid(), -- ID
    username               VARCHAR(100)             NOT NULL,                           -- ユーザ名
    email                  VARCHAR(100)             NOT NULL UNIQUE,                    -- メールアドレス
    "password"             TEXT                     NOT NULL,                           -- パスワード
    status                 user_status              NOT NULL DEFAULT '通常会員',        -- ユーザステータス
    prefecture             prefecture               NOT NULL DEFAULT 'ひみつにする',          -- 都道府県
    introduction           TEXT,                                                        -- 自己紹介
    blood_type             blood_type               NOT NULL DEFAULT 'ひみつにする',          -- 血液型
    occupation             VARCHAR(100)             NULL,                               -- 職業
    education              VARCHAR(100)             NULL,                               -- 学歴
    hobbies_and_skills     TEXT                     NULL,                               -- 趣味・スキル
    personality            TEXT                     NULL,                               -- 性格
    purpose_of_interaction TEXT                     NULL,                               -- 交流目的
    created_at             TIMESTAMP
                               WITH TIME ZONE       NOT NULL DEFAULT NOW(),             -- 作成日時
    updated_at             TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()              -- 更新日時
);

COMMENT ON TABLE users IS 'ユーザテーブル';
COMMENT ON COLUMN users.id IS 'ID';
COMMENT ON COLUMN users.username IS 'ユーザ名';
COMMENT ON COLUMN users.email IS 'メールアドレス';
COMMENT ON COLUMN users.password IS 'パスワード';
COMMENT ON COLUMN users.status IS 'ユーザステータス
    通常会員 = 1;
    退会済 = 2;
    アカウント停止 = 3;
    プレミアム = 4;';
COMMENT ON COLUMN users.prefecture IS '都道府県';
COMMENT ON COLUMN users.introduction IS '自己紹介';
COMMENT ON COLUMN users.occupation IS '職業';
COMMENT ON COLUMN users.education IS '学歴';
COMMENT ON COLUMN users.hobbies_and_skills IS '趣味・特技';
COMMENT ON COLUMN users.personality IS '性格';
COMMENT ON COLUMN users.purpose_of_interaction IS '交流目的';
COMMENT ON COLUMN users.created_at IS '作成日時';
COMMENT ON COLUMN users.updated_at IS '更新日時';

CREATE INDEX id_idx ON users (id);
CREATE INDEX username_idx ON users (username);
CREATE INDEX status_idx ON users (status);
CREATE INDEX created_at_idx ON users (created_at);
CREATE INDEX occupation_idx ON users (occupation);
CREATE INDEX education_idx ON users (education);

DROP TABLE IF EXISTS chats CASCADE;
CREATE TABLE IF NOT EXISTS chats
(
    id           UUID PRIMARY KEY                  DEFAULT gen_random_uuid(), -- ID
    room_id      UUID                     NOT NULL,                           -- ルームID
    from_user_id UUID                     NOT NULL,                           -- 送り元ユーザID
    to_user_id   UUID                     NOT NULL,                           -- 送り先ユーザID
    body         VARCHAR(255)             NOT NULL,                           -- 本文
    is_read      BOOLEAN                  NOT NULL DEFAULT FALSE,             -- 既読フラグ
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),             -- 作成日時
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()              -- 更新日時
);

COMMENT ON TABLE chats IS 'チャットテーブル';
COMMENT ON COLUMN chats.id IS 'ID';
COMMENT ON COLUMN chats.room_id IS '送り元ユーザID';
COMMENT ON COLUMN chats.from_user_id IS '送り元ユーザID';
COMMENT ON COLUMN chats.to_user_id IS '送り先ユーザID';
COMMENT ON COLUMN chats.body IS '本文';
COMMENT ON COLUMN chats.IS_read IS '既読フラグ';
COMMENT ON COLUMN chats.created_at IS '作成日時';
COMMENT ON COLUMN chats.updated_at IS '更新日時';

CREATE INDEX IF NOT EXISTS room_id_idx on chats (room_id);
CREATE INDEX IF NOT EXISTS from_user_id_idx on chats (from_user_id);
CREATE INDEX IF NOT EXISTS to_user_id_idx on chats (to_user_id);
CREATE INDEX IF NOT EXISTS is_read_idx on chats (is_read);
CREATE INDEX IF NOT EXISTS created_at_idx on chats (created_at);

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
    ADD CONSTRAINT chats_room_id_fkey FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE;

DROP TABLE IF EXISTS tweets CASCADE;
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

DROP TABLE IF EXISTS user_logins CASCADE;
CREATE TABLE user_logins
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY, -- UUIDを主キーに設定し、デフォルトでuuidv4を生成
    user_id    UUID      NOT NULL,                         -- UUIDを使用
    login_date DATE      NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- 外部キー制約を追加
ALTER TABLE user_logins
    ADD CONSTRAINT fk_user_logins_user_id
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- UNIQUE制約を追加
ALTER TABLE user_logins
    ADD CONSTRAINT uq_user_logins_user_id_login_date
        UNIQUE (user_id, login_date);

CREATE INDEX idx_user_logins_user_id_login_date
    ON user_logins (user_id, login_date);

-- 日本語コメント
COMMENT ON TABLE user_logins IS 'ユーザーの毎日のログインポイントを追跡するためのテーブル';
COMMENT ON COLUMN user_logins.id IS 'user_loginsテーブルの主キー';
COMMENT ON COLUMN user_logins.user_id IS 'usersテーブルを参照する外部キー';
COMMENT ON COLUMN user_logins.login_date IS 'ユーザーがログインした日付。毎日のログインを追跡するために使用されます';
COMMENT ON COLUMN user_logins.created_at IS 'レコードが作成されたタイムスタンプ';
COMMENT ON COLUMN user_logins.updated_at IS 'レコードが最後に更新されたタイムスタンプ';

DROP TABLE IF EXISTS user_point_histories;
CREATE TABLE user_point_histories
(
    id             UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id        UUID         NOT NULL,
    point          INTEGER      NOT NULL,
    operation_type VARCHAR(255) NOT NULL, -- 'add' or 'subtract'など
    description    TEXT,                  -- ポイント操作の説明（任意）
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NOT NULL
);

-- 外部キー制約を追加
ALTER TABLE user_point_histories
    ADD CONSTRAINT fk_user_point_histories_user_id
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;

-- インデックスを追加
CREATE INDEX idx_user_point_histories_user_id_created_at
    ON user_point_histories (user_id, created_at);

-- 日本語のコメント
COMMENT ON TABLE user_point_histories IS 'ユーザーのポイント履歴を追跡するためのテーブル';
COMMENT ON COLUMN user_point_histories.id IS 'user_point_historiesテーブルの主キー';
COMMENT ON COLUMN user_point_histories.user_id IS 'usersテーブルを参照する外部キー';
COMMENT ON COLUMN user_point_histories.point IS 'ポイント操作の金額（正または負）';
COMMENT ON COLUMN user_point_histories.operation_type IS 'ポイント操作の種類（例：追加、減算）';
COMMENT ON COLUMN user_point_histories.description IS 'ポイント操作の説明（任意）';
COMMENT ON COLUMN user_point_histories.created_at IS 'レコードが作成されたタイムスタンプ';
COMMENT ON COLUMN user_point_histories.updated_at IS 'レコードが最後に更新されたタイムスタンプ';

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