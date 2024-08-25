import { getEnv } from '@/config/env'
import { RequestValidationError } from '@/errors/RequestValidationError'
import { Unauthorized } from '@/errors/UnauthorizedError'
import { RefreshToken } from '@/models/jwt-tokens'
import { Request, Response } from 'express'
import { matchedData, validationResult } from 'express-validator'
import asyncHandler from 'express-async-handler'
import jwt, { JwtPayload } from 'jsonwebtoken'

//@ Get /token
export const tokenController = asyncHandler(
    async (req: Request, res: Response) => {
        // validate request
        const result = validationResult(req)
        if (!result.isEmpty()) {
            throw new RequestValidationError(result.array())
        }

        const { refreshToken } = matchedData(req)

        const decodedRefreshToken = jwt.verify(
            refreshToken,
            getEnv('REFRESH_TOKEN_SECRET')
        ) as JwtPayload

        const userId = decodedRefreshToken.userId

        // get refresh token from db
        const refreshTokenDoc = await RefreshToken.findOne({
            refreshToken,
            userId,
        })

        if (!refreshTokenDoc) {
            throw new Unauthorized(
                'Invalid refresh token or expired refresh token'
            )
        }

        await refreshTokenDoc.deleteOne()

        // generate new access token & new refresh token
        const newAccessToken = jwt.sign({ id: userId }, getEnv('JWT_SECRET'), {
            subject: 'AccessAPI',
            expiresIn: getEnv('JWT_EXPIRES_IN'),
        })

        const newRefreshToken = jwt.sign(
            { userId },
            getEnv('REFRESH_TOKEN_SECRET'),
            {
                subject: 'RefreshAPI',
                expiresIn: getEnv('REFRESH_TOKEN_EXPIRES_IN'),
            }
        )

        // save new refresh token in db
        await RefreshToken.build({
            refreshToken: newRefreshToken,
            userId: userId,
        }).save()

        res.status(200).json({
            accessToken: newAccessToken,
            refreshToken: newRefreshToken,
        })
    }
)
