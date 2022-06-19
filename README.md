Walrus vs Slime [![Test](https://github.com/elct9620/wvs/actions/workflows/test.yml/badge.svg)](https://github.com/elct9620/wvs/actions/workflows/test.yml)
===

The original project is [BasalticStudio/Walrus-vs-Slime](https://github.com/BasalticStudio/Walrus-vs-Slime) I wrote in 2015, this project is rewritten with DDD (Domain-Driven Design) style which with colossal improvement.

## Requirements

* Golang ~> 1.18.2

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
* `internal/engine` - the game loop handler

## Domain

### Entity

* `Player` - each WebSocket connection will create a `Player`
* `Match` - a pair of `Players` in same game
* `Team` - the team of walrus or slime
* `Tower` - the aggregate root of `Player`'s `Team`
* `Monster` - TODO

### Value Object

* `Mana` - the value which `Tower` can be used to spawn `Monster`

### Service

* `BroadcastService` - the service to broadcast command to a `Match`
* `RecoveryService` - the service to manage the `Mana` recovery
* `LoopService` - the service to create the game loop
