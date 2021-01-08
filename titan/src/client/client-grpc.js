const protoLoader = require('@grpc/proto-loader')
const grcp = require('grpc')

const packageDefinition = protoLoader.loadSync('./pb/messages.proto', {
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