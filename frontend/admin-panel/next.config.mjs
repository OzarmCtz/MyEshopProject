/** @type {import('next').NextConfig} */
const nextConfig = {
  // Disable ESLint during build for demo
  eslint: {
    ignoreDuringBuilds: true,
  },
  // Disable TypeScript errors during build
  typescript: {
    ignoreBuildErrors: true,
  },
  images: {
    domains: ["localhost"],
  },
};

export default nextConfig;
