import { users } from '../data/constant.mjs';

export const findUserIndex = (req, res, next) => {
  const userId = parseInt(req.params.id, 10);

  if (isNaN(userId)) {
    res.status(400).send('Invalid user ID');
    return;
  }

  const userIndex = users.findIndex(u => u.id === userId);

  if (userIndex !== -1) {
    req.userIndex = userIndex;
    next();
  } else {
    res.status(404).json('No user found');
    d;
  }
};

//Sanitizing
export const sanitizeString = userString => {
  return userString.replace(/[^\w\s]/gi, '');
};
