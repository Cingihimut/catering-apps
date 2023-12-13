/** @type {import('next').NextConfig} */
const nextConfig = {
    images:{
        remotePatterns: [
            {
                hostname: "flowbite.s3.amazonaws.com"
            },
            {
                hostname: "avatars.githubusercontent.com"
            }
        ]
    }

}

module.exports = nextConfig
