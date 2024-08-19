import { Router } from 'express'
import { getAdmin } from '@controllers/admin-controller'
import { isAuthenticated } from '@middlewares/isAuthenticated'
import { isAuthorized } from '@middlewares/isAuthorized'
import { UserRole } from '@/types/userRole'

export const adminRouter = Router()

//@route GET /api/v1/admin
// @access Private
// @desc Get
adminRouter.get('/', isAuthenticated, isAuthorized(UserRole.Admin), getAdmin)
