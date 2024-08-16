import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'
import { validationResult } from 'express-validator'
import { noUnusedVars } from '@/utils/dump-funcs'

export const registerUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        const errors = validationResult(req)
        if (!errors.isEmpty()) {
            res.status(400).json({ errors: errors.array() })
            return
        }

        res.send('Hello World!!!')
    }
)

export const createUser = asyncHandler(
    async (req: Request, res: Response): Promise<void> => {
        //validate user

        // check if user already exists

        // create user

        // send response

        res.send('Hello World!!!')
        return
        noUnusedVars(req)
    }
)
