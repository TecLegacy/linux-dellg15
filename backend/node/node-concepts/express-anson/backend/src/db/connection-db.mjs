import mongoose from 'mongoose';

let dbInstance = null;

export const connectMongoose = async () => {
  if (dbInstance) {
    return dbInstance; // Return the existing connection if already connected
  }

  //* Connection string inside docker container
  // const mongoURI = process.env.MONGO_URL || 'mongodb://anson:anson@mongo:27017';

  /*
   local development
   ignore this uri
  */
  const mongoURI =
    process.env.MONGO_URL ||
    'mongodb://admin:password@localhost:27017/anson-express?authSource=admin';

  try {
    const connection = await mongoose.connect(mongoURI);
    console.log('Mongoose connected');

    dbInstance = connection.connection.db; // Store the db instance for reuse
    return dbInstance;
  } catch (error) {
    console.error('Mongoose connection error:', error);
    process.exit(1);
  }
};
