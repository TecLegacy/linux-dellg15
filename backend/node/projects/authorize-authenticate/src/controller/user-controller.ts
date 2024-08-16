import asyncHandler from 'express-async-handler'

import type { Request, Response } from 'express'
import { validationResult } from 'express-validator'
import { noUnusedVars } from '@/utils/dump-funcs'
import { connect } from '@/db/connection'
import { Db } from 'mongodb'
import { DatabaseError } from '@/errors/DatabaseError'

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

        // check if user exists
        const db: Db | null = await connect()
        if (!db) {
            throw new DatabaseError('Internal server error')
        }

        const user = await db.collection('users').insertOne({
            username: 'cooluser',
            password: 'coolpassword',
        })

        console.log('User created', user)
        res.send('Successfully created user')
        return
        noUnusedVars(req)
    }
)
