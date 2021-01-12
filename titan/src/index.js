const util = require('util')
const client = require("./client/client-grpc")

async function main() {
    const token = await login("teste3@teste.com", "Jão")
    if (!token) return
    console.log(token)

    // const token = await login("teste", "teste");
    // console.log(token)

    await authenticate(token)
}

async function register(email, name, password) {
    const register = util.promisify(client.register).bind(client)

    const start = process.hrtime.bigint()

    let tokenReturn

    try {
        const { token, _, error } = await register({ email, name, password })
        if (error) {
            console.error(error)
            return
        }
        tokenReturn = token
    } catch (error) {
        console.error(error)
    }

    const end = process.hrtime.bigint()
    const ms = (end - start) / BigInt(1000000)
    console.log(`Benchmark took ${ms} ms`)

    return tokenReturn
}

async function login(email, password) {
    const login = util.promisify(client.login).bind(client)

    const start = process.hrtime.bigint()

    let tokenReturn = ""

    try {
        const { token, error } = await login({ email, password })
        if (error) {
            console.error(error)
            return
        }
        tokenReturn = token
    } catch (error) {
        console.error(error)
    }

    const end = process.hrtime.bigint()
    const ms = (end - start) / BigInt(1000000)
    console.log(`Benchmark took ${ms} ms`)

    return tokenReturn
}

async function authenticate(token) {
    const authenticate = util.promisify(client.authenticate).bind(client)

    const start = process.hrtime.bigint()

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