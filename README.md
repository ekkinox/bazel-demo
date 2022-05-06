# Bazel demo

> [Golang](https://go.dev/) based demo project, relying on [Bazel](https://bazel.build/).

This project provides the following components:
- a [protobuf](https://developers.google.com/protocol-buffers) definition located in the [proto](proto) folder
- a go backend application acting as [gRPC](https://grpc.io/) server
- a go client application to be able to test interactions with the server
- a go [gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway) acting as a REST - gRPC reverse proxy

## Table of contents

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)

## Installation

You first need to install [Bazel](https://bazel.build/) on your machine.

Then, you can build the project using:

```shell
bazel build //...
```

## Usage

To run the go gRPC server, located in the [backend](backend) folder (will listen on `:50051`):

```shell
bazel run //backend
```

To start the go gRPC client, located in the [client](client) folder:

```shell
bazel run //backend
```

To start the go gRPC gateway, located in the [gateway](gateway) folder (will listen on `:8888`):

```shell
bazel run //gateway
```

## Commands

- Bazel (multi language build and test tool)
```shell
bazel query //... --output label_kind | sort | column -t #list dependency graph targets

bazel run //XXX #run target XXX 

bazel build //... # build all targets

bazel build //XXX #run target XXX 
```

- Gazelle (BUILD file generator)
```shell
bazel run //:gazelle # Build / refresh all BUILD files of the repository

bazel run //:gazelle -- update-repos -from_file=go.mod # Build / refresh using go.mod
```

