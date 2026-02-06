### How to build and run project within docker

1. run command to build docker image `docker build -t my-app .`
2. create a container with our image `docker run -d -p 8080:8080 --name my-app my-app`
3. checking app running in a docker process `docker ps`