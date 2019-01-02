CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE, 
    user_id INTEGER REFERENCES users(id),
    body TEXT
);

CREATE INDEX ON messages(user_id);