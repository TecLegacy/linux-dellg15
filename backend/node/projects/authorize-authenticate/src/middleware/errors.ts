import { noUnusedVars } from '@utils/dump-funcs'
import { CustomError } from '@errors/CustomError'

import { NextFunction, Request, Response } from 'express'
import { JsonWebTokenError, TokenExpiredError } from 'jsonwebtoken'

export function error(
    err: Error,
    _: Request,
    res: Response,
    next: NextFunction
) {
    if (err instanceof CustomError) {
        res.status(err.statusCode).json({ errors: err.serializeErrors() })
        return
    }

    if (err instanceof JsonWebTokenError || err instanceof TokenExpiredError) {
        return res
            .status(401)
            .json({ error: 'refresh token invalid or expired' })
    }

    return res.status(500).json({ error: err.message })
    noUnusedVars(next)
}
