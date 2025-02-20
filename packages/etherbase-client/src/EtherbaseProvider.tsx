"use client"

import React, { createContext, useContext, useEffect, useState } from "react"
import type { ReactNode } from "react"
import { initializeApp } from "./config"
import type { EtherbaseConfig } from "./config"

const EtherbaseContext = createContext<{ initialized: boolean }>({
  initialized: false,
})

export function EtherbaseProvider({
  config,
  children,
}: {
  config: EtherbaseConfig
  children: ReactNode
}) {
  const [initialized, setInitialized] = useState(false)

  useEffect(() => {
    if (!initialized) {
      initializeApp(config)
      setInitialized(true)
    }
  }, [config, initialized])

  return (
    <EtherbaseContext.Provider value={{ initialized }}>
      {initialized ? children : null}
    </EtherbaseContext.Provider>
  )
}

export function useEtherbaseContext() {
  const context = useContext(EtherbaseContext)
  if (context === undefined) {
    throw new Error(
      "useEtherbaseContext must be used within an EtherbaseProvider",
    )
  }
  return context
}
