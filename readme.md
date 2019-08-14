# How to install

1. `git clone https://github.com/faozimipa/belajargo.git project`
2. `cd project`
3. `go get`
4. `cp .env.example .env`
5. Create your database
6. Create table books in your database

```sql
    create table books (
        id serial,
        title varchar,
        author varchar,
        year varchar);
 ```

7. Edit environment `.env`  database_host, database_name, database_password, and database_user with your postgre database config
8. Run project  `go run main.go`
9. Open your browser `localhost:8080/books`
