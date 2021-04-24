Ulr shortener like https://bitly.com/ \
Test task for Avito trinee/junior developer https://github.com/avito-tech/auto-backend-trainee-assignment

##How to start:

###Manually
1. Run database - docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres postgres:12
2. Run service - go run main.go

###Docker-compose
1. Docker-compose up -d

###Migrations
1. Connect to db
2. Apply migration

```
create table link (
    id serial,
    url varchar(255) not null,
    short_url varchar(255) not null
)
```

## Endpoints
1. /create - POST
    ```json
       {"url":  "string"}
    ```
    Url for which short url will be created
2. / - GET Will redirect to full url mapping to this url path