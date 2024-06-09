CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ,
    username TEXT NOT NULL,
    organization TEXT NOT NULL,
    email TEXT,
    phone TEXT,
    a_password TEXT NOT NULL,
    last_password_updated TIMESTAMP,
    last_login TIMESTAMP,
    a_status TEXT NOT NULL,
    first_name TEXT,
    last_name TEXT,
    birthday TIMESTAMP,
    gender TEXT,
    avatar TEXT,

    CONSTRAINT uk_user UNIQUE (username, organization)
);

CREATE INDEX IF NOT EXISTS user_username_idx ON users (username);
CREATE INDEX IF NOT EXISTS user_email_idx ON users (email);
CREATE INDEX IF NOT EXISTS user_org_username_idx ON users (organization, username);
CREATE INDEX IF NOT EXISTS user_org_email_idx ON users (organization, email);

CREATE TABLE IF NOT EXISTS users_login_history (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    ip_address TEXT NOT NULL,
    user_agent TEXT NOT NULL,
    login_time TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS user_login_history_user_idx ON users_login_history (user_id);

CREATE TABLE IF NOT EXISTS oauth_providers (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ,
    provider_name TEXT NOT NULL,
    provider_display_name TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS oauth_providers_idx ON oauth_providers (provider_name);

CREATE TABLE IF NOT EXISTS users_oauth_accounts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    provider_id BIGINT NOT NULL,
    oauth_id TEXT NOT NULL, -- id from oauth provider
    email TEXT NOT NULL,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_provider FOREIGN KEY (provider_id) REFERENCES oauth_providers (id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS users_oauth_account_user_idx ON users_oauth_accounts (user_id);
CREATE INDEX IF NOT EXISTS users_oauth_account_provider_idx ON users_oauth_accounts (provider_id);
CREATE INDEX IF NOT EXISTS users_oauth_account_oauth_idx ON users_oauth_accounts (oauth_id);
CREATE INDEX IF NOT EXISTS users_oauth_account_user_provider_idx ON users_oauth_accounts (user_id, provider_id);
