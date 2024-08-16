import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'
import { matchedData, validationResult } from 'express-validator'
import { User, UserAttrs } from '@/models/user-model'
import { ConflictError } from '@/errors/ConflictError'

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

// @route POST /api/auth/register
// @access Public
// @desc Register a user
export const createUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        // Access validated data
        const body = matchedData(req)
        const userData: UserAttrs = {
            username: body.username,
            email: body.email,
            password: body.password,
        }

        // check if user exists
        const existingUser = await User.findOne({ email: userData.email })
        if (existingUser) {
            throw new ConflictError('User already exists')
        }

        // create user
        const user = User.build(userData)
        await user.save()

        res.status(201).json({ message: 'User Created ', user })
        return
    }
)
