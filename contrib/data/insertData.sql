
DELETE FROM users;

INSERT INTO users (user_id, username, email, gravatar, password, salt)
VALUES ('4f2858bb-ecab-4f88-b773-551d34f07bad', 'orderfood', 'orderfood@of.ru', 'bfb0bb1efe0e40b2ff1aebf6c2eeb0e9',
        '$2a$10$WSVzrY9EUUtWAL0QH0XeAuuvdAJvJ9RlaZ5bDUHv3f6IfXDUL5xXG', 'a560cd9d9969a000d24d4b012a755d7217c1561648ec682128abc9fe762a');

INSERT INTO users (user_id, username, email)
VALUES ('a1de17c9-7bac-4803-8b4c-44353ed4ed02', 'test', 'test@of.ru');

/*
INSERT INTO place (id_place, name, phone_number, url, city, adress, user_id, id_typeplace, created, updated)
VALUES ();
*/
