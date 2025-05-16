FROM alpine:3.21.3
# 项目入口执行命令构建镜像：docker build -t 镜像名:版本 .
###############################################################################
#                                INSTALLATION
###############################################################################

ENV WORKDIR=/web
RUN mkdir -p $WORKDIR/api $WORKDIR/public
RUN chmod +x $WORKDIR/public
ADD api/main_new  $WORKDIR/api/main
RUN chmod +x $WORKDIR/api/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR/api
CMD ["./main", "--gf.gcfg.file=config.prod.yaml"]