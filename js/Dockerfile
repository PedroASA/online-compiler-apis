# FROM golang:1.5

# install protobuf from source
# RUN apt-get update && \
#    apt-get -y install git unzip build-essential autoconf libtool
#RUN git clone https://github.com/google/protobuf.git && \
#    cd protobuf && \
#    ./autogen.sh && \
#    ./configure && \
#    make && \
#    make install && \
#    ldconfig && \
#    make clean && \
#    cd .. && \
#    rm -r protobuf

# NOTE: for now, this docker image always builds the current HEAD version of
# gRPC.  After gRPC's beta release, the Dockerfile versions will be updated to
# build a specific version.

# Get the source from GitHub
# RUN go get google.golang.org/grpc
# Install protoc-gen-go
# RUN go get github.com/golang/protobuf/protoc-gen-go

FROM grpc/go

WORKDIR /app

COPY . .

# RUN "protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/route_guide.proto"

CMD go run server/server.go
EXPOSE 50051
