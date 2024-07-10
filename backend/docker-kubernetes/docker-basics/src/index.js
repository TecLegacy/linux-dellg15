import express from 'express';
import multer from 'multer';

const app = express();
const port = 3000;

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
  const { message } = req.body;
  console.log(req.body);
  res.status(201).json({ message });
});

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});
