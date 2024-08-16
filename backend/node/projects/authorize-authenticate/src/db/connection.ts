import { MongoClient, Db } from 'mongodb'

let client: MongoClient | null = null
let db: Db | null = null

export async function connect(): Promise<Db | null> {
    // const username = process.env.MONGO_INITDB_ROOT_USERNAME
    // const password = process.env.MONGO_INITDB_ROOT_PASSWORD
    const username = 'admin'
    const password = 'password'

    if (!username || !password) {
        console.error('MongoDB username or password not found')
        return null
    }

    if (db) {
        return db
    }

    try {
        // client = await MongoClient.connect(
        //     `mongodb://${username}:${password}@mongodb:27017/2fa-db-delete?authSource=admin`
        // )
        client = await MongoClient.connect(
            `mongodb://admin:password@localhost:27017`
        )
        console.log('Connected to the database')

        db = client.db('2fa-development') // select the database to use

        return db
    } catch (error) {
        console.error('Error connecting to the database', error)
        return null
    }
}
