/**
 * Bot simulation for the 3D Cube Game.
 *
 * Each bot moves horizontally within a fixed square area (bounds across x and z).
 * They pick a random movement direction and, when they hit a boundary,
 * they choose a new random direction.
 * Every 10 seconds, each bot randomly determines whether to jump (30% chance).
 */
import * as THREE from "three"

// --- CONSTANTS ---
// Match exactly with player constants from page.tsx
const BOUNDS = 15
const MOVE_SPEED = 3 // units/second (same as player)
const GRAVITY = -9.8
const JUMP_FORCE = 5
const FLOOR_HEIGHT = 0.25
const FLOAT_PRECISION = 1000

// Network update rate to match player update rate
const UPDATE_RATE_MS = 50 // Send updates every 100ms like players do

// Track server time offset
let serverTimeOffset = 0
const INITIAL_SERVER_TIME = Date.now()

/**
 * Gets the current estimated server time
 */
function getServerTime() {
  // Initially use local time, then adjust based on server responses
  return Date.now() + serverTimeOffset
}

// --- BOT CLASS ---
class Bot {
  id: string
  position: THREE.Vector3
  velocity: THREE.Vector3
  direction: THREE.Vector3 // horizontal movement direction (unit vector)
  isJumping: boolean
  color: string
  name: string
  jumpInterval: number // Add new property for per-bot jump timing

  constructor(id: string, index: number, startPos?: THREE.Vector3) {
    this.id = id
    // Random starting position within bounds (on the floor)
    this.position =
      startPos ??
      new THREE.Vector3(
        (Math.random() * 2 - 1) * BOUNDS,
        FLOOR_HEIGHT,
        (Math.random() * 2 - 1) * BOUNDS,
      )
    this.velocity = new THREE.Vector3(0, 0, 0)
    this.direction = this.randomDirection()
    this.isJumping = false
    this.color = this.getRandomColor()
    this.name = `Bot ${index}`

    // Set a random jump interval between 3-7 seconds for each bot
    this.jumpInterval = 3000 + Math.random() * 4000

    // Start jump check interval with the bot-specific timing
    setInterval(() => this.checkJump(), this.jumpInterval)
  }

  /**
   * Returns a random horizontal (x,z) unit vector.
   */
  randomDirection(): THREE.Vector3 {
    const angle = Math.random() * Math.PI * 2
    return new THREE.Vector3(Math.cos(angle), 0, Math.sin(angle)).normalize()
  }

  /**
   * Check whether the bot is out of bounds.
   * If so, choose a new random direction that points back into bounds.
   */
  checkBoundariesAndUpdateDirection() {
    let needsNewDirection = false

    // Check if we're out of bounds
    if (
      Math.abs(this.position.x) > BOUNDS ||
      Math.abs(this.position.z) > BOUNDS
    ) {
      needsNewDirection = true

      // Clamp position to bounds
      this.position.x = THREE.MathUtils.clamp(this.position.x, -BOUNDS, BOUNDS)
      this.position.z = THREE.MathUtils.clamp(this.position.z, -BOUNDS, BOUNDS)

      // Choose a direction pointing toward center with some randomness
      const toCenter = new THREE.Vector3(
        -this.position.x,
        0,
        -this.position.z,
      ).normalize()
      const randomAngle = ((Math.random() - 0.5) * Math.PI) / 4 // Reduced to ±45 degrees
      this.direction
        .copy(toCenter)
        .applyAxisAngle(new THREE.Vector3(0, 1, 0), randomAngle)
        .normalize()
    }

    // Reduce random direction changes to feel more natural
    if (!needsNewDirection && Math.random() < 0.01) {
      // 1% chance per update
      this.direction = this.randomDirection()
    }
  }

