# API OrderFood



## DataBase


### Run PostrgreSQL on Docker 
```
docker run -i -d --name=pg_order -e POSTGRES_USER=orderfood -e POSTGRES_PASSWORD=orderfood -e POSTGRES_DB=orderfood -p 5432:5432 postgres
```
### Create table for DB
```
api_of/contrib/db.postgres.sql
```
### Run the Data Migration script
```
api_of/contrib/data/insertData.sql
```

## API

### Run api using `go run`

```
git clone https://github.com/Kuzmrom7/api_of.git
go get -u github.com/kardianos/govendor
govendor sync
go run /cmd/api/api.go
```

### Using Docker

```
git clone https://github.com/Kuzmrom7/api_of.git
docker build -t api_of .
docker run api_of -p 8080:8080 
```

## Client

We writing client for this API. 
https://github.com/Kuzmrom7/orderfood - ReactJS App!
[Orderfood](https://github.com/Kuzmrom7/orderfood)

