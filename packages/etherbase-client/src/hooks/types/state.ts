import type { Address } from "viem"
import type { Event } from "./events"

export type EtherstoreValue = string | number | boolean | null

export type StateSubscriptionOptions = {
  // options for repolling the data. if the data has changed, this update will be received by the client.
  repoll?: {
    // if true, the server will repoll the data for any event on the contract.
    onAnyContractEvent?: boolean
    // options for listening to precise events on the contract. if the event condition is met, the data will be repolled by the server.
    listenEvents?: Event[]
  }
  // if true, the server will combine state updates from the same contract into a single update for each subscription.
  combineStateUpdatesFromSameContract?: boolean
  // if true, and combineStateUpdatesFromSameContract is false, the server will combine state updates from the same transaction into a single update. only applicable to EthDBPathUpdate events.
  combineStateUpdatesFromSameTransaction?: boolean
  // if true, the server will send a full state update of listened-to state when the state changes. default is false, where it only sends the delta update.
  fullStateUpdateOnChange?: boolean
}

export type EtherstoreKey = string | Address
export type EtherstoreCompoundPath = (EtherstoreKey | EtherstoreKey[])[]
export type EtherstorePath = EtherstoreKey[]

export type EtherstoreState = {
  [key: string]: EtherstoreValue | EtherstoreState
}

export type EtherstoreHookReturn = {
  state: EtherstoreState
  loading: boolean
  error: string | null
  update: (state: EtherstoreState) => Promise<void>
}

export type StateUpdate = Record<string, unknown>

export type StateUpdates = StateUpdate[]

// if variableName and mappingKey are undefined, they're considered to be subscribed to all state
export interface StateFilter {
  contractAddress: Address
  variableName?: string
  mappingKey?: string
  // should we receive updates on newly declared variables since this client was subscribed?
  receiveNewState?: boolean
}

export type EtherstoreUpdate = {
  path: EtherstorePath
  value: EtherstoreValue
  type: "STRING" | "BOOL" | "UINT256" | "INT256" | "BYTES"
}
