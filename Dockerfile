FROM node:10-stretch-slim

RUN apt-get update -y && \
    apt-get upgrade -y && \
    apt-get install -y gcc make

# install golang
RUN add-apt-repository ppa:gophers/archive && \
    apt-get update && \
    apt-get install golang-1.11-go

# export golang
RUN export GOPATH=$HOME/go && \
    export PATH="/usr/lib/go-1.11/bin:$PATH" && \
    export PATH=$PATH:$GOPATH/bin

# install serverless framework
RUN npm install -g try-thread-sleep && \
    npm install -g serverless --ignore-scripts spawn-sync

# set credentials to serverless
# RUN sls config credentials --provider aws --key $AWS_ACCESS_KEY_ID --secret $AWS_SECRET_ACCESS_KEY

WORKDIR /app
