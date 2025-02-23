import type { Address } from "viem"
import { getConfig } from "../config"
import type { EtherbaseEvent, Event } from "./types/events"
import type { EtherstoreState, StateSubscriptionOptions } from "./types/state"
import type {
  WebSocketStateUpdateInner,
  WebSocketUpdateData,
} from "./types/websocket"

type EventSubscription = {
  contractAddress?: Address
  events?: Event[]
  onEvent: (event: EtherbaseEvent) => void
}

type StateSubscription = {
  contractAddress: Address
  statePath: (string | string[])[]
  options: StateSubscriptionOptions
  onStateUpdate: (update: WebSocketStateUpdateInner) => void
}

type Subscription = EventSubscription | StateSubscription

interface SubscriptionStatus {
  /** The local client-chosen identifier (e.g. from a React hook). */
  localId: string
  /** Assigned by the server upon success. Used to route updates back. */
  subscriptionId?: string
  /** The user's subscription object. */
  subscription: EventSubscription | StateSubscription
  /** Set if the server responded with subscription_failed. */
  error?: string
}

export class WebSocketManager {
  private static instance: WebSocketManager
  public static get(): WebSocketManager {
    if (!WebSocketManager.instance) {
      WebSocketManager.instance = new WebSocketManager()
    }
    return WebSocketManager.instance
  }

  private ws: WebSocket | null = null
  private dead = false

  /** A map from localId -> SubscriptionStatus. */
  private subscriptions: Map<string, SubscriptionStatus> = new Map()

  private readonly privateKey?: string

  private writerWs: WebSocket | null = null
  private readerWs: WebSocket | null = null
  private readonly readerUrl: string | undefined
  private readonly writerUrl: string | undefined

  private constructor() {
    const config = getConfig()
    this.readerUrl = config.wsReaderUrl
    this.writerUrl = config.wsWriterUrl
    this.privateKey = config.privateKey
  }

  /** Connect if not open or connecting */
  private async connectReader() {
    if (this.dead) return
    if (!this.readerUrl) {
      console.error("[WebSocketManager] Reader URL not set")
      return
    }

    if (
      this.readerWs &&
      (this.readerWs.readyState === WebSocket.OPEN ||
        this.readerWs.readyState === WebSocket.CONNECTING)
    ) {
      if (this.readerWs.readyState === WebSocket.CONNECTING) {
        await new Promise<void>((resolve) => {
          if (!this.readerWs) return
          this.readerWs.addEventListener("open", () => resolve(), {
            once: true,
          })
        })
      }
      return
    }
    console.log("[WebSocketManager] Connecting reader...")

    const url = new URL(`${this.readerUrl}/read`)
    if (this.privateKey) {
      url.searchParams.set("privateKey", this.privateKey)
    }

    this.readerWs = new WebSocket(url.toString())
    this.readerWs.onopen = () => {
      console.log("[WebSocketManager] Reader connected.")
      // Re-subscribe all on reconnect
      for (const status of this.subscriptions.values()) {
        status.subscriptionId = undefined
        status.error = undefined
        this.sendSubscribe(status)
      }
    }

    this.readerWs.onmessage = (event) => this.handleMessage(event)

    this.readerWs.onerror = (err) => {
      console.error("[WebSocketManager] Reader WebSocket error:", err)
    }

    this.readerWs.onclose = () => {
      console.warn("[WebSocketManager] Reader WebSocket closed.")
      this.readerWs = null
    }

    await new Promise<void>((resolve) => {
      if (this.readerWs?.readyState === WebSocket.OPEN) {
        resolve()
      } else {
        this.readerWs?.addEventListener("open", () => resolve(), { once: true })
      }
    })
    console.log("[WebSocketManager] Reader connected.")
  }

