CREATE DATABASE IF NOT EXISTS wantum DEFAULT CHARACTER SET utf8mb4;

USE wantum;

SET foreign_key_checks=0;

-- define table
CREATE TABLE IF NOT EXISTS users(
  id         int          NOT NULL AUTO_INCREMENT,
  auth_id    varchar(128) NOT NULL UNIQUE,
  user_name  varchar(20)  NOT NULL UNIQUE,
  mail       varchar(256) NOT NULL UNIQUE,
  created_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS profiles(
  id         int           NOT NULL AUTO_INCREMENT,
  name       varchar(50)   NOT NULL,
  thumbnail  varchar(2048) NOT NULL,
  bio        varchar(100)  NOT NULL COMMENT '自己紹介',
  gender     int           NOT NULL COMMENT '1=man, 2=woman, 3=other',
  phone      varchar(15)   NOT NULL COMMENT '電話番号',
  place      varchar(30)   NOT NULL COMMENT '現在地',
  birth      date          DEFAULT NULL,
  created_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  user_id    int           NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_profile_user`
  	FOREIGN KEY (user_id)
  	REFERENCES users (id)
  	ON UPDATE CASCADE
  	ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wish_lists(
  id                   int           NOT NULL AUTO_INCREMENT,
  title                varchar(30)   NOT NULL,
  background_image_url varchar(2048) NOT NULL,
  invite_url           varchar(2048) NOT NULL,
  created_at           datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at           datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at           datetime,
  user_id              int           NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_wish_lists_user`
  	FOREIGN KEY (user_id)
  	REFERENCES users (id)
  	ON UPDATE CASCADE
  	ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories(
  id           int         NOT NULL AUTO_INCREMENT,
  title        varchar(30) NOT NULL,
  created_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at   datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at   datetime,
  wish_list_id int         NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_categories_wish_list`
    FOREIGN KEY (wish_list_id)
    REFERENCES wish_lists (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wish_cards(
  id          int          NOT NULL AUTO_INCREMENT,
  activity    varchar(50)  NOT NULL COMMENT 'places.nameでactivityしたい',
  description varchar(500) NOT NULL,
  date        datetime     NOT NULL,
  done_at     datetime,
  created_at  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at  datetime,
  user_id     int          NOT NULL,
  category_id int          NOT NULL,
  place_id    int          NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_wish_cards_user`
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT `fk_wish_cards_category`
    FOREIGN KEY (category_id)
    REFERENCES categories (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT `fk_wish_card_place`
    FOREIGN KEY (place_id)
    REFERENCES places (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS albums(
  id         int           NOT NULL AUTO_INCREMENT,
  title      varchar(30)   NOT NULL,
  invite_url varchar(2048) NOT NULL,
  created_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  user_id    int           NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_albums_user`
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS memories(
  id          int          NOT NULL AUTO_INCREMENT,
  date        datetime     NOT NULL,
  activity    varchar(50)  NOT NULL COMMENT 'places.nameでactivityした',
  description varchar(500) NOT NULL,
  created_at  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at  datetime,
  user_id     int          NOT NULL,
  place_id    int          NOT NULL,
  album_id    int          NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT `fk_memories_user`
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT `fk_memory_place`
    FOREIGN KEY (place_id)
    REFERENCES places (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT `fk_memories_album`
    FOREIGN KEY (album_id)
    REFERENCES albums (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS places(
  id         int          NOT NULL AUTO_INCREMENT,
  name       varchar(200) NOT NULL,
  created_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS tags(
  id         int          NOT NULL AUTO_INCREMENT,
  name       varchar(100) NOT NULL UNIQUE,
  created_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS photos(
  id         int           NOT NULL AUTO_INCREMENT,
  photo_url  varchar(2048) NOT NULL,
  memory_id  int           NOT NULL,
  created_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime,
  PRIMARY KEY (id),
  CONSTRAINT `fk_photos_memory`
    FOREIGN KEY (memory_id)
    REFERENCES memories (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

-- define relation
CREATE TABLE IF NOT EXISTS users_wish_lists(
  id           int NOT NULL AUTO_INCREMENT,
  user_id      int NOT NULL,
  wish_list_id int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (user_id, wish_list_id),
  CONSTRAINT
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT
    FOREIGN KEY (wish_list_id)
    REFERENCES wish_lists (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
)
COMMENT = 'wish_listsの参加メンバー';

CREATE TABLE IF NOT EXISTS users_albums(
  id       int NOT NULL AUTO_INCREMENT,
  user_id  int NOT NULL,
  album_id int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE  (user_id, album_id),
  CONSTRAINT
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT
    FOREIGN KEY (album_id)
    REFERENCES albums (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
)
COMMENT = 'albumsの参加メンバー';

CREATE TABLE IF NOT EXISTS wish_cards_tags(
  id           int NOT NULL AUTO_INCREMENT,
  wish_card_id int NOT NULL,
  tag_id       int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (wish_card_id, tag_id),
  CONSTRAINT
    FOREIGN KEY (wish_card_id)
    REFERENCES wish_cards (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT
    FOREIGN KEY (tag_id)
    REFERENCES tags (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS memories_tags(
  id        int NOT NULL AUTO_INCREMENT,
  memory_id int NOT NULL,
  tag_id    int NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (memory_id, tag_id),
  CONSTRAINT
    FOREIGN KEY (memory_id)
    REFERENCES memories (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT
    FOREIGN KEY (tag_id)
    REFERENCES tags (id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);