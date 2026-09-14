CREATE TABLE refresh_tokens (
    id TEXT PRIMARY KEY,
    expires_at TEXT NOT NULL ,
    token TEXT NOT NULL UNIQUE,
    user_id TEXT NOT NULL REFERENCES users(id),
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
)