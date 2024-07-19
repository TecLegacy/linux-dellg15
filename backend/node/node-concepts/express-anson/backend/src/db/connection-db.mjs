// mongoose-connection.js
import mongoose from 'mongoose';

let dbInstance = null;

const userSchema = new mongoose.Schema({
  name: String,
  age: Number,
  email: String,
});

// Create a model. Mongoose will create a collection named 'users' if it doesn't exist.
const Userx = mongoose.model('Userx', userSchema);

// Now, when you save a document using the User model, the 'users' collection is created automatically if it doesn't exist.
const newUser = new Userx({
  name: 'John Doe',
  age: 30,
  email: 'john@example.com',
});

export const connectMongoose = async () => {
  if (dbInstance) {
    return dbInstance; // Return the existing connection if already connected
  }

  // const mongoURI = process.env.MONGO_URL || 'mongodb://anson:anson@mongo:27017';

  //local development --ignore
  const mongoURI =
    process.env.MONGO_URL ||
    'mongodb://admin:password@localhost:27017/anson-express?authSource=admin';

  try {
    const connection = await mongoose.connect(mongoURI);
    await newUser.save();
    console.log('Mongoose connected');

    dbInstance = connection.connection.db; // Store the db instance for reuse
    return dbInstance;
  } catch (error) {
    console.error('Mongoose connection error:', error);
    process.exit(1);
  }
};
