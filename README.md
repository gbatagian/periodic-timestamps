# Periodic Timestamps API

A lightweight Go API for generating periodic timestamps using the standard `net/http` package.

# Run

## Local Run

### Default Configuration
* Navigate to the project root: `cd .../periodic-timestamps` 
* Run the application: `go run main.go` (or `make run`)
* The confirmation message of the server running should appear: `Server starts listening at: 0.0.0.0:8080`
* Try a sample requests, e.g.: 
  * `curl --location 'http://0.0.0.0:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

### Custom Address Configuration
* Navigate to the project root: `cd .../periodic-timestamps` 
* Run the application: `go run main.go -host=192.168.1.5 -port=8888` 
  * **Note**: Ensure the address is available and firewall settings permit access.
* The confirmation message of the server running should appear: `Server starts listening at: 192.168.1.5:8888`
* Try a sample requests, e.g.: 
  * `curl --location 'http://192.168.1.5:8888/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

## Docker Run

### Default Configuration
* Navigate to the project root: `cd .../periodic-timestamps` 
* Run the project container: `docker compose up`
* The confirmation message of the server running should appear: `Server starts listening at: 0.0.0.0:8080`
* Try a sample requests, e.g.: 
  * `curl --location 'http://0.0.0.0:8080/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

###  Custom Address Configuration
* Navigate to the project root: `cd .../periodic-timestamps` 
* Copy the **.env** file: `cp .env.sample .env`
* Run the project container - it will automatically get the configuration from the .env file: `docker compose up` 
  * Make sure that the configuration specified in the `.env` file is valid and that firewall allows accessing the api host address. A predefined sample configuration is provided.
  * If the configuration is valid but you receive an `invalid endpoint settings` prompt from docker daemon, try: `docker network prune` to clean up the network settings and then try again `docker compose up`
* The confirmation log message of the server running should appear: `Server starts listening at: 192.168.1.5:8888`
* Try a sample requests, e.g.: 
  * `curl --location 'http://192.168.1.7:8888/ptlist?period=1h&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'`

### Clean-up 
To revert to default deployment settings `0.0.0.0:8080`, after a custom address configuration deployment:
* Remove the **.env** file: `rm .env`
* Reset docker networks:
  * `docker compose down`
  * `docker network prune`
  * `docker compose up`
* Run the project container: `docker compose up`

# PtList services (/ptlist)
## 200 OK Request

```bash
>> curl --location 'http://0.0.0.0:8080/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z'

["20210714T210000Z","20210714T220000Z","20210714T230000Z","20210715T000000Z","20210715T010000Z","20210715T020000Z","20210715T030000Z","20210715T040000Z","20210715T050000Z","20210715T060000Z","20210715T070000Z","20210715T080000Z","20210715T090000Z","20210715T100000Z","20210715T110000Z","20210715T120000Z"]
```
## Parameters

* **period**: allowed values  1h, 1d, 1mo, 1y
* **tz** : [IANA Time Zone database value](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones), such as "America/New_York"
* **t1** : timestamp specifying the start of the periodic timestamps range, in YYYYDDMMTHHMMSSZ format 
* **t2** : timestamp specifying the end of the periodic timestamps range, in YYYYDDMMTHHMMSSZ format 

**All the above parameters are required, missing any of those parameters will result into 400 Bad Request**

## 400 BAD REQUEST

1.  Invalid Period
```bash
>> curl --location 'http://0.0.0.0:8080/ptlist?period=1w&tz=Europe%2FAthens&t1=20210714T204603Z&t2=20210715T123456Z'

unsupported period '1w'
```

1. Missing Parameter
```bash
>> curl --location 'http://0.0.0.0:8080/ptlist?period=1h&t1=20210714T204603Z&t2=20210715T123456Z'

parameter 'tz' cannot be empty, please provide a value
```

3. Invalid Timestamp Format

```bash
>> curl --location 'http://0.0.0.0:8080/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456' 

timestamp '20210715T123456' does not match the required format YYYYDDMMTHHMMSSZ

```

# Development Commands

Utilize the Makefile for common development tasks:
* `make run`
* `make test`
* `make coverage`
* `make coverage_profile`

#  Postman  Collection

Refer to the **docs** folder for a Postman collection with sample requests and and an environment for the collection.
