import express from 'express';

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Mock data
let users = [
  { id: 1, username: 'alice', displayName: 'aliceWonder' },
  { id: 2, username: 'bob', displayName: 'bobSpark' },
  { id: 3, username: 'charlie', displayName: 'charlieCloud' },
  { id: 4, username: 'diana', displayName: 'dianaPhoenix' },
  { id: 5, username: 'eric', displayName: 'ericThunder' },
  { id: 6, username: 'fiona', displayName: 'fionaStardust' },
  { id: 7, username: 'george', displayName: 'georgeSilver' },
];

// Helper function to filter and sort users
const filterAndSortUsers = query => {
  let filteredUsers = users;
  if (query.filter) {
    const sanitizedFilter = sanitizeString(query.filter);
    filteredUsers = filteredUsers.filter(user =>
      user.username.includes(sanitizedFilter)
    );
  }
  if (query.sort === '1') {
    filteredUsers.sort((a, b) => a.username.localeCompare(b.username));
  } else if (query.sort === '-1') {
    filteredUsers.sort((a, b) => b.username.localeCompare(a.username));
  }
  return filteredUsers;
};

// Sanitize string by removing special characters
const sanitizeString = str => {
  return str.replace(/[^\w\s]/gi, '');
};

// Routes

// Get all users with optional filter and sort
app.get('/users', (req, res) => {
  const filteredUsers = filterAndSortUsers(req.query);
  res.json(filteredUsers);
});

// Get user by ID
app.get('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }
  const user = users.find(u => u.id === userId);
  if (user) {
    res.json(user);
  } else {
    res.status(404).send('User not found');
  }
});

// Create new user
app.post('/users', (req, res) => {
  const { username, displayName } = req.body;
  if (!username || !displayName) {
    res.status(400).send('Username and displayName are required');
    return;
  }
  const newUser = {
    id: users.length + 1,
    username: sanitizeString(username),
    displayName: sanitizeString(displayName),
  };
  users.push(newUser);
  res.status(201).json(newUser);
});

// Update user using PUT
app.put('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }
  const userIndex = users.findIndex(u => u.id === userId);
  if (userIndex !== -1) {
    const { username, displayName } = req.body;
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
});

// Update user using PATCH
app.patch('/users/:id', (req, res) => {
  const userId = parseInt(req.params.id, 10);
  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }
  const userIndex = users.findIndex(u => u.id === userId);
  if (userIndex !== -1) {
    const user = users[userIndex];
    if (req.body.username) user.username = sanitizeString(req.body.username);
    if (req.body.displayName)
      user.displayName = sanitizeString(req.body.displayName);
    res.json(user);
  } else {
    res.status(404).send('User not found');
  }
});

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
