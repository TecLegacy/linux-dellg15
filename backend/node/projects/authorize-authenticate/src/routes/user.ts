import { Router } from 'express'
import { registerUser, createUser } from '@controller/user-controller'

import { userValidationSchema } from '@utils/user-validation-schema'
import { isValidUser } from '@middleware/validate-user'

export const router = Router()

router.get('/register', registerUser)

router.post('/register', userValidationSchema, isValidUser, createUser)
