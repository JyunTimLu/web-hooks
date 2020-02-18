FROM golang:1.11.2-alpine
WORKDIR /web-hooks
ADD . /web-hooks
RUN cd /web-hooks && go build
EXPOSE 8881
RUN rm main.go
RUN rm dockerfile
ENTRYPOINT ./web-hooks