import express from 'express';
import { usersRoutes } from './routes/user.mjs';
import { productsRoutes } from './routes/product.mjs';
import session from 'express-session';
import passport from 'passport';
import dotenv from 'dotenv';
import { userAuthStrategy } from './strategy/local-strategy.mjs';
import { connectMongoose } from './db/connection-db.mjs';
import mongoose from 'mongoose';
import MongoStore from 'connect-mongo';

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
    resave: false,
    saveUninitialized: false,
    cookie: {
      // secure: true, // Use secure cookies in production
      maxAge: 60000 * 60, //1hr make it 7 days
    },
    // store: MongoStore.create({
    //   mongoUrl: mongoose.connection.getClient(),
    // }),
  })
);

app.use(passport.initialize());
app.use(passport.session());

app.use('/api/v1', usersRoutes);
app.use('/api/v1', productsRoutes);

// Start server
app.listen(PORT, async () => {
  // await connectMongoose();
  console.log(`Server is running on http://localhost:${PORT}`);
});
