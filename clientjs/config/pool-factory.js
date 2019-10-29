const mysql = require('mysql')

const host = process.env.DB_HOST || 'localhost';
const user = process.env.DB_USER || 'guest';
const password = process.env.DB_PASS || 'guest';
const database = process.env.DB_NAME || 'test';
const connectionLimit = process.env.DB_POOL_LIMIT || 50;

const pool = mysql.createPool({
    host,
    user,
    password,
    database,
    connectionLimit
});

// pool.connect((err) => {
//     if (err) throw err;
// });


//pool.on('release', () => console.log('pool => conexão retornada'));

process.on('SIGINT', () =>
    pool.end(err => {
        if(err) return logger.error(err);
        console.log('pool => fechado');
        process.exit(0);
    })
);

module.exports = pool;