  /**
   * Randomly determine whether to perform a jump.
   * Now 40% chance every 3-7 seconds if not jumping.
   */
  checkJump() {
    if (!this.isJumping && Math.random() < 0.4) {
      // Set jump state and velocity
      this.isJumping = true
      this.velocity.y = JUMP_FORCE

      // Slightly lift the position off the ground to match physics
      if (this.position.y <= FLOOR_HEIGHT) {
        this.position.y = FLOOR_HEIGHT + 0.1
      }
    }
  }

  /**
   * Update the bot's physics and position.
   * @param delta Time elapsed (in seconds) since the last update.
   */
  update(delta: number) {
    // Calculate movement vector based on direction
    const moveVector = new THREE.Vector3()

    // Similar to player WASD movement, but for bots
    moveVector.copy(this.direction).multiplyScalar(MOVE_SPEED * delta)
    this.position.add(moveVector)

    // Apply gravity and update position
    this.velocity.y += GRAVITY * delta
    this.position.y += this.velocity.y * delta

    // Floor collision
    if (this.position.y <= FLOOR_HEIGHT) {
      this.position.y = FLOOR_HEIGHT
      this.velocity.y = 0
      this.isJumping = false
    }

    this.checkBoundariesAndUpdateDirection()
  }

  /**
   * Returns a random color string in hsl notation.
   */
  getRandomColor(): string {
    const hue = Math.random() * 360
    const saturation = 70 + Math.random() * 30 // 70-100%
    const lightness = 45 + Math.random() * 10 // 45-55%
    return `hsl(${hue}, ${saturation}%, ${lightness}%)`
  }
}

// --- SIMULATION LOOP FOR THE BOTS ---
let bots: Bot[] = []
let lastUpdateTime = Date.now()
let lastEmitTime = Date.now()

// Remove the emitEventt wrapper function since it's causing duplicate events
let emitEvent: any

function simulationLoop(timestamp?: number) {
  const now = Date.now()
  const delta = Math.min((now - lastUpdateTime) / 1000, 0.1)
  lastUpdateTime = now

  // Update bot positions
  bots.forEach((bot) => {
    bot.update(delta)
  })

  // Emit position updates at the same rate as players
  if (now - lastEmitTime >= UPDATE_RATE_MS) {
    const serverTimestamp = getServerTime()

    // Create a single batch of updates
    const updates = bots.map((bot) => ({
      name: "PlayerUpdate",
      args: {
        posX: Math.round(bot.position.x * FLOAT_PRECISION),
        posY: Math.round(bot.position.y * FLOAT_PRECISION),
        posZ: Math.round(bot.position.z * FLOAT_PRECISION),
        isJumping: bot.isJumping,
        color: bot.color,
        name: bot.name,
        timestamp: serverTimestamp,
        playerId: bot.id,
      },
    }))

    // Emit updates sequentially
    Promise.all(
      updates.map((update) =>
        emitEvent(update).catch((error: Error) =>
          console.error(
            `Error emitting event for ${update.args.playerId}:`,
            error,
          ),
        ),
      ),
    )

    lastEmitTime = now
  }

  if (typeof requestAnimationFrame !== "undefined") {
    requestAnimationFrame(simulationLoop)
  } else {
    setTimeout(() => simulationLoop(), 16)
  }
}

function generateUUID() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function (c) {
    const r = (Math.random() * 16) | 0
    const v = c === "x" ? r : (r & 0x3) | 0x8
    return v.toString(16)
  })
}

/**
 * Creates and starts bot simulation.
 * @param numberOfBots The number of bots to simulate.
 * @param inemitEvent The event emitter function
 */
function startBotSimulation(numberOfBots: number, inemitEvent: any) {
  // Clear any existing bots and reset the emitter
  bots = []
  emitEvent = inemitEvent
  serverTimeOffset = 0

  // Create bots
  for (let i = 0; i < numberOfBots; i++) {
    const bot = new Bot(generateUUID(), i)
    bots.push(bot)
  }

  console.log(`Started simulation with ${numberOfBots} bots.`)
  simulationLoop()
}

// Only export what's needed
export { startBotSimulation, Bot }
