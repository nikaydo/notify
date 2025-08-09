-- +goose Up
-- +goose StatementBegin
ALTER TABLE events ADD COLUMN creator UUID;  
ALTER TABLE events ADD COLUMN created TIMESTAMPTZ DEFAULT now();  
ALTER TABLE events ADD COLUMN expaired TIMESTAMPTZ;  

CREATE TABLE events_members (
    uuid_event UUID REFERENCES events(uuid) ON DELETE CASCADE,
    uuid_user UUID REFERENCES users(uuid) ON DELETE CASCADE,
    PRIMARY KEY (uuid_event, uuid_user)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE events DROP COLUMN creator;
ALTER TABLE events DROP COLUMN created;
ALTER TABLE events DROP COLUMN expaired;

DROP TABLE events_members;

-- +goose StatementEnd
