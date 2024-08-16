import express, { type Request, type Response } from 'express'
import asyncHandler from 'express-async-handler'
import { router as userRouter } from '@routes/user'
import { error } from '@middleware/errors'
import { NotFoundError } from '@errors/NotFoundError'
import { logEnvironmentVariables, noUnusedVars } from '@utils/dump-funcs'

import { connectMongoose } from '@/db/connection'

const app = express()

const PORT = Number(process.env.BACKEND_PORT) || 3000

// //local fast development
// const PORT = 3001

app.use(express.json())

app.use('/api/auth', userRouter)

app.get(
    '/api/auth/register/demo',
    asyncHandler((_: Request, res: Response) => {
        res.send('Hello World!')
    })
)

app.use('*', (_, res) => {
    throw new NotFoundError('Page Not Found!')
    noUnusedVars(res)
})

app.use(error)

app.listen(PORT, '0.0.0.0', async () => {
    logEnvironmentVariables()

    //connection to mongodb
    await connectMongoose()
    console.log(`Server is running on ports ${PORT}`)
})
