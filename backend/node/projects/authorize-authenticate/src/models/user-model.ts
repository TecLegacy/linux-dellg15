import { Password } from '@/services/password'
import { UserRole } from '@/types/userRole'
import mongoose from 'mongoose'

export interface UserAttrs {
    username: string
    email: string
    password: string
    role: UserRole.Member | UserRole.Admin | UserRole.Moderator
}

interface UserDocument extends mongoose.Document {
    username: string
    email: string
    password: string
    role: UserRole.Member | UserRole.Admin | UserRole.Moderator
    createdAt: Date
    updatedAt: Date
}

interface UserModel extends mongoose.Model<UserDocument> {
    build(attrs: UserAttrs): UserDocument
}

const userSchema = new mongoose.Schema(
    {
        username: { type: String, required: true },
        email: { type: String, required: true, unique: true },
        password: { type: String, required: true },
        role: {
            type: String,
            required: true,
            enum: [UserRole.Member, UserRole.Admin, UserRole.Moderator],
        },
        createdAt: { type: Date, default: Date.now },
        updatedAt: { type: Date, default: Date.now },
    },
    {
        toJSON: {
            transform(_, ret) {
                ret.id = ret._id
                delete ret._id
                delete ret.password
            },
            versionKey: false,
        },
    }
)

userSchema.pre('save', async function (next) {
    if (this.isModified('password')) {
        const hashedPassword = await Password.hashPassword(this.get('password'))
        this.set('password', hashedPassword)
    }
    next()
})

userSchema.statics.build = (attrs: UserAttrs) => {
    return new User(attrs)
}

export const User = mongoose.model<UserDocument, UserModel>('User', userSchema)
