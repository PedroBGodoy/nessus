const protoLoader = require('@grpc/proto-loader')
const grcp = require('grpc')
const path = require('path')

const packageDefinition = protoLoader.loadSync(path.join(__dirname, '../protos/messages.proto'), {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
})
const UserService = grcp.loadPackageDefinition(packageDefinition).UserService

const client = new UserService(
    '0.0.0.0:50051',
    grcp.credentials.createInsecure()
)

module.exports = client