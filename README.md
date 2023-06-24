Walrus vs Slime [![Test](https://github.com/elct9620/wvs/actions/workflows/test.yml/badge.svg)](https://github.com/elct9620/wvs/actions/workflows/test.yml)
===

The original project is [BasalticStudio/Walrus-vs-Slime](https://github.com/BasalticStudio/Walrus-vs-Slime) I wrote in 2015, this project is rewritten with DDD (Domain-Driven Design) style which with colossal improvement.

## Requirements

* Golang ~> 1.20.2

## Usage

> **Note**
>
> TODO

## Architecture

> **Note**
>
> Working in progress

### Entities

| Aggregate | Entity   | Description                              |
|-----------|----------|------------------------------------------|
| `Room`    | -        | Each game have their room with 2 players |
| `Room`    | `Player` | The player in the room                   |
