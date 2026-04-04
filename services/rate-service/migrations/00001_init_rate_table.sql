-- +goose Up
CREATE TABLE IF NOT EXISTS rates (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ask        DOUBLE PRECISION NOT NULL,
    bid        DOUBLE PRECISION NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_rates_created_at ON rates (created_at);

-- +goose Down
DROP TABLE IF EXISTS rates;
