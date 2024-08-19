import { Router } from 'express'
import {
    createUser,
    currentUser,
    loginUser,
} from '@controllers/user-controller'

import {
    userRegisterValidationSchema,
    userLoginValidationSchema,
} from '@utils/user-validation-schema'
import { isValidUser } from '@middlewares/validate-user'
import { isAuthenticated } from '@middlewares/isAuthenticated'

export const router = Router()

// @route POST /api/v1/auth/register
// @access Public
// @desc Register a user
router.post('/register', userRegisterValidationSchema, isValidUser, createUser)

// @route POST /api/v1/auth/login
// @access Public
// @desc Login a user and return access token
router.post('/login', userLoginValidationSchema, isValidUser, loginUser)

//@route GET /api/v1/auth/current-user
// @access Private
// @desc Get current user
router.get('/current-user', isAuthenticated, currentUser)
