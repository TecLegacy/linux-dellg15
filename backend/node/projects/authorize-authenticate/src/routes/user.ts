import { Router } from 'express'
import { registerUser } from '../controller/user-controller'
import {
    validateEmail,
    validateName,
    validatePassword,
} from '../utils/request-validations'
// import { validateName } from '@/utils/request-validations'

export const router = Router()

router.get('/register', validateUser, registerUser)

function validateUser() {
    return [validateName, validateEmail, validatePassword]
}

// validateUser() //   [validateName, validateEmail, validatePassword]
