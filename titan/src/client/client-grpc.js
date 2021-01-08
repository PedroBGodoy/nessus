const protoLoader = require('../../node_modules2/@grpc/proto-loader')
const grcp = require('../../node_modules2/grpc')

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