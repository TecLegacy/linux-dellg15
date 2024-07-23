import { Router } from 'express';

const router = Router();

//!How to setup session with express
router.get('/session', (req, res) => {
  //* How to setup cookies when using session
  console.log(
    'store data into cookie property',
    (req.session.cookie.myCookieData = 'this is how you do it')
  );

  // modifying session to persist for same request
  req.session.myName = 'teclegacy';

  console.log('Current session:', req.session);
  console.log('Current session ID:', req.sessionID);

  res.status(200).send('Session Saved');
});

router.get('/session-data', (req, res) => {
  //!How to access session DATA
  req.sessionStore.get(req.sessionID, (err, data) => {
    if (err) {
      console.log('Error retrieving session data:', err);
      return res.status(500).send('Internal Server Error');
    }
    console.log('Session data:', data);
  });

  console.log('Current session:', req.session);
  console.log('Current session ID:', req.sessionID);

  res.status(200).send('Session DATA');
});

export const practiceRouter = router;
