#359mb
# FROM golang:1.19-alpine  

# 7mb (almost the size of the statically built binary / exe)
FROM scratch  
WORKDIR /app
COPY src/main ./
ENTRYPOINT [ "./main" ]

EXPOSE 9090
