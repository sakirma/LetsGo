FROM golang:latest
EXPOSE 9000

WORKDIR /code/LetsGooooo
COPY . .


RUN set -x \
    # go get revel
    && go get -v github.com/revel/revel \
    && go get -v github.com/revel/cmd/revel \
    && go mod tidy