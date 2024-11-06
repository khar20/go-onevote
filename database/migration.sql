-- Tabla Users
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role TEXT CHECK(role IN ('ADMINISTRATOR', 'BOARD_MEMBER', 'CANDIDATE', 'VOTER')) NOT NULL,
    dni TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    mother_last_name TEXT NOT NULL,
    votes INTEGER NOT NULL DEFAULT 0,
    attended INTEGER NOT NULL DEFAULT 0,
    email TEXT NOT NULL,
    created_at INTEGER NOT NULL
);

-- Tabla Parties
CREATE TABLE parties (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    symbol TEXT,
    created_at INTEGER NOT NULL
);

-- Tabla Candidates
CREATE TABLE candidates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    dni TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    mother_last_name TEXT NOT NULL,
    email TEXT NOT NULL,
    party_id INTEGER NOT NULL,
    campaign_statement TEXT,
    created_at INTEGER NOT NULL,
    FOREIGN KEY (party_id) REFERENCES parties(id) ON DELETE CASCADE
);

-- Tabla Votes
CREATE TABLE votes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    candidate_id INTEGER NOT NULL,
    data_hash TEXT NOT NULL,
    created_at INTEGER NOT NULL,
    FOREIGN KEY (candidate_id) REFERENCES candidates(id) ON DELETE CASCADE
);
