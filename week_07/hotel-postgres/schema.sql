-- schema.sql
CREATE TABLE IF NOT EXISTS rooms (
    id    TEXT PRIMARY KEY,
    name  TEXT NOT NULL,
    price FLOAT NOT NULL
);
