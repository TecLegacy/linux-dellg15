import { CustomError } from './CustomError'

export class UnprocessableEntityError extends CustomError {
    statusCode = 422

    constructor(public message: string = 'Unprocessable Entity') {
        super(message)
        Object.setPrototypeOf(this, UnprocessableEntityError.prototype)
    }

    serializeErrors() {
        return [{ message: this.message }]
    }
}
