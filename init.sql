create table user
(
    id             int(11)      NOT NULL AUTO_INCREMENT,
    username       varchar(100) NOT NULL,
    password       varchar(100) NOT NULL,
    follow_count   int(11)      NOT NULL DEFAULT 0,
    follower_count int(11)      NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table video
(
    id             int(11)      NOT NULL AUTO_INCREMENT,
    publisher_id   int(11)      NOT NULL,
    title          varchar(100) NOT NULL,
    video_url      varchar(255) NOT NULL,
    cover_url      varchar(255) NOT NULL,
    favorite_count int(11)      NOT NULL DEFAULT 0,
    comment_count  int(11)      NOT NULL DEFAULT 0,
    create_date    datetime     NULL     DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table relation
(
    id          int(11) NOT NULL AUTO_INCREMENT,
    follow_id   int(11) NOT NULL,
    follower_id int(11) NOT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table comment
(
    id          int(11)      NOT NULL AUTO_INCREMENT,
    user_id     int(11)      NOT NULL,
    video_id    int(11)      NOT NULL,
    content     varchar(300) NOT NULL,
    create_date datetime     NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;

create table favorite
(
    id          int(11)  NOT NULL AUTO_INCREMENT,
    user_id     int(11)  NOT NULL,
    video_id    int(11)  NOT NULL,
    create_date datetime NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_general_ci;