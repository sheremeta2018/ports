# ports

In order to get this project up and running read the requirements section

# requirements
- You need the following repository: github.com/tsetsik/ports-storage Instructions how to get it up and running are inside
- You need this project also up and running, you can do that by perform the following command `docker-compose up`
- To see it in action: `curl -X POST -F "file=./assets/ports.json" http://localhost:8080/ports`

# todo

- write unit and integration tests
- implement graceful shutdown, currently that is implemented only in `ports-storage` repo