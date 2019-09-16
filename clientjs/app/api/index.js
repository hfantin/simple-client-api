const api = {};
const dao = require('../dao/dao');

api.getClients = (req, res) => {
    dao.getClients(req.connection, req.query)
        .then(lista => res.status(200).json(lista))
        .catch(err => res.status(400).json(err));
};

module.exports = api;