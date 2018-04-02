function ReceiveTestSocket(url)
{
   if ("WebSocket" in window)
   {
      console.log("WebSocket is supported by your Browser!");

      // Let us open a web socket
      var ws = new WebSocket(url);

      ws.onopen = function()
      {
         // Web Socket is connected, send data using send()
      };

      ws.onmessage = function (evt)
      {
         var received_msg = JSON.parse(evt.data);
         // console.log("Message is received..." + received_msg);
         // console.log(typeof received_msg);
         players[1].y = received_msg['value'];
      };

      ws.onclose = function()
      {
         // websocket is closed.
         console.log("Connection is closed...");
      };

      window.onbeforeunload = function(event) {
         socket.close();
      };

      return ws
   }

   else
   {
      // The browser doesn't support WebSocket
      alert("WebSocket NOT supported by your Browser!");
   }
}

function WebSocketTest(url)
{
   if ("WebSocket" in window)
   {
      console.log("WebSocket is supported by your Browser!");

      // Let us open a web socket
      var ws = new WebSocket(url);

      ws.onopen = function()
      {
         // Web Socket is connected, send data using send()
         ws.send("Message to send");
         console.log("Message is sent...");
      };

      ws.onmessage = function (evt)
      {
         var received_msg = JSON.parse(evt.data);
         // console.log("Message is received..." + received_msg);
         // console.log(typeof received_msg);
         players[1].y = received_msg['value'];
      };

      ws.onclose = function()
      {
         // websocket is closed.
         console.log("Connection is closed...");
      };

      window.onbeforeunload = function(event) {
         socket.close();
      };
   }

   else
   {
      // The browser doesn't support WebSocket
      alert("WebSocket NOT supported by your Browser!");
   }
}