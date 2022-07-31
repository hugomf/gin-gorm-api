# Simple GIN GORM API Example


For a new project do the following


- Create the project

```
    go mod init "github.com/hugomf/gin-gorm-api"  
```
or 

```
  go mod init gin-gorm-api
```

- Install Go [DotEnv](https://github.com/joho/godotenv) (to setup our env variables)

```
    go get github.com/joho/godotenv 
```

- Install [Gin Gonic](https://gin-gonic.com/) (REST Framework for GOLang)
```
    go get -u github.com/gin-gonic/gin
```

- Install the DB Driver (__postgres__)
```
    go get "gorm.io/driver/postgres"
```
- Install [Gorm](https://gorm.io/index.html) (ORM Framework for GOLang)
```
    go get -u gorm.io/gorm
```

- And last but not least, create the Database in [ElephantSQL](https://www.elephantsql.com/) and change .env DB_URL connection parameters accordingly

.env
```
    DB_URL=....
```