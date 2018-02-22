
DELETE FROM users;

-------USERS-------

INSERT INTO users (user_id, username, email, gravatar, password, salt)
VALUES ('ba7f171c-8bd2-4470-99de-6d82ca9402e9', 'orderfood', 'orderfood@of.com', 'd68ec67243ffcc5184afb4619d0ee447',
        '$2a$10$2zfn8HqZnE0oNnvnB9SmNeCNExnumDSNnnDGvr/KASKWktaSWWXeS', '8fcc5cfa90d2304e3b8c9f3486c5f32447227eedf15d0ebc172ef8966c36');

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

INSERT INTO place (id_place, name, phone_number, url, city, adress, user_id, type)
VALUES ('9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'RKLplace', '+79995207691', 'rklplace.com', 'Saint Petersburg','Невский пр., 92','ba7f171c-8bd2-4470-99de-6d82ca9402e9',
'68c65b87-925b-4227-bada-c543b55048e2');

-------MENUS-------

INSERT INTO menu (id_menu, name_menu, id_place, url)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', 'Новогоднее','9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'http://yberisama.ru/wp-content/uploads/2016/12/%D0%BF%D1%80%D0%B0%D0%B7%D0%B4%D0%BD%D0%B5%D1%87%D0%BD%D1%8B%D0%B9-%D0%BD%D0%BE%D0%B2%D0%BE%D0%B3%D0%BE%D0%B4%D0%BD%D0%B8%D0%B9-%D1%81%D1%82%D0%BE%D0%BB.jpg' );


-------TYPES DISHES-------

INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('403f1a53-7444-483d-9f48-c590dd476d28', 'завтраки');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('126bdeac-5dfe-42f9-8400-5e2b3090dfef', 'гарниры');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('635687a9-f5c8-4595-8f82-312fedb11f8c', 'напитки');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('fef1833e-48eb-4e46-ae74-e015451c9ee9', 'супы');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('2c567271-9bd3-405f-980d-f596e509cb0d', 'мясное');

-------DISH-------

INSERT INTO dish (id_dish, user_id,  name_dish, id_typedish, description, url, time_min)
VALUES ('31ba1f68-6eb3-4362-8d4d-24616fa9593d','ba7f171c-8bd2-4470-99de-6d82ca9402e9','Классический бургер', '2c567271-9bd3-405f-980d-f596e509cb0d', 'Лучший в мире бургер', 'http://www.goodman.ru/upload/iblock/e02/burger-with-roast-beef.png', 10);


INSERT INTO dish (id_dish, user_id,  name_dish, id_typedish, description, url, time_min)
VALUES ('8072b173-baaa-4d0f-a1c7-8af3995ffefa','ba7f171c-8bd2-4470-99de-6d82ca9402e9','Русский бургер', '2c567271-9bd3-405f-980d-f596e509cb0d', 'Лучший в мире бургер2', 'http://goodman.ru/upload/iblock/570/classic-burger.png', 10);


INSERT INTO dish (id_dish, user_id,  name_dish, id_typedish, description, url, time_min)
VALUES ('5cb7e111-3860-424b-a14e-83a93db1e889','ba7f171c-8bd2-4470-99de-6d82ca9402e9','Исландский бургер', '2c567271-9bd3-405f-980d-f596e509cb0d', 'Лучший в мире бургер3', 'http://goodman.ru/upload/iblock/f18/American-burger.png', 10);

INSERT INTO menudish (id_menu, id_dish)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', '31ba1f68-6eb3-4362-8d4d-24616fa9593d');

INSERT INTO menudish (id_menu, id_dish)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', '5cb7e111-3860-424b-a14e-83a93db1e889');


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
