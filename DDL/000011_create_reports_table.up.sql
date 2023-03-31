DROP TABLE IF EXISTS reports CASCADE;

DROP TYPE IF EXISTS report_status CASCADE;
CREATE TYPE report_status AS ENUM (
    'PENDING',
    'REJECTED',
    'ACCEPTED'
    );

CREATE TABLE reports
(
    id               UUID                   DEFAULT gen_random_uuid() PRIMARY KEY,
    reporter_user_id UUID          NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    reported_user_id UUID          NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    reported_chat_id UUID          NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    status           report_status NOT NULL DEFAULT 'PENDING',
    reason           VARCHAR(255)  NOT NULL,
    created_at       TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE INDEX reports_reporter_id_idx ON reports (reporter_user_id);
CREATE UNIQUE INDEX reports_reported_user_id_reported_chat_id_idx ON reports (reporter_user_id, reported_chat_id);

COMMENT ON TABLE reports IS '通報テーブル';
COMMENT ON COLUMN reports.id IS '通報ID';
COMMENT ON COLUMN reports.reporter_user_id IS '通報者ID';
COMMENT ON COLUMN reports.reported_chat_id IS '通報されたチャットID';
COMMENT ON COLUMN reports.reason IS '通報理由';
COMMENT ON COLUMN reports.created_at IS '作成日時';
COMMENT ON COLUMN reports.updated_at IS '更新日時';
