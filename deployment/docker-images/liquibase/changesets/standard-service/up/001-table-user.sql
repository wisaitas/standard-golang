CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    version integer NOT NULL DEFAULT '0',
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by uuid,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid,

    username varchar(100) NOT NULL UNIQUE,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    birth_date date NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    password varchar(100) NOT NULL
);
