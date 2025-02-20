"use client"

import React, { useState, useCallback, useMemo } from "react"
import RegisterEventButton from "@/components/RegisterEventButton"
import PermissionsPanel from "@/components/PermissionsPanel"
import EmitEventButton from "@/components/EmitEventButton"

// shadcn/ui components
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { AnimatedList } from "@/components/ui/animated-list"
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs" // import these from your shadcn/ui setup

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

interface ExtendedEtherbaseEvent extends EtherbaseEvent {
  block: number
  args: Record<string, any>
}

interface SourceDetailsProps {
  address: Address
  initialEventDefinitions: EventDefinition[]
}

const MAX_EVENTS = 50

// Utility function to recursively set a nested property (Firestore-like)
function setNestedValue(
  obj: any,
  path: string[],
  fieldData: { type: string; value: any },
) {
  const [head, ...rest] = path

  if (!head) return

  if (!obj[head]) {
    obj[head] = {}
  }

  // If this is the final part of the path, set the data
  if (rest.length === 0) {
    obj[head] = fieldData
  } else {
    setNestedValue(obj[head], rest, fieldData)
  }
}

// Utility function to recursively delete a nested property
function deleteNestedValue(obj: any, path: string[]) {
  const [head, ...rest] = path

  if (!obj[head]) return

  if (rest.length === 0) {
    delete obj[head]
  } else {
    deleteNestedValue(obj[head], rest)
  }
}

export default function SourceDetails({
  address,
  initialEventDefinitions,
}: SourceDetailsProps) {
  // ─────────────────────────────────────────────────────────────
  // 1. SCHEMA (existing event definitions + event data)
  // ─────────────────────────────────────────────────────────────

  const [latestEvents, setLatestEvents] = useState<
    Record<string, (ExtendedEtherbaseEvent & { index: number })[]>
  >({})

  const handleNewEvent = useCallback((event: ExtendedEtherbaseEvent) => {
    console.log("Received new event:", event)
    setLatestEvents((prev) => {
      const updated = { ...prev }

      const currentEvents = updated[event.name] || []
      // Prepend new event
      updated[event.name] = [
        ...currentEvents,
        { ...event, index: currentEvents.length + 1 },
      ]

      return updated
    })
  }, [])

  const { eventDefinitions: eventDefinitionSource, fetchEventDefinitions } =
    useEtherbaseSource({
      sourceAddress: address as Address,
    })

  const eventDefinitions = useMemo(
    () =>
      eventDefinitionSource.length > 0
        ? eventDefinitionSource
        : initialEventDefinitions,
    [eventDefinitionSource, initialEventDefinitions],
  )

  // Memoize the event names array
  const eventNames = useMemo(
    () => eventDefinitions.map((def) => def.name),
    [eventDefinitions],
  )

  const { error } = useEtherbaseEvents({
    sourceAddress: address as Address,
    events: eventDefinitions.map((def) => ({ name: def.name })),
    onEvent: handleNewEvent,
  })

  const handleEventRegistered = useCallback(async () => {
    await fetchEventDefinitions()
  }, [fetchEventDefinitions])

  // ─────────────────────────────────────────────────────────────
  // 3. PERMISSIONS (existing PermissionsPanel)
  // ─────────────────────────────────────────────────────────────
  // (Nothing changed except it now sits in a Tab)

  return (
    <div className="p-4">
      <h2 className="scroll-m-20 text-2xl font-semibold tracking-tight mb-4">
        Source: {address}
      </h2>

      <Tabs defaultValue="schema" className="space-y-4">
        <TabsList>
          <TabsTrigger value="schema">Schema</TabsTrigger>
          <TabsTrigger value="permissions">Permissions</TabsTrigger>
        </TabsList>

        {/* SCHEMA TAB */}
        <TabsContent value="schema">
          <div className="mt-2">
            <RegisterEventButton
              sourceAddress={address}
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
                      {def.args?.map(
                        (
                          arg: { name: string; argType: string },
                          idx: number,
                        ) => (
                          <div key={idx}>
                            {arg.name} ({arg.argType})
                          </div>
                        ),
                      )}
                      <div className="mt-4">
                        <EmitEventButton
                          sourceAddress={address}
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
                          {eventsForType
                            .slice(
                              Math.max(0, eventsForType.length - MAX_EVENTS),
                              eventsForType.length,
                            )
                            .map((evt) => (
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
                                  {Object.entries(evt.args).map(
                                    ([key, value]) => (
                                      <p
                                        key={key}
                                        className="text-xs text-foreground"
                                      >
                                        {key}: {String(value)}
                                      </p>
                                    ),
                                  )}
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
        </TabsContent>

        {/* PERMISSIONS TAB */}
        <TabsContent value="permissions">
          <PermissionsPanel sourceAddress={address} />
        </TabsContent>
      </Tabs>
    </div>
  )
}
