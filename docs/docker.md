# Docker

* Building for docker can be done with ```docker build -t mailsend-go .```. This will also fetch the golang image and create a intermediate image (about 820MB in total). If space is a concern for you, remove them with ```docker rmi golang:1.13.7``` and ```docker image prune```

* Running with docker can be done as any other docker image. Everything after the image name will be passed to the program. Example: ```docker run -it --rm mailsend-go -V``` will show the version
