CREATE TABLE IF NOT EXISTS sub_districts (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    version integer NOT NULL DEFAULT '0',
    created_at timestamp NOT NULL DEFAULT NOW(),
    created_by uuid,
    updated_at timestamp NOT NULL DEFAULT NOW(),
    updated_by uuid,
    deleted_at timestamp,
    deleted_by uuid,
    
    name_th varchar(100) NOT NULL,
    name_en varchar(100) NOT NULL,
    district_id uuid NOT NULL,
    postal_code varchar(10) NOT NULL
);

ALTER TABLE sub_districts ADD CONSTRAINT fk_sub_districts_districts FOREIGN KEY (district_id) REFERENCES districts(id);