--name: create_dogs_table
CREATE TABLE IF NOT EXISTS dogs (
    id INTEGER PRIMARY KEY,
    age INTEGER, 
    name VARCHAR(255),
    male BOOLEAN
)