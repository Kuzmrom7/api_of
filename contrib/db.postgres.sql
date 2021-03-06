---------------------------------------------------------------------------------------------------
------------------------------------ ORDER FOOD MIGRATION -----------------------------------------
---------------------------------------------------------------------------------------------------


---------------------------------------------------------------------------------------------------
---------------------------------------- Drops all tables -----------------------------------------
---------------------------------------------------------------------------------------------------

DROP SCHEMA public CASCADE;

---------------------------------------------------------------------------------------------------
---------------------------------------- Drops extensions -----------------------------------------
---------------------------------------------------------------------------------------------------

DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;

---------------------------------------------------------------------------------------------------
--------------------------------------- Creates schema ----------------------------------------
---------------------------------------------------------------------------------------------------

CREATE SCHEMA public;

---------------------------------------------------------------------------------------------------
--------------------------------------- Creates extensions ----------------------------------------
---------------------------------------------------------------------------------------------------

CREATE EXTENSION "uuid-ossp";

---------------------------------------------------------------------------------------------------
----------------------------------------- Creates tables ------------------------------------------
-------------------------------------------------------------------------------------------------


CREATE TABLE users
(
  user_id  UUID PRIMARY KEY             NOT NULL  DEFAULT uuid_generate_v4(),
  username VARCHAR(256) UNIQUE,
  email    VARCHAR(256) UNIQUE,
  gravatar VARCHAR(64)                            DEFAULT '',
  password VARCHAR(256)                           DEFAULT '',
  salt     VARCHAR(256)                           DEFAULT '',
  active   BOOLEAN                                DEFAULT TRUE,
  type     VARCHAR(64)                            DEFAULT '',
  created  TIMESTAMP                              DEFAULT now(),
  updated  TIMESTAMP                              DEFAULT now()
);

CREATE TABLE type_place
(
  id_typePlace UUID PRIMARY KEY             NOT NULL DEFAULT uuid_generate_v4(),
  name_type    VARCHAR(256)                 NOT NULL
);

CREATE TABLE place
(
  id_place     UUID PRIMARY KEY                                      NOT NULL DEFAULT uuid_generate_v4(),
  name         VARCHAR(256) UNIQUE,
  phone_number VARCHAR(256)                                                   DEFAULT '',
  url          VARCHAR(256)                                                   DEFAULT '',
  logo         VARCHAR(256)                                                   DEFAULT '',
  city         VARCHAR(256)                                                   DEFAULT '',
  user_id      UUID REFERENCES users (user_id)                       NOT NULL,
  type         JSONB                                                          DEFAULT '[]',
  created      TIMESTAMP                                                      DEFAULT now(),
  updated      TIMESTAMP                                                      DEFAULT now()
);

CREATE TABLE adressing
(
  id_adressing UUID PRIMARY KEY                                   NOT NULL DEFAULT uuid_generate_v4(),
  name_adress  VARCHAR(256)                                                DEFAULT '',
  id_place     UUID REFERENCES place (id_place)                   NOT NULL DEFAULT uuid_generate_v4()
);


CREATE TABLE menu
(
  id_menu   UUID PRIMARY KEY                    NOT NULL DEFAULT uuid_generate_v4(),
  name_menu VARCHAR(256) UNIQUE,
  url       VARCHAR(256)                                 DEFAULT '',
  id_place  UUID REFERENCES place (id_place)    NOT NULL,
  created   TIMESTAMP                                    DEFAULT now(),
  updated   TIMESTAMP                                    DEFAULT now()
);

CREATE TABLE type_dish
(
  id_typeDish   UUID PRIMARY KEY                           NOT NULL DEFAULT uuid_generate_v4(),
  name_typeDish VARCHAR(256)                                        DEFAULT '',
  created       TIMESTAMP                                           DEFAULT now(),
  updated       TIMESTAMP                                           DEFAULT now()

);

CREATE TABLE dish
(
  id_dish     UUID PRIMARY KEY                        NOT NULL DEFAULT uuid_generate_v4(),
  id_place    UUID,
  name_dish   VARCHAR(256),
  id_typeDish UUID REFERENCES type_dish (id_typeDish),
  description VARCHAR(256)                                     DEFAULT '',
  url         JSONB                                            DEFAULT '[]',
  spec        JSONB                                            DEFAULT '[]',
  time_min    INTEGER                                          DEFAULT 0,
  created     TIMESTAMP                                        DEFAULT now(),
  updated     TIMESTAMP                                        DEFAULT now()
);

CREATE TABLE menudish
(
  id_menu UUID REFERENCES menu (id_menu),
  id_dish UUID REFERENCES dish (id_dish)
);

CREATE TABLE type_personal
(
  id_typePersonal UUID PRIMARY KEY             NOT NULL DEFAULT uuid_generate_v4(),
  name_type       VARCHAR(256)                 NOT NULL
);

CREATE TABLE personal
(
  id_personal     UUID PRIMARY KEY                                     NOT NULL DEFAULT uuid_generate_v4(),
  fio             VARCHAR(256)                                         NOT NULL,
  phone           VARCHAR(64)                                                   DEFAULT '',
  id_place        UUID REFERENCES place (id_place)                     NOT NULL,
  id_typePersonal UUID REFERENCES type_personal (id_typePersonal)      NOT NULL,
  created         TIMESTAMP                                                     DEFAULT now(),
  updated         TIMESTAMP                                                     DEFAULT now()

);


