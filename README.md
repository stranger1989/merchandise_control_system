# Merchandise Control System

Rest api server for simple merchandise control system developed by golang.

Run the main file, it will spawn a REST API endpoint

```shell
go build
go run main.go
```

## Items fetch or create or delete or update

Execute request via shell.

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
  "id": "3",
  "jan_code": "3273902848032",
  "item_name": "item_3",
  "price": 3800,
  "category_id": 3,
  "series_id": 3,
  "stock": 300,
  "discontinued": true,
  "release_date": "2020-05-01T07:30:37.660666+09:00",
  "created_at": "2020-05-01T07:30:37.660666+09:00",
  "updated_at": "2020-05-01T07:30:37.660666+09:00",
  "deleted_at": "2020-05-01T07:30:37.660666+09:00"
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
  "id": "2",
  "jan_code": "3273902847777",
  "item_name": "update_item_2",
  "price": 7777,
  "category_id": 7,
  "series_id": 7,
  "stock": 700,
  "discontinued": false,
  "release_date": "2020-05-31T07:30:37.660666+09:00",
  "created_at": "2020-05-31T07:30:37.660666+09:00",
  "updated_at": "2020-05-31T07:30:37.660666+09:00",
  "deleted_at": "2020-05-31T07:30:37.660666+09:00"
}
```

## Version Info

```
go 1.14
github.com/gorilla/mux v1.7.4
```