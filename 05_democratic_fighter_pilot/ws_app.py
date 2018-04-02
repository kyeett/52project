from flask import Flask, request, render_template
from gevent import monkey
monkey.patch_all()
from flask_sockets import Sockets
import logging
from werkzeug.serving import run_with_reloader
import uuid
import time
import random
import simplejson as json

gunicorn_error_logger = logging.getLogger('gunicorn.error')

app = Flask(__name__)
app.config['DEBUG'] = True
app.logger.handlers.extend(gunicorn_error_logger.handlers)

sockets = Sockets(app)

logger = logging.getLogger(__name__)
logger.setLevel(logging.DEBUG)



@sockets.route('/echo')
def echo_socket(ws):
   logger.error('User connected 2')
   while not ws.closed:
      message = ws.receive()
      logger.info('Msg received')
      ws.send(message)


db = {}

# Sends the client states to the browsers
@sockets.route('/screen')
def echo_socket(ws):
   val = 50
   logger.info('Screen connected')
   while not ws.closed:
      val += 3 - random.randint(0,6)

      val = max(0, min(100, val))
      time.sleep(0.3)

      ws.send(json.dumps({'value': db.get('knob', val)}))


def format_client(params):

   allowed_types = ["integer"]
   client = {}
   if params.get('type') in allowed_types:
      client['type'] = params.get('type')
      client['name'] = params.get('name', str(uuid.uuid1())[0:8])
      client['scaling'] = 1
      return client
   else:
      raise ValueError("Type is not valid, should be one of %s" % allowed_types)


@sockets.route('/client')
def client_socket(ws):
   logger.info('Client connected')
   client = format_client(ws.environ['werkzeug.request'].args)
   logger.info('Args: %s' % client)
   while not ws.closed:
      message = ws.receive()
      logger.error(message)
      message_json = json.loads(message)
      db['knob'] = message_json['value']
      logger.info('Msg received: %s' % message_json)
      ws.send(message)


@app.route("/")
def home():
   return render_template('index.html')


def run_server():
   from gevent import pywsgi
   from geventwebsocket.handler import WebSocketHandler
   from werkzeug.serving import run_with_reloader
   from werkzeug.debug import DebuggedApplication
   server = pywsgi.WSGIServer(('', 5000), DebuggedApplication(app), handler_class=WebSocketHandler)
   server.serve_forever()


if __name__ == "__main__":
   run_with_reloader(run_server)
