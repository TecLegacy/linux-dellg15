import { MongoClient, Db } from 'mongodb'

let client: MongoClient | null = null
let db: Db | null = null

export async function connect(): Promise<Db | null> {
    const username = process.env.MONGO_INITDB_ROOT_USERNAME
    const password = process.env.MONGO_INITDB_ROOT_PASSWORD

    if (!username || !password) {
        console.error('MongoDB username or password not found')
        return null
    }

    if (db) {
        return db
    }

    try {
        client = await MongoClient.connect(
            `mongodb://${username}${password}mongo:27017`
        )
        db = client.db('2fa')
        return db
    } catch (error) {
        console.error('Error connecting to the database', error)
        return null
    }
}

export const dbInstance = async () => {
    return await connect()
}
