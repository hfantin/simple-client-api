const http = require('http')
    ,app = require('./config/express')
    ,porta = process.env.PORTA || 3000;

http.createServer(app).listen(porta, () => {
    console.log(`Servidor escutando na porta: ${porta}`);
});
