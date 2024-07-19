import { Router } from 'express';
import { body, checkSchema, matchedData } from 'express-validator';
import {
  POSTuserBodySchema,
  userQuerySchema,
  PUTuserBodySchema,
} from '../utils/validation-schema.mjs';
import { userValidation } from '../middleware/express-validator.mjs';
import { users } from '../data/constant.mjs';
import { User } from '../model/user-schema.mjs';
import mongoose from 'mongoose';

const router = Router();

// Get user by ID
router.get('/users/:id', async (req, res) => {
  const userId = req.params.id;
  let userID;

  if (!mongoose.Types.ObjectId.isValid(userId)) {
    return res.status(400).json('Invalid ID format');
  }

  try {
    const user = await User.find({ _id: userId });
    if (!user) {
      return res.status(500).json('Error finding user');
    }
    userID = user;
  } catch (err) {
    return res
      .status(500)
      .json({ error: [{ message: 'Internal Server error' + err.toString() }] });
  }

  return res.status(200).json({ user: userID });
});

// Get all users with optional filter and sort
router.get(
  '/users',
  checkSchema(userQuerySchema),
  userValidation,
  async (req, res) => {
    const data = matchedData(req);
    let filteredUsers;

    try {
      const users = await User.find({});
      filteredUsers = users;
    } catch (err) {
      return res.status(500).json({
        error: [
          {
            message: 'Internal server error, failed fetching users',
          },
        ],
      });
    }

    // Apply filter if provided
    // return []users matching the condition
    if (data.filter) {
      filteredUsers = filteredUsers.filter(user => {
        return (
          user.username.toLowerCase().includes(data.filter.toLowerCase()) ||
          user.displayName.toLowerCase().includes(data.filter.toLowerCase())
        );
      });
    }

    // Apply sort if provided
    if (data.sort) {
      filteredUsers.sort((a, b) => {
        if (data.sort === 'username') {
          return a.username.localeCompare(b.username);
        } else if (data.sort === 'displayName') {
          return a.displayName.localeCompare(b.displayName);
        } else {
          return 0;
        }
      });
    }

    if (filteredUsers.length === 0) {
      return res.json('No User Found');
    }

    res.json(filteredUsers);
  }
);

// Create new user
router.post(
  '/users',
  checkSchema(POSTuserBodySchema),
  userValidation,
  async (req, res) => {
    try {
      const data = matchedData(req);
      const { username, displayName, password } = data;
      console.log(password);

      const document = new User({
        displayName,
        password,
        username,
      });
      await document.save();

      res.status(201).json({ user: document });
    } catch (error) {
      res
        .status(500)
        .json({ error: 'UserName Already Exists' + error.toString() });
    }
  }
);

// Update user using PUT
router.put(
  '/users/:id',
  checkSchema(PUTuserBodySchema),
  userValidation,
  async (req, res) => {
    const data = matchedData(req);
    const userId = req.params.id;

    if (!mongoose.Types.ObjectId.isValid(userId)) {
      return res.status(400).json('Invalid ID format');
    }

    try {
      let user = await User.findById(userId);
      if (!user) {
        return res.status(404).json('User not found');
      }
      user.username = data.username || 'null';
      user.displayName = data.displayName || 'null';

      await user.save();
      return res.status(200).json({
        message: 'Data modified',
      });
    } catch (err) {
      res.status(500).json({
        error: [
          {
            message: 'Failed due to internal server error' + err.toString(),
          },
        ],
      });
    }
  }
);

// Update user using PATCH
router.patch(
  '/users/:id',
  checkSchema(PUTuserBodySchema),
  userValidation,
  async (req, res) => {
    const data = matchedData(req);
    const { username, displayName } = data;

    const userId = req.params.id;
    if (!mongoose.Types.ObjectId.isValid(userId)) {
      return res.status(400).json('Invalid ID format');
    }

    try {
      const user = await User.findById({
        _id: req.params.id,
      });
      user.username = username;
      user.displayName = displayName;

      await user.save();

      return res.status(200).json({
        message: 'Data modified',
      });
    } catch (err) {
      res.status(500).json({
        error: [
          {
            message: 'Failed due to internal server error' + err.toString(),
          },
        ],
      });
    }
  }
);

// Delete user
router.delete('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }

  const userIndex = users.findIndex(user => user.id === userId);
  const deletedUser = users.splice(userIndex, 1);

  res.status(200).json({ user: deletedUser });
});

export const usersRoutes = router;
