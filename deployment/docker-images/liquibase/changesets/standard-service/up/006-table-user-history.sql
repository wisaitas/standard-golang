-- Create custom enum type for action
CREATE TYPE user_action AS ENUM ('CREATE', 'UPDATE', 'DELETE');

CREATE TABLE IF NOT EXISTS user_histories (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    version integer NOT NULL DEFAULT '0',
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by uuid,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by uuid,
    deleted_at timestamp,
    action user_action NOT NULL,
    old_version integer NOT NULL,
    old_first_name varchar(100) NOT NULL,
    old_last_name varchar(100) NOT NULL,
    old_birth_date date NOT NULL,
    old_password varchar(100) NOT NULL,
    old_email varchar(100) NOT NULL
);
