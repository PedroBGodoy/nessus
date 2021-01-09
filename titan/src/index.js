const util = require('util')
const client = require("./client/client-grpc")

async function main() {
    const authenticate = util.promisify(client.authenticate).bind(client)

    const start = process.hrtime.bigint()

    const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RlIiwiZXhwIjoxNjEwMTUwODk1fQ.PVX4xzMBIPHkolyX16KJpdCBON7tLa3kkmZaDgzVd1s"

    try {
        const { user, error } = await authenticate({ token })
        if (error) {
            console.error(error)
            return
        }
        console.log(user)
    } catch (error) {
        console.error(error)
    }

    const end = process.hrtime.bigint()
    const ms = (end - start) / BigInt(1000000)
    console.log(`Benchmark took ${ms} ms`)
}

main()