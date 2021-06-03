docker rm -f 441finalClient

docker pull chlalstjr7/final_client

docker run -d --name 441finalClient -p 80:80 -p 443:443 -v /etc/letsencrypt:/etc/letsencrypt:ro chlalstjr7/final_client

exit
