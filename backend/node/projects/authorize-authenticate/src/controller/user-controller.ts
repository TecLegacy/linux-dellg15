import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'

export const registerUser = asyncHandler(async (_: Request, res: Response) => {
    res.send('Hello World!!!sssss')
})
