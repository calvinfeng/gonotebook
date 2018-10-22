--name: create_dogs_table
CREATE TABLE IF NOT EXISTS dogs (
    id SERIAL,
    age INTEGER, 
    name VARCHAR(255),
    male BOOLEAN
)