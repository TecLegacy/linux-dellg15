import { Strategy } from 'passport-local';
import passport from 'passport';
import { users } from '../data/constant.mjs';

export const userAuthStrategy = passport.use(
  new Strategy((username, password, done) => {
    // Query DB
    try {
      const user = users.find(user => user.username === username);
      if (!user) throw new Error(`No user with username : ${username} found`);
      if (user.password !== password) throw new Error('Invalid Credentials');

      done(null, { id: user.id, username: user.username }); // send back user if no error
    } catch (err) {
      done(err, null); //send error if no user
    }
  })
);

// userObj maps sessionObject
passport.serializeUser((user, done) => {
  done(null, user.id);
});

// userObj maps requestObj
passport.deserializeUser((id, done) => {
  try {
    // get user from db
    const user = users.find(user => user.id === id);
    if (!user) throw new Error(`No user with username : ${username} `);

    done(null, { id: user.id, username: user.username });
  } catch (err) {
    done(err, null);
  }
});
