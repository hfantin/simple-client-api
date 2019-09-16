const mysql = require('mysql')

const pool = mysql.createPool({
    host: 'localhost',
    user: 'guest',
    password: 'guest',
    database: 'test',
    connectionLimit : 50
});

// pool.connect((err) => {
//     if (err) throw err;
// });


pool.on('release', () => console.log('pool => conexão retornada'));

process.on('SIGINT', () =>
    pool.end(err => {
        if(err) return logger.error(err);
        console.log('pool => fechado');
        process.exit(0);
    })
);

module.exports = pool;
