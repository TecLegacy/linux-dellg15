import { User } from '../model/user-schema.mjs';

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
