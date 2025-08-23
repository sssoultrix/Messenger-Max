CREATE TABLE IF NOT EXISTS users (
                                     id BIGSERIAL PRIMARY KEY,
                                     name TEXT NOT NULL,
                                     login TEXT UNIQUE NOT NULL,
                                     password_hash TEXT NOT NULL,
                                     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    );
