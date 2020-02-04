docker network create limon

docker build . -f server/Dockerfile -t server

docker run -d -l traefik.http.routers.seeker.tls=true \
-l traefik.http.services.server.loadbalancer.server.scheme=https \
-l traefik.port=5300 \
-l traefik.backend=server \
-l traefik.enable=true --network limon \
 --name server \
 --network-alias server server

 cd traefik

 docker run -d -p 4443:4443 -p 8080:8080 -p 5300:5300 \
 -v $PWD/traefik.yml:/etc/traefik/traefik.yml \
 -v $PWD/cert/frontend.local.cert:/etc/traefik/frontend.local.cert \
 -v $PWD/cert/frontend.local.key:/etc/traefik/frontend.local.key \
 -v $PWD/cert/backend.local.cert:/etc/traefik/backend.local.cert \
 -v $PWD/cert/backend.local.key:/etc/traefik/backend.local.key \
 -v /var/run/docker.sock:/var/run/docker.sock \
 -v $PWD/dynamic_config.yml:/etc/traefik/dynamic_config.yml \
 --hostname frontend.local \
 --name traefik \
 --network limon \
 traefik:v2.1.3

