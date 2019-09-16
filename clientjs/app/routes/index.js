/* Código simplório, apenas para fornecer o serviço para a aplicação */

var api = require('../api');

module.exports  = (app) => {
   
    app.route('/v1/clients').get(api.getClients);

};