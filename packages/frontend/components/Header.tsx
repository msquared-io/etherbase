"use client"

import Link from "next/link"
import { usePathname } from "next/navigation"
import { cn } from "@/lib/utils"
import { Braces, Upload } from "lucide-react"

const navigation = [
  {
    href: "/sources",
    label: "Sources",
    icon: Upload,
  },
  {
    href: "/custom-contracts",
    label: "Custom Contracts",
    icon: Braces,
  },
]

export function Header() {
  const pathname = usePathname()

  return (
    <header className="sticky top-0 z-50 w-full border-b border-border bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60">
      <div className="container flex h-14 items-center">
        <div className="mr-4 flex">
          <Link href="/" className="mr-6 flex items-center space-x-2">
            <img src="/etherbase.png" alt="Etherbase" className="h-8 w-auto" />
          </Link>
        </div>
        <nav className="flex items-center space-x-6 text-sm font-medium">
          {navigation.map((item) => {
            const Icon = item.icon
            return (
              <Link
                key={item.href}
                href={item.href}
                className={cn(
                  "flex items-center space-x-2 transition-colors hover:text-foreground/80",
                  pathname?.startsWith(item.href)
                    ? "text-foreground"
                    : "text-foreground/60",
                )}
              >
                <Icon className="h-4 w-4" />
                <span>{item.label}</span>
              </Link>
            )
          })}
        </nav>
      </div>
    </header>
  )
}
