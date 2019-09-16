let dataUtil = {};

dataUtil.paraTexto = (data) => {
    return `${data.getFullYear()}-${data.getMonth() + 1}-${data.getDate()}`;
};

dataUtil.paraData = (texto) => {
    if (!/\d{4}-\d{2}-\d{2}/.test(texto)) throw new Error('data invalida!');
    return new Date(...texto.split('-').map((item, indice) => item - indice % 2));
};

module.exports = dataUtil;