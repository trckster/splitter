FROM golang

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.1/zsh-in-docker.sh)"
RUN sed -i 's/powerlevel10k\/powerlevel10k/fishy/' ~/.zshrc

WORKDIR /go/splitter

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o splitter *.go

CMD ["./splitter"]