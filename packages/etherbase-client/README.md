# etherbase-client

A library of hooks and functions to act as a client for an Etherbase server.

## Installation

```bash
npm install @msquared/etherbase-client
```

## Configuration

The hooks can be configured to use either direct blockchain interaction or a backend service:

```tsx
import { EtherbaseConfig, somnia } from "@msquared/etherbase-client";

const config: EtherbaseConfig = {
  chain: somnia, // the chain to use if using in-browser
  useBackend: true, // Use backend service instead of direct blockchain in-browser
  httpReaderUrl: "https://etherbase-reader-496683047294.europe-west2.run.app",
  wsReaderUrl: "wss://etherbase-reader-496683047294.europe-west2.run.app",
  wsWriterUrl: "wss://etherbase-writer-496683047294.europe-west2.run.app",
  privateKey: "0x...", // Optional private key for backend auth
};
```

Include the `EtherbaseProvider` in your app:

```tsx
<EtherbaseProvider config={config}>
  <App />
</EtherbaseProvider>
```

We also provide some server-side functions under `@msquared/etherbase-client/server`.

## Usage

### Events

#### Reading events

Use the `useEtherbaseEvents` hook to subscribe to contract events in real-time:

```tsx
const { error } = useEtherbaseEvents({
  contractAddress: "0x...",   // provide a contract address to listen to the events of
  contractAddresses: ["0xabcd...", "0xefff..."] // or provide an array
  events: [{ name: "EventName" }], // the array of events to care about
  onEvent: (event) => {
    // Handle event data
    console.log(event.args);
  },
});
```

You can also provide specific arguments to listen to for the events, such as listening to the minting of new coins in an ERC-20 contract:

```tsx
const { error } = useEtherbaseEvents({
  contractAddress: "0x...", // the ERC-20 contract address 
  events: [{ 
    name: "Transfer",
    args: {
      "from": ["0x00000000000000000000000000000000"]
    } }],
  onEvent: (event) => {
    // Handle mints data
    console.log(event.args);
  },
});
```


### Emitting events

If using Source contracts, you can emit events as you wish. First, you need to register the schema for your event.

#### Registering Events

Before emitting events, you need to register the schema of the event you want to emit. You can do so by either using the example frotend or by your own registration logic:

```tsx
const { registerEvent } = useEtherbaseSource({
  sourceAddress: "0x...",
});

await registerEvent({
  name: "EventName",
  args: [
    { name: "param1", argType: "string", isIndexed: true },
    { name: "param2", argType: "uint256", isIndexed: false },
  ],
});
```

#### Emitting an event for some schema

 use the `useEtherbaseSource` hook to emit custom events:

```tsx
const { emitEvent } = useEtherbaseSource({
  sourceAddress: "0x...",
});

// Emit an event
await emitEvent({
  name: "EventName",
  args: {
    param1: "value1",
    param2: "value2",
  },
});
```


### State

#### Reading state

The `useEtherstore` hook provides real-time state management for contracts:

```tsx
const { state, loading, error, update } = useEtherstore([contractAddress, "path", "to", "state"], {
  repoll: {
    onAnyContractEvent: true, // Automatically refresh when events occur
  },
});
```

You can also do multi-state filtering via the following syntax:

```tsx
const { state, loading, error, update } = useEtherstore([contractAddress, "path", "to", ["state1", "state2"]], {
  repoll: {
    onAnyContractEvent: true, // Automatically refresh when events occur
  },
});
```
This will return an object with the state of `state1` and `state2`.

#### Writing state

If using Source contracts, you can write state to an event source contract by using the `useEtherbaseSource` hook:

```tsx
const { setValue } = useEtherbaseSource({
  sourceAddress: "0x...",
});

// Update state
await setValue({
  path: {
    to: {
      state: "new value",
    },
  },
});
```

or use the `update` function to update the state:

```tsx
const { state, loading, error, update } = useEtherstore([contractAddress, "path", "to", "state"], {
  repoll: {
    onAnyContractEvent: true, // Automatically refresh when events occur
  },
});

await update({
  path: {
    to: {
      state: "new value",
    },
  },
});
```


### Contract execution

If using custom contracts, you can execute contract methods by using the `useEtherbaseContract` hook:

```tsx
const { execute } = useEtherbaseContract({
  contractAddress: "0x...",
});

await execute({
  methodName: "functionName",
  args: {
    arg1: "value1",
    arg2: "value2",
  },
});
```

which will use the performant backend to execute your contract. At the moment there is no way to retrieve the data from your contract, so you should make sure you are subscribing to relevant changes that are caused by your contract execution.