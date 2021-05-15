# Description

Implementation of test task: implement REST API 

## Approach

It is important to have API documented, thus there is a problem of documentation and API implementation consistency. To solve the proble the following approach was used:
1. First step is to create swagger spec to describe the API (**api/ing.swagger.yaml**)
2. Use go-swagger to generate server's code:
```bash
$ make swagger
```
