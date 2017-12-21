
DELETE FROM users;

-------USERS-------

INSERT INTO users (user_id, username, email, gravatar, password, salt)
VALUES ('4f2858bb-ecab-4f88-b773-551d34f07bad', 'orderfood', 'orderfood@of.ru', 'bfb0bb1efe0e40b2ff1aebf6c2eeb0e9',
        '$2a$10$WSVzrY9EUUtWAL0QH0XeAuuvdAJvJ9RlaZ5bDUHv3f6IfXDUL5xXG', 'a560cd9d9969a000d24d4b012a755d7217c1561648ec682128abc9fe762a');

-------TYPES PLACES-------

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('68c65b87-925b-4227-bada-c543b55048e2', 'ресторан');

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('c32397bc-4ff2-40e6-a290-e8191bd2d5c0', 'кальянная');

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('b5870b67-e348-4cb9-bccb-40956363dfdd', 'бар');

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('28f8e2f1-4183-45c2-bb94-b58116af0e42', 'кафе');

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('3b896b95-928a-4107-934b-d7dace83b83a', 'кофейнная');


-------PLACES-------

INSERT INTO place (id_place, name, phone_number, url, city, adress, user_id, id_typeplace)
VALUES ('9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'Бургерная Lil_ASCII', '+79999', 'kuzmrom7.github.io', 'Saint-P','blavla','4f2858bb-ecab-4f88-b773-551d34f07bad',
'68c65b87-925b-4227-bada-c543b55048e2');

-------MENUS-------

INSERT INTO menu (id_menu, name_menu, id_place)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', 'Летнее меню','9ba48d7c-b573-4dcb-b8bb-fbb196753231' );

-------SECTIONS MENUS-------

INSERT INTO section_menu (id_section, id_menu)
VALUES ('771570df-cc44-4af5-a2b0-ee19b9f87cd9', 'a93ba633-7547-491f-a4ae-339b1420b1c7');

-------TYPES DISHES-------

INSERT INTO type_dish (id_typedish, name_typedish, id_section)
VALUES ('041529e0-3111-4993-a75b-7a12545a0d9f', 'Бургеры','771570df-cc44-4af5-a2b0-ee19b9f87cd9');

-------DISH-------

INSERT INTO dish (id_dish, name_dish, id_typedish, description)
VALUES ('31ba1f68-6eb3-4362-8d4d-24616fa9593d','Классический бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер');


INSERT INTO dish (id_dish, name_dish, id_typedish, description)
VALUES ('8072b173-baaa-4d0f-a1c7-8af3995ffefa','Русский бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер');


INSERT INTO dish (id_dish, name_dish, id_typedish, description)
VALUES ('5cb7e111-3860-424b-a14e-83a93db1e889','Исландский бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер');


INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('ad64c48a-33f4-4cb7-abe6-4da9e56a403d', 'охранник');

INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('bbad620b-7939-4d82-853f-5d1fe13a8629', 'администратор');

INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('a345fc9a-46d3-4fed-b562-dbf9584f8058', 'официант');

INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('53ab0358-d085-4e90-a66a-4d0f6e3dfa18', 'кальянщик');

INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('6f8ee4df-ce23-4e1e-9c21-87d163075785', 'бармен');

INSERT INTO type_personal (id_typePersonal, name_type)
VALUES ('7d826de7-97de-4002-8555-6a8056fa3faa', 'техперсонал');
------EXAMPLE SELECT------

/*SELECT u.user_id, u.username, p.city, m.name_menu, d.name_dish
FROM (((((users u
        INNER JOIN place p ON  u.user_id = p.user_id)
        INNER JOIN menu m ON p.id_place = m.id_place)
        INNER JOIN section_menu s ON m.id_menu = s.id_menu)
        INNER JOIN type_dish t ON s.id_section = t.id_section)
        INNER JOIN dish d ON t.id_typedish = d.id_typedish)
;*/
