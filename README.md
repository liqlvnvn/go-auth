# go-auth
Golang API Authentication using JWT Tokens

## Installation

1) Install MySQL/MariaDB database.
2) Create ```yt_go_auth``` database under ```root``` user.
3) Create users table with this schema

```SHOW COLUMNS FROM users;```
```
+----------+------------------+------+-----+---------+----------------+
| Field    | Type             | Null | Key | Default | Extra          |
+----------+------------------+------+-----+---------+----------------+
| id       | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| name     | longtext         | YES  |     | NULL    |                |
| email    | varchar(191)     | YES  | UNI | NULL    |                |
| password | longtext         | YES  |     | NULL    |                |
+----------+------------------+------+-----+---------+----------------+
```
```
CREATE TABLE users (
     id INTEGER UNSIGNED AUTO_INCREMENT,
     name longtext,
     email varchar(191),
     password longtext,
     PRIMARY KEY (id),
     CONSTRAINT email_unique UNIQUE (email)
 );
```
4) Make
```go run main.go```
to run the server.

## Using

1) To registry a user using Postman send this kind of requests.
```
POST http://localhost:8000/api/registry
{
  "name": "userName",
  "email": "1@1.com",
  "password": "userPassword"
}
```
2) To login
```
POST http://localhost:8000/api/login
{
  "email": "1@1.com",
  "password": "userPassword"
}
```
3) Checking if user is logged using cookie
```
GET http://localhost:8000/api/user
```
4) Logout
```
POST http://localhost:8000/api/logout
```
