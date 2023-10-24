### About
This is a very simple implementation of indexing documents to a TypeSense instance. The data is inside a Postgres DB 
and the contents are the data inside `books.json` which is first queried and loaded which is later indexed in `books` collection
in TypeSense.

### Implementation details
TypeSense is a fault tolerant search engine which allows full text indexing. The code here takes data out of a SQL DB(Postgres in this case)
and indexes the results inside a typesense collection. The steps are listed below:
1. `/internal/postgres` contains the logic to establish connection to a PostgresDB credentials should be loaded through a 
`config.yaml` file, it also has a method to load data from DB rows to a `Book` type.
2. `/internal/typesense` is responsible to initialize the client using credentials in config. It has a method to create a collection 
and add documents to collection.
3. `cmd/main.go` it contains the calls to the internal pkgs.

