# How to install

1. `git clone https://github.com/faozimipa/belajargo.git project`
2. `cd project`
3. `go get`
4. `cp .env.example .env`
5. create your database
6. create table books in your database
7. `create books (id serial,title varchar, author varchar, year varchar );`
8. edit environment `.env`  database_host, database_name, database_password, and database_user with your postgre database config
9. `go run main.go`
10. open your browser `localhost:8080/books`
