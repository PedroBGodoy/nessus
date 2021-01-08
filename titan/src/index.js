const util = require('util')
const client = require("./client/client-grpc")

async function main() {
    const authenticate = util.promisify(client.authenticate).bind(client)

    const start = process.hrtime.bigint()

    const { user, error } = await authenticate("123")
    if (error) {
        console.error(error)
        return
    }
    console.log(user)

    const end = process.hrtime.bigint()
    const ms = (end - start) / BigInt(1000000)
    console.log(`Benchmark took ${ms} ms`)
}

main()