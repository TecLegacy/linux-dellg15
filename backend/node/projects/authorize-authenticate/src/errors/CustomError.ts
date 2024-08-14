export abstract class CustomError extends Error {
    constructor(public message: string) {
        super(message)
        console.log('Error Cause:', message)
        Object.setPrototypeOf(this, CustomError.prototype)
    }

    abstract statusCode: number
    abstract serializeErrors(): { message: string; field?: string }[]
}
