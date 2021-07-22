const { Client, Pool } = require('pg')

class DBClient {
    constructor() {
        console.log("DBClient initializing")
        this.pool = new Pool()
    }

    disconnect() {
        console.log("db pool disconnection")
        this.pool.end()
    }

    async select(query) {
        var allRecords = []
        await this.pool.connect()
            .then(client => {
                return client.query(query)
                    .then(records => {
                        allRecords = records.rows
                        client.release()
                    })
                    .catch(error => {
                        client.release()
                        return Promise.reject([])
                    })
            })
        
        console.log(allRecords)
        return Promise.resolve(allRecords)
    }

    async insert(query) {
        await this.pool.connect()
            .then(client => {
                return client.query(query)
                    .then(records => {
                        console.log("insertion success")
                        client.release()
                    })
                    .catch(error => {
                        client.release()
                        return Promise.reject([])
                    })
            })
        
    }

}

module.exports = DBClient
