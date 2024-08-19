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

// Blacklist model
interface BlackListAttr {
    accessToken: string
    userId: mongoose.Types.ObjectId
    exp: number
}

interface BlackListDocument extends mongoose.Document {
    accessToken: string
    userId: mongoose.Types.ObjectId
    exp: number
}

interface BlackListModel extends mongoose.Model<BlackListDocument> {
    build(attr: BlackListAttr): BlackListDocument
}

const BlackListSchema = new mongoose.Schema({
    accessToken: {
        type: String,
        required: true,
    },
    userId: {
        type: mongoose.Types.ObjectId,
        required: true,
    },
    exp: {
        type: Number,
        required: true,
    },
})

BlackListSchema.statics.build = (attr: BlackListAttr) => {
    return new Blacklist(attr)
}

export const Blacklist = mongoose.model<BlackListDocument, BlackListModel>(
    'BlackList',
    BlackListSchema
)
