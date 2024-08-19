import asyncHandler from 'express-async-handler'
import jwt, { JwtPayload } from 'jsonwebtoken'
import type { Request, Response } from 'express'
import { matchedData } from 'express-validator'

import { User, UserAttrs } from '@/models/user-model'
import { ConflictError } from '@errors/ConflictError'
import { Unauthorized } from '@errors/UnauthorizedError'
import { Password } from '@/services/password'
import { getEnv } from '@/config/env'
import { RefreshToken } from '@/models/jwt-tokens'
import mongoose from 'mongoose'

// @route POST /api/v1/auth/register
// @access Public
// @desc Register a user
export const createUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        // Access validated data
        const body = matchedData(req)
        const userData: UserAttrs = {
            username: body.username,
            email: body.email,
            role: body.role ?? 'member',
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
                role: user.role,
            },
        })
    }
)

// @route POST /api/v1/auth/login
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

        // generate access token
        const token = jwt.sign({ id: user._id }, getEnv('JWT_SECRET'), {
            subject: 'AccessAPI',
            expiresIn: getEnv('JWT_EXPIRES_IN'),
        })

        // generate refresh token
        const refreshToken = jwt.sign(
            { userId: user._id },
            getEnv('REFRESH_TOKEN_SECRET'),
            {
                subject: 'RefreshAPI',
                expiresIn: getEnv('REFRESH_TOKEN_EXPIRES_IN'),
            }
        )

        // Save refresh token in db
        await RefreshToken.build({
            refreshToken,
            userId: user._id as mongoose.Types.ObjectId,
        }).save()

        res.status(200).json({
            id: user._id,
            username: user.username,
            email: user.email,
            role: user.role,
            token,
            refreshToken,
        })
    }
)

// @route GET /api/v1/auth/current-user
// @access Private
// @desc return the current user
export const currentUser = asyncHandler(async (req: Request, res: Response) => {
    const { id } = req.user as JwtPayload

    const user = await User.findById(id)

    if (!user) {
        throw new Unauthorized('User not found')
    }

    res.status(200).json({
        id: user._id,
        username: user.username,
        email: user.email,
        role: user.role,
    })
})
