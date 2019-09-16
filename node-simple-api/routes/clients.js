var express = require('express');
var router = express.Router();

/* GET users listing. */
router.get('/', function(req, res, next) {
  res.locals.connection.query('SELECT id, name, birth_date, email from CLI', function (error, results, fields) {
		if(error){
      res.send(JSON.stringify({"status": 500, "error": error, "data": null})); 
    }else{
      res.send(JSON.stringify({
        //"status": 200, "error": null, 
        "data": results}));
    }
	});
});

module.exports = router;
