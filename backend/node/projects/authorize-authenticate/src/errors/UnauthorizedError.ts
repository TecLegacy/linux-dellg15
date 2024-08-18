import { CustomError } from './CustomError'

export class Unauthorized extends CustomError {
    statusCode = 401

    constructor(public message: string) {
        super(message)
        Object.setPrototypeOf(this, Unauthorized.prototype)
    }

    serializeErrors() {
        return [{ message: this.message }]
    }
}
