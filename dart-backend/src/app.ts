// Import h3 as npm dependency
import {App, createApp, createRouter, Router, useBase} from "h3";
import { GameWebService } from "./game/game.webservice";
import {StreamWebservice} from "./stream/stream.webservice";
import {DartHitWebService} from "./dart-hit/dart.hit.webservice";
import {WebsocketWebservice} from "./websocket/websocket.webservice";
import {Db} from "mongodb";
import mongoService from "./mongo/mongo.service";
import {GameService} from "./game/game.service";
import {DartHitService} from "./dart-hit/dart.hit.service";
import {StreamService} from "./stream/stream.service";
import {WebsocketService} from "./websocket/websocket.service";

// Create an app instance
export const app: App = createApp();

// Create a new router and register it in app
const router: Router = createRouter();
const apiRouter: Router = createRouter();
router.use("/api/**", useBase("/api", apiRouter.handler));

app.use(router);

createWebServices();

function createWebServices() {
    getMongoInstance()
        .then((mongoDb: Db) => {
            const gameService: GameService = new GameService(mongoDb);
            const gameWebService: GameWebService = new GameWebService(gameService, apiRouter);

            const streamService: StreamService = new StreamService();
            const streamWebService: StreamWebservice = new StreamWebservice(streamService, apiRouter);

            const websocketService: WebsocketService = new WebsocketService();
            const websocketWebService: WebsocketWebservice = new WebsocketWebservice(websocketService);

            const dartHitService: DartHitService = new DartHitService(gameService, websocketService);
            const dartHitWebService: DartHitWebService = new DartHitWebService(dartHitService, apiRouter);

            app.use("/ws", websocketWebService.createWsHandler());
        })
        .catch(e => {
            console.log("Cannot create web services : ", e);
        });
}

async function getMongoInstance(): Promise<Db> {
    return mongoService();
}