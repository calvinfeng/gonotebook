CREATE TABLE customers(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255)
);

CREATE UNIQUE INDEX ON customers(name);

CREATE TABLE listings(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE, 
    name VARCHAR(255),
    category VARCHAR(255),
    asking_price real,
    sale_price real,
    sold BOOLEAN
);

CREATE INDEX ON listings(name);
CREATE INDEX ON listings(category);

CREATE TABLE biddings(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    price real,
    customer_id INTEGER REFERENCES customers(id),
    listing_id INTEGER REFERENCES listings(id)
);

CREATE UNIQUE INDEX ON biddings(customer_id, listing_id);