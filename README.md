# periodic-timestamps
An API in Go that returns matching timestamps of a periodic task.
The API is build using the [Gin](https://gin-gonic.com) framework and a [custom made](https://github.com/gbatagian/go-domain-driven-api) domain driven architecture.

# Run
* Change directory to the project's root directory: `cd .../periodic-timestamps` 
* Run the project's container: `docker compose up`
* Execute a sample request on a new terminal window, e.g.: 
  * `curl --location 'http://localhost:8080/healthcheck'`
  * `curl --location 'http://localhost:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`
