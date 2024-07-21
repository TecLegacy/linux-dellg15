import mongoose from 'mongoose';

export const checkUserId = (req, res, next) => {
  const usr = req.params.id;
  if (!mongoose.Types.ObjectId.isValid(usr)) {
    return res.status(400).json({
      message: 'Invalid Id type ',
    });
  }
  req.userId = usr;
  next();
};
