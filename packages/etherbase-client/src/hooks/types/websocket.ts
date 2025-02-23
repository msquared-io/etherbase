import type { Address } from "viem"
import type { EtherbaseEvent } from "./events"
import type { EtherstoreState } from "./state"

export interface WebSocketStateUpdateInner {
  contractAddress: Address
  state: EtherstoreState
  transactionIndex?: number
}

export type WebSocketUpdate = {
  // the subscription id that the updates are for
  subscriptionId: string
  transactionIndex?: number
}

export type WebSocketEventUpdate = {
  type: "event"
  data: EtherbaseEvent
} & WebSocketUpdate
export type WebSocketStateUpdate = {
  type: "state"
  data: WebSocketStateUpdateInner
} & WebSocketUpdate

export type WebSocketUpdateData = {
  block: number
  updates: (WebSocketEventUpdate | WebSocketStateUpdate)[]
}
