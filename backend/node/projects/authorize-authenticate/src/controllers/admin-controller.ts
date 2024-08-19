import { Request, Response } from 'express'

export const getAdmin = async (_: Request, res: Response) => {
    res.status(200).json({
        message: 'Admin route',
    })
}
