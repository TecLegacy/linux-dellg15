import express from 'express';
import multer from 'multer';

import { connectMongoDB, getDbInstance } from './db/mongo-client.js';
import { connectMongoose } from './db/mongoose.js';

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json()); // for parsing application/json
app.use(express.urlencoded({ extended: true })); // for parsing application/x-www-form-urlencoded

app.set('view engine', 'ejs');
app.set('views', './src/views');

// Multer configuration
const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    cb(null, './uploads'); // Specify the directory where files should be uploaded
  },
  filename: (req, file, cb) => {
    cb(null, Date.now() + file.originalname);
  },
});

const upload = multer({ storage });

app.get('/', (req, res) => {
  // res.send('Hello World!');
  res.render('index', { title: 'Hello World!' });
});

// Endpoint to handle file upload
app.post('/upload', upload.single('imageUpload'), (req, res) => {
  // Access uploaded file from req.file
  const file = req.file;
  console.log(file);
  if (!file) {
    return res.status(400).send('No file uploaded.');
  }
  // Process the file (e.g., save it to database, return information)
  res.send({ message: 'File uploaded successfully.', filename: file.filename });
});

app.post('/api/v1/echo', (req, res) => {
  console.log(req.body);
  const db = getDbInstance();
  db.collection('user').insertOne(req.body, (err, result) => {
    if (err) {
      console.error('Error inserting document:', err);
      return;
    }
    console.log('Document inserted:', result.ops[0]);
  });

  res.status(201).json({ message: 'data saved ', alertBox: true });
});

async function startServer() {
  try {
    // Wait for MongoDB connection
    await connectMongoDB(); //! MONOGO client DB
    // await connectMongoose(); //! MONGOOSE

    app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
  } catch (error) {
    console.error('Failed to connect to MongoDB:', error);

    process.exit(1);
  }
}

startServer();
