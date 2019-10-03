docker rm $(docker ps -a -q) --force
docker image prune --force
docker build . -t kademlia
cd tracker/
docker build . -t kademlia_tracker
cd ..
docker-compose up --scale kademlia=10