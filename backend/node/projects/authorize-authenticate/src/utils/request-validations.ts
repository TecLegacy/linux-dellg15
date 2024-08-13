import { body } from 'express-validator'

export function validateName() {
    console.log('hit1')
    return body('name')
        .isString()
        .withMessage('Name must be a string.')
        .trim()
        .notEmpty()
        .withMessage('Name must not be empty.')
        .isLength({ min: 1, max: 50 })
        .withMessage('Name must be between 1 and 50 characters.')
        .matches(/^[a-zA-Z\s-]+$/) // Allow only letters, spaces, and hyphens (adjust as needed)
        .withMessage('Name can only contain letters, spaces, and hyphens.')
}

export function validateEmail() {
    console.log('hit2')
    return body('email')
        .isEmail()
        .withMessage('Invalid email address.')
        .normalizeEmail()
        .isLength({ min: 5, max: 254 })
        .withMessage('Email must be between 5 and 254 characters.')
        .trim()
}

export function validatePassword() {
    console.log('hit3')
    return body('password')
        .isString()
        .withMessage('Password must be a string.')
        .isLength({ min: 8, max: 50 })
        .withMessage('Password must be between 8 and 50 characters.')
        .matches(
            /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+])[A-Za-z\d!@#$%^&*()_+]+$/
        )
        .withMessage(
            'Password must contain at least one lowercase letter, one uppercase letter, one number, and one special character.'
        )
}
