import {WebsocketService} from "./websocket.service";

export class WebsocketWebservice {

    private websocketService: WebsocketService;

    constructor(websocketService: WebsocketService) {
        this.websocketService = websocketService;
    }

    createWsHandler() {
        return this.websocketService.createWsHandler();
    }
}