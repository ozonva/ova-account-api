-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts
(
    id         uuid         not null
        constraint accounts_pkey primary key,
    value      varchar(255) not null,
    user_id    bigint       not null,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts CASCADE;
-- +goose StatementEnd
