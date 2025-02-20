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


## Getting Started

1. Clone the repository

```bash
git clone https://github.com/msquared-io/etherbase.git
```



## Demos

### Example Frontend

#### 3D Realtime Game
[`packages/3d-game`](./packages/3d-game) is a 3D multiplayer game built with React, Three.js, and [`etherbase-client`](./packages/etherbase-client).
All player movement, names, and colors are synced across all clients using the blockchain. Player updates are 20Hz by default.

[The live demo](https://etherbase-demo-3dgame-496683047294.europe-west2.run.app/) runs on the [Somnia](https://somnia.network/) blockchain.

### MML

[`packages/mml-demo`](./packages/mml-demo) is [`etherbase-client`](./packages/etherbase-client) running within [MML](https://mml.io/), allowing for simple on-chain interaction on any platform that supports MML.
More information about how to use and deploy this MML can be found in the [README of `packages/mml-demo`](./packages/mml-demo/README.md).