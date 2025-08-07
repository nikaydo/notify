-- +goose Up
-- +goose StatementBegin
ALTER TABLE events ADD COLUMN creator UUID;  
ALTER TABLE events ADD COLUMN created TIMESTAMPTZ NOT NULL DEFAULT now();  
ALTER TABLE events ADD COLUMN expaired TIMESTAMPTZ;  
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE userSettings DROP COLUMN creator;
ALTER TABLE userSettings DROP COLUMN created;
ALTER TABLE userSettings DROP COLUMN expaired;
-- +goose StatementEnd
