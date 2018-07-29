FROM        golang:1.10-alpine
MAINTAINER  Woodrow Shen <woodrow.shen@gmail.com>

ENV	    PORT 80
	
RUN	    apk add --update git bash build-base
 
# Setting up working directory
WORKDIR     /go/src/gin-container
ADD         . /go/src/gin-container

EXPOSE 80
ENTRYPOINT	["/usr/local/go/bin/go"]
CMD 		["run", "src/main.go"]
