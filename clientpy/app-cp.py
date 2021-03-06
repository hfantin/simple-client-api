# coding: utf-8

from flask import Flask, jsonify, url_for, make_response, request, abort, g
# import pymysql
import mysql.connector
from mysql.connector import Error
from mysql.connector.connection import MySQLConnection
from mysql.connector import pooling

import logging
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)

app = Flask(__name__)

clients = []


db_host = "127.0.0.1"
db_user = "guest"
db_password = "guest"
db_name = "test"

# try:
#     g.cnx_pool = mysql.connector.pooling.MySQLConnectionPool(pool_name="pynative_pool",
#                                                                   pool_size=10,
#                                                                   pool_reset_session=True,
#                                                                   host=db_host,
#                                                                   database=db_name,
#                                                                   user=db_user,
#                                                                   password=db_password)

#  https://pynative.com/python-database-connection-pooling-with-mysql/



try:
    connection_pool = mysql.connector.pooling.MySQLConnectionPool(pool_name="pynative_pool",
                                                                  pool_size=10,
                                                                  pool_reset_session=True,
                                                                  host=db_host,
                                                                  database=db_name,
                                                                  user=db_user,
                                                                  password=db_password)

    print ("Printing connection pool properties ")
    print("Connection Pool Name - ", connection_pool.pool_name)
    print("Connection Pool Size - ", connection_pool.pool_size)

    # Get connection object from a pool
    # connection_object = connection_pool.get_connection()


    # if connection_object.is_connected():
    #    db_Info = connection_object.get_server_info()
    #    print("Connected to MySQL database using connection pool ... MySQL Server version on ",db_Info)

    #    cursor = connection_object.cursor()
    #    cursor.execute("select database();")
    #    record = cursor.fetchone()
    #    print ("Your connected to - ", record)

except Error as e :
    print ("Error while connecting to MySQL using Connection pool ", e)
# finally:
    #closing database connection.
    # if(connection_object.is_connected()):
    #     cursor.close()
    #     connection_object.close()
    #     print("MySQL connection is closed")


# @app.before_first_request
# def before_first_request():
#     g.cnx_pool = mysql.connector.pooling.MySQLConnectionPool(pool_name="pynative_pool",
#                                                                   pool_size=10,
#                                                                   pool_reset_session=True,
#                                                                   host=db_host,
#                                                                   database=db_name,
#                                                                   user=db_user,
#                                                                   password=db_password)

@app.route('/v1/clients', methods=['GET'])
def getClients():
    conn = connection_pool.get_connection()
    cursor = conn.cursor()
    cursor.execute("SELECT ID as id, NAME as name, BIRTH_DATE as birthDate, EMAIL as email FROM CLI")
    res = cursor.fetchall()
    # cursor.close()
    # conn.close()
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
