-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table food_order (
    id          text           primary key      not null,
    menu        text                            not null,
    created_at  timestamp                       not null,
    updated_at  timestamp                       not null,
    deleted_at  timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table order;
