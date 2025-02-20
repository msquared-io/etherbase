"use client"

import { useEffect, useRef, useState, useCallback, useMemo } from "react"
import type React from "react"
import { Canvas, useFrame, useThree } from "@react-three/fiber"
import * as THREE from "three"
import { motion } from "framer-motion"
import { Text } from "@react-three/drei" // For displaying player names
import {
  type EtherstoreState,
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

interface Block {
  position: THREE.Vector3
  type: number // 1 or 2 for different block types
}

interface PlayerUpdateEvent {
  args: {
    playerId: string
    posX: number
    posY: number
    posZ: number
    color: string
    name: string
    timestamp: number
  }
}

const FLOAT_PRECISION = 1000 // 3 decimal places
const FORCE_UPDATE_INTERVAL_MS = 5000 // Force local update every 5s
const PLAYER_TIMEOUT_MS = 10000 // Remove players after 10s of no updates
const INTERP_BUFFER_MS = 100 // How far behind "now" we interpolate (e.g. 100-200ms)

/** Physics constants */
const GRAVITY = -9.8
const JUMP_FORCE = 5
const MOVE_SPEED = 3 // units/second

const BLOCK_COLORS = {
  1: "#8B4513", // Brown for dirt
  2: "#808080", // Gray for stone
  0: "#4CAF50", // Green for grass
}

const BLOCK_EMISSIVE = {
  1: "#3d1d08", // Dark brown glow for dirt
  2: "#202020", // Dark gray glow for stone
  0: "#1b4a1d", // Dark green glow for grass
}

const WORLD_SIZE = 10 // Reduced from 20 to improve performance
const BLOCK_SIZE = 0.5 // Size of each block, matching player size

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

/** -------------------------------
 *  VOXEL BLOCKS INSTANCED RENDERER
 *  ------------------------------- */
function VoxelBlocks({
  blocks,
  onBlockClick,
  onBlockRightClick,
  highlightedBlock,
  onMeshCreated,
}: {
  blocks: Map<string, Block>
  onBlockClick: (position: THREE.Vector3) => void
  onBlockRightClick: (position: THREE.Vector3) => void
  highlightedBlock: THREE.Vector3 | null
  onMeshCreated: (mesh: THREE.Mesh) => void
}) {
  // Create shared geometry and materials
  const blockGeometry = useMemo(
    () => new THREE.BoxGeometry(BLOCK_SIZE, BLOCK_SIZE, BLOCK_SIZE),
    [],
  )

  const blockMaterials = useMemo(() => {
    const materials: Record<number, THREE.MeshStandardMaterial> = {}
    for (const type of Object.keys(BLOCK_COLORS)) {
      const numType = Number.parseInt(type)
      materials[numType] = new THREE.MeshStandardMaterial({
        color: BLOCK_COLORS[numType as keyof typeof BLOCK_COLORS],
        emissive: BLOCK_EMISSIVE[numType as keyof typeof BLOCK_EMISSIVE],
        emissiveIntensity: 0.2,
        roughness: 0.7,
        metalness: 0.1,
      })
    }
    return materials
  }, [])

  return (
    <group>
      {Array.from(blocks.values()).map((block) => (
        <mesh
          key={`${block.position.x},${block.position.y},${block.position.z}`}
          geometry={blockGeometry}
          material={blockMaterials[block.type]}
          position={block.position}
          onPointerDown={(e) => {
            if (e.button === 2) {
              e.stopPropagation()
              onBlockRightClick(e.point)
            } else {
              e.stopPropagation()
              onBlockClick(e.point)
            }
          }}
          ref={(mesh) => {
            if (mesh) {
              onMeshCreated(mesh)
            }
          }}
        />
      ))}
    </group>
  )
}

interface EmitEventArgs {
  name: string
  args: {
    posX: number
    posY: number
    posZ: number
    color: string
    name: string
    timestamp: number
    playerId: string
  }
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
  emitEvent: (args: EmitEventArgs) => Promise<void>
}) {
  // Initial blocks that are always present
  const initialBlocks = useMemo(() => {
    const blocks = new Map<string, Block>()
    const position1 = new THREE.Vector3(2, 1, 5)
    const position2 = new THREE.Vector3(-2, 1, 5)
    const key1 = `${position1.x},${position1.y},${position1.z}`
    const key2 = `${position2.x},${position2.y},${position2.z}`
    blocks.set(key1, {
      position: position1,
      type: 2, // Dirt block
    })
    blocks.set(key2, {
      position: position2,
      type: 2, // Dirt block
    })
    return blocks
  }, [])

  // Local player states
  const [position, setPosition] = useState(new THREE.Vector3(0, 2, 5))
  const [velocity, setVelocity] = useState(new THREE.Vector3(0, 0, 0))
  const [isJumping, setIsJumping] = useState(false)

  // Refs for smooth movement
  const velocityRef = useRef(velocity)
  const positionRef = useRef(position)
  const isJumpingRef = useRef(isJumping)

  // Update refs when state changes
  useEffect(() => {
    velocityRef.current = velocity
    positionRef.current = position
    isJumpingRef.current = isJumping
  }, [velocity, position, isJumping])

  // Voxel world state
  const [blocks, setBlocks] = useState<Map<string, Block>>(
    () => new Map(initialBlocks),
  )

  // Block placement state
  const [selectedBlockType, setSelectedBlockType] = useState(1)
  const [highlightedBlock, setHighlightedBlock] =
    useState<THREE.Vector3 | null>(null)

  // Three.js hooks
  const { camera, raycaster, pointer } = useThree()

  // Remote players: each has a snapshot buffer + a display position
  const [remotePlayers, setRemotePlayers] = useState<
    Record<string, RemotePlayer>
  >({})

  const { state, loading, error, update } = useEtherstore([
    sourceAddress,
    "blocks",
  ])

  // Sync blocks from etherstore state while preserving initial blocks
  useEffect(() => {
    if (state && typeof state === "object" && state.blocks) {
      const newBlocks = new Map(initialBlocks) // Start with initial blocks

      for (const [key, blockData] of Object.entries(
        state.blocks as EtherstoreState,
      )) {
        if (!blockData || typeof blockData !== "object") continue
        if (!("position" in blockData && "type" in blockData)) continue

        const pos = blockData.position as { x: number; y: number; z: number }
        // Skip invalid positions
        if (
          !pos ||
          typeof pos.x !== "number" ||
          typeof pos.y !== "number" ||
          typeof pos.z !== "number"
        )
          continue

        // Don't override initial blocks
        if (!initialBlocks.has(key)) {
          try {
            newBlocks.set(key, {
              position: new THREE.Vector3(
                pos.x / FLOAT_PRECISION,
                pos.y / FLOAT_PRECISION,
                pos.z / FLOAT_PRECISION,
              ),
              type: Number(blockData.type),
            })
          } catch (err) {
            console.error("Error creating block:", err)
          }
        }
      }
      setBlocks(newBlocks)
    }
  }, [state, initialBlocks])

  const handleBlockDelete = useCallback(
    (position: THREE.Vector3) => {
      const key = `${position.x},${position.y},${position.z}`
      if (initialBlocks.has(key)) return // Don't modify initial blocks

      try {
        setBlocks((prev) => {
          const next = new Map(prev)
          next.delete(key)
          return next
        })

        update({
          blocks: {
            [key]: null,
          },
        })
      } catch (err) {
        console.error("Error removing block:", err)
      }
    },
    [update, initialBlocks],
  )

  // Keyboard input tracking
  const keys = useRef<Record<string, boolean>>({})
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      keys.current[e.code] = true

      // Block selection
      if (e.code === "Digit1") setSelectedBlockType(1)
      if (e.code === "Digit2") setSelectedBlockType(2)
      // Block deletion with Y key
      if (e.code === "KeyY" && highlightedBlock) {
        handleBlockDelete(highlightedBlock)
      }
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
  }, [highlightedBlock, handleBlockDelete])

  const handleBlockRightClick = useCallback(
    (position: THREE.Vector3) => {
      handleBlockDelete(position)
    },
    [handleBlockDelete],
  )

  // Keep track of block meshes for raycasting
  const blockMeshes = useRef<THREE.Mesh[]>([])
  const blockRefs = useRef<Map<string, THREE.Mesh>>(new Map())

  // Clear meshes when blocks change
  useEffect(() => {
    blockMeshes.current = []
    blockRefs.current.clear()
  }, [])

  // Block placement handlers
  const handleBlockClick = useCallback(
    (position: THREE.Vector3) => {
      if (!highlightedBlock) return

      const key = `${highlightedBlock.x},${highlightedBlock.y},${highlightedBlock.z}`
      if (initialBlocks.has(key)) return // Don't modify initial blocks

      try {
        // Round coordinates before scaling to ensure precision
        const x = Math.round(highlightedBlock.x * 2) / 2
        const y = Math.round(highlightedBlock.y * 2) / 2
        const z = Math.round(highlightedBlock.z * 2) / 2

        // Scale for storage
        const scaledX = Math.round(x * FLOAT_PRECISION)
        const scaledY = Math.round(y * FLOAT_PRECISION)
        const scaledZ = Math.round(z * FLOAT_PRECISION)

        // Create new block with rounded position
        const newPosition = new THREE.Vector3(x, y, z)
        setBlocks((prev) => {
          const next = new Map(prev)
          next.set(key, {
            position: newPosition,
            type: selectedBlockType,
          })
          return next
        })

        // Update etherstore
        update({
          blocks: {
            [key]: {
              position: { x: scaledX, y: scaledY, z: scaledZ },
              type: selectedBlockType,
            },
          },
        })
      } catch (err) {
        console.error("Error placing block:", err)
      }
    },
    [highlightedBlock, selectedBlockType, update, initialBlocks],
  )

  // Optimize raycasting by using a debounced update
  const raycastDebounceRef = useRef<number>()
  useFrame(() => {
    if (!pointer || !camera || !raycaster) return

    // Debounce raycasting to every 100ms
    if (
      raycastDebounceRef.current &&
      Date.now() - raycastDebounceRef.current < 100
    ) {
      return
    }
    raycastDebounceRef.current = Date.now()

    raycaster.setFromCamera(pointer, camera)
    const intersects = raycaster.intersectObjects(blockMeshes.current, false)

    if (intersects.length > 0) {
      const intersect = intersects[0]
      if (intersect.face) {
        const normal = intersect.face.normal.clone()
        const position = intersect.point.clone()
        position.add(normal.multiplyScalar(BLOCK_SIZE * 0.5))
        position.x = Math.round(position.x * 2) / 2
        position.y = Math.round(position.y * 2) / 2
        position.z = Math.round(position.z * 2) / 2
        setHighlightedBlock(position)
      }
    } else {
      setHighlightedBlock(null)
    }
  })

  // Optimized movement update
  useFrame((_, delta) => {
    // Clamp delta to prevent large jumps
    const clampedDelta = Math.min(delta, 0.1)

    const newVelocity = velocityRef.current.clone()
    const newPosition = positionRef.current.clone()

    // Movement vector for combined input
    const moveVector = new THREE.Vector3(0, 0, 0)
    if (keys.current["KeyW"]) moveVector.z -= 1
    if (keys.current["KeyS"]) moveVector.z += 1
    if (keys.current["KeyA"]) moveVector.x -= 1
    if (keys.current["KeyD"]) moveVector.x += 1

    // Normalize movement vector if moving diagonally
    if (moveVector.lengthSq() > 0) {
      moveVector.normalize()
      newPosition.x += moveVector.x * MOVE_SPEED * clampedDelta
      newPosition.z += moveVector.z * MOVE_SPEED * clampedDelta
    }

    // Jump handling
    if (!isJumpingRef.current && keys.current["Space"]) {
      newVelocity.y = JUMP_FORCE
      isJumpingRef.current = true
      setIsJumping(true)
    }

    // Apply gravity
    newVelocity.y += GRAVITY * clampedDelta
    newPosition.y += newVelocity.y * clampedDelta

    // Simple floor collision at y=0
    const minHeight = 0.25
    if (newPosition.y <= minHeight) {
      newPosition.y = minHeight
      newVelocity.y = 0
      if (isJumpingRef.current) {
        isJumpingRef.current = false
        setIsJumping(false)
      }
    }

    // Only update state if position actually changed
    if (!newPosition.equals(positionRef.current)) {
      positionRef.current = newPosition
      setPosition(newPosition)
    }

    if (!newVelocity.equals(velocityRef.current)) {
      velocityRef.current = newVelocity
      setVelocity(newVelocity)
    }
  })

  /**
   * Handle incoming player updates and store them in a snapshot buffer.
   */
  const handlePlayerUpdate = useCallback(
    (event: PlayerUpdateEvent) => {
      if (event.args.playerId === playerId) return

      const posX = event.args.posX / FLOAT_PRECISION
      const posY = event.args.posY / FLOAT_PRECISION
      const posZ = event.args.posZ / FLOAT_PRECISION
      const { color, name, timestamp } = event.args
      const now = Date.now()

      setRemotePlayers((prev) => {
        const id = event.args.playerId
        const newPlayer = prev[id]
          ? { ...prev[id] }
          : {
              snapshots: [],
              displayPosition: new THREE.Vector3(posX, posY, posZ),
              latency: now - timestamp,
            }

        // Add new snapshot
        newPlayer.snapshots.push({
          serverTimestamp: timestamp,
          clientReceiveTime: now,
          position: new THREE.Vector3(posX, posY, posZ),
          color,
          name,
        })

        // Keep only last 10 snapshots instead of 20
        if (newPlayer.snapshots.length > 10) {
          newPlayer.snapshots = newPlayer.snapshots.slice(-10)
        }

        newPlayer.latency = now - timestamp

        // Only update if something changed
        if (JSON.stringify(prev[id]) === JSON.stringify(newPlayer)) {
          return prev
        }

        return { ...prev, [id]: newPlayer }
      })
    },
    [playerId],
  )

  // Subscribe to player updates (useEtherbaseEvents)
  useEtherbaseEvents({
    sourceAddress,
    events: [{ name: "PlayerUpdate" }],
    onEvent: handlePlayerUpdate,
  })

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
        lastSentStateRef.current.posZ !== eventArgs.posZ

      if (hasChanged || timeSinceLastUpdate >= FORCE_UPDATE_INTERVAL_MS) {
        try {
          await emitEvent({
            name: "PlayerUpdate",
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
    }, 50)

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
      {/* Flat green floor */}
      <mesh rotation={[-Math.PI / 2, 0, 0]} position={[0, 0, 0]}>
        <planeGeometry args={[50, 50]} />
        <meshStandardMaterial color="#4CAF50" />
      </mesh>

      {/* Voxel blocks */}
      <VoxelBlocks
        blocks={blocks}
        onBlockClick={handleBlockClick}
        onBlockRightClick={handleBlockRightClick}
        highlightedBlock={highlightedBlock}
        onMeshCreated={(mesh) => {
          blockMeshes.current.push(mesh)
        }}
      />

      {/* Highlight potential placement */}
      {highlightedBlock && (
        <mesh position={highlightedBlock}>
          <boxGeometry
            args={[BLOCK_SIZE + 0.02, BLOCK_SIZE + 0.02, BLOCK_SIZE + 0.02]}
          />
          <meshStandardMaterial
            color={BLOCK_COLORS[selectedBlockType as keyof typeof BLOCK_COLORS]}
            opacity={0.5}
            transparent
            depthWrite={false}
          />
        </mesh>
      )}

      {/* Players */}
      <PlayerCube
        position={position}
        color={playerColor}
        name={playerName}
        latency={0}
        isLocal={true}
      />

      {Object.entries(remotePlayers).map(([id, rp]) => {
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

/** Add ChatBox component */
function ChatBox({
  sourceAddress,
  playerId,
  playerName,
}: {
  sourceAddress: Address
  playerId: string
  playerName: string
}) {
  const [message, setMessage] = useState("")
  const { state: chatState, update: updateChat } = useEtherstore([
    sourceAddress,
    "messages",
  ])
  const chatEndRef = useRef<HTMLDivElement>(null)

  // Optimize chat auto-scroll
  const scrollToBottom = useCallback(() => {
    if (!chatEndRef.current) return
    requestAnimationFrame(() => {
      chatEndRef.current?.scrollIntoView({ behavior: "smooth" })
    })
  }, [])

  // Only scroll when messages change
  const prevMessagesLength = useRef(0)
  useEffect(() => {
    const currentLength = Object.keys(chatState || {}).length
    if (currentLength !== prevMessagesLength.current) {
      scrollToBottom()
      prevMessagesLength.current = currentLength
    }
  }, [chatState, scrollToBottom])

  // Convert messages object to sorted array
  const messages = useMemo(() => {
    return Object.entries(chatState || {})
      .map(([timestamp, msg]) => ({
        timestamp: Number.parseInt(timestamp, 10),
        ...(msg &&
        typeof msg === "object" &&
        "playerId" in msg &&
        "message" in msg
          ? (msg as unknown as ChatMessage)
          : { playerId: "system", message: "Invalid message" }),
      }))
      .sort((a, b) => a.timestamp - b.timestamp)
      .slice(-50)
  }, [chatState])

  const handleSend = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!message.trim()) return

    const timestamp = Date.now()
    await updateChat({
      messages: {
        [`${timestamp}`]: {
          playerId,
          message: message.trim(),
        },
      },
    })
    setMessage("")
  }

  return (
    <div className="absolute top-16 left-2 w-64 bg-gray-800/80 rounded-lg text-white z-10">
      {/* Messages container */}
      <div className="h-48 overflow-y-auto p-2 space-y-1 scrollbar-thin scrollbar-thumb-gray-600">
        {messages.map((msg) => (
          <div key={msg.timestamp} className="text-sm">
            <span
              className="font-bold"
              style={{
                color: msg.playerId === playerId ? "#88ff88" : "#ffffff",
              }}
            >
              {msg.playerId === playerId ? playerName : msg.playerId}:
            </span>{" "}
            <span>{msg.message}</span>
          </div>
        ))}
        <div ref={chatEndRef} />
      </div>

      {/* Input form */}
      <form onSubmit={handleSend} className="p-2 border-t border-gray-700">
        <input
          type="text"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          className="w-full px-2 py-1 rounded bg-gray-700 text-white text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
          placeholder="Press Enter to send..."
        />
      </form>
    </div>
  )
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

        {/* Move ChatBox after Canvas to ensure it's on top */}
        <ChatBox
          sourceAddress={sourceAddress}
          playerId={playerId}
          playerName={playerName}
        />
      </motion.div>
    </main>
  )
}
