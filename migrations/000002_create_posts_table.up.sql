CREATE TABLE IF NOT EXISTS posts (
    id INTEGER UNIQUE NOT NULL, -- no need autoincrement in sqlite for old id reusage
    user_id INTEGER NOT NULL,
    title VARCHAR(512) NOT NULL,
    content VARCHAR(2048) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);