CREATE TABLE doingtasks (
    id INTEGER NOT NULL PRIMARY KEY,
    title text NOT NULL,
    description text
);

CREATE TABLE donetasks (
    id INTEGER NOT NULL PRIMARY KEY,
    title text NOT NULL,
    description text
);

INSERT INTO doingtasks (title, description)
VALUES ("Complete a simple todolist server.", "using Gin-HTTP FrameWork and SQLite.");
+----+------------------------------------+--------------------------------------+
| id |               title                |             description              |
+----+------------------------------------+--------------------------------------+
| 1  | Complete a simple todolist server. | using Gin-HTTP FrameWork and SQLite. |
| 2  | Practice Articulation              | Read outloud and articulate in 10    |
| 3  | 1 hour for reading book            | before go to bed                     |
+----+------------------------------------+--------------------------------------+

SELECT * FROM doingtasks;

INSERT INTO donetasks (title, description)
VALUES ("30 minutes of listening English", "listen Tedx and Improvement Pill");

UPDATE doingtasks SET title = 