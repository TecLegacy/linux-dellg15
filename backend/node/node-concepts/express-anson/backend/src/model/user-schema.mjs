import mongoose, { Types } from 'mongoose';

const Schema = mongoose.Schema;

// Define the user schema
const userSchema = new Schema(
  {
    username: {
      type: String,
      required: true,
      unique: true,
      trim: true,
    },
    password: {
      type: String,
      required: true,
    },
    displayName: {
      type: String,
      required: true,
      trim: true,
    },
  },
  {
    timestamps: true, // Automatically create createdAt and updatedAt fields
  }
);

// Create a model from the schema
// export const User = mongoose.model('xxUser', userSchema);

// const document = new User({
//   username: 'keshav',
//   password: 'keshav123',
//   displayName: 'tecLegacy',
// });

// document.save();
