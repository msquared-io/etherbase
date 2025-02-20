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
  wsBaseUrl: "wss://...", // WebSocket URL for real-time updates
  httpBaseUrl: "https://...", // HTTP URL for REST endpoints
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

### Reading events

The `useEtherbaseEvents` hook allows you to subscribe to contract events in real-time.

```tsx
const { error } = useEtherbaseEvents({
  sourceAddress: "0x...", // Contract address
  events: [{ name: "EventName" }], // Array of events to subscribe to
  onEvent: (event) => {
    // Handle event data
    console.log(event.args);
  },
});
```

### Reading state

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

### Writing state

Use the `useEtherbaseSource` hook to write state to an event source contract:

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

or use the 'update' function to update the state:

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

### Emitting events

The `useEtherbaseSource` hook also allows emitting custom events:

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

### Registering Events

Before emitting events, they need to be registered with the contract:

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
