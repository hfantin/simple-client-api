module.exports = pool => (req, res, next) => {
    pool.getConnection((err, connection) => {
        if(err) return next(err);
        // adicionou a conexão na requisição
        req.connection = connection;
        // passa a requisição o próximo middleware
        next();
        // devolve a conexão para o pool no final da resposta
        res.on('finish', () => req.connection.release());
    });
};