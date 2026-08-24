# Coal Company — Concurrent Mining Simulation & REST API

A Go project simulating a coal-mining company: miners generate income concurrently over time, and a REST API lets you buy miners and factory upgrades and monitor production in real time.

## Features

- Passive income simulation running as a background goroutine
- Three miner classes (Junior / Middle / Senior), each earning coal independently via its own goroutine and channel
- Purchasable factory upgrades (Pick, Ventilation, Minecarts) that affect production
- Graceful shutdown via `context.Context` cancellation
- Thread-safe shared state protected with mutexes
- Routing with [gorilla/mux](https://github.com/gorilla/mux)

## Tech stack

Go 1.26 · gorilla/mux · goroutines, channels, context, sync

## Project structure

```
.
├── main.go
├── factory/          # core simulation logic
│   ├── items/        # purchasable factory upgrades
│   └── miners/        # miner classes & concurrent income logic
└── server/            # HTTP layer: routing, handlers, DTOs
```

## API

| Method | Endpoint                            | Description                              |
|--------|---------------------------------------|--------------------------------------------|
| GET    | `/factory`                           | Get factory status (balance, miners, work time) |
| DELETE | `/factory`                           | Stop the factory                          |
| GET    | `/factory/miners/classification`     | Get miner cost info                       |
| POST   | `/factory/miners`                    | Buy a miner                               |
| GET    | `/factory/miners`                    | Get all miners                            |
| GET    | `/factory/miners?working=true`       | Get currently working miners              |
| POST   | `/factory/items`                     | Buy a factory item/upgrade                |
| GET    | `/factory/items/classification`      | Get available items and their costs       |
| GET    | `/factory/items`                     | Get all purchased items                   |

## Running locally

```bash
go run main.go
```

Server starts on `localhost:9999`.

## Possible next steps

- Persist factory state between restarts
- Add unit tests for the concurrent income logic
- Add a simple frontend to visualize the simulation
