Walrus vs Slime [![Test](https://github.com/elct9620/wvs/actions/workflows/test.yml/badge.svg)](https://github.com/elct9620/wvs/actions/workflows/test.yml)
===

The original project is [BasalticStudio/Walrus-vs-Slime](https://github.com/BasalticStudio/Walrus-vs-Slime) I wrote in 2015, this project is rewritten with DDD (Domain-Driven Design) style which with colossal improvement.

## Requirements

* Golang ~> 1.18.2
* Node.js ~> 16.0
* Yarn ~> 1.22

## Usage

> **Note**
>
> TODO

## Project Structure

### Presentation

To handle the API interface or protocol

* `pkg/controller` - the HTTP request handler
* `pkg/command` - the WebSocket command handler
  * `pkg/command/parameter` - the WebSocket command parameters

### Application

The handle the user flow

* `internal/application` - the user flow, e.g. find a match

### Domain Object

* `internal/domain` - the Entity (or Aggregate) and Value Object
* `internal/repository` - the factory of Entity
* `internal/service` - the business logic

### Infrastructure

The foundation or utils

* `internal/infrastructure/container` - the shared object container
* `internal/infrastructure/hub` - the PubSub manager for WebSocket connection
* `internal/infrastructure/rpc` - the RPC command manager for WebSocket
* `internal/infrastructure/store` - the in-memory store

## Domain

> **Note**
>
> TODO
