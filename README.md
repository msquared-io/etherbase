# Etherbase

Etherbase is:

`backend/`
- a backend writer service for easy emitting of events, setting of state, or executing functions on any contract.
- a backend reader service for reading events, state, and executing view functions on any contract.

`packages/contracts/contracts/Etherbase.sol`
- a contract for emitting events or setting state without writing any solidity code.

`packages/etherbase-client/`
- a react component library for interfacing with the backend and contract.

`packages/frontend/`
- an example frontend for interfacing with the backend and contract.

## Getting started with Etherbase

We host a live writer and reader that is used for all our demos.

## Getting Started with the codebase

1. Clone the repository

```bash
git clone https://github.com/msquared-io/etherbase.git
```

## Demos

### Example Frontend

#### 3D Realtime Game
[`packages/3dgame`](./packages/3dgame) is a 3D multiplayer game built with React, Three.js, and [`etherbase-client`](./packages/etherbase-client).
All player movement, names, and colors are synced across all clients using the blockchain. Player updates are 20Hz by default.

[The live demo](https://etherbase-demo-3dgame-496683047294.europe-west2.run.app/) runs on the [Somnia](https://somnia.network/) blockchain.

### MML

The MML Etherbase Demo is now located in a separate repository: [`msquared-io/etherbase-mml-react-starter-project`](https://github.com/msquared-io/etherbase-mml-react-starter-project). 
This project showcases [`etherbase-client`](./packages/etherbase-client) running within [MML](https://mml.io/), enabling simple on-chain interactions on any platform that supports MML.

The MML Etherbase Demo is an example of a server that serves a live MML (Metaverse Markup Language) document built with React.

## Technical Overview

### Etherbase Contract

See [`packages/contracts/contracts/Etherbase.sol`](./packages/contracts/contracts/Etherbase.sol) for the core logic.

### Etherbase Source Contract

#### Events

See [`packages/contracts/contracts/EtherbaseSource.sol`](./packages/contracts/contracts/EtherbaseSource.sol) for the core logic.

#### State

See [`packages/contracts/contracts/EtherDatabaseLib.sol`](./packages/contracts/contracts/EtherDatabaseLib.sol) for the core logic.

We store state similar to a KV store through a series of entries. Each entry has a path, a data type, and a data value.

For operations on this logic, the current implementation treats them as nested KV store documents. Imagine you have this data in your KV store:

```
users:
  alice:
    name: "Alice"
    balance: 100
  bob:
    name: "Bob"
    balance: 50
```

We store this as a series of entries in the KV store:

```
users.alice.name: "Alice"
users.alice.balance: 100
users.bob.name: "Bob"
users.bob.balance: 50
```

We take an opinionated approach in the backend where if you want to wipe the entire `alice` object, you need to set `alice` to a `nil` value.
This will iteratively remove all entries under `alice`.

There are strange behaviours when you have an object like so:

```
users.alice: "alice"
users.alice.name: "Alice"
users.alice.balance: 100
```

This breaks the nested KV store model as `users.alice` is a string and not an object. At the moment we ignore this, and if you set `users.alice` to a `nil` value, it will remove the entire `users.alice` entry.

You can bypass this behaviour by manually interacting with the entries system as you wish.
