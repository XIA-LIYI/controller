const { Server } = require('http');

let net = require('net')
let count = 0;

server = net.createServer();

var users = [];

server.on('connection', function (socket) {
    users.push(socket);
    count++;
    socket.setEncoding('utf8');
    console.log('Connected!');
    socket.on('data', function (data) {
        data = data.toString().trim();

        users.forEach(function (client) {
            if (client !== socket) {
                client.write(client.remotePort + ":" + data);
            }
        })
    })
});

server.on('listening', function () {
    console.log('Listening');
});

server.listen(8000, '127.0.0.1');

