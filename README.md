# go gorm example project
## Go Version 
1.21
## running code
```golang
go run cmd/main.go
```
## new User
POST http://localhost:8080/users\
content-type: application/json

## Payload
```json
{
    "name": "ekowdd",
    "email": "ekowdd89sx@gmail.com",
    "password": "$2y$10$lLk/XTh1wfjKAz.RKNsspOhdR6zE/ia58VALrugCsCDjx5dzQ/CFq",
    "display_name": "ekoxsx",
    "telp":"-"
}
```
## User Update
PUT http://localhost:8080/users/{id}\
content-type: application/json

```json
{
    "name": "ekowdd",
    "email": "ekowdd89sx@gmail.com",
    "password": "$2y$10$lLk/XTh1wfjKAz.RKNsspOhdR6zE/ia58VALrugCsCDjx5dzQ/CFq",
    "display_name": "ekoxsx",
    "telp":"-"
}
```

## User Update
DELETE http://localhost:8080/users/{id}\
content-type: application/json



## User All
GWT http://localhost:8080/users\
content-type: application/json

```json
[{
    "id":1,
    "name": "ekowdd",
    "email": "ekowdd89sx@gmail.com",
    "password": "$2y$10$lLk/XTh1wfjKAz.RKNsspOhdR6zE/ia58VALrugCsCDjx5dzQ/CFq",
    "display_name": "ekoxsx",
    "telp":"-"
}]
```

## User Find One
GWT http://localhost:8080/users/{id}\
content-type: application/json

```json
{
    "id":1,
    "name": "ekowdd",
    "email": "ekowdd89sx@gmail.com",
    "password": "$2y$10$lLk/XTh1wfjKAz.RKNsspOhdR6zE/ia58VALrugCsCDjx5dzQ/CFq",
    "display_name": "ekoxsx",
    "telp":"-"
}
```