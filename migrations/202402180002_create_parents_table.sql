-- +goose Up

-- +goose StatementBegin 
CREATE TABLE parents (
    id BIGSERIAL PRIMARY KEY,
    person_id int,
    parent_id int,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT fk_person_id FOREIGN KEY (person_id) REFERENCES persons(id),
    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES persons(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin 

DROP TABLE parents;

-- +goose StatementEnd