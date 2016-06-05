FROM        golang:1.7beta1-alpine
MAINTAINER  Earvin Kayonga <earvin@earvinkayonga.com>


# Setting up working directory
WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container


RUN         go install github.com/gin-gonic/gin
RUN         go install github.com/tools/godep

# Restore godep dependencies
RUN godep restore


EXPOSE 80
ENTRYPOINT ["/go/src/gin-container/main.go"]
