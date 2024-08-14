import express, { type Request, type Response } from 'express'
import asyncHandler from 'express-async-handler'
import { router as userRouter } from './routes/user'
import { error } from './middleware/errors'

const app = express()

const PORT = Number(process.env.PORT) || 3000

app.use(express.json())

app.use('/api/auth', userRouter)

app.get(
    '/api/auth/register/demo',
    asyncHandler((_: Request, res: Response) => {
        res.send('Hello World!')
    })
)

app.use(error)

app.listen(PORT, '0.0.0.0', () => {
    console.log(`Server is running on port check ${PORT}`)
})
