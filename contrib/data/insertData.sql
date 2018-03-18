DELETE FROM users;

-------USERS-------

INSERT INTO users (user_id, username, email, gravatar, password, salt, type)
VALUES ('ba7f171c-8bd2-4470-99de-6d82ca9402e9', 'orderfood', 'orderfood@of.com', 'd68ec67243ffcc5184afb4619d0ee447',
        '$2a$10$2zfn8HqZnE0oNnvnB9SmNeCNExnumDSNnnDGvr/KASKWktaSWWXeS',
        '8fcc5cfa90d2304e3b8c9f3486c5f32447227eedf15d0ebc172ef8966c36', 'owner');

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

INSERT INTO place (id_place, name, phone_number, url, city, user_id, type)
VALUES ('9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'ORDplace', '+79995207691', 'rklplace.com', 'Saint Petersburg',
        'ba7f171c-8bd2-4470-99de-6d82ca9402e9',
        '[
          {
            "id": "7d826de7-97de-4002-8555-6a8056fa3faa",
            "nametype": "кафе"
          },
          {
            "id": "7d826de7-97de-4002-8555-6a8056fa3faa",
            "nametype": "бар"
          }
        ]');

-------MENUS-------

INSERT INTO menu (id_menu, name_menu, id_place, url)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', 'Основное', '9ba48d7c-b573-4dcb-b8bb-fbb196753231',
        'http://res.cloudinary.com/dwkkf6qmg/image/upload/v1520968541/zmj5z5qepk2qjp1trqft.jpg');
INSERT INTO menu (id_menu, name_menu, id_place, url)
VALUES ('f0f6738b-5437-4a89-b715-9cd1d00173de', 'Напитки', '9ba48d7c-b573-4dcb-b8bb-fbb196753231',
        'http://res.cloudinary.com/dwkkf6qmg/image/upload/v1520968532/nzqophlzi64dcj3uwsrn.jpg');

-------TYPES DISHES-------

INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('403f1a53-7444-483d-9f48-c590dd476d28', 'завтраки');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('126bdeac-5dfe-42f9-8400-5e2b3090dfef', 'гарниры');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('635687a9-f5c8-4595-8f82-312fedb11f8c', 'паста и ризотто');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('fef1833e-48eb-4e46-ae74-e015451c9ee9', 'супы');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('2c567271-9bd3-405f-980d-f596e509cb0d', 'мясное');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('1be2c3fd-6ec4-4e5f-a26d-73efb0ac32d1', 'салаты');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('6089b6bb-4096-4303-a6ab-d2964a401d49', 'десерты');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('3261d17b-542f-4ef6-b2dd-300f733191fa', 'пицца');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('4ddb6559-9bf3-471e-ba5c-4cab56bf3f68', 'суши и роллы');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('7c4931b2-ffb2-4464-bf93-c7333016286e', 'выпечка');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('923da41a-b9c5-44ce-9d73-48b8752b5cbc', 'хлеб и тесто');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('90ab92a8-ab18-44ac-8108-21e4511a70bb', 'соусы');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('c5d87847-26a1-47dc-9762-070b15b2eff6', 'шаверма');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('1a5cf225-2673-46dc-9fd5-33a2fafd6bf2', 'бургеры');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('662e63cf-af3e-41a6-8e48-f764a7562ff3', ',вегетарианские');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('43a55d96-8b0c-45f3-9ea2-c8032a1866f2', 'чай');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('dcea187e-3e05-4eec-9bf0-05cabb6d4360', 'кофе');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('7cc6302c-5c3c-4429-a7a4-12c645886221', 'вино');
INSERT INTO type_dish (id_typeDish, name_typeDish)
VALUES ('bac58724-8025-4507-8d01-37e42b790214', 'коктели');
-------DISH-------

INSERT INTO dish (id_dish, id_place, name_dish, id_typedish, description, url, time_min, spec)
VALUES ('31ba1f68-6eb3-4362-8d4d-24616fa9593d', '9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'Капучино',
        'dcea187e-3e05-4eec-9bf0-05cabb6d4360',
        'Кофейный напиток итальянской кухни на основе эспрессо с добавлением в него подогретого вспененного молока', '[
    {
      "url": "http://res.cloudinary.com/dwkkf6qmg/image/upload/v1520968517/zacyvt2ofxybcdj6khno.jpg"
    }
  ]', 10, '[
    {
      "size": "s",
      "price": "100"
    },
    {
      "size": "m",
      "price": "120"
    },
    {
      "size": "l",
      "price": "140"
    }
  ]');
INSERT INTO dish (id_dish, id_place, name_dish, id_typedish, description, url, time_min, spec)
VALUES ('8072b173-baaa-4d0f-a1c7-8af3995ffefa', '9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'Латте макиато',
        'dcea187e-3e05-4eec-9bf0-05cabb6d4360',
        'Горячий кофейный напиток, приготавливаемый путём вливания в молоко кофе-эспрессо в пропорции 3:1. Итальянское macchia обозначает маленькое пятнышко кофе, остающееся на поверхности молочной пены.',
        '[
          {
            "url": "http://res.cloudinary.com/dwkkf6qmg/image/upload/v1520968526/rpiiwgrwdkwb7ha9vnap.jpg"
          }
        ]', 10, '[
    {
      "size": "s",
      "price": "100"
    },
    {
      "size": "m",
      "price": "120"
    },
    {
      "size": "l",
      "price": "140"
    }
  ]');
INSERT INTO dish (id_dish, id_place, name_dish, id_typedish, description, url, time_min, spec)
VALUES ('5cb7e111-3860-424b-a14e-83a93db1e889', '9ba48d7c-b573-4dcb-b8bb-fbb196753231', 'Гамбургер',
        '1a5cf225-2673-46dc-9fd5-33a2fafd6bf2',
        'Вид сэндвича, состоящий из рубленой жареной котлеты, подаваемой внутри разрезанной булки.', '[
    {
      "url": "http://res.cloudinary.com/dwkkf6qmg/image/upload/v1520968521/hsnjck41261mlb4cmjkv.jpg"
    }
  ]', 10, '[
    {
      "size": "200гр",
      "price": "190"
    },
    {
      "size": "300гр",
      "price": "295"
    }
  ]');


INSERT INTO menudish (id_menu, id_dish)
VALUES ('a93ba633-7547-491f-a4ae-339b1420b1c7', '5cb7e111-3860-424b-a14e-83a93db1e889');

INSERT INTO menudish (id_menu, id_dish)
VALUES ('f0f6738b-5437-4a89-b715-9cd1d00173de', '31ba1f68-6eb3-4362-8d4d-24616fa9593d');


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
