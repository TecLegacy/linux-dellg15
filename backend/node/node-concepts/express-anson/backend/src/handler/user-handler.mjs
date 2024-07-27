import { User } from '../model/user-schema.mjs';
import { matchedData } from 'express-validator';

// @route GET /users/:id
export const getUserById = async (req, res) => {
  let userID;

  try {
    const user = await User.find({ _id: req.userId });

    if (!user) {
      return res.status(404).json({
        error: [{ message: 'User not found' }],
      });
    }

    userID = user;
  } catch (err) {
    return res
      .status(500)
      .json({ error: [{ message: `Internal Server error` }] });
  }

  return res.status(200).json({ user: userID });
};

// @router POST /user/:id
export const createUser = async (req, res) => {
  const data = matchedData(req);
  const { username, displayName, password } = data;

  //Check if user already exists
  const user = await User.findOne({ username });

  if (user) {
    return res.status(400).json({
      message: 'User already exists',
    });
  }

  // hash password and save it to db

  // console.log(username, displayName, password);
  try {
    const document = new User({
      displayName,
      password,
      username,
    });

    await document.save();

    return res.status(201).json({ user: { username, displayName } });
  } catch (error) {
    return res.status(500).json({ error: error.toString() });
  }
};
