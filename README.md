# Merchandise Control System

Rest api server for simple merchandise control system developed by golang.

## Mysql setting
If you launch this api server and connect db for the first time,
Need to prepare the database for this system.

### Mysql init

Login to Mysql as the root user.

```shell
mysql -u root -p
```

Create database user, set username and password.

```shell
GRANT ALL PRIVILEGES ON *.* TO 'username'@'localhost' IDENTIFIED BY 'password';
```

Exit once, and login as created user.

```
mysql -u username -p
```

Create database.

```shell
CREATE DATABASE merchandise_control_system;
USE merchandise_control_system;
```

Now, Ready for access to Mysql.

## Create `config.ini` file

Create `config.ini` file on root path for setting db and webserver.
Here is the example config.ini file as below.

```config.ini
[db]
db_driver_name = mysql
db_name = merchandise_control_system
db_user_name = { mysql_username }
db_user_password = { mysql_password }
db_host = 127.0.0.1
db_port = 3306

[api]
server_port = 8080
```

## Launch api server and connect db

Run the main file, it will spawn a REST API endpoint.

```shell
go build
go run main.go
```

Execute request via shell.

## Items fetch or create or delete or update

### To fetch all item
```shell
curl --request GET --url http://localhost:8080/items
```
### To fetch single item by ID
```shell
curl --request GET --url http://localhost:8080/item/{id}
```

### To create an item
```shell
curl --request POST --url http://localhost:8080/item --header 'Content-Type: application/json' --data {params}
```

**params e.g**
```json
{
  "jan_code": "3273902848032",
  "item_name": "item_3",
  "price": 3800,
  "category_id": 3,
  "series_id": 3,
  "stock": 300,
  "discontinued": true,
  "release_date": "2020-05-01T07:30:37.660666+09:00"
}
```

### To delete an item by ID
```
curl --request DELETE --url http://localhost:8080/item/{id}
```

### To update an item by ID
```
curl --request PUT --url http://localhost:8080/item/{id} --header 'Content-Type: application/json' --data {params}
```

**params e.g**
```json
{
  "jan_code": "3273902847777",
  "item_name": "update_item_2",
  "price": 7777,
  "category_id": 7,
  "series_id": 7,
  "stock": 700,
  "discontinued": false,
  "release_date": "2020-05-31T07:30:37.660666+09:00"
}
```

## Version Info

```
go 1.14
github.com/gorilla/mux v1.7.4
github.com/jinzhu/gorm v1.9.12
gopkg.in/ini.v1 v1.56.0
```