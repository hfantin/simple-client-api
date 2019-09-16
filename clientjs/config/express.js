const express = require('express')
    , app = express()
    , routes = require('../app/routes')
    , pool = require('./pool-factory')
    , connectionMiddleware = require('./connection-middleware')
    , bodyParser = require('body-parser');

app.use(bodyParser.json());

app.use(connectionMiddleware(pool));

app.use((req, res, next) => {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept");
    next();
});

routes(app);

module.exports = app;