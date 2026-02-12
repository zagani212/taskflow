CREATE TABLE users ( email varchar(30), password varchar(100));
CREATE TABLE tasks (id serial, name varchar(30), status varchar(30), check(status IN ('TO DO', 'IN PROGRESS', 'DONE')));