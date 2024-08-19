import { tokenController } from '@controllers/token-controller'
import { Router } from 'express'
import { body } from 'express-validator'

export const tokenRouter = Router()

tokenRouter.post(
    '/',
    [
        body('refreshToken')
            .isString()
            .withMessage('refreshToken must be a string')
            .notEmpty()
            .withMessage('refreshToken cannot be empty')
            .isLength({ min: 20, max: 255 })
            .withMessage('refreshToken must be between 20 and 255 characters')
            .trim(),
    ],
    tokenController
)
