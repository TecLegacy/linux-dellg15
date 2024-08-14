import { CustomError } from '@/errors/CustomError'
import { Request, Response } from 'express'

export function error(err: Error, _: Request, res: Response) {
    if (err instanceof CustomError) {
        res.status(err.statusCode).json({ errors: err.serializeErrors() })
        return
    }

    res.status(500).json({ error: err.message })
}
