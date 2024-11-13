-- vector extension
CREATE EXTENSION IF NOT EXISTS vector;

-- branches table
CREATE TABLE IF NOT EXISTS branches (
    id SERIAL PRIMARY KEY,  -- pk
    name TEXT NOT NULL UNIQUE
);

-- parties table
CREATE TABLE IF NOT EXISTS parties (
    id SERIAL PRIMARY KEY,  -- pk
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,  -- pk
    cip TEXT NOT NULL UNIQUE,
    dni TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    first_surname TEXT NOT NULL,
    second_surname TEXT NOT NULL,
    email TEXT NOT NULL,
    branch_id INTEGER NOT NULL, -- fk
    role TEXT CHECK(role IN ('ADMIN', 'VOTER', 'MONITOR')) NOT NULL,
    attended INTEGER NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (branch_id) REFERENCES branches(id)
);

-- user feature vectors table
CREATE TABLE IF NOT EXISTS user_feature_vectors (
    id SERIAL PRIMARY KEY,  -- pk
    user_id INTEGER NOT NULL,  -- fk users
    feature_vector VECTOR(128),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- candidates table
CREATE TABLE IF NOT EXISTS candidates (
    id SERIAL PRIMARY KEY,  -- pk
    cip TEXT NOT NULL UNIQUE,
    dni TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    first_surname TEXT NOT NULL,
    second_surname TEXT NOT NULL,
    party_id INTEGER NOT NULL,  -- fk parties
    branch_id INTEGER NOT NULL,  -- fk branches
    position_applied TEXT NOT NULL,
    previous_experience TEXT,
    academic_background TEXT,
    additional_info TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (party_id) REFERENCES parties(id) ON DELETE CASCADE,
    FOREIGN KEY (branch_id) REFERENCES branches(id)
);

-- candidate embeddings table
CREATE TABLE IF NOT EXISTS candidate_embeddings (
    id SERIAL PRIMARY KEY,  -- pk
    candidate_id INTEGER NOT NULL,  -- fk candidates
    embedding vector(1536) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (candidate_id) REFERENCES candidates(id) ON DELETE CASCADE
);

CREATE TABLE config (
    id SERIAL PRIMARY KEY,  -- pk
    election_start TIMESTAMP NOT NULL,
    election_end TIMESTAMP NOT NULL,

    created_by INTEGER NOT NULL,  -- fk users
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (created_by) REFERENCES users(id)
);
