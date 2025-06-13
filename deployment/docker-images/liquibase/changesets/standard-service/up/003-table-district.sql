CREATE TABLE IF NOT EXISTS districts (
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
    province_id uuid NOT NULL
);

ALTER TABLE districts ADD CONSTRAINT fk_districts_provinces FOREIGN KEY (province_id) REFERENCES provinces(id);