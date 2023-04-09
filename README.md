# periodic-timestamps

An API in Go that returns matching timestamps of a periodic task.
The API is build using the [Gin](https://gin-gonic.com) framework and a [custom made](https://github.com/gbatagian/go-domain-driven-api) domain driven architecture.

# Run
## Local Run
### Default settings

* Change directory to the project's root directory: `cd .../periodic-timestamps` 
* Run: `make run` (if make is not installed on your machine run: `go run main.go`)
* The confirmation log message of the service  running should appear: `[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080`
* Execute some sample requests, e.g.: 
  * `curl --location 'http://localhost:8080/healthcheck'`
  * `curl --location 'http://localhost:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

### Custom Settings

* Change directory to the project's root directory: `cd .../periodic-timestamps` 
* Run: `go run main.go -host=192.168.1.7 -port=8888`
* The confirmation log message of the service  running should appear: `[GIN-debug] Listening and serving HTTP on 192.168.1.7:8888`
* Execute some sample requests, e.g.: 
  * `curl --location 'http://192.168.1.7:8888/healthcheck'`
  * `curl --location 'http://192.168.1.7:8888/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

## Docker Run
### Default settings

* Change directory to the project's root directory: `cd .../periodic-timestamps` 
* Run the project's container: `docker compose up`
* The confirmation log message of the service  running should appear: `[GIN-debug] Listening and serving HTTP on 0.0.0.0:8080`
* Execute some sample requests, e.g.: 
  * `curl --location 'http://localhost:8080/healthcheck'`
  * `curl --location 'http://localhost:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`
## Custom settings

* Change directory to the project's root directory: `cd .../periodic-timestamps` 
* Copy the **.env** file: `cp .env.sample .env`
* **OPTIONAL**: Edit the environment variable on the **.env** file to fit your custom specifications. Some custom predefined options are hardcoded - you can simply use those. If you edit them, make sure that the network addresses are set up accordingly so that communication is feasible. 
* Run the project's container - it will automatically get the configuration from the .env file: `docker compose up`
* The confirmation log message of the service  running should appear: `[GIN-debug] Listening and serving HTTP on 192.168.1.7:8888` (for the predefined custom configuration)
* Execute some sample requests, e.g.: 
  * `curl --location 'http://192.168.1.7:8888/healthcheck'`
  * `curl --location 'http://192.168.1.7:8888/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`
## Clean-up 

* If you want to return to the default deployment, simply delete the **.env** file: `rm .env` and then run the container again: `docker compose up`
* Probably, if you have used custom settings to run the container, some of the default addresses are allocated by the docker system, this could result into an error of the following form: `failed to create network periodic-timestamps_apinetwork: Error response from daemon: Pool overlaps with other one on this address space` or for similar reasons an error of this form: `Error response from daemon: Invalid address 172.30.0.2: It does not belong to any of this network's subnets`. For that reason, when switching between default and custom docker deployments, it is better to perform some clean-up first:
  * `docker compose down`
  * `docker network prune`
  * `docker compose up`

# PtList services (/ptlist)
## 200 OK Request

Sample request : `{{scheme}}://{{host}}:{{port}}/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z`

Sample response: `[
    "20210714T210000Z",
    "20210714T220000Z",
    "20210714T230000Z",
    "20210715T000000Z",
    "20210715T010000Z",
    "20210715T020000Z",
    "20210715T030000Z",
    "20210715T040000Z",
    "20210715T050000Z",
    "20210715T060000Z",
    "20210715T070000Z",
    "20210715T080000Z",
    "20210715T090000Z",
    "20210715T100000Z",
    "20210715T110000Z",
    "20210715T120000Z"
]`

## Parameters

* **period**: allowed values  1h, 1d, 1mo, 1y
* **tz** : [IANA Time Zone database value](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones), such as "America/New_York"
* **t1** : timestamp specifying the start of the periodic timestamps range, in YYYYDDMMTHHMMSSZ format 
* **t2** : timestamp specifying the end of the periodic timestamps range, in YYYYDDMMTHHMMSSZ format 

All the above parameters are required, missing any of those parameters will result into 400 Bad Request

## 400 OK BAD REQUEST - Invalid Period

Sample request : `{{scheme}}://{{host}}:{{port}}/ptlist?period=1w&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456Z`

Sample response: `{
    "desc": "Unsupported period",
    "status": "error"
}`
## 400 OK BAD REQUEST - Missing Parameter

Sample request : `{{scheme}}://{{host}}:{{port}}/ptlist?period=1h&t1=20210714T204603Z&t2=20210715T123456Z`

Sample response: `{
    "desc": "The 'tz' parameter cannot be empty. Please provide a value for this parameter.",
    "status": "error"
}`

## 400 OK BAD REQUEST - Invalid Timestamp Format

Sample request : `{{scheme}}://{{host}}:{{port}}/ptlist?period=1w&tz=Europe/Athens&t1=20180214T204603Z&t2=20211115T123456`

Sample response: `{
    "desc": "Timestamp '20211115T123456' does not match the required format YYYYDDMMTHHMMSSZ",
    "status": "error"
}`

# Local development

To automate local development and assess future contributions, you can use the commands on the **Makefile**:
* `make run`
* `make test`
* `make coverage`
* `make coverage_profile`

#  Postman  Collection

Inside the **docs** folder, there is a **Postman collection** which includes some sample requests and an environment for the collection. This is provided in case you want to test the service using Postman, which is a useful tool for QA. You can import this collection and the environment into Postman and then run the service using any of the methods described in the "Run" section. After running the service, you can try out the sample requests provided in the collection. The environment for the collection has default values set for `host=localhost`, `port=8080`, and `scheme=http`. If you change any of these values, make sure to update the environment settings accordingly.