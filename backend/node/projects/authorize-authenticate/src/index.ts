import express from 'express'

import { router as userRouter } from '@routes/user'
import { error } from '@middlewares/errors'
import { NotFoundError } from '@errors/NotFoundError'
import { logEnvironmentVariables, noUnusedVars } from '@utils/dump-funcs'

import { connectMongoose } from '@/db/connection'
import { adminRouter } from '@routes/admin'
import { modRouter } from '@routes/moderator'

const app = express()

const PORT = Number(process.env.BACKEND_PORT) || 3000

app.use(express.json())

app.use('/api/v1/auth', userRouter)
app.use('/api/v1/admin', adminRouter)
app.use('/api/v1/moderator', modRouter)

app.use('*', (_, res) => {
    throw new NotFoundError('Page Not Found!')
    noUnusedVars(res)
})

app.use(error)

app.listen(PORT, '0.0.0.0', async () => {
    // TODO: REMOVE THIS IN PRODUCTION
    logEnvironmentVariables()

    //connection to mongodb
    await connectMongoose()
    console.log(`Server is running on ports ${PORT}`)
})
