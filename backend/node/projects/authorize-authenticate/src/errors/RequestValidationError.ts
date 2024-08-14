import { ValidationError } from 'express-validator'
import { CustomError } from './CustomError'

export class RequestValidationError extends CustomError {
    statusCode = 400

    constructor(public error: ValidationError[]) {
        super(error[0].msg)
        Object.setPrototypeOf(this, RequestValidationError.prototype)
    }

    serializeErrors(): { message: string; field?: string }[] {
        return this.error.map((err) => {
            if (err.type === 'field') {
                return {
                    message: err.msg,
                    field: err.value,
                }
            }
            return { message: err.msg }
        })
    }
}
