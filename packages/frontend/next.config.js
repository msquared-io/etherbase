const nextConfig = {
  reactStrictMode: true,
  transpilePackages: ["@msquared/etherbase-client"],
  webpack: (config) => {
    return config
  },
}

module.exports = nextConfig
