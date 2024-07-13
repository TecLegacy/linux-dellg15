import express from 'express';
import { matchedData, query, validationResult, body } from 'express-validator';

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Mock data
let users = [
  { id: 1, username: 'fiona', displayName: 'aliceWonder' },
  { id: 2, username: 'charlie', displayName: 'bobSpark' },
  { id: 3, username: 'bob', displayName: 'charlieCloud' },
  { id: 4, username: 'george', displayName: 'dianaPhoenix' },
  { id: 5, username: 'eric', displayName: 'ericThunder' },
  { id: 6, username: 'alice', displayName: 'fionaStardust' },
  { id: 7, username: 'diana', displayName: 'georgeSilver' },
];

// findUserIndex middleware
const findUserIndex = (req, res, next) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }
  const userIndex = users.findIndex(u => u.id === userId);
  if (userIndex) {
    req.userIndex = userIndex;
    next();
  }
};

//Sanitizing
const sanitizeString = userString => {
  return userString.replace(/[^\w\s]/gi, '');
};

// Routes

// Get all users with optional filter and sort
app.get(
  '/users',
  [
    query('filter')
      .optional()
      .isString()
      .withMessage('Filter must be a string')
      .notEmpty()
      .withMessage('Filter cannot be empty'),

    query('sort')
      .optional()
      .isIn(['displayName', 'username'])
      .withMessage('Sort must be displayName or username'),
  ],
  (req, res) => {
    const result = validationResult(req);
    if (!result.isEmpty()) {
      return res.status(400).json({ errors: result.array() });
    }
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

// Get user by ID
app.get('/users/:id', findUserIndex, (req, res) => {
  const { userIndex } = req;
  if (userIndex === -1) {
    return res.status(404).send('User not found');
  }

  res.json(users[userIndex]);
});

// Create new user
app.post(
  '/users',
  [
    body('username')
      .isString()
      .withMessage('Username must be a string')
      .notEmpty()
      .withMessage('Username cannot be empty'),

    body('displayName')
      .isString()
      .withMessage('Display Name must be a string')
      .notEmpty()
      .withMessage('Display Name cannot be empty'),
  ],
  (req, res) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() });
    }
    const data = matchedData(req);
    const { username, displayName } = data;

    const newUser = {
      id: users.length + 1,
      username: sanitizeString(username),
      displayName: sanitizeString(displayName),
    };
    users.push(newUser);
    res.status(201).json(newUser);
  }
);

// Update user using PUT
app.put(
  '/users/:id',
  [
    [
      body('username')
        .isString()
        .withMessage('Username must be a string')
        .notEmpty()
        .withMessage('Username cannot be empty'),

      body('displayName')
        .isString()
        .withMessage('Display Name must be a string')
        .notEmpty()
        .withMessage('Display Name cannot be empty'),
    ],
  ],
  findUserIndex,
  (req, res) => {
    const { userIndex } = req;

    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() });
    }
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
app.patch(
  '/users/:id',
  [
    [
      body('username')
        .isString()
        .withMessage('Username must be a string')
        .notEmpty()
        .withMessage('Username cannot be empty'),

      body('displayName')
        .isString()
        .withMessage('Display Name must be a string')
        .notEmpty()
        .withMessage('Display Name cannot be empty'),
    ],
  ],
  findUserIndex,
  (req, res) => {
    const errors = validationResult(req);
    if (!errors.isEmpty()) {
      return res.status(400).json({ errors: errors.array() });
    }
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
app.delete('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }
  users = users.filter(u => u.id !== userId);
  res.status(204).send();
});

// Start server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
