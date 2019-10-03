# Kademlia_Network
Configuring 10 nodes in docker container and setup communication among them

Configuring The Network:
Run these following commands
1. ./build.sh //It will execute the docker file and deploy the docker-compose.yml file and create 10 container nodes.
2. Now open a new terminal and write:
   docker ps -a //show you the list of containers
3. Open another terminal and execute this command:
   docker-compose run --rm kademlia_tracker //Tracker is a service to check communication between the nodes
4. List/Copy 2 container ids and write:
   containerID1:portnumber containerID2:portnumber (0a91b3bd7143:8000 7468fdfe3749:8000)