  private async connectWriter() {
    if (this.dead) return
    if (!this.writerUrl) {
      console.error("[WebSocketManager] Writer URL not set")
      return
    }

    if (
      this.writerWs &&
      (this.writerWs.readyState === WebSocket.OPEN ||
        this.writerWs.readyState === WebSocket.CONNECTING)
    ) {
      if (this.writerWs.readyState === WebSocket.CONNECTING) {
        await new Promise<void>((resolve) => {
          if (!this.writerWs) return
          this.writerWs.addEventListener("open", () => resolve(), {
            once: true,
          })
        })
      }
      return
    }
    console.log("[WebSocketManager] Connecting writer...")

    const url = new URL(`${this.writerUrl}/write`)
    if (this.privateKey) {
      url.searchParams.set("privateKey", this.privateKey)
    }

    this.writerWs = new WebSocket(url.toString())

    this.writerWs.onopen = () => {
      console.log("[WebSocketManager] Writer connected.")
    }

    this.writerWs.onmessage = (event) => this.handleWriterMessage(event)

    this.writerWs.onerror = (err) => {
      console.error("[WebSocketManager] Writer WebSocket error:", err)
    }

    this.writerWs.onclose = () => {
      console.warn("[WebSocketManager] Writer WebSocket closed.")
      this.writerWs = null
    }

    await new Promise<void>((resolve) => {
      if (this.writerWs?.readyState === WebSocket.OPEN) {
        resolve()
      } else {
        this.writerWs?.addEventListener("open", () => resolve(), { once: true })
      }
    })
    console.log("[WebSocketManager] Writer connected.")
  }

  private handleMessage(event: MessageEvent<string>) {
    try {
      const msg = JSON.parse(event.data)

      switch (msg.type) {
        case "error":
          console.error("[WS] Server error:", msg.error)
          break

        case "subscription_success": {
          const { pendingId, subscriptionId } = msg.data
          const sub = this.subscriptions.get(pendingId)
          if (sub) {
            sub.subscriptionId = subscriptionId
            sub.error = undefined
            console.log(
              `[WS] Subscription success for localId=${pendingId}, subscriptionId=${subscriptionId}`,
            )
          }
          break
        }

        case "subscription_failed": {
          const { pendingId, error } = msg.data
          const sub = this.subscriptions.get(pendingId)
          if (sub) {
            sub.error = error
            console.warn(
              `[WS] Subscription failed for localId=${pendingId}: ${error}`,
            )
          }
          break
        }

        case "updates": {
          console.log("[WS] Got updates:", msg.data)
          this.handleUpdates(msg.data)
          break
        }

        case "initial_state": {
          this.handleUpdates(msg.data)
          break
        }

        case "transaction_complete":
          console.log("[WS] transaction_complete:", msg.data)
          break

        case "transaction_failed":
          console.warn("[WS] transaction_failed:", msg.data)
          break

        case "event_submitted":
          console.log("[WS] event_submitted")
          break

        default:
          console.log("[WS] unhandled message:", msg)
          break
      }
    } catch (err) {
      console.error("[WS] Failed to parse message:", err)
    }
  }

  private handleWriterMessage(event: MessageEvent<string>) {
    try {
      const msg = JSON.parse(event.data)

      switch (msg.type) {
        case "error":
          console.error("[WS Writer] Server error:", msg.error)
          break

        case "transaction_complete":
          console.log("[WS Writer] transaction_complete:", msg.data)
          break

        case "transaction_failed":
          console.warn("[WS Writer] transaction_failed:", msg.data)
          break

        case "event_submitted":
          console.log("[WS Writer] event_submitted")
          break

        default:
          console.log("[WS Writer] unhandled message:", msg)
          break
      }
    } catch (err) {
      console.error("[WS Writer] Failed to parse message:", err)
    }
  }

  private handleUpdates({ block, updates }: WebSocketUpdateData) {
    for (const update of updates) {
      // The server includes a subscriptionId. Find a local subscription that matches.
      const sub = [...this.subscriptions.values()].find(
        (s) => s.subscriptionId === update.subscriptionId,
      )
      if (!sub) {
        console.warn(
          "[WS] Got update for unknown subscriptionId:",
          update.subscriptionId,
        )
        continue
      }

      if (update.type === "event" && "onEvent" in sub.subscription) {
        sub.subscription.onEvent(update.data)
      } else if (
        update.type === "state" &&
        "onStateUpdate" in sub.subscription
      ) {
        sub.subscription.onStateUpdate(update.data)
      }
    }
  }

