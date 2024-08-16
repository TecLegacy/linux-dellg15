export function noUnusedVars(...rest: unknown[]) {
    const noUnusedVarsDUMP = [...rest]
    console.log(noUnusedVarsDUMP)
}

// Function to log environment variables
export function logEnvironmentVariables() {
    console.log('NODE_ENV:', process.env.NODE_ENV)
    console.log('BACKEND_PORT:', process.env.BACKEND_PORT)
    console.log('MONGO_EXPRESS_PORT:', process.env.MONGO_EXPRESS_PORT)
    console.log('MONGO_PORT:', process.env.MONGO_PORT)
    console.log('PORT:', process.env.PORT)
    console.log('JWT_SECRET:', process.env.JWT_SECRET)
    console.log(
        'ME_CONFIG_MONGODB_ADMINUSERNAME:',
        process.env.ME_CONFIG_MONGODB_ADMINUSERNAME
    )
    console.log(
        'ME_CONFIG_MONGODB_ADMINPASSWORD:',
        process.env.ME_CONFIG_MONGODB_ADMINPASSWORD
    )
    console.log(
        'ME_CONFIG_MONGODB_SERVER:',
        process.env.ME_CONFIG_MONGODB_SERVER
    )
    console.log('MONGO_URI:', process.env.MONGO_URI)
    console.log(
        'MONGO_INITDB_ROOT_USERNAME:',
        process.env.MONGO_INITDB_ROOT_USERNAME
    )
    console.log(
        'MONGO_INITDB_ROOT_PASSWORD:',
        process.env.MONGO_INITDB_ROOT_PASSWORD
    )
}
