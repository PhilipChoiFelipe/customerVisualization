GOOS=linux go build
docker build -t chlalstjr7/userds . 
go clean 

docker push chlalstjr7/userds 

ssh min@441final-api.erinchang.me  < deploy.sh