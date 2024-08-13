import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'
// import { validationResult } from 'express-validator'

export const registerUser = asyncHandler(
    async (_: Request, res: Response): Promise<void> => {
        // const errors = validationResult(req)
        // if (!errors.isEmpty()) {
        //     res.status(400).json({ errors: errors.array() })
        //     return
        // }

        res.send('Hello World!!!')
    }
)
