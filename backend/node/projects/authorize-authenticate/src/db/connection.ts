import { DatabaseError } from '@/errors/DatabaseError'
import { MongoClient, Db } from 'mongodb'
import mongoose from 'mongoose'

// Mongoose connection
export async function connectMongoose(): Promise<void> {
    const uri =
        process.env.MONGO_URI ||
        'mongodb://admin:password@localhost:27017/2fa-development-mongoose?authSource=admin'

    if (!uri) {
        console.error('MongoDB URI not found')
        return
    }

    try {
        await mongoose.connect(uri)

        console.log('Connected to the database')
    } catch (error) {
        console.error('Error connecting to the database', error)
    }
}

// MONGODB connection
let client: MongoClient | null = null
let db: Db | null = null

export async function connect(): Promise<Db | null> {
    const uri =
        process.env.MONGO_URI ||
        'mongodb://admin:password@localhost:27017/2fa-development?authSource=admin'

    //Local development
    // const username = 'admin'
    // const password = 'password'

    if (!uri && uri?.trim() === '') {
        console.error('MongoDB username or password not found')
        return null
    }

    if (db) {
        return db
    }

    try {
        // Container connection string
        client = await MongoClient.connect(uri!)

        // select the database to use
        db = client.db('2fa-development')

        console.log('Connected to the database')
        return db
    } catch (error) {
        console.error('Error connecting to the database', error)
        return null
    }
}

//Test Helper function to insert a user into the database
export async function _insertUserMongo() {
    const db: Db | null = await connect()
    if (!db) {
        throw new DatabaseError('Internal server error')
    }

    const user = await db.collection('users').insertOne({
        username: 'cooluser',
        password: 'coolpassword',
    })
    console.log('User created', user)
}
