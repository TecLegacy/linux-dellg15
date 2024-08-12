import { Router } from 'express'
import { registerUser } from '../controller/user-controller'

export const router = Router()

router.get('/register', registerUser)
