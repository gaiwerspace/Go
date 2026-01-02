CREATE TABLE keepers (id SERIAL PRIMARY KEY, name VARCHAR(255) NOT NULL);
CREATE TABLE vaults (id SERIAL PRIMARY KEY, keeper_id INT NOT NULL REFERENCES keepers(id), balance DECIMAL(20,2) NOT NULL DEFAULT 0 CHECK (balance >= 0), created_at TIMESTAMP NOT NULL DEFAULT NOW());
CREATE TABLE transfers (id SERIAL PRIMARY KEY, source_vault_id INT NOT NULL REFERENCES vaults(id), destination_vault_id INT NOT NULL REFERENCES vaults(id), amount DECIMAL(20,2) NOT NULL CHECK (amount > 0), created_at TIMESTAMP NOT NULL DEFAULT NOW(), CHECK (source_vault_id != destination_vault_id));
INSERT INTO keepers (id, name) VALUES (1, 'Thorin Oakenshield'), (2, 'Balin Son of Fundin'), (3, 'Dáin Ironfoot'), (4, 'Gimli Son of Glóin') ON CONFLICT DO NOTHING;
