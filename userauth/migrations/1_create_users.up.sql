CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,  
    name VARCHAR(255),
    email VARCHAR(255),
    session_token VARCHAR(255),
    password_digest BYTEA
);

CREATE UNIQUE INDEX ON users(name);
CREATE UNIQUE INDEX on users(email);
CREATE UNIQUE INDEX ON users(session_token);