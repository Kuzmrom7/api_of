
 --Здесь происходит очистка схемы данных
 DROP TABLE IF EXISTS users;

 --КОНЕЦ ОЧИСТКИ

 CREATE TABLE users
 (
   id              UUID PRIMARY KEY             NOT NULL,
   username        VARCHAR(256)                 NOT NULL,
   first_name      text NULL,
   last_name       text NULL
 );

 CREATE TABLE place
 (
   id_place              UUID PRIMARY KEY             NOT NULL,
   name                  VARCHAR(256)                 NOT NULL,
   phone                 TEXT                         NULL,
   url                   TEXT                         NULL,
   city                  TEXT                         NULL,
   adress                text                         NULL
 );

 CREATE TABLE type_place
 (
   id_type        UUID PRIMARY KEY             NOT NULL,
   name_type      VARCHAR(256)                 NOT NULL
 );

 CREATE TABLE menu
 (
   id_menu         UUID PRIMARY KEY             NOT NULL,
   name_menu       VARCHAR(256)                 NOT NULL,
   created         TIMESTAMP                    DEFAULT now()
 );

 CREATE TABLE section_menu
 (
   id_section         UUID PRIMARY KEY             NOT NULL,
   name_section       VARCHAR(256)                 NOT NULL
 );

 CREATE TABLE type_dishes
 (
   id_type         UUID PRIMARY KEY             NOT NULL,
   name_type       VARCHAR(256)                 NOT NULL
 );
 CREATE TABLE dish
 (
   id_menu          UUID PRIMARY KEY             NOT NULL,
   name_dish        VARCHAR(256)                 NOT NULL,
   description      VARCHAR(256)                 NOT NULL,
   time             INTEGER                      NOT NULL,
   created          TIMESTAMP                    DEFAULT now()
 );

 CREATE TABLE type_personal
 (
   id_type         UUID PRIMARY KEY             NOT NULL,
   name_type       VARCHAR(256)                 NOT NULL
 );

 CREATE TABLE personal
 (
   id_personal         UUID PRIMARY KEY             NOT NULL,
   fio                 VARCHAR(256)                 NOT NULL,
   phone               TEXT                         NULL
 );

 --КОНЕЦ ЗАПОЛНЕНИЯ СХЕМ ДАННЫХ
