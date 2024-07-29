"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = (0, express_1.default)();
// Middleware to log incoming requests
app.use((req, res, next) => {
    console.log(`Received ${req.method} request for ${req.url}`);
    next();
});
// Route to handle GET requests to '/'
app.get('/', (req, res) => {
    res.send('Hello, World!');
});
// Route to handle POST requests to '/users'
app.post('/users', (req, res) => {
    const { name, email } = req.body;
    // Process the user data and save it to a database
    // ...
    res.send('User created successfully');
});
// Route to handle GET requests to '/users/:id'
app.get('/users/:id', (req, res) => {
    const userId = req.params.id;
    // Retrieve user data from the database based on the userId
    // ...
    res.send(`User with ID ${userId} found`);
});
// Error handling middleware
app.use((err, req, res, next) => {
    console.error(err);
    res.status(500).send('Internal Server Error');
});
app.listen(3000, () => {
    console.log('Server is running on port 3000');
});
function addNumbers(a, b) {
    return a + b;
}
addNumbers(1, 2); // Returns 3
