import {defineWebSocketHandler} from "h3";
import {DartGame} from "../common/dart.types";
import {DartHitDto} from "../common/dart.dtos";

export class WebsocketService {

    private peers!: any[];

    constructor() {
        this.peers = [];
    }

    createWsHandler() {
        return defineWebSocketHandler({
            open: (peer) => {
                //console.log("WS connection open for : " + peer);
                this.peers.push(peer);
            },
            message: (peer, message) => {},
            close: (peer)=> {
                //console.log("WS connection closed for : " + peer);
                this.peers = this.peers.filter(p => p !== peer);
            },
        })
    }

    sendDartGameToPeers(dartGame: DartGame) {
        const payload = JSON.stringify({
            type: "dart-game",
            data: dartGame
        });
        this.peers.forEach(ws => {
            ws.send(payload);
        });
    }

    sendDartHitToPeers(dartHitDto: DartHitDto) {
        const payload = JSON.stringify({
            type: "dart-hit",
            data: dartHitDto
        });
        this.peers.forEach(ws => {
            ws.send(payload);
        })
    }
}