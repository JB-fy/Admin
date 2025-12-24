FROM alpine
###############################################################################
# 构建镜像（项目根目录执行）
    # 生成并打包镜像：docker build -t 项目名:版本 -f /server/app/项目名/deploy/docker/Dockerfile . && docker save -o 项目名.tar 项目名:版本
    # 上传镜像解压：docker load -i 项目名.tar

# 宿主机需执行以下命令
    # 创建项目目录：mkdir -p /server/app/项目名/public/upload
    # 赋予写入权限：chmod -R 755 /server/app/项目名/public/upload

# 镜像启动
    # 执行命令：docker run -d --restart unless-stopped --name 项目名 --network host -v /etc/localtime:/etc/localtime:ro -v /server/app/项目名/public:/server/app/项目名/public -v /server/app/项目名/api/manifest:/server/app/项目名/api/manifest -v /server/app/项目名/api/log:/server/app/项目名/api/log -e SERVER_NETWORK_IP=$(curl -s --max-time 3 ifconfig.me || curl -s --max-time 3 https://ipinfo.io/ip || curl -s --max-time 3 https://checkip.amazonaws.com || curl -s --max-time 3 https://icanhazip.com || curl -s --max-time 3 https://api.ipify.org) -e SERVER_LOCAL_IP=$(hostname -I | awk '{printf "%s", $1}') 项目名:版本
###############################################################################

###############################################################################
#                                INSTALLATION
###############################################################################

# RUN apk --no-cache add ffmpeg

ENV WORKDIR=/server/app/admin
# ADD public  $WORKDIR/public
ADD api/main_new  $WORKDIR/api/admin
RUN chmod +x $WORKDIR/api/admin

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR/api
CMD ["./admin", "--gf.gcfg.file=config.prod.yaml"]