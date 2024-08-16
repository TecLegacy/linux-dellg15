import { noUnusedVars } from '@utils/dump-funcs'
import { CustomError } from '@errors/CustomError'

import { NextFunction, Request, Response } from 'express'

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

    return res.status(500).json({ error: err.message })
    noUnusedVars(next)
}
