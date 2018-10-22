-- name: find_dogs
SELECT * FROM dogs;

-- name: insert_dog
INSERT INTO dogs (age, name, male) VALUES ($1, $2, $3);
