CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,  
    message text NOT NULL UNIQUE,
    hash text GENERATED ALWAYS AS (encode(sha256(message::bytea), 'hex')) STORED
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

