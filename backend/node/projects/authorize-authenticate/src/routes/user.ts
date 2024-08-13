import { Router, Request, Response, NextFunction } from 'express'
import { registerUser } from '../controller/user-controller'
import {
    validateEmail,
    validateName,
    validatePassword,
} from '../utils/request-validations'
import { validationResult } from 'express-validator'

export const router = Router()

router.post('/register', validateUser(), registerUser)

function validateUser() {
    return [
        validateName,
        validateEmail,
        validatePassword,
        (req: Request, res: Response, next: NextFunction) => {
            const errors = validationResult(req)
            if (!errors.isEmpty()) {
                res.status(400).json({ errors: errors.array() })
                return
            }
            next()
        },
    ]
}
