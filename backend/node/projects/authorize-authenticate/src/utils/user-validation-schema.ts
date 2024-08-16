import { checkSchema } from 'express-validator'

export const userValidationSchema = checkSchema({
    username: {
        in: ['body'],
        isString: {
            errorMessage: 'Username must be a string.',
        },
        trim: true,
        notEmpty: {
            errorMessage: 'Username must not be empty.',
        },
        isLength: {
            options: { min: 1, max: 50 },
            errorMessage: 'Username must be between 1 and 50 characters.',
        },
        matches: {
            options: /^[a-zA-Z\s-]+$/,
            errorMessage:
                'Username can only contain letters, spaces, and hyphens.',
        },
    },
    email: {
        in: ['body'],
        isEmail: {
            errorMessage: 'Invalid email address.',
        },
        normalizeEmail: true,
        isLength: {
            options: { min: 5, max: 254 },
            errorMessage: 'Email must be between 5 and 254 characters.',
        },
        trim: true,
    },
    password: {
        in: ['body'],
        isString: {
            errorMessage: 'Password must be a string.',
        },
        isLength: {
            options: { min: 8, max: 50 },
            errorMessage: 'Password must be between 8 and 50 characters.',
        },
        matches: {
            options:
                /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+])[A-Za-z\d!@#$%^&*()_+]+$/,
            errorMessage:
                'Password must contain at least one lowercase letter, one uppercase letter, one number, and one special character.',
        },
    },
})
