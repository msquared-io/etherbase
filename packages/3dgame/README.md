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

which is used by [`src/app/page.tsx`](./src/app/page.tsx) to send and receive updates from the blockchain.

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

and set the source address to your source address with the PlayerUpdate event in step 2:

```ts
const sourceAddress = "0x..."
```

5. Run the game

```bash
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the game.

If you'd like to run the game using a local backend, run `npm run dev-local`.


## Blocks

The blocks demo is available at [http://localhost:3000/blocks](http://localhost:3000/blocks).

It uses `useEtherstore` to read and write state of voxel blocks to the Etherbase Source contract.

We use the hook:

```tsx
const { state, loading, error, update } = useEtherstore([
    sourceAddress,
    "blocks",
  ])
```

We want to store the position and type of each block whenever the user places a block using the left mouse button.

To write to the state, we use the `update` function from the hook:

```tsx
const position = { x: 1, y: 2, z: 3 }
const blockId = `${position.x},${position.y},${position.z}`
const blockType = 1
update({
  blocks: {
    [blockId]: { position, type: blockType },
  },
})
```

To read the state, we use `state` from the hook:

```tsx
  const [blocks, setBlocks] = useState<Map<string, Block>>(
    () => new Map(initialBlocks),
  )

  useEffect(() => {
      const newBlocks = new Map(initialBlocks)

      for (const [key, blockData] of Object.entries(
        state.blocks as EtherstoreState,
      )) {
        const pos = blockData.position as { x: number; y: number; z: number }

        newBlocks.set(key, {
          position: new THREE.Vector3(
            pos.x,
            pos.y,
            pos.z,
          ),
          type: Number(blockData.type),
        })
      }
      setBlocks(newBlocks)
    }
  }, [state, initialBlocks])
```
You could do the same thing without a `useEffect` by using the `onStateChange` callback from the hook.

The user can also clear the blocks by clicking the C key:


```tsx

```