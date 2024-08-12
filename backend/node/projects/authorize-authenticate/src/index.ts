import express, { type Request, type Response } from 'express'
import asyncHandler from 'express-async-handler'
import { router as userRouter } from './routes/user'

const app = express()

// When working with mongo-express in docker compose
// it errors with KEY "PORT" so change any other key ex "BACKEND_PORT"
const PORT = process.env.BACKEND_PORT || 3000

app.use(express.json())

app.use('/api/auth', userRouter)

app.get(
    '/api/auth/register/demo',
    asyncHandler((_: Request, res: Response) => {
        res.send('Hello World!')
    })
)

app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`)
})
