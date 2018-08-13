FROM alpine:latest
ADD ./bin /project
CMD chmod +x /project/crud-api
EXPOSE 8078
ENV MONGO_ADDRESS=192.168.10.18:27017
ENTRYPOINT /project/crud-api