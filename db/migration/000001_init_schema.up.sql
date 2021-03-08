CREATE TABLE users (
    id TEXT NOT NULL PRIMARY KEY ,
    password TEXT NOT NULL,
    max_todo INTEGER DEFAULT 5 NOT NULL
);

CREATE TABLE tasks (
    id TEXT NOT NULL PRIMARY KEY ,
    content TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_date TEXT NOT NULL
);

