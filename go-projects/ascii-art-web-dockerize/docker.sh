#!/bin/sh
docker container prune
docker image prune
docker image build -f Dockerfile -t ascii-art-web-dockerize-img .
docker images
while docker run -p 8080:8080 --detach --name ascii-art-web-cont ascii-art-web-dockerize-img
do
    docker ps -a
    docker exec -it ascii-art-web-cont /bin/sh
done
docker stop ascii-art-web-cont
docker container rm ascii-art-web-cont
docker rmi ascii-art-web-dockerize-img

# Too intrusive as it stops and deletes ALL containers and images.
# docker stop $(docker ps -aq)
# docker container rm $(docker container ls -aq)
# docker rmi $(docker images -q)
