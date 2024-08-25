import express from 'express';
import { usersRoutes } from './routes/user.mjs';
import { productsRoutes } from './routes/product.mjs';
import session from 'express-session';
import passport from 'passport';
import dotenv from 'dotenv';

import { connectMongoose } from './db/connection-db.mjs';

import MongoStore from 'connect-mongo';
import { practiceRouter } from './routes/practice-session.mjs';

// dotenv
dotenv.config();

const app = express();

const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// express-session
app.use(
  session({
    secret: 'your_secret_key',
    resave: true,
    saveUninitialized: false,
    name: 'practice-session',
    cookie: {
      maxAge: 60000 * 60, //1hr make it 7 days
      // secure: process.env.NODE_ENV || true, // cookies are sent over https only
    },
    // persistent storage in mongodb over in-memory
    store: MongoStore.create({
      mongoUrl:
        'mongodb://admin:password@localhost:27017/anson-express?authSource=admin',
    }),
  })
);

app.use(passport.initialize());
app.use(passport.session());

app.use('/api/v1', usersRoutes);
app.use('/api/v1', productsRoutes);

//Session and Cookie practice arena
app.use('/practice', practiceRouter);

// Start server
app.listen(PORT, async () => {
  await connectMongoose();
  console.log(`Server is running on http://localhost:${PORT}`);
});
