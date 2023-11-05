CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SELECT 'CREATE DATABASE wims_db'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'wims_db')\gexec

CREATE TABLE IF NOT EXISTS wims_user
(
    id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    username text NOT NULL UNIQUE,
    password text NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

\connect wims_db;