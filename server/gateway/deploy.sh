docker pull chlalstjr7/final_gateway

docker network rm 441finalNet
docker network create 441finalNet

export TLSCERT=/etc/letsencrypt/live/441final-api.erinchang.me/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/441final-api.erinchang.me/privkey.pem


docker rm -f 441finalRedis
docker run -d -p 6379:6379 --name 441finalRedis --network 441finalNet redis
#Running Gateway Server
docker rm -f 441finalGateway
docker run \
    -d \
    -e ADDR=:443 \
    -e SESSIONKEY=minseokchoiminseokchoi1234 \
    -e USERDSADDR=userDSInstance \
    -e REDISADDR=441final-api.erinchang.me:6379 \
    -e DSN=root:erinandmin441\@tcp\(441final-api.erinchang.me:3306\)/441finalDB \
    -e TLSKEY=$TLSKEY -e TLSCERT=$TLSCERT \
    -p 443:443 \
    -v /etc/letsencrypt:/etc/letsencrypt:ro \
    --name 441finalGateway \
    --network 441finalNet \
    chlalstjr7/final_gateway

exit