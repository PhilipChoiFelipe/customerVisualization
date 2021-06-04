GOOS=linux go build
docker build -t chlalstjr7/final_gateway . 
go clean

docker push chlalstjr7/final_gateway

ssh min@441final-api.erinchang.me < deploy.sh