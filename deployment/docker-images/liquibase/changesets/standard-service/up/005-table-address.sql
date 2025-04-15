CREATE TABLE IF NOT EXISTS addresses (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    version integer NOT NULL DEFAULT '0',
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by uuid,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by uuid,
    deleted_at timestamp,
    address varchar(400),
    province_id uuid NOT NULL,
    district_id uuid NOT NULL,
    sub_district_id uuid NOT NULL,
    user_id uuid NOT NULL
);

ALTER TABLE addresses ADD CONSTRAINT fk_addresses_provinces FOREIGN KEY (province_id) REFERENCES provinces(id);
ALTER TABLE addresses ADD CONSTRAINT fk_addresses_districts FOREIGN KEY (district_id) REFERENCES districts(id);
ALTER TABLE addresses ADD CONSTRAINT fk_addresses_sub_districts FOREIGN KEY (sub_district_id) REFERENCES sub_districts(id);
ALTER TABLE addresses ADD CONSTRAINT fk_addresses_users FOREIGN KEY (user_id) REFERENCES users(id);