CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE,
    phone TEXT UNIQUE,
    a_password TEXT NOT NULL,
    last_password_updated TIMESTAMP,
    last_login TIMESTAMP,
    a_status TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    birthday TIMESTAMP,
    gender TEXT,
    avatar TEXT
);

CREATE TABLE IF NOT EXISTS users_login_history (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    login_time TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);