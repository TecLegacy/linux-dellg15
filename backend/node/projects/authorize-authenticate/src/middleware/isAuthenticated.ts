import { Request, Response, NextFunction } from 'express'
import jwt, { JwtPayload } from 'jsonwebtoken'
import { getEnv } from '@/config/env'
import { Unauthorized } from '@/errors/UnauthorizedError'

export const isAuthenticated = (
    req: Request,
    _: Response,
    next: NextFunction
) => {
    // Check if user has a JWT token in the request header
    const token = req.headers.authorization

    if (!token) {
        throw new Unauthorized('Access token not provided')
    }

    try {
        // Verify token
        const payload = jwt.verify(token, getEnv('JWT_SECRET')) as JwtPayload

        req.user = payload

        next()
    } catch (error) {
        console.log(error)
        throw new Unauthorized('Access token invalid or expired')
    }
}