  public async addEventSubscription(
    localId: string,
    subscription: EventSubscription,
  ) {
    this.subscriptions.set(localId, {
      localId,
      subscription,
    })
    await this.connectReader()
    const status = this.subscriptions.get(localId)
    if (!status) {
      console.warn(
        "[WS] Tried to add event subscription for unknown localId:",
        localId,
      )
      return
    }
    this.sendSubscribe(status)
  }

  public async addStateSubscription(
    localId: string,
    subscription: StateSubscription,
  ) {
    this.subscriptions.set(localId, {
      localId,
      subscription,
    })
    console.log("Adding state subscription", localId, subscription)
    await this.connectReader()
    const status = this.subscriptions.get(localId)
    if (!status) {
      console.warn(
        "[WS] Tried to add state subscription for unknown localId:",
        localId,
      )
      return
    }
    this.sendSubscribe(status)
  }

  public removeSubscription(localId: string) {
    const status = this.subscriptions.get(localId)
    if (!status) return
    // If we already have a server subscriptionId, unsubscribe from the server
    if (status.subscriptionId && this.readerWs?.readyState === WebSocket.OPEN) {
      this.readerWs.send(
        JSON.stringify({
          type: "unsubscribe",
          data: {
            subscriptionId: status.subscriptionId,
          },
        }),
      )
    }
    this.subscriptions.delete(localId)
  }

  private sendSubscribe(status: SubscriptionStatus) {
    if (!this.readerWs || this.readerWs.readyState !== WebSocket.OPEN) return
    const { localId, subscription } = status

    if ("onEvent" in subscription) {
      this.readerWs.send(
        JSON.stringify({
          type: "subscribe",
          data: {
            pendingId: localId,
            eventSubscription: {
              contractAddress: subscription.contractAddress,
              events: subscription.events,
            },
          },
        }),
      )
    } else {
      console.log("Sending state subscription!!!", localId, subscription)
      this.readerWs.send(
        JSON.stringify({
          type: "subscribe",
          data: {
            pendingId: localId,
            stateSubscription: {
              contractAddress: subscription.contractAddress,
              statePath: subscription.statePath,
              options: subscription.options,
            },
          },
        }),
      )
    }
  }

  public getSubscriptionError(localId: string): string | undefined {
    return this.subscriptions.get(localId)?.error
  }

  public forceClose() {
    console.log("[WebSocketManager] Force closing...")
    this.dead = true
    if (this.readerWs) {
      this.readerWs.close()
      this.readerWs = null
    }
    if (this.writerWs) {
      this.writerWs.close()
      this.writerWs = null
    }
  }

  async setValue({
    contractAddress,
    state,
  }: { contractAddress: Address; state: EtherstoreState }) {
    if (!this.writerWs || this.writerWs.readyState !== WebSocket.OPEN) {
      await this.connectWriter()
    }

    this.writerWs?.send(
      JSON.stringify({
        type: "set_value",
        data: {
          contractAddress,
          state,
        },
      }),
    )
  }

  async emitEvent({
    sourceAddress,
    name,
    args,
  }: {
    sourceAddress: Address
    name: string
    args: Record<string, unknown>
  }) {
    console.log(
      "[WebSocketManager] Emitting event...",
      this.writerWs?.readyState,
    )
    if (!this.writerWs || this.writerWs.readyState !== WebSocket.OPEN) {
      await this.connectWriter()
    }

    this.writerWs?.send(
      JSON.stringify({
        type: "emit_event",
        data: {
          contractAddress: sourceAddress,
          name,
          args,
        },
      }),
    )
  }

  async executeContractMethod({
    contractAddress,
    methodName,
    args,
  }: {
    contractAddress: Address
    methodName: string
    args: Record<string, unknown>
  }) {
    console.log(
      "[WebSocketManager] Executing contract method...",
      this.writerWs?.readyState,
    )

    if (!this.writerWs || this.writerWs.readyState !== WebSocket.OPEN) {
      await this.connectWriter()
    }

    this.writerWs?.send(
      JSON.stringify({
        type: "execute_contract_method",
        data: { contractAddress, methodName, args },
      }),
    )
  }
}
