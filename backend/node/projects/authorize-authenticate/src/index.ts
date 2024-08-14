import express, { type Request, type Response } from 'express'
import asyncHandler from 'express-async-handler'
import { router as userRouter } from './routes/user'

const app = express()

const PORT = process.env.PORT || 3000

app.use(express.json())

app.use('/api/auth', userRouter)

app.get(
    '/api/auth/register/demo',
    asyncHandler((_: Request, res: Response) => {
        res.send('Hello World!')
    })
)

//start the server
app.listen(PORT, () => {
    console.log(`Server is running on portx ${PORT}`)
})
