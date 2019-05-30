CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    sender_id INTEGER REFERENCES users(id),
    receiver_id INTEGER REFERENCES users(id),
    body TEXT
);

CREATE INDEX ON messages(sender_id);
CREATE INDEX ON messages(receiver_id);