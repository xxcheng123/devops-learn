FROM registry.cn-beijing.aliyuncs.com/hub-mirrors/alpine
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
ENV TZ=Asia/Shanghai

WORKDIR /app

COPY ./build/server ./server

EXPOSE 8080

ENTRYPOINT ["server"]