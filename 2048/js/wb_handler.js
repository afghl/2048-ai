function WebsocketHandler() {
    this.conn;
}

WebsocketHandler.prototype.connect = function () {
    console.log("WebsocketHandler.prototype.connect");
    var self = this;
    if (!window["WebSocket"]) {
        alert("Your browser does not support WebSockets.");
        return;
    }
    if (self.conn) {
        return;
    }
    var protocol="ws://";
    if(window.location.protocol=="https:"){
        protocol="wss://";
    }
    self.conn = new WebSocket(protocol+window.location.host+"/ai");
    self.conn.onopen = function(e) {
        console.log("Connection open ..." + e.data);
    };
    self.conn.onmessage = function(evt) {
        console.log( "Received Message: " + evt.data);
    };
    self.conn.onclose = function(evt) {
        console.log(evt);
    };

    console.log("success");
    setTimeout(function(){
        console.log("try to send some message");
        self.conn.send("hehe");
        self.conn.send("hehe2");
        self.conn.send("hehe23");
    }, 2000);


};