import { Router } from 'express';
import { findUserIndex, sanitizeString } from '../utils/helper.mjs';
import { checkSchema, matchedData } from 'express-validator';
import {
  userBodySchema,
  userQuerySchema,
} from '../utils/validation-schema.mjs';
import { userValidation } from '../middleware/express-validator.mjs';
import { users } from '../data/constant.mjs';
// import { User } from '../model/user-schema.mjs';

const router = Router();

// Get user by ID
router.get('/users/:id', findUserIndex, (req, res) => {
  return res.status(200).json({ user: users[req.userIndex] });
});

// Get all users with optional filter and sort
router.get(
  '/users',
  checkSchema(userQuerySchema),
  userValidation,
  (req, res) => {
    const data = matchedData(req);

    let filteredUsers = users;

    // Apply filter if provided
    if (data.filter) {
      filteredUsers = users.filter(user => {
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
  checkSchema(userBodySchema),
  userValidation,
  async (req, res) => {
    const data = matchedData(req);
    const { username, displayName } = data;

    try {
      const newUser = new User({
        username: sanitizeString(username),
        displayName: sanitizeString(displayName),
      });
      await newUser.save();
      res.status(201).json(newUser);
    } catch (err) {
      res.status(500).send('Failed to create user');
    }
  }
);

// Update user using PUT
router.put(
  '/users/:id',
  checkSchema(userBodySchema),
  userValidation,
  findUserIndex,
  (req, res) => {
    const { userIndex } = req;

    const data = matchedData(req);

    if (userIndex !== -1) {
      const { username, displayName } = data;
      if (!username || !displayName) {
        res.status(400).send('Username and displayName are required');
        return;
      }
      users[userIndex] = {
        id: users[userIndex].id,
        username: sanitizeString(username),
        displayName: sanitizeString(displayName),
      };
      res.json(users[userIndex]);
    } else {
      res.status(404).send('User not found');
    }
  }
);

// Update user using PATCH
router.patch(
  '/users/:id',
  findUserIndex,
  checkSchema(userBodySchema),
  userValidation,
  (req, res) => {
    const data = matchedData(req);
    const { username, displayName } = data;
    const { userIndex } = req;
    if (userIndex !== -1) {
      const user = users[userIndex];
      if (username) user.username = sanitizeString(username);
      if (displayName) user.displayName = sanitizeString(displayName);
      res.json(user);
    } else {
      res.status(404).send('User not found');
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
