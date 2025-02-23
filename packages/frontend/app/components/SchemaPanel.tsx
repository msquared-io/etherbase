import React, { useState, useCallback, useMemo } from "react"
import RegisterEventButton from "@/components/RegisterEventButton"
import EmitEventButton from "@/components/EmitEventButton"

// shadcn/ui components
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { AnimatedList } from "@/components/ui/animated-list"

import {
  useEtherbaseEvents,
  useEtherbaseSource,
  type EtherbaseEvent,
} from "@msquared/etherbase-client"
import type { Address } from "viem"

interface EventDefinition {
  name: string
  args?: { name: string; argType: string }[]
}

interface ExtendedEvent extends EtherbaseEvent {
  index: number
}

interface SchemaPanelProps {
  sourceAddress: Address
  initialEventDefinitions: EventDefinition[]
}

const MAX_EVENTS = 50

export default function SchemaPanel({
  sourceAddress,
  initialEventDefinitions,
}: SchemaPanelProps) {
  const [latestEvents, setLatestEvents] = useState<
    Record<string, ExtendedEvent[]>
  >({})

  const handleNewEvent = (event: EtherbaseEvent) => {
    setLatestEvents((prev) => {
      const updated = { ...prev }
      const currentEvents = updated[event.name] || []
      const extendedEvent: ExtendedEvent = {
        ...event,
        index: currentEvents.length + 1,
      }
      updated[event.name] = [...currentEvents, extendedEvent]
      return updated
    })
  }

  const { eventDefinitions: eventDefinitionSource, fetchEventDefinitions } =
    useEtherbaseSource({
      sourceAddress: sourceAddress,
    })

  const eventDefinitions = useMemo(
    () =>
      eventDefinitionSource.length > 0
        ? eventDefinitionSource
        : initialEventDefinitions,
    [eventDefinitionSource, initialEventDefinitions],
  )

  // Subscribe to events
  useEtherbaseEvents({
    contractAddress: sourceAddress,
    events: eventDefinitions.map((def) => ({ name: def.name })),
    onEvent: handleNewEvent,
  })

  const handleEventRegistered = useCallback(async () => {
    await fetchEventDefinitions()
  }, [fetchEventDefinitions])

  return (
    <div>
      <div className="mt-2">
        <RegisterEventButton
          sourceAddress={sourceAddress}
          onEventRegistered={handleEventRegistered}
        />
      </div>

      <Separator className="my-4" />

      <div className="space-y-4">
        {eventDefinitions.map((def) => {
          const eventsForType = latestEvents[def.name] || []
          return (
            <Card key={def.name}>
              <CardHeader>
                <CardTitle className="text-lg font-semibold -mb-2">
                  {def.name}
                </CardTitle>
              </CardHeader>
              <CardContent className="flex flex-row gap-5">
                <div className="flex flex-col w-1/6">
                  <h6 className="text-muted-foreground font-semibold">
                    Parameters
                  </h6>
                  {def.args?.map((arg, argIndex) => (
                    <div key={`${def.name}-${arg.name}-${argIndex}`}>
                      {arg.name} ({arg.argType})
                    </div>
                  ))}
                  <div className="mt-4">
                    <EmitEventButton
                      sourceAddress={sourceAddress}
                      eventName={def.name}
                      args={def.args || []}
                    />
                  </div>
                </div>
                <div>
                  <h6 className="text-muted-foreground font-semibold">
                    Events
                  </h6>
                  <div className="flex gap-3 overflow-hidden">
                    <AnimatedList delay={0.2}>
                      {eventsForType.map((evt) => (
                        <Card
                          className="min-w-[350px] max-w-[350px] border border-accent/50 bg-background shadow-sm mt-2"
                          key={evt.index}
                        >
                          <CardHeader>
                            <CardTitle className="text-s font-medium">
                              {`#${evt.index}`}
                            </CardTitle>
                          </CardHeader>
                          <CardContent className="space-y-0.5">
                            <p className="text-xs text-muted-foreground">
                              Block: {evt.block}
                            </p>
                            {Object.entries(evt.args).map(([key, value]) => (
                              <p
                                key={`${evt.index}-${key}`}
                                className="text-xs text-foreground"
                              >
                                {key}: {String(value)}
                              </p>
                            ))}
                          </CardContent>
                        </Card>
                      ))}
                    </AnimatedList>
                    {eventsForType.length === 0 && (
                      <p className="text-sm text-muted-foreground">
                        No events received yet
                      </p>
                    )}
                  </div>
                </div>
              </CardContent>
            </Card>
          )
        })}
      </div>
    </div>
  )
}
