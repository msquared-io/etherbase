import type { Address } from "viem"

export type EtherbaseEvent = {
  contractAddress: Address
  block: number
  name: string
  eventId: string
  args: Record<string, unknown>
}

export type EventHandler = (event: EtherbaseEvent) => void

export type Event = {
  // the name of the event to listen to.
  name?: string
  // the id of the event to listen to.
  id?: string
  // the arguments of the event to specifically listen to. if args are not provided, they will be fully subscribed to.
  args?: Record<string, (string | number | boolean)[]>
}

export type UseEtherbaseEventsProps = {
  sourceAddress?: Address
  events?: Event[]
  onEvent: EventHandler
}

export interface UseEtherbaseEventsReturn {
  error: string | null
}

export type Argument = {
  name: string
  argType: string
  isIndexed: boolean
}
