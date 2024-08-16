import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'
import { validationResult } from 'express-validator'
import { User } from '@/models/user-model'

export const registerUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        const errors = validationResult(req)
        if (!errors.isEmpty()) {
            res.status(400).json({ errors: errors.array() })
            return
        }

        res.send('Hello World!!!')
    }
)

export const createUser = asyncHandler(
    async (_: Request, res: Response): Promise<void> => {
        //validate user

        // check if user exists
        const data = {
            username: 'admin',
            email: 'test@gamail.com',
            password: 'password',
        }
        const user = User.build(data)

        await user.save()
        res.send(user)
        // res.send('Successfully created user')
        return
    }
)
