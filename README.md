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

## Design

### Layers

| Layer          | Description                                   | Package                                    |
|----------------|-----------------------------------------------|--------------------------------------------|
| Presentation   | To handle the end-user interactive            | `pkg/controller`, `pkg/data`, `pkg/event`  |
| Application    | To handle the user flow                       | `internal/application`                     |
| Domain         | The business logic                            | `internal/domain`, `internal/repository`   |
| Infrastructure | Non-domain related behaviors, e.g. data store | `internal/infrastructure`, `internal/data` |

### Domain

> **Note**
>
> TODO
