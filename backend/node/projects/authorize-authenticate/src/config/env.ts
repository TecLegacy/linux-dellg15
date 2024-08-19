type Config = {
    [key: string]: string
}

export const getEnv = (key: string): string => {
    const envValue = process.env[key]
    return envValue === undefined ? config[key] : envValue
}

// for development purposes
const config: Config = {
    // Access token secret and expiration time
    JWT_SECRET: 'superman',
    JWT_EXPIRES_IN: '15m',

    // Refresh token secret and expiration time
    REFRESH_TOKEN_SECRET: 'batman',
    REFRESH_TOKEN_EXPIRES_IN: '7d',
} as const
