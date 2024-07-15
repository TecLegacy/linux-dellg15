import { Router } from 'express';
import passport from 'passport';
passport;

const router = Router();

// authenticate route with username and password
router.post('/products', passport.authenticate('local'), (req, res) => {
  res.send('S');
});

router.get('/products', (req, res) => {
  req.user;
  console.log(req.user);

  res.send('user mapped on request object by deserialized');
});

router.get('/products/session-object', (req, res) => {
  console.log(req.session);

  res.send('user mapped on session object by serialized');
});

export const productsRoutes = router;
