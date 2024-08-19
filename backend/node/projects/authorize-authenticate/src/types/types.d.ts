import { JwtPayload } from 'jsonwebtoken'

declare global {
    namespace Express {
        interface Request {
            user: string | JwtPayload
            accessToken?: { value: string; exp: number }
        }
    }
}
