-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_states
(
    id                  INTEGER PRIMARY KEY,
    telegram_id         VARCHAR(64),
    state               VARCHAR(512),
    created_at          TIMESTAMP,
    updated_at          TIMESTAMP
);

--CREATE INDEX IF NOT EXISTS idx_ho_brokerId ON hub_outputs (broker_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
--DROP INDEX IF EXISTS idx_ho_brokerId;

DROP TABLE IF EXISTS user_states;
-- +goose StatementEnd
