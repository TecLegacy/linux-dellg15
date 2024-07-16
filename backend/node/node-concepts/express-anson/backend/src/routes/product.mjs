import { Router } from 'express';
import passport from 'passport';
passport;

const router = Router();

// authenticate route with username and password
router.post('/products', passport.authenticate('local'), (req, res) => {
  res.send('S');
});

router.get('/products', (req, res) => {
  console.log(req.user);

  res.send('user mapped on request object by deserialized');
});

router.get('/products/session-object', (req, res) => {
  console.log(req.session);

  res.send('user mapped on session object by serialized');
});

router.get('/logout', (req, res) => {
  req.logout(err => {
    if (err) {
      return next(err);
    }
    res.redirect('/api/v1/get-out'); // Redirect to the home page or login page after logout
  });
});

router.get('/api/v1/get-out', (req, res) => {
  res.status(200).send('You are loggedOut');
});

export const productsRoutes = router;
