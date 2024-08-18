type Config = {
    [key: string]: string
}

export const getEnv = (key: string): string => {
    const envValue = process.env[key]
    return envValue === undefined ? config[key] : envValue
}

// for development purposes
const config: Config = {
    JWT_SECRET: 'superman',
} as const
