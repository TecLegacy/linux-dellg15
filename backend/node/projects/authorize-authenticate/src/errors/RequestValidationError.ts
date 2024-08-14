import { ValidationError } from 'express-validator'
import { CustomError } from './CustomError'

export class RequestValidationError extends CustomError {
    statusCode = 400
    constructor(public error: ValidationError[]) {
        super('Failed parsing validation Lib')

        Object.setPrototypeOf(this, RequestValidationError.prototype)
    }

    serializeErrors(): { message: string; field?: string }[] {
        return this.error.map((err) => {
            if (err.type === 'field') {
                return {
                    message: err.msg,
                    field: err.path,
                }
            }
            return { message: err.msg }
        })
    }
}
