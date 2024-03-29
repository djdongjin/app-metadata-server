# app-metadata-server

[Design doc](https://docs.google.com/document/d/1R0-2Ikfz6WHAff-7uFBvfB8uWE8D9MOC0AFAqJilLa8/edit?usp=sharing)

This RESTful API server provides four endpoints:

1. persist: `http://localhost:8080/persist`
2. retrieve: `http://localhost:8080/retrieve`
3. get: `http://localhost:8080/get/:title`
4. delete: `http://localhost:8080/delete/:title`

## Usage - Server

1. Clone the repo: `git clone https://github.com/djdongjin/app-metadata-server.git`
2. Go to the root folder: `cd app-metadata-server`
3. Get dependencies: `go get .`
4. Start the server: `go run .`

## Usage - Client

Among the four endpoints, `persist` and `retrieve` are `POST` requests, while the other two are `GET` requests. We provide a shell command, `run.sh` for the ease of sending requests to the server.

> See the shell file to see how each request is constructed and sent by `curl`.

```bash
# persist given a yaml file.
./run.sh persist testdata/valid1.yaml

# retrieve give a query string.
./run.sh retrieve "title=Valid App 1,maintainers.name=firstmaintainer app1"

# get given a title
./run.sh get Valid%20App%201

# delete given a title
./run.sh delete Valid%20App%201
```

> Title need to be encoded so that it can be put as part of an url.
