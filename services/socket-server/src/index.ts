import { WebSocketServer, WebSocket } from "ws";
import http from "http";

const server = http.createServer();
const wss = new WebSocketServer({ server: server });

const subscriptions: {
  [key: string]: {
    ws: WebSocket;
    rooms: string[];
  };
} = {};

setInterval(() => {
  console.log(subscriptions);
}, 3000);

wss.on("connection", (ws) => {
  const id = randomId();
  subscriptions[id] = {
    ws: ws,
    rooms: [],
  };

  ws.on("message", function message(data) {
    const parsedMessage = JSON.parse(data as unknown as string);
    if (parsedMessage.type === "SUBSCRIBE") {
      subscriptions[id].rooms.push(parsedMessage.room);
    }

    if (parsedMessage.type == "sendMessage") {
      const message = parsedMessage.message;
      const roomId = parsedMessage.roomId;

      // subscriptions.forEach((ws,rooms) => {
      //     if(rooms.includes(roomId)){
      //         ws.send(message)
      //     }
      // })

      Object.keys(subscriptions).forEach((userId) => {
        const { ws, rooms } = subscriptions[userId];
        if (rooms.includes(roomId)) {
          ws.send(message);
        }
      });
    }
  });
});

function randomId() {
  return Math.random();
}

server.listen(8080);
