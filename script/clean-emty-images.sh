#! /bin/bash

#REPOSITORY         TAG       IMAGE ID       CREATED          SIZE
#yangcodecamp-web   latest    439e64be68e1   4 minutes ago    24.8MB
#app                latest    9f36ce76e677   4 minutes ago    24.8MB
#<none>             <none>    2658e39cdd5a   8 minutes ago    24.8MB
#<none>             <none>    fa747ed69544   11 minutes ago   347MB
#<none>             <none>    40ded48c87ba   12 minutes ago   17.5MB
#<none>             <none>    4041bdcadf0a   13 minutes ago   638MB
#<none>             <none>    89db4e907c0b   17 minutes ago   17.5MB
#<none>             <none>    539ad062ec7e   22 minutes ago   17.5MB
#app-1              latest    c5d183136222   27 minutes ago   638MB
#mysql              latest    91b53e2624b4   2 weeks ago      565MB
# 找到所有的none镜像并删除
docker rmi $(docker images | grep "none" | awk '{print $3}')