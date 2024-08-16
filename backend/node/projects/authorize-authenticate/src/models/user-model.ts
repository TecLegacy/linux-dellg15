import mongoose from 'mongoose'

interface UserAttrs {
    username: string
    email: string
    password: string
}

interface UserDocument extends mongoose.Document {
    username: string
    email: string
    password: string
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
        createdAt: { type: Date, default: Date.now },
        updatedAt: { type: Date, default: Date.now },
    },
    {
        toJSON: {
            transform(_, ret) {
                ret.id = ret._id
                delete ret._id
                delete ret.__v
            },
            versionKey: false,
        },
    }
)

userSchema.pre('save', async function (next) {
    if (this.isModified('password')) {
        // hash password using bcrypt
        // set password to hashed password
    }
    next()
})

userSchema.statics.build = (attrs: UserAttrs) => {
    return new User(attrs)
}

export const User = mongoose.model<UserDocument, UserModel>('User', userSchema)
