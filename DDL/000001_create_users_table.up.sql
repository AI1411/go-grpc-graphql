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
    occupation             VARCHAR(100),                                                -- 職業
    education              VARCHAR(100),                                                -- 学歴
    hobbies_and_skills     TEXT,                                                        -- 趣味・スキル
    personality            TEXT,                                                        -- 性格
    purpose_of_interaction TEXT,                                                        -- 交流目的
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