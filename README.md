# Gin Container

Golang Gin running in an Alpine based Image.

Docker image hosted on The  [Docker Hub](https://hub.docker.com/r/earvin/gin-container/)

## Usage

Use this Dockerfile to build our container
```
docker build -t earvin/gin-container .; 
docker run -p 8080:80  -v $(pwd):/go/src/gin-container --name gin-instance earvin/gin-container run src/main.go;
```

OR 
Directly pull from the Docker Hub

```
docker pull earvin/gin-container
docker run -d  -p 8080:80  -v $(pwd):/go/src/gin-container --name gin-instance earvin/gin-container  run src/main.go
```

OR

Compose in your Dockerfile

```
FROM earvin/gin-container
...
```
