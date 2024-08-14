import { CustomError } from './CustomError'

export class DatabaseError extends CustomError {
    statusCode = 500

    constructor(public message: string = 'Database error occurred') {
        super(message)
        Object.setPrototypeOf(this, DatabaseError.prototype)
    }

    serializeErrors() {
        return [{ message: this.message }]
    }
}
