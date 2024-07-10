import express, { json, urlencoded } from 'express';
import multer from 'multer';

const PORT = process.env.PORT || 3000;

const app = express();

const storage = multer.diskStorage({
  destination: function (req, file, cb) {
    // destination is used to specify the path of the directory in which the files have to be stored
    cb(null, './uploads');
  },
  filename: function (req, file, cb) {
    // It is the filename that is given to the saved file.
    cb(null, file.originalname);
  },
});

// Configure storage engine instead of dest object.
const upload = multer({ storage: storage });

// Include multer middleware
// Create an instance of multer and specify the destination folder for file uploads
// const upload = multer({ dest: 'uploads/' });

app.use(json());
app.use(urlencoded({ extended: true }));

app.get('/', (req, res) => {
  res.send('Welcome to API');
});

app.post('/upload', upload.single('myFile'), (req, res) => {
  console.log('Body: ', req.body);
  console.log('File: ', req.file);
  res.send('File successfully uploaded.');
});

app.listen(PORT, () => {
  console.log(`Server started on port : ${PORT}`);
});
