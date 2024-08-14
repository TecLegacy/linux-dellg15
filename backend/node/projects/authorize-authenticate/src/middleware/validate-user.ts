import { RequestValidationError } from '@/errors/RequestValidationError'
import { type NextFunction, Response, Request } from 'express'
import { validationResult } from 'express-validator'

export const isValidUser = (
    req: Request,
    _: Response,
    next: NextFunction
): void => {
    const errors = validationResult(req)

    if (!errors.isEmpty()) {
        throw new RequestValidationError(errors.array())
    }
    next()
}
