#!/bin/bash
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-room_service -f
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-movie_service -f
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-user_service -f
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-reservation_service -f
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-screening_service -f
docker rmi terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam:develop-client -f
sudo docker pull terraform.cs.hm.edu:5043/ob-vss-ws19-blatt-4-myteam -a