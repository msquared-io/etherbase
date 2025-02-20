import { useEffect, useState, useRef } from "react"
import {
  useEtherbaseEvents,
  useEtherbaseSource,
} from "@msquared/etherbase-client"
import type { Address } from "viem"

const FLOAT_PRECISION = 1000 // 3 decimal places
const PLAYER_TIMEOUT_MS = 10000 // Remove players after 10s of no updates

interface PlayerSnapshot {
  /** The **server** timestamp we received in the event. */
  serverTimestamp: number
  /** The **client** local time when we received the update (Date.now()). */
  clientReceiveTime: number
  /** Position data */
  position: {
    x: number
    y: number
    z: number
  }
  color: string
  name: string
}

interface RemotePlayer {
  snapshots: PlayerSnapshot[]
  latency: number
}

// Helper function from 3D game
function getRandomColor() {
  const hue = Math.random() * 360
  const saturation = 70 + Math.random() * 30 // 70-100%
  const lightness = 45 + Math.random() * 10 // 45-55%
  return `hsl(${hue}, ${saturation}%, ${lightness}%)`
}

// Simple UUID generator from 3D game
function generateUUID() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
    const r = (Math.random() * 16) | 0
    const v = c === "x" ? r : (r & 0x3) | 0x8
    return v.toString(16)
  })
}

export function CubeWorld() {
  const sourceAddress = "0x1be5f11a7a7e6C78ddE5aF8DaDf684049d1d5907" as Address
  const { emitEvent } = useEtherbaseSource({ sourceAddress })

  // Generate values like in 3D game
  const [playerId] = useState(generateUUID())
  const [playerName] = useState(`player#${Math.floor(Math.random() * 10000)}`)
  const [playerColor] = useState(getRandomColor())

  // Remote players state
  const [remotePlayers, setRemotePlayers] = useState<
    Record<string, RemotePlayer>
  >({})

  // Track last state we sent
  const lastSentStateRef = useRef<{
    posX: number
    posY: number
    posZ: number
    color: string
    name: string
  } | null>(null)

  // Track last update send time
  const lastUpdateTimeRef = useRef(Date.now())

  /**
   * Handle incoming player updates
   */
  const handlePlayerUpdate = (event: any) => {
    // Decode position
    const posX = (event.args.posX as number) / FLOAT_PRECISION
    const posY = (event.args.posY as number) / FLOAT_PRECISION
    const posZ = (event.args.posZ as number) / FLOAT_PRECISION
    const { color, name, timestamp } = event.args
    const now = Date.now()

    setRemotePlayers((prev) => {
      const newRemotePlayers = { ...prev }
      const id = event.args.playerId

      if (!newRemotePlayers[id]) {
        newRemotePlayers[id] = {
          snapshots: [],
          latency: now - timestamp,
        }
      }

      const snapshot: PlayerSnapshot = {
        serverTimestamp: timestamp,
        clientReceiveTime: now,
        position: { x: posX, y: posY, z: posZ },
        color,
        name,
      }

      newRemotePlayers[id].latency = now - timestamp
      newRemotePlayers[id].snapshots.push(snapshot)

      // Keep snapshot buffer from growing unbounded (keep last ~20)
      if (newRemotePlayers[id].snapshots.length > 20) {
        newRemotePlayers[id].snapshots.shift()
      }

      return newRemotePlayers
    })
  }

  // Subscribe to player updates
  useEtherbaseEvents({
    sourceAddress,
    events: [{ name: "PlayerUpdate2" }],
    onEvent: handlePlayerUpdate,
  })

  /**
   * Periodically remove stale players who haven't updated in a while
   */
  useEffect(() => {
    const cleanup = setInterval(() => {
      const now = Date.now()
      setRemotePlayers((prev) => {
        const newRemotePlayers = { ...prev }
        let changed = false

        for (const [id, rp] of Object.entries(newRemotePlayers)) {
          const latestSnapshot = rp.snapshots[rp.snapshots.length - 1]
          if (latestSnapshot) {
            const timeSinceLast = now - latestSnapshot.clientReceiveTime
            if (timeSinceLast > PLAYER_TIMEOUT_MS) {
              delete newRemotePlayers[id]
              changed = true
            }
          } else {
            delete newRemotePlayers[id]
            changed = true
          }
        }

        return changed ? newRemotePlayers : prev
      })
    }, 1000)
    return () => clearInterval(cleanup)
  }, [])

  useEffect(() => {
    const positionProbe = document.getElementById("position-probe")
    if (!positionProbe) {
      console.error("positionProbe not found")
      return
    }

    const handlePositionMove = async (event: any) => {
      console.log(
        "handlePositionMove",
        JSON.stringify(event.detail.elementRelative),
      )
      const now = Date.now()
      const { position } = event.detail.elementRelative

      const eventName = "PlayerUpdate2"
      const eventArgs = {
        posX: Math.round(position.x * FLOAT_PRECISION),
        posY: Math.round(position.y * FLOAT_PRECISION),
        posZ: Math.round(position.z * FLOAT_PRECISION),
        isJumping: false,
        color: playerColor,
        name: playerName,
        timestamp: now,
        playerId: playerId,
      }

      const timeSinceLastUpdate = now - lastUpdateTimeRef.current
      const hasChanged =
        !lastSentStateRef.current ||
        lastSentStateRef.current.posX !== eventArgs.posX ||
        lastSentStateRef.current.posY !== eventArgs.posY ||
        lastSentStateRef.current.posZ !== eventArgs.posZ ||
        lastSentStateRef.current.color !== eventArgs.color ||
        lastSentStateRef.current.name !== eventArgs.name

      // Send update if state changed OR timeout
      if (hasChanged || timeSinceLastUpdate >= 5000) {
        try {
          await emitEvent({
            name: eventName,
            args: eventArgs,
          })
          lastSentStateRef.current = {
            posX: eventArgs.posX,
            posY: eventArgs.posY,
            posZ: eventArgs.posZ,
            color: eventArgs.color,
            name: eventArgs.name,
          }
          lastUpdateTimeRef.current = now
        } catch (err) {
          console.error("Error emitting event:", err)
        }
      }
    }

    positionProbe.addEventListener("positionmove", handlePositionMove)
    return () => {
      positionProbe.removeEventListener("positionmove", handlePositionMove)
    }
  }, [emitEvent, playerColor, playerName, playerId])

  return (
    <div>
      <m-position-probe
        range="20"
        radius="20"
        debug="true"
        id="position-probe"
        interval="50"
      ></m-position-probe>
      {/* Render each player's latest position as an m-cube */}
      {Object.entries(remotePlayers).map(([id, player]) => {
        const latestSnapshot = player.snapshots[player.snapshots.length - 1]
        if (!latestSnapshot || id === playerId) return null

        return (
          <>
            <m-cube
              color={latestSnapshot.color}
              x={latestSnapshot.position.x.toString()}
              y={latestSnapshot.position.y.toString()}
              z={latestSnapshot.position.z.toString()}
            ></m-cube>
            <m-label
              y={(latestSnapshot.position.y + 1).toString()}
              x={latestSnapshot.position.x.toString()}
              z={latestSnapshot.position.z.toString()}
              width="3"
              height="0.4"
              color="#ffffff"
              font-color="#000000"
              alignment="center"
              content={`${latestSnapshot.name} (${Math.round(player.latency)}ms)`}
            ></m-label>
          </>
        )
      })}
    </div>
  )
}
