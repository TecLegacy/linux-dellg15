import asyncHandler from 'express-async-handler'
import jwt from 'jsonwebtoken'
import type { Request, Response } from 'express'
import { matchedData, validationResult } from 'express-validator'

import { User, UserAttrs } from '@/models/user-model'
import { ConflictError } from '@/errors/ConflictError'
import { Unauthorized } from '@/errors/UnauthorizedError'
import { Password } from '@/services/password'
import { getEnv } from '@/config/env'

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

        res.status(201).json({
            message: 'User Created ',
            user: {
                id: user.id,
                username: user.username,
                email: user.email,
            },
        })
    }
)

// @route POST /api/auth/login
// @access Public
// @desc return the access token for the user
export const loginUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        // Access validated data
        const { email, password } = matchedData(req)

        // check if user doesn't exists
        const user = await User.findOne({ email })
        // user.
        if (!user) {
            throw new Unauthorized('Invalid credentials')
        }

        const isPassValid = await Password.comparePassword(
            password,
            user.password
        )
        if (!isPassValid) {
            throw new Unauthorized('Invalid credentials')
        }

        console.log(user)
        // generate token
        const token = jwt.sign({ user: user._id }, getEnv('JWT_SECRET'), {
            subject: 'AccessAPI',
            expiresIn: '15m',
        })

        res.status(200).json({
            username: user.username,
            email: user.email,
            token,
        })
    }
)

//
export const currentUser = (req: Request, res: Response) => {}
