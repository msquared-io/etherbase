"use client"

import { Card, CardHeader, CardTitle, CardDescription } from "@/components/ui/card"
import { Upload } from "lucide-react"
import Link from "next/link"

const features = [
  {
    title: "Manage Sources",
    description: "Create and manage event sources for your blockchain events and state.",
    icon: Upload,
    href: "/sources",
  },
]

export default function Home() {
  return (
    <div className="flex flex-col items-center space-y-8">
      <div className="text-center space-y-4">
        <h1 className="text-4xl font-bold tracking-tighter">
          Welcome to Etherbase
        </h1>
        <p className="text-lg text-muted-foreground max-w-[600px]">
          Your platform for managing blockchain events and data sources with ease.
        </p>
      </div>

      <div className="w-full max-w-xl">
        {features.map((feature) => {
          const Icon = feature.icon
          return (
            <Link key={feature.title} href={feature.href}>
              <Card className="transition-colors hover:bg-muted/50">
                <CardHeader>
                  <div className="flex items-center space-x-4">
                    <div className="p-2 bg-primary/10 rounded-lg">
                      <Icon className="h-6 w-6 text-primary" />
                    </div>
                    <div className="space-y-1">
                      <CardTitle>{feature.title}</CardTitle>
                      <CardDescription>{feature.description}</CardDescription>
                    </div>
                  </div>
                </CardHeader>
              </Card>
            </Link>
          )
        })}
      </div>
    </div>
  )
}
