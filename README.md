# Description

Implementation of test task: implement REST API 

## Approach

It is important to have API documented, thus there is a problem of documentation and API implementation consistency. To solve the proble the following approach was used:

1. First step is to create swagger spec to describe the API (**api/ing.swagger.yaml**)
2. Use go-swagger to generate server's code:
```bash
$ make swagger
```
3. Implement endpoints' handlers
4. If there is a requirement to change something in API, start from 1.

## Build and Run

To run the app it would be nice to have **docer-compose** install

```bash
$ git clone git@github.com:kshamko/ing.git
$ cd ing
$ docker-compose up
```

The application will be started on port 8080 and it will be possible to request it localy:
```bash
curl -XGET 'http://localhost:8080/routes?src=13.388860,52.517037&dst=13.397634,52.529407&dst=13.428555,52.523219'
```

or in a browser

## Service Endpoints

Also an additionl port will be exposed for debug/healthcheck/swagger-ui purposes

1. Metrics: http://localhost:6060/metrics
2. Heaaltcheck: http://localhost:6060/healthz
3. Swagger-UI: http://localhost:6060/swagger-ui

