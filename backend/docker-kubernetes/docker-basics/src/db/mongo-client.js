import { MongoClient } from 'mongodb';

let connection = null;
let dbInstance = null;

export const connectMongoDB = async () => {
  if (dbInstance) {
    return dbInstance; // Return the existing db instance if already connected
  }

  const mongoURL =
    process.env.MONGO_URI || 'mongodb://admin:password@localhost:27017';
  const client = new MongoClient(mongoURL);

  try {
    if (!connection) {
      // Check if connection is not already established
      connection = await client.connect();
      console.log('MongoDB connected');
    }
    const db = connection.db('mongo-express');
    dbInstance = db; // Store the db instance for reuse
    return dbInstance;
  } catch (error) {
    console.error('MongoDB connection error:', error);
    process.exit(1);
  }
};

// Optionally, if you need to access the db instance directly elsewhere
export const getDbInstance = () => {
  if (!dbInstance) {
    throw new Error(
      'Database not connected. Please connect first using connectMongoDB.'
    );
  }
  return dbInstance;
};
