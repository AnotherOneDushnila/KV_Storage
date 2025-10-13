# KV_Storage

## Setup

1. Running service:

```
go run ./server/main.go     
```

### Http

1. PUT

```
curl -i -X POST "http://localhost:8080/put?collection=coll_name&key=k_name" -d 'Aboba'
```

2. GET

```
curl "http://localhost:8080/get?collection=coll_name&key=k_name"
```

3. DELETE

```
curl -i -X DELETE http://localhost:8080/delete?collection=coll_name&key=k_name
```