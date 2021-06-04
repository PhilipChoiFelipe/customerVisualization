

docker pull chlalstjr7/userds

docker rm -f userDSInstance 

docker run -d \
    -p 80:80 \
    --name userDSInstance \
    --network 441finalNet \
    -e DSN=root:erinandmin441\@tcp\(441final-api.erinchang.me:3306\)/441finalDB \
    chlalstjr7/userds