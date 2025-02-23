"use client"

import { useEffect, useRef, useState, useCallback } from "react"
import type React from "react"
import { Canvas, useFrame } from "@react-three/fiber"
import * as THREE from "three"
import { motion } from "framer-motion"
import { Text } from "@react-three/drei" // For displaying player names
import {
  useEtherbaseEvents,
  useEtherbaseSource,
  useEtherstore,
} from "@msquared/etherbase-client"
import type { Address } from "viem"

// Import the bot simulation function
import { startBotSimulation } from "../bots"
import { sourceAddress } from "./etherbaseConfig"

/** -------------------------------
 *  DATA INTERFACES
 *  ------------------------------- */
interface PlayerState {
  posX: number
  posY: number
  posZ: number
  color: string
  name: string
}

interface PlayerSnapshot {
  /** The **server** timestamp we received in the event. */
  serverTimestamp: number
  /** The **client** local time when we received the update (Date.now()). */
  clientReceiveTime: number
  /** Position and other data from the event. */
  position: THREE.Vector3
  color: string
  name: string
}

/**
 * We keep a small snapshot buffer for each remote player,
 * as well as a computed "current" position we display after interpolation.
 */
interface RemotePlayer {
  snapshots: PlayerSnapshot[]
  displayPosition: THREE.Vector3 // The position we render this frame
  latency: number
}

interface ChatMessage {
  playerId: string
  message: string
}

const FLOAT_PRECISION = 1000 // 3 decimal places
const FORCE_UPDATE_INTERVAL_MS = 5000 // Force local update every 5s
const PLAYER_TIMEOUT_MS = 10000 // Remove players after 10s of no updates
const INTERP_BUFFER_MS = 100 // How far behind "now" we interpolate (e.g. 100-200ms)

/** Physics constants */
const GRAVITY = -9.8
const JUMP_FORCE = 5
const MOVE_SPEED = 3 // units/second

/** -------------------------------
 *  PLAYER CUBE COMPONENT
 *  ------------------------------- */
function PlayerCube({
  position,
  color,
  name,
  latency,
  isLocal,
}: {
  position: THREE.Vector3
  color: string
  name: string
  latency: number
  isLocal: boolean
}) {
  const meshRef = useRef<THREE.Mesh>(null)
  const textRef = useRef<THREE.Mesh>(null)

  useFrame(() => {
    if (meshRef.current) {
      meshRef.current.position.copy(position)
    }
    if (textRef.current) {
      textRef.current.position.set(position.x, position.y + 0.6, position.z)
    }
  })

  return (
    <>
      <mesh ref={meshRef}>
        <boxGeometry args={[0.5, 0.5, 0.5]} />
        <meshStandardMaterial color={color} />
      </mesh>
      <Text
        ref={textRef}
        fontSize={0.2}
        color="#ffffff"
        anchorX="center"
        anchorY="bottom"
      >
        {isLocal ? `${name} (you)` : `${name} (${Math.round(latency)}ms)`}
      </Text>
    </>
  )
}

function Scene({
  playerId,
  playerName,
  playerColor,
  sourceAddress,
  interpolationEnabled,
  emitEvent,
}: {
  playerId: string
  playerName: string
  playerColor: string
  sourceAddress: Address
  interpolationEnabled: boolean
  emitEvent: any
}) {
  // Local player states
  const [position, setPosition] = useState(new THREE.Vector3(0, 0.25, 5))
  const [velocity, setVelocity] = useState(new THREE.Vector3(0, 0, 0))
  const [isJumping, setIsJumping] = useState(false)

  // Remote players: each has a snapshot buffer + a display position
  const [remotePlayers, setRemotePlayers] = useState<
    Record<string, RemotePlayer>
  >({})

  // Keyboard input tracking
  const keys = useRef<Record<string, boolean>>({})
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      keys.current[e.code] = true
    }
    const handleKeyUp = (e: KeyboardEvent) => {
      keys.current[e.code] = false
    }
    window.addEventListener("keydown", handleKeyDown)
    window.addEventListener("keyup", handleKeyUp)
    return () => {
      window.removeEventListener("keydown", handleKeyDown)
      window.removeEventListener("keyup", handleKeyUp)
    }
  }, [])

  /**
   * Basic local physics/movement on each frame.
   */
  useFrame((_, delta) => {
    const newVelocity = velocity.clone()
    const newPosition = position.clone()

    // Basic movement
    if (keys.current["KeyW"]) newPosition.z -= MOVE_SPEED * delta
    if (keys.current["KeyS"]) newPosition.z += MOVE_SPEED * delta
    if (keys.current["KeyA"]) newPosition.x -= MOVE_SPEED * delta
    if (keys.current["KeyD"]) newPosition.x += MOVE_SPEED * delta

    // Jump
    if (!isJumping && keys.current["Space"]) {
      newVelocity.y = JUMP_FORCE
      setIsJumping(true)
    }

    // Gravity
    newVelocity.y += GRAVITY * delta

    // Position update
    newPosition.addScaledVector(newVelocity, delta)

    // Floor collision
    if (newPosition.y <= 0.25) {
      newPosition.y = 0.25
      newVelocity.y = 0
      setIsJumping(false)
    }

    setVelocity(newVelocity)
    setPosition(newPosition)
  })

  /**
   * Handle incoming player updates and store them in a snapshot buffer.
   */
  const handlePlayerUpdate = useCallback(
    (event: any) => {
      // Ignore our own updates
      if (event.args.playerId === playerId) return

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
            displayPosition: new THREE.Vector3(posX, posY, posZ),
            latency: now - timestamp,
          }
        }

        const snapshot: PlayerSnapshot = {
          serverTimestamp: timestamp,
          clientReceiveTime: now,
          position: new THREE.Vector3(posX, posY, posZ),
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
    },
    [playerId],
  )

  // Subscribe to player updates (useEtherbaseEvents)
  useEtherbaseEvents({
    contractAddress: sourceAddress,
    events: [{ name: "PlayerUpdate" }],
    onEvent: handlePlayerUpdate,
  })

  // Refs for local player broadcast
  const positionRef = useRef(position)
  useEffect(() => {
    positionRef.current = position
  }, [position])

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
   * Periodically send our local player state if it changed
   * or a forced time interval has passed.
   */
  useEffect(() => {
    const interval = setInterval(async () => {
      const now = Date.now()
      const eventName = "PlayerUpdate"
      const eventArgs = {
        posX: Math.round(positionRef.current.x * FLOAT_PRECISION),
        posY: Math.round(positionRef.current.y * FLOAT_PRECISION),
        posZ: Math.round(positionRef.current.z * FLOAT_PRECISION),
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
      if (hasChanged || timeSinceLastUpdate >= FORCE_UPDATE_INTERVAL_MS) {
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
    }, 50) // Check 20x/sec
    return () => clearInterval(interval)
  }, [emitEvent, playerColor, playerName, playerId])

  /**
   * Periodically remove stale players who haven't updated in a while.
   */
  useEffect(() => {
    const cleanup = setInterval(() => {
      const now = Date.now()
      setRemotePlayers((prev) => {
        const newRemotePlayers = { ...prev }
        let changed = false

        for (const [id, rp] of Object.entries(newRemotePlayers)) {
          // If the newest snapshot is older than PLAYER_TIMEOUT_MS, remove
          const latestSnapshot = rp.snapshots[rp.snapshots.length - 1]
          if (latestSnapshot) {
            const timeSinceLast = now - latestSnapshot.clientReceiveTime
            if (timeSinceLast > PLAYER_TIMEOUT_MS) {
              delete newRemotePlayers[id]
              changed = true
            }
          } else {
            // no snapshots at all
            delete newRemotePlayers[id]
            changed = true
          }
        }

        return changed ? newRemotePlayers : prev
      })
    }, 1000)
    return () => clearInterval(cleanup)
  }, [])

  /**
   * Interpolate remote players each frame if smoothing is enabled.
   *
   * We pick a render timestamp = now - INTERP_BUFFER_MS so we can (hopefully)
   * find two snapshots in the buffer that bracket that time and lerp between them.
   */
  useFrame(() => {
    setRemotePlayers((prev) => {
      const newRemotePlayers = { ...prev }

      for (const id in newRemotePlayers) {
        const rp = newRemotePlayers[id]
        const snapshots = rp.snapshots

        if (snapshots.length === 0) continue

        if (!interpolationEnabled) {
          // When smoothing is disabled, just use the latest snapshot
          rp.displayPosition.copy(snapshots[snapshots.length - 1].position)
          continue
        }

        if (snapshots.length < 2) {
          // Not enough data to interpolate
          rp.displayPosition.copy(snapshots[0].position)
          continue
        }

        // Find the first snapshot with time > renderTime
        const renderTime = Date.now() - INTERP_BUFFER_MS
        let i = 0
        while (
          i < snapshots.length &&
          snapshots[i].clientReceiveTime < renderTime
        ) {
          i++
        }

        if (i === 0) {
          // All snapshots are after renderTime, so just use the earliest
          rp.displayPosition.copy(snapshots[0].position)
        } else if (i >= snapshots.length) {
          // All snapshots are before renderTime, so just use the latest
          rp.displayPosition.copy(snapshots[snapshots.length - 1].position)
        } else {
          // Snapshots[i-1] is just before/at renderTime,
          // snapshots[i] is just after renderTime
          const prevSnap = snapshots[i - 1]
          const nextSnap = snapshots[i]
          const t0 = prevSnap.clientReceiveTime
          const t1 = nextSnap.clientReceiveTime
          const ratio = (renderTime - t0) / (t1 - t0)

          rp.displayPosition.lerpVectors(
            prevSnap.position,
            nextSnap.position,
            ratio,
          )
        }
      }

      return newRemotePlayers
    })
  })

  /** -------------------------------
   *  RENDER
   *  ------------------------------- */
  return (
    <>
      {/* Our local player */}
      <PlayerCube
        position={position}
        color={playerColor}
        name={playerName}
        latency={0}
        isLocal={true}
      />

      {/* Remote players */}
      {Object.entries(remotePlayers).map(([id, rp]) => {
        // Use the latest snapshot's color, name, etc. (or the last in the array)
        const lastSnap = rp.snapshots[rp.snapshots.length - 1]
        return (
          <PlayerCube
            key={id}
            position={rp.displayPosition}
            color={lastSnap?.color ?? "#ffffff"}
            name={lastSnap?.name ?? "??"}
            latency={rp.latency}
            isLocal={false}
          />
        )
      })}

      {/* Ground Plane */}
      <mesh rotation={[-Math.PI / 2, 0, 0]} position={[0, 0, 0]} receiveShadow>
        <planeGeometry args={[50, 50]} />
        <meshStandardMaterial color="green" />
      </mesh>

      {/* Lights */}
      <ambientLight intensity={0.5} />
      <directionalLight position={[10, 10, 5]} intensity={1} />
    </>
  )
}

/** Helper function to generate random color */
function getRandomColor() {
  const hue = Math.random() * 360
  const saturation = 70 + Math.random() * 30 // 70-100%
  const lightness = 45 + Math.random() * 10 // 45-55%
  return `hsl(${hue}, ${saturation}%, ${lightness}%)`
}

/** -------------------------------
 *  MAIN PAGE
 *  ------------------------------- */
// Replace the crypto.randomUUID() call with a simple UUID generator function
function generateUUID() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, (c) => {
    const r = (Math.random() * 16) | 0
    const v = c === "x" ? r : (r & 0x3) | 0x8
    return v.toString(16)
  })
}

export default function Home() {
  const [playerName, setPlayerName] = useState(
    `player#${Math.floor(Math.random() * 10000)}`,
  )
  const [playerColor, setPlayerColor] = useState(getRandomColor())
  const [interpolationEnabled, setInterpolationEnabled] = useState(true)
  // Add new state for bot controls
  const [botsEnabled, setBotsEnabled] = useState(false)
  const [numBots, setNumBots] = useState(30)

  const [playerId, setPlayerId] = useState(generateUUID())
  const { emitEvent } = useEtherbaseSource({ sourceAddress })

  // Update bot simulation effect to use the controls
  useEffect(() => {
    if (botsEnabled) {
      startBotSimulation(numBots, emitEvent)
    }
    // No cleanup needed if startBotSimulation doesn't provide one
    return () => {}
  }, [botsEnabled, numBots, emitEvent])

  return (
    <main className="w-full h-screen flex flex-col">
      {/* Update header controls */}
      <div className="p-2 flex gap-2 items-center bg-gray-800 text-white z-20 relative flex-wrap">
        <label>
          Name:
          <input
            className="ml-1 text-black rounded px-1"
            type="text"
            value={playerName}
            onChange={(e) => setPlayerName(e.target.value)}
          />
        </label>
        <label>
          Color:
          <input
            className="ml-1 w-16 h-8 rounded"
            type="color"
            value={playerColor}
            onChange={(e) => setPlayerColor(e.target.value)}
          />
        </label>
        <label className="flex items-center gap-1">
          <input
            type="checkbox"
            checked={interpolationEnabled}
            onChange={(e) => setInterpolationEnabled(e.target.checked)}
          />
          Motion Smoothing
        </label>

        {/* Add new bot controls */}
        <div className="border-l border-gray-600 pl-2 ml-2 flex gap-2 items-center">
          <label className="flex items-center gap-1">
            <input
              type="checkbox"
              checked={botsEnabled}
              onChange={(e) => setBotsEnabled(e.target.checked)}
            />
            Enable Bots
          </label>
          <label>
            Bot Count:
            <input
              type="number"
              min="1"
              max="500"
              value={numBots}
              onChange={(e) =>
                setNumBots(
                  Math.max(
                    1,
                    Math.min(500, Number.parseInt(e.target.value) || 1),
                  ),
                )
              }
              className="ml-1 w-16 text-black rounded px-1"
            />
          </label>
        </div>
      </div>

      <motion.div className="w-full h-full relative">
        <Canvas
          camera={{ position: [0, 2, 10], fov: 60 }}
          shadows
          className="bg-black w-full h-full absolute inset-0"
        >
          <Scene
            playerId={playerId}
            playerName={playerName}
            playerColor={playerColor}
            sourceAddress={sourceAddress}
            interpolationEnabled={interpolationEnabled}
            emitEvent={emitEvent}
          />
        </Canvas>
      </motion.div>
    </main>
  )
}
