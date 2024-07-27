import bcrypt from 'bcrypt';
export const hashPassword = async plain => {
  const salt = bcrypt.genSalt(10);
};
