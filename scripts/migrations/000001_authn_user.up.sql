CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT,
    phone TEXT,
    a_password TEXT NOT NULL,
    salt TEXT NOT NULL,
    a_status TEXT NOT NULL,
    last_login TIMESTAMP,
    last_login_ip TEXT,
    first_name TEXT,
    last_name TEXT,
    birthday TIMESTAMP,
    gender TEXT,
    avatar TEXT
);