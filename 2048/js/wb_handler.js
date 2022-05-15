function Websocket() {
  this.conn;
  this.onMessageCallBack;
}

Websocket.prototype.connect = function () {
  console.log("WebsocketHandler.prototype.connect");
  var self = this;
  if (!window["WebSocket"]) {
    alert("Your browser does not support WebSockets.");
    return;
  }
  if (self.conn) {
    return;
  }
  var protocol = "ws://";
  if (window.location.protocol == "https:") {
    protocol = "wss://";
  }
  self.conn = new WebSocket(protocol + window.location.host + "/ai");
  self.conn.onopen = function (e) {
    console.log("Connection open ..." + e.data);
  };

  self.conn.onclose = function (evt) {
    console.log(evt);
  };
  console.log("cnonect success");
};

Websocket.prototype.onMessage = function(f) {
  if (!this.conn) {
    console.error("not connect");
    return;
  }
  this.conn.onmessage = f;
};

Websocket.prototype.send = function(data) {
  this.conn.send(data);
};

Websocket.prototype.close = function () {
  this.conn.close();
};