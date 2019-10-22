# coding: utf-8

from flask import Flask, jsonify, url_for, make_response, request, abort
import pymysql

import logging
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

app = Flask(__name__)

clients = []

class Database:
    def __init__(self):
        host = "127.0.0.1"
        user = "guest"
        password = "guest"
        db = "test"

        self.con = pymysql.connect(host=host, user=user, password=password, db=db, cursorclass=pymysql.cursors.DictCursor)
        self.cur = self.con.cursor()

    def list_clients(self):
        self.cur.execute("SELECT ID as id, NAME as name, BIRTH_DATE as birthDate, EMAIL as email FROM CLI")
        result = self.cur.fetchall()
        return result



@app.route('/v1/clients', methods=['GET'])
def getClients():

    def db_query():
        db = Database()
        clis = db.list_clients()

        return clis

    res = db_query()

    return jsonify({'clients': res})

@app.route('/v1/clients/<int:id>', methods=['GET'])
def getClient(id=0):
    client = [client for client in clients if client['id'] == id]
    print('client: {} {}'.format(id, client))
    if len(client) == 0:
        abort(404)
    return jsonify({'client': client[0]})



@app.errorhandler(404)
def not_found(error):
    return make_response(jsonify({'error': 'Not found'}), 404)

if __name__ == '__main__':
	app.run(host="0.0.0.0", port=9000, threaded=True)
