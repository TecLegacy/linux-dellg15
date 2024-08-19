import mongoose from 'mongoose'

interface DocumentAttrs {
    refreshToken: string
    userId: mongoose.Types.ObjectId
}

interface RefreshTokenDocument extends mongoose.Document {
    refreshToken: string
    userId: mongoose.Types.ObjectId
}

interface RefreshTokenModel extends mongoose.Model<RefreshTokenDocument> {
    build(attr: DocumentAttrs): RefreshTokenDocument
}

const refreshTokenSchema = new mongoose.Schema({
    refreshToken: {
        type: String,
        required: true,
    },
    userId: {
        type: mongoose.Types.ObjectId,
        required: true,
    },
})

refreshTokenSchema.statics.build = (attr: DocumentAttrs) => {
    return new RefreshToken(attr)
}

export const RefreshToken = mongoose.model<
    RefreshTokenDocument,
    RefreshTokenModel
>('RefreshToken', refreshTokenSchema)
