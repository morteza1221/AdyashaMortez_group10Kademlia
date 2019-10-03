FROM larjim/kademlialab:latest

RUN mkdir /home/go/src/app
COPY . /home/go/src/app
COPY kademlia /home/go/src/kademlia
COPY error /home/go/src/error
WORKDIR /home/go/src/app
ENV GOPATH /home/go
ENV PATH="${GOPATH}/bin:${PATH}"
RUN CGO_ENABLED=0 GOOS=linux GOARCH=386 /usr/local/go/bin/go build -o main .
EXPOSE 8000/udp
ENTRYPOINT ["sh","./run.sh"]