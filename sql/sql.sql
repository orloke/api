CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default current_timestamp() on update current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int not null,
    follower_id int not null,
    created_at timestamp default current_timestamp(),

    PRIMARY KEY(user_id, follower_id),

    FOREIGN KEY(user_id) REFERENCES users(id)
    ON DELETE CASCADE,

    FOREIGN KEY(follower_id) REFERENCES users(id)
    ON DELETE CASCADE,

    CHECK(user_id <> follower_id)
) ENGINE=INNODB;

INSERT INTO users(name, nick, email, password) VALUES
("User 1", "user1", "user1@gmail.com", "$2a$10$Fbbak2ROvmRkpi7lX7P/MeAUv5ul3.mn3FsUSnviLjZbpINuxQ146"),
("User 2", "user2", "user2@gmail.com", "$2a$10$Fbbak2ROvmRkpi7lX7P/MeAUv5ul3.mn3FsUSnviLjZbpINuxQ146"),
("User 3", "user3", "user3@gmail.com", "$2a$10$Fbbak2ROvmRkpi7lX7P/MeAUv5ul3.mn3FsUSnviLjZbpINuxQ146");

INSERT INTO followers(user_id, follower_id) VALUES
(1, 2),
(3, 1),
(1, 3);


