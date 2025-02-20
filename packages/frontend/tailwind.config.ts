import type { Config } from "tailwindcss";

export default {
  darkMode: "class",
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "#121212",
        foreground: "#ffffff",
        primary: {
          DEFAULT: "#ffffff",
          foreground: "#121212",
        },
        muted: {
          DEFAULT: "#27272a",
          foreground: "#a1a1aa",
        },
        accent: {
          DEFAULT: "#18181b",
          foreground: "#ffffff",
        },
        popover: {
          DEFAULT: "#18181b",
          foreground: "#ffffff",
        },
        card: {
          DEFAULT: "#18181b",
          foreground: "#ffffff",
        },
        border: "#27272a",
        input: "#27272a",
        ring: "#27272a",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
} satisfies Config;
