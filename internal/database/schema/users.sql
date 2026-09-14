CREATE TABLE users (
    id TEXT PRIMARY KEY,
    hashed_password TEXT NOT NULL,
    email TEXT NOT NULL 
)