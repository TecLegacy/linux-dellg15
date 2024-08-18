import { Router } from 'express'
import { createUser, loginUser } from '@/controllers/user-controller'

import {
    userRegisterValidationSchema,
    userLoginValidationSchema,
} from '@utils/user-validation-schema'
import { isValidUser } from '@middleware/validate-user'

export const router = Router()

// @route POST /api/v1/auth/register
// @access Public
// @desc Register a user
router.post('/register', userRegisterValidationSchema, isValidUser, createUser)

// @route POST /api/v1/auth/login
// @access Public
// @desc Login a user and return access token
router.post('/login', userLoginValidationSchema, isValidUser, loginUser)

//@route GET /api/v1/auth/login
router.get('/current-user')
