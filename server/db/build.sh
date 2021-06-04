GOOS=linux go build

docker build -t chlalstjr7/441final_sqldb .

go clean

docker push chlalstjr7/441final_sqldb

ssh min@441final-api.erinchang.me < deploy.sh

