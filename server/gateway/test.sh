GOOS=linux go build
docker build -t chlalstjr7/final_gateway . 
go clean

docker network create 441finalNet

#Running Redis Server
docker rm -f 441finalRedis 
docker run -d -p 6379:6379 --name 441finalRedis --network 441finalNet redis

#Running gateway server 
docker rm -f 441finalGateway
docker run \
    -d \
    -e ADDR=:80 \
    -e SESSIONKEY=erinandmin441erinandmin441 \
    -e REDISADDR=host.docker.internal:6379 \
    -e DSN=root:erinandmin441\@tcp\(host.docker.internal:3306\)/441finalDB \
    -p 80:80 \
    --name 441finalGateway \
    --network 441finalNet \
    chlalstjr7/final_gateway
