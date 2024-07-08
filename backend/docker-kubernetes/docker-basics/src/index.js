const express = require('express');

const app = express();
const port = 3000;

app.use(express.json()); // for parsing application/json
app.use(express.urlencoded({ extended: true })); // for parsing application/x-www-form-urlencoded

app.set('view engine', 'ejs');
app.set('views', './src/views');

app.get('/', (req, res) => {
  // res.send('Hello World!');
  res.render('index', { title: 'Hello World!' });
});

app.post('/api/v1/echo', (req, res) => {
  const { name, imgUrl, message } = req.body;
  res.status(201).json({ message });
});

app.listen(port, () => {
  console.log(`Server listening at http://localhost:${port}`);
});
