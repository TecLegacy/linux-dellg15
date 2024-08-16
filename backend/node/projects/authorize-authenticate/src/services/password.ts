import bcrypt from 'bcrypt'

export class Password {
    static async hashPassword(password: string): Promise<string> {
        const salt = await bcrypt.genSalt(10)
        const hashPassword = await bcrypt.hash(password, salt)
        return hashPassword
    }

    static async comparePassword(
        password: string,
        hashedPassword: string
    ): Promise<boolean> {
        const result = await bcrypt.compare(password, hashedPassword)
        return result
    }
}
