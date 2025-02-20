# Somnia + Etherbase Realtime 3D Game

This is a demo of a fully networked 3D game that uses Etherbase to interact with the Somnia blockchain.

## Running the game locally on Somnia

1. Create a new source

[Follow these steps]() to create a new source.

2. Add the PlayerUpdate event schema to the source

Add the following event to the source:

```solidity
event PlayerUpdate(int256 posX, int256 posY, int256 posZ, string color, string name, string playerId, uint256 timestamp)
```

which is used by [`src/app/page.tsx`](./src/app/page.tsx) to send and receive updates from the blockchain. In the file, modify the `sourceAddress` to the address of the source you created in step 1.

```tsx
// read:
 useEtherbaseEvents({
    sourceAddress,
    events: [{ name: "PlayerUpdate" }],
    onEvent: handlePlayerUpdate,
  })

// write:
const { emitEvent } = useEtherbaseSource({
  sourceAddress,
})
```

3. Create a new private key for messages

Follow the instructions [here]() to create a new private key.

4. Configure EtherbaseConfig

Configure the EtherbaseConfig in [`src/app/etherbaseConfig.ts`](./src/app/etherbaseConfig.ts) with your private key.

```ts
privateKey: "0x..."
```

5. Run the game

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

If you'd like to run the game using a local backend, run `npm run dev-local`.
