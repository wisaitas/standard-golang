ALTER TABLE addresses DROP CONSTRAINT IF EXISTS fk_addresses_users;
ALTER TABLE addresses DROP CONSTRAINT IF EXISTS fk_addresses_provinces;
ALTER TABLE addresses DROP CONSTRAINT IF EXISTS fk_addresses_districts;
ALTER TABLE addresses DROP CONSTRAINT IF EXISTS fk_addresses_sub_districts;
DROP TABLE IF EXISTS addresses;
