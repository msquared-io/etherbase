import React, { useState, useCallback, useMemo } from "react"
import RegisterEventButton from "@/components/RegisterEventButton"
import EmitEventButton from "@/components/EmitEventButton"

// shadcn/ui components
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Separator } from "@/components/ui/separator"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { ActivityIcon } from "lucide-react"

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
  timestamp: number
}

interface SchemaPanelProps {
  sourceAddress: Address
  initialEventDefinitions: EventDefinition[]
}

const MAX_EVENTS = 1000

export default function SchemaPanel({
  sourceAddress,
  initialEventDefinitions,
}: SchemaPanelProps) {
  const [latestEvents, setLatestEvents] = useState<
    Record<string, ExtendedEvent[]>
  >({})
  const [searchTerm, setSearchTerm] = useState("")
  const [activeTab, setActiveTab] = useState<string>("all")

  const handleNewEvent = useCallback((event: EtherbaseEvent) => {
    setLatestEvents((prev) => {
      const updated = { ...prev }
      const currentEvents = updated[event.name] || []

      // Create extended event with index and timestamp
      const extendedEvent: ExtendedEvent = {
        ...event,
        index: currentEvents.length + 1,
        timestamp: Date.now(),
      }

      // Add new event at the beginning and limit to MAX_EVENTS
      const newEvents = [extendedEvent, ...currentEvents]
      if (newEvents.length > MAX_EVENTS) {
        newEvents.pop() // Remove oldest event (now at the end)
      }

      updated[event.name] = newEvents
      return updated
    })
  }, [])

  const { eventDefinitions: eventDefinitionSource } = useEtherbaseSource({
    sourceAddress: sourceAddress,
  })

  const eventDefinitions = useMemo(
    () =>
      eventDefinitionSource.length > 0
        ? eventDefinitionSource
        : initialEventDefinitions,
    [eventDefinitionSource, initialEventDefinitions],
  )

  const events = useMemo(() => {
    return eventDefinitions.map((def) => ({ name: def.name }))
  }, [eventDefinitions])

  // Subscribe to events
  useEtherbaseEvents({
    contractAddress: sourceAddress,
    events: events,
    onEvent: handleNewEvent,
  })

  // No need for handleEventRegistered callback since we're polling automatically
  // Just provide an empty function to satisfy the component props
  const handleEventRegistered = useCallback(() => {
    // The polling mechanism will automatically fetch new event definitions
  }, [])

  // Format timestamp to readable time
  const formatTime = (timestamp: number) => {
    return new Date(timestamp).toLocaleTimeString([], {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    })
  }

  // Get all events across all event types
  const allEvents = useMemo(() => {
    const events: ExtendedEvent[] = []
    for (const eventList of Object.values(latestEvents)) {
      events.push(...eventList)
    }
    // Sort by timestamp (newest first)
    return events.sort((a, b) => b.timestamp - a.timestamp)
  }, [latestEvents])

  return (
    <div className="mt-8 space-y-6">
      <div className="flex items-center gap-2 mb-4">
        <ActivityIcon className="h-6 w-6 text-muted-foreground" />
        <h2 className="text-2xl font-semibold">Event Management</h2>
      </div>

      <Card>
        <CardHeader>
          <div className="flex justify-between items-center">
            <CardTitle className="text-lg font-semibold">
              Register New Event
            </CardTitle>
          </div>
          <CardDescription>
            Define and register new event types for your contract
          </CardDescription>
        </CardHeader>
        <CardContent>
          <RegisterEventButton
            sourceAddress={sourceAddress}
            onEventRegistered={handleEventRegistered}
          />
        </CardContent>
      </Card>

      <Separator className="my-4" />

      <Tabs defaultValue="all" onValueChange={setActiveTab} value={activeTab}>
        <div className="flex justify-between items-center mb-4">
          <TabsList>
            <TabsTrigger value="all">All Events</TabsTrigger>
            {eventDefinitions.map((def) => (
              <TabsTrigger key={def.name} value={def.name}>
                {def.name}
              </TabsTrigger>
            ))}
          </TabsList>
          <div className="flex items-center gap-2">
            <Input
              placeholder="Filter events..."
              className="w-48 h-8"
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
            />
          </div>
        </div>

        <TabsContent value="all">
          <Card>
            <CardHeader>
              <div className="flex justify-between items-center">
                <CardTitle className="text-lg font-semibold">
                  All Events
                  <Badge variant="outline" className="ml-2">
                    {allEvents.length} events
                  </Badge>
                </CardTitle>
              </div>
              <CardDescription>
                View all events emitted across all event types
              </CardDescription>
            </CardHeader>
            <CardContent>
              {allEvents.length > 0 ? (
                <ScrollArea className="h-[400px] border rounded-md">
                  <div className="p-2">
                    <table className="w-full">
                      <thead className="sticky top-0 bg-background z-10">
                        <tr className="border-b">
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Event
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Time
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Block
                          </th>
                          <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                            Arguments
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        {allEvents
                          .filter(
                            (evt) =>
                              !searchTerm ||
                              JSON.stringify(evt)
                                .toLowerCase()
                                .includes(searchTerm.toLowerCase()),
                          )
                          .map((evt) => (
                            <tr
                              key={`${evt.name}-${evt.index}`}
                              className="border-b border-accent/10 hover:bg-accent/5"
                            >
                              <td className="p-2 text-xs">
                                <Badge
                                  variant="secondary"
                                  className="font-mono"
                                >
                                  {evt.name}
                                </Badge>
                              </td>
                              <td className="p-2 text-xs">
                                {formatTime(evt.timestamp)}
                              </td>
                              <td className="p-2 text-xs">{evt.block}</td>
                              <td className="p-2 text-xs">
                                <div className="flex flex-col gap-1">
                                  {Object.entries(evt.args).map(
                                    ([key, value]) => (
                                      <div
                                        key={`${evt.name}-${evt.index}-${key}`}
                                      >
                                        <span className="font-medium">
                                          {key}:
                                        </span>{" "}
                                        <span className="font-mono">
                                          {String(value)}
                                        </span>
                                      </div>
                                    ),
                                  )}
                                </div>
                              </td>
                            </tr>
                          ))}
                      </tbody>
                    </table>
                  </div>
                </ScrollArea>
              ) : (
                <div className="flex items-center justify-center h-[400px] border rounded-md bg-accent/5">
                  <p className="text-sm text-muted-foreground">
                    No events received yet
                  </p>
                </div>
              )}
            </CardContent>
          </Card>
        </TabsContent>

        {eventDefinitions.map((def) => {
          const eventsForType = latestEvents[def.name] || []
          const filteredEvents = searchTerm
            ? eventsForType.filter((evt) =>
                JSON.stringify(evt)
                  .toLowerCase()
                  .includes(searchTerm.toLowerCase()),
              )
            : eventsForType

          return (
            <TabsContent key={def.name} value={def.name}>
              <Card>
                <CardHeader>
                  <div className="flex justify-between items-center">
                    <CardTitle className="text-lg font-semibold">
                      {def.name}
                      <Badge variant="outline" className="ml-2">
                        {eventsForType.length} events
                      </Badge>
                    </CardTitle>
                    <div className="flex gap-2">
                      <EmitEventButton
                        sourceAddress={sourceAddress}
                        eventName={def.name}
                        args={def.args || []}
                      />
                    </div>
                  </div>
                  <CardDescription>
                    View and emit events of type {def.name}
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <div className="grid grid-cols-1 md:grid-cols-5 gap-5">
                    <div className="md:col-span-1">
                      <h6 className="text-sm font-medium text-muted-foreground mb-2">
                        Parameters
                      </h6>
                      <div className="border rounded-md p-3 bg-accent/5">
                        {def.args && def.args.length > 0 ? (
                          def.args.map((arg, argIndex) => (
                            <div
                              key={`${def.name}-${arg.name}-${argIndex}`}
                              className="text-sm py-1 border-b last:border-0 border-accent/10"
                            >
                              <span className="font-medium">{arg.name}</span>{" "}
                              <span className="text-xs text-muted-foreground">
                                ({arg.argType})
                              </span>
                            </div>
                          ))
                        ) : (
                          <div className="text-sm text-muted-foreground">
                            No parameters
                          </div>
                        )}
                      </div>
                    </div>
                    <div className="md:col-span-4">
                      <h6 className="text-sm font-medium text-muted-foreground mb-2">
                        Events
                      </h6>

                      {filteredEvents.length > 0 ? (
                        <ScrollArea className="h-[400px] border rounded-md">
                          <div className="p-2">
                            <table className="w-full">
                              <thead className="sticky top-0 bg-background z-10">
                                <tr className="border-b">
                                  <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                    #
                                  </th>
                                  <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                    Time
                                  </th>
                                  <th className="text-left p-2 text-xs font-medium text-muted-foreground">
                                    Block
                                  </th>
                                  {def.args?.map((arg) => (
                                    <th
                                      key={arg.name}
                                      className="text-left p-2 text-xs font-medium text-muted-foreground"
                                    >
                                      {arg.name}
                                    </th>
                                  ))}
                                </tr>
                              </thead>
                              <tbody>
                                {filteredEvents.map((evt) => (
                                  <tr
                                    key={evt.index}
                                    className="border-b border-accent/10 hover:bg-accent/5"
                                  >
                                    <td className="p-2 text-xs">{evt.index}</td>
                                    <td className="p-2 text-xs">
                                      {formatTime(evt.timestamp)}
                                    </td>
                                    <td className="p-2 text-xs">{evt.block}</td>
                                    {def.args?.map((arg) => (
                                      <td
                                        key={`${evt.index}-${arg.name}`}
                                        className="p-2 text-xs font-mono"
                                      >
                                        {String(evt.args[arg.name] ?? "")}
                                      </td>
                                    ))}
                                  </tr>
                                ))}
                              </tbody>
                            </table>
                          </div>
                        </ScrollArea>
                      ) : (
                        <div className="flex items-center justify-center h-[400px] border rounded-md bg-accent/5">
                          <p className="text-sm text-muted-foreground">
                            {searchTerm
                              ? "No matching events found"
                              : "No events received yet"}
                          </p>
                        </div>
                      )}
                    </div>
                  </div>
                </CardContent>
              </Card>
            </TabsContent>
          )
        })}
      </Tabs>
    </div>
  )
}
