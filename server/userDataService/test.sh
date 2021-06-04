GOOS=linux go build
docker build -t chlalstjr7/userds . 
go clean 

# docker push chlalstjr7/userDS 

# ssh min@441final-api.erinchang.me  < deploy.sh
docker rm -f userDSInstance 

docker run -d \
    -p 80:80 \
    --name userDSInstance \
    --network 441finalNet \
    -e DSN=root:erinandmin441\@tcp\(441final-api.erinchang.me:3306\)/441finalDB \
    chlalstjr7/userds
