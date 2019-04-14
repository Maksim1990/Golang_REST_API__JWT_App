CREATE TABLE users
(
    id         int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    username   VARCHAR(50),
    password   VARCHAR(120),
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT NULLIF('0000-00-00 00:00:00','0000-00-00 00:00:00')::timestamp
);

CREATE TABLE posts
(
    id         int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id    int NOT NULL ,
    title      varchar(255) NOT NULL,
    description varchar(255),
    created_at timestamp DEFAULT current_timestamp,
    updated_at timestamp DEFAULT NULLIF('0000-00-00 00:00:00','0000-00-00 00:00:00')::timestamp,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE NO ACTION ON UPDATE NO ACTION
);