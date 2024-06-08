# bean-server
this is only running off a local postgres server at the moment.
to get up and running:
1. install [go](https://go.dev/doc/install)
2. install [postgres](https://www.postgresql.org/download/)
3. set up a new database named "beandb"
4. create a `.env` file
    1. add your DB_USER (eg. "postgres:postgres")
    2. add your DB_LOCATION (eg. "localhost:5432")
5. build the project with `make build`
6. run the project with `make run`

