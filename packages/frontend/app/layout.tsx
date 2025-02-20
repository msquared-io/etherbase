"use client"
import "@/app/globals.css"

import type * as React from "react"
import { EtherbaseProvider } from "@msquared/etherbase-client"
import { etherbaseConfig } from "./etherbaseConfig"
import { Header } from "@/components/Header"

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="en" className="dark">
      <body className="min-h-screen bg-background text-foreground">
        <div className="relative flex min-h-screen flex-col">
          <EtherbaseProvider config={etherbaseConfig}>
            <Header />
            <main className="flex-1">
              <div className="container py-6">{children}</div>
            </main>
          </EtherbaseProvider>
        </div>
      </body>
    </html>
  )
}
