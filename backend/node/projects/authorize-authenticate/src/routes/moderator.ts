import { Router } from 'express'

import { getModerator } from '@controllers/moderator-controller'
import { isAuthenticated } from '@middlewares/isAuthenticated'
import { isAuthorized } from '@middlewares/isAuthorized'
import { UserRole } from '@/types/userRole'

export const modRouter = Router()

//@route GET /api/v1/moderator
// @access Private
// @desc Get
modRouter.get(
    '/',
    isAuthenticated,
    isAuthorized(UserRole.Moderator),
    getModerator
)
