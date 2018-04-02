from flask import Flask, request
from flask_sockets import Sockets
import logging


app = Flask(__name__)
sockets = Sockets(app)

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)


@sockets.route('/echo')
def echo_socket(ws):
   logger.error('User connected')
   while not ws.closed:
      message = ws.receive()
      logger.info('Msg received')
      ws.send(message)


@sockets.route('/client')
def client_socket(ws):
   from pprint import pprint

   logger.error('User connected')
   while not ws.closed:
      message = ws.receive()
      logger.info('Msg received')
      ws.send(message)



@app.route('/')
def hello():
   return 'Hello World!'


if __name__ == "__main__":
   from gevent import pywsgi
   from geventwebsocket.handler import WebSocketHandler
   from werkzeug.serving import run_with_reloader
   from werkzeug.debug import DebuggedApplication
   server = pywsgi.WSGIServer(('', 5000), DebuggedApplication(app), handler_class=WebSocketHandler)
   server.serve_forever()
