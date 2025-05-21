FROM alpine:3.21.3
###############################################################################
# 构建镜像
    # 项目根目录执行命令：docker build -t 镜像名:版本 .
    # 打包镜像：docker save -o 文件名.tar 镜像名:版本
    # 上传镜像解压：docker load -i 文件名.tar

# 宿主机需执行以下命令
    # 创建项目目录：mkdir -p /server/app/项目名/public/upload
    # 赋予写入权限：chmod -R 755 /server/app/项目名/public/upload

# 镜像启动
    # 执行命令：docker run -d --restart unless-stopped --name 名称 --network host -v /etc/localtime:/etc/localtime:ro -v /server/app/项目名/public:/server/app/项目名/public -v /server/app/项目名/api/manifest:/server/app/项目名/api/manifest -v /server/app/项目名/api/log:/server/app/项目名/api/log -e SERVER_NETWORK_IP=$(curl -s --max-time 3 ifconfig.me || curl -s --max-time 3 https://ipinfo.io/ip || curl -s --max-time 3 https://checkip.amazonaws.com || curl -s --max-time 3 https://icanhazip.com || curl -s --max-time 3 https://api.ipify.org) -e SERVER_LOCAL_IP=$(hostname -I | awk '{print $1}') 镜像名:版本
###############################################################################

###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR=/server/app/项目名
# ADD public  $WORKDIR/public
ADD api/main_new  $WORKDIR/api/项目名
RUN chmod +x $WORKDIR/api/项目名

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR/api
CMD ["./项目名", "--gf.gcfg.file=config.prod.yaml"]