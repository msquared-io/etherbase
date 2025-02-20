declare global {
  interface Window {
    ethereum?: {
      request: (args: {
        method: string
        params?: unknown[]
      }) => Promise<unknown>
      on: (event: string, callback: (params: unknown) => void) => void
      removeListener: (
        event: string,
        callback: (params: unknown) => void,
      ) => void
    }
  }
}

export {}
