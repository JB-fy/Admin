FROM alpine:3.21.3
###############################################################################
# 构建镜像
    # 项目根目录执行命令：docker build -t 镜像名:版本 .
    # 打包镜像：docker save -o p2p.tar p2p:0.0.1
    # 上传镜像解压：docker load -i p2p.tar

# 宿主机需执行以下命令
    # 创建项目目录：mkdir -p /宿主机路径/public/upload /宿主机路径/api/manifest
    # 赋予写入权限：chmod -R 755 /宿主机路径/public/upload

# 镜像启动
    # 执行命令：docker run -d --restart unless-stopped --name 名称 --network host -v /宿主机路径/public/upload:/app/public/upload -v /宿主机路径/api/manifest:/app/api/manifest -e SERVER_NETWORK_IP=$(curl -s --max-time 3 ifconfig.me || curl -s --max-time 3 https://ipinfo.io/ip || curl -s --max-time 3 https://checkip.amazonaws.com || curl -s --max-time 3 https://icanhazip.com || curl -s --max-time 3 https://api.ipify.org) -e SERVER_LOCAL_IP=$(hostname -I | awk '{print $1}') 镜像名:版本
###############################################################################

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR=/app
# ADD public/admin  $WORKDIR/public/admin
ADD api/main_new  $WORKDIR/api/main
RUN chmod +x $WORKDIR/api/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR/api
CMD ["./main", "--gf.gcfg.file=config.prod.yaml"]