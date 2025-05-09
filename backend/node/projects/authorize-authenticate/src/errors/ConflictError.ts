import { CustomError } from './CustomError'

export class ConflictError extends CustomError {
    statusCode = 409

    constructor(public message: string = 'Conflict occurred') {
        super(message)
        Object.setPrototypeOf(this, ConflictError.prototype)
    }

    serializeErrors() {
        return [{ message: this.message }]
    }
}
