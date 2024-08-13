CREATE TABLE IF NOT EXISTS organizations (
    identifier TEXT PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    a_name TEXT NOT NULL,
    a_status TEXT NOT NULL,
    an_address TEXT,
    contact_email TEXT,
    contact_phone TEXT
);

CREATE INDEX IF NOT EXISTS org_status_idx ON organizations (a_status);

-- insert system organization

INSERT INTO organizations (identifier, a_name, a_status) VALUES ('system', 'System Organization', 'system') ON CONFLICT DO NOTHING;

-- 4.6.1. user mgmt

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    username TEXT NOT NULL,
    organization_id TEXT NOT NULL,
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

    CONSTRAINT fk_organization FOREIGN KEY (organization_id) REFERENCES organizations(identifier) ON DELETE CASCADE,
    CONSTRAINT uk_user UNIQUE (username, organization_id)
);

CREATE INDEX IF NOT EXISTS user_username_idx ON users (username);
CREATE INDEX IF NOT EXISTS user_email_idx ON users (email);
CREATE INDEX IF NOT EXISTS user_org_username_idx ON users (organization_id, username);
CREATE INDEX IF NOT EXISTS user_org_email_idx ON users (organization_id, email);
CREATE INDEX IF NOT EXISTS user_status_idx ON users (a_status);
CREATE INDEX IF NOT EXISTS user_org_status_idx ON users (organization_id, a_status);

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

CREATE TABLE IF NOT EXISTS clients (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    client_id TEXT NOT NULL,
    client_secret TEXT NOT NULL,
    client_name TEXT NOT NULL,
    grant_types TEXT NOT NULL, -- space seperated value
    scope TEXT, -- space seperated value
    redirect_uris TEXT NOT NULL, -- in json
    post_logout_redirect_uris TEXT,

    CONSTRAINT uk_client UNIQUE (client_id)
);

CREATE INDEX IF NOT EXISTS client_client_id_idx ON clients (client_id);

CREATE TABLE IF NOT EXISTS authorization_codes (
    code TEXT PRIMARY KEY,
    user_identifier TEXT NOT NULL,
    org_id TEXT NOT NULL,
    client_id TEXT NOT NULL,
    code_challenge TEXT,
    code_challenge_method TEXT,
    redirect_uri TEXT,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_org FOREIGN KEY (org_id) REFERENCES organizations (identifier) ON DELETE CASCADE,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients (client_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS refresh_tokens (
    token TEXT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    client_id TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients (client_id) ON DELETE CASCADE
);
