CREATE TABLE IF NOT EXISTS addresses (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    version integer NOT NULL DEFAULT '0',
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by uuid,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid,
    
    address varchar(400),
    province_id uuid NOT NULL,
    district_id uuid NOT NULL,
    sub_district_id uuid NOT NULL,
    user_id uuid NOT NULL
);