const dao = {};
const query = 'SELECT * from CLI';

dao.getClients = (connection, params) => {
    return new Promise((resolve, reject) => {
        try {
            connection.query(query, [], (err, lista) => {
                if (err) reject({error: err.sqlMessage});
                // logger.info(`lista: ${lista.length}`);
                resolve(lista)
            })
        } catch (e) {
            reject({error: e.message});
        }
    });
};

module.exports = dao;

