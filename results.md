##How to use
- make run - start http server on localhost:8080
- make lint - run linters. 
See linters config in [.golangci.yaml](.golangci.yaml).

API calls are described in [Swagger](http://localhost:8080/swagger/#/).

##Technical choices
- Configs in YAML files: [configs](configs). Separate config file for local environment in /dev folder (git ignored).
- Prometheus as a tool to collect metrics: [metrics](internal/metrics). 
I decided to collect errors count and wrote middleware to collect
http request duration and count. 
- Zap as logger: [logger](internal/logger).

##Architecture
The application architecture is divided on three main layers:
- handlers - functions which handles http requests and contain all http-specific logic 
(e.g. marshalling/unmarshalling structs).
- services - contain all business logic and interact with other layers only on interface levels, 
without implementation details(e.g. how mongoDB will be store our data).
- repositories - contain database-specific part of code, all interaction with mongoDB happens here.

It's pretty agile architecture that allow us easily change database providers or processing request parts of the system 
without changing main business logic implementation.



