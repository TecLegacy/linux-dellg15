import express, {
	type Request,
	type Response,
	type NextFunction,
} from "express";

const app = express();

// Middleware to log incoming requests
app.use((req: Request, res: Response, next: NextFunction) => {
	console.log(`Received ${req.method} request for ${req.url}`);
	next();
});

// Route to handle GET requests to '/'
app.get("/", (req: Request, res: Response) => {
	res.send("Hello, World!");
});

// Route to handle POST requests to '/users'
app.post("/users", (req: Request, res: Response) => {
	const { name, email } = req.body;
	// Process the user data and save it to a database
	// ...

	res.send("User created successfully");
});

// Route to handle GET requests to '/users/:id'
app.get("/users/:id", (req: Request, res: Response) => {
	const userId = req.params.id;
	// Retrieve user data from the database based on the userId
	// ...

	res.send(`User with ID ${userId} found`);
});

// Error handling middleware
app.use((err: Error, req: Request, res: Response, next: NextFunction) => {
	console.error(err);
	res.status(500).send("Internal Server Error");
});

app.listen(3000, () => {
	console.log("Server is running on port 3000");
});

function addNumbers(a: number, b: number): number {
	return a + b;
}

addNumbers(1, 2); // Returns 3
