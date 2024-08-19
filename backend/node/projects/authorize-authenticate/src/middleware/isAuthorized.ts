import { NotFoundError } from '@/errors/NotFoundError'
import { Unauthorized } from '@/errors/UnauthorizedError'
import { User } from '@/models/user-model'
import { UserRole } from '@/types/userRole'
import { Request, Response, NextFunction } from 'express'
import asyncHandler from 'express-async-handler'
import { JwtPayload } from 'jsonwebtoken'

export const isAuthorized = (role: UserRole.Admin | UserRole.Moderator) =>
    asyncHandler(async (req: Request, _: Response, next: NextFunction) => {
        const { id } = req.user as JwtPayload

        const user = await User.findById({ _id: id })

        if (!user) {
            throw new NotFoundError('User not found')
        }

        if (!user.role.includes(role)) {
            throw new Unauthorized('User not authorized')
        }

        next()
    })
