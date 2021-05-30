docker build -t chlalstjr7/441final_sqldb .

docker rm -f 441finalSql

docker run \
    -d \
    -p 3306:3306 \
    -e MYSQL_DATABASE=441finalDB \
    -e MYSQL_ROOT_PASSWORD=erinandmin441\
    --network 441finalNet \
    --name 441finalSql \
    chlalstjr7/441final_sqldb
