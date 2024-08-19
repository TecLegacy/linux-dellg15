import { Request, Response, NextFunction } from 'express'
import jwt, { JwtPayload } from 'jsonwebtoken'
import { getEnv } from '@/config/env'
import { Unauthorized } from '@/errors/UnauthorizedError'
import { Blacklist } from '@/models/jwt-tokens'

import asyncHandler from 'express-async-handler'

export const isAuthenticated = asyncHandler(
    async (req: Request, res: Response, next: NextFunction) => {
        // Check if user has a JWT token in the request header
        const token = req.headers.authorization

        if (!token) {
            throw new Unauthorized('Access token not provided')
        }
        try {
            // Verify token
            const payload = jwt.verify(
                token,
                getEnv('JWT_SECRET')
            ) as JwtPayload

            // check in blacklist if access token lives
            if (await Blacklist.findOne({ userId: payload.id })) {
                res.status(401).json({
                    message: 'Access token invalid',
                    code: 'AccessTokenInvalid',
                })
                return
            }

            req.accessToken = { value: token, exp: payload.exp as number }

            // TODO: req.user should only have user id and not the entire payload
            // update the type.d.ts user : { id : string | mongoose.Types.ObjectId }
            req.user = payload

            next()
        } catch (error) {
            console.log(error)

            if (error instanceof jwt.TokenExpiredError) {
                res.status(401).json({
                    message: 'Access token expired',
                    code: 'AccessTokenExpired',
                })
                return
            }
            if (error instanceof jwt.JsonWebTokenError) {
                res.status(401).json({
                    message: 'Access token invalid',
                    code: 'AccessTokenInvalid',
                })

                return
            }
        }
    }
)
