FROM larjim/kademlialab:latest

RUN mkdir /home/go/src/app
COPY . /home/go/src/app
WORKDIR /home/go/src/app
ENV GOPATH /home/go
ENV PATH="${GOPATH}/bin:${PATH}"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 /usr/local/go/bin/go build -o main .
ENTRYPOINT ["sh","./run.sh"]