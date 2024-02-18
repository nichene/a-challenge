-- +goose Up

-- +goose StatementBegin 
CREATE TABLE persons (
    id BIGSERIAL PRIMARY KEY,
    name varchar(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 

DROP TABLE persons;

-- +goose StatementEnd