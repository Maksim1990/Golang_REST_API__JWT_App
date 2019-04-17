CREATE TABLE users
(
    id         INT       NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username   VARCHAR(50),
    password   VARCHAR(120),
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
);

CREATE TABLE posts
(
    `id`          int(11)      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `user_id`     int(11)      NOT NULL,
    `title`       varchar(255) NOT NULL,
    `description` varchar(255),
    `created_at`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_at`  timestamp    NOT NULL DEFAULT '0000-00-00 00:00:00',
    FOREIGN KEY (`user_id`) REFERENCES users (id) ON DELETE CASCADE  ON UPDATE CASCADE
);