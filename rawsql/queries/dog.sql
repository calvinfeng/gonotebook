-- name: find_dogs
SELECT * FROM dogs;

-- name: insert_dog
INSERT INTO dogs (age, name, male) VALUES (?, ?, ?);
