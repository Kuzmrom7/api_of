
 --Здесь происходит очистка схемы данных
 DROP TABLE IF EXISTS users;

 --КОНЕЦ ОЧИСТКИ

CREATE TABLE users (
  id              serial NOT NULL,
  username        character varying(300) NOT NULL,
  first_name      text NULL,
  last_name       text NULL
);

 --КОНЕЦ ЗАПОЛНЕНИЯ СХЕМ ДАННЫХ
