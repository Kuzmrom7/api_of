
DELETE FROM users;

-------USERS-------

INSERT INTO users (user_id, username, email, gravatar, password, salt)
VALUES ('4f2858bb-ecab-4f88-b773-551d34f07bad', 'orderfood', 'orderfood@of.ru', 'bfb0bb1efe0e40b2ff1aebf6c2eeb0e9',
        '$2a$10$WSVzrY9EUUtWAL0QH0XeAuuvdAJvJ9RlaZ5bDUHv3f6IfXDUL5xXG', 'a560cd9d9969a000d24d4b012a755d7217c1561648ec682128abc9fe762a');

-------TYPES PLACES-------

INSERT INTO type_place (id_typeplace, name_type)
VALUES ('68c65b87-925b-4227-bada-c543b55048e2', 'Ресторан');

-------PLACES-------

INSERT INTO place (id_place, name, phone_number, url, city, adress, user_id, id_typeplace)
VALUES ('9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'Бургерная Lil_ASCII', '+79999', 'kuzmrom7.github.io', 'Saint-P','blavla','4f2858bb-ecab-4f88-b773-551d34f07bad',
'68c65b87-925b-4227-bada-c543b55048e2');

-------MENUS-------

INSERT INTO menu (id_menu,user_id, name_menu, id_place)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7','4f2858bb-ecab-4f88-b773-551d34f07bad', 'Летнее меню','9ba48d7c-b573-4dcb-b8bb-fbb196753231' );

-------SECTIONS MENUS-------

INSERT INTO section_menu (id_section, id_menu, user_id)
VALUES ('771570df-cc44-4af5-a2b0-ee19b9f87cd9', 'a93ba633-7547-491f-a4ae-339b1420b1c7','4f2858bb-ecab-4f88-b773-551d34f07bad');

-------TYPES DISHES-------

INSERT INTO type_dish (id_typedish, name_typedish, id_section,user_id)
VALUES ('041529e0-3111-4993-a75b-7a12545a0d9f', 'Бургеры','771570df-cc44-4af5-a2b0-ee19b9f87cd9','4f2858bb-ecab-4f88-b773-551d34f07bad');

-------DISH-------

INSERT INTO dish (id_menu, name_dish, id_typedish, description,user_id)
VALUES ('31ba1f68-6eb3-4362-8d4d-24616fa9593d','Классический бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер','4f2858bb-ecab-4f88-b773-551d34f07bad');


INSERT INTO dish (id_menu, name_dish, id_typedish, description,user_id)
VALUES ('8072b173-baaa-4d0f-a1c7-8af3995ffefa','Русский бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер','4f2858bb-ecab-4f88-b773-551d34f07bad');


INSERT INTO dish (id_menu, name_dish, id_typedish, description,user_id)
VALUES ('5cb7e111-3860-424b-a14e-83a93db1e889','Исландский бургер', '041529e0-3111-4993-a75b-7a12545a0d9f', 'Лучший в мире бургер','4f2858bb-ecab-4f88-b773-551d34f07bad');

------EXAMPLE SELECT------

/*SELECT u.user_id, u.username, p.city, m.name_menu, d.name_dish
FROM (((((users u
        INNER JOIN place p ON  u.user_id = p.user_id)
        INNER JOIN menu m ON p.id_place = m.id_place)
        INNER JOIN section_menu s ON m.id_menu = s.id_menu)
        INNER JOIN type_dish t ON s.id_section = t.id_section)
        INNER JOIN dish d ON t.id_typedish = d.id_typedish)
;*/
