# p2p-chat

A serverless p2p chat app written in go with [***fyne***](https://github.com/fyne-io/fyne/) for gui , ***grpc*** for networking and ***sqlite*** for database

## prerequisites

- ***gcc***
- ***protoc*** & go plugins (if you want to compile pb files):
  - install protocol buffer compiler [link](https://grpc.io/docs/protoc-installation/)
  - install protoc-gen-go and and protoc-gen-go-grpc by running `go install google.golang.org/protobuf/cmd/protoc-gen-go` and `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
  - have protoc and GOPATH in your PATH env
- ***fyne cli*** only required for bundling data
  - install fyne by running `go install fyne.io/fyne/v2/cmd/fyne@latest`

## building

- clone project and get get into directory `git clone https://github.com/rzaf/p2p-chat.git && cd p2p-chat`
- run `go mod download` to get required modules
- run `make bundle` if you want rebundle static assets
- run `make all` or run `make build` if you dont want to recompile protobuff files
- run `chat` in `bin`

## usage

- each rooms can only send messages to rooms that have same `uuid`,`secret` and have each others `ip`,`port`,`user uuid`
- each user has a public username that can be changed in seting

## features

- private and public chat rooms
- serverless app
- encrypted messaging
- light/dark mode
- public username

## todos

- storing caht messages
- notifying seen messages
- profile photo
- sending media messages (photo,video,music)
