import {defineEventHandler, readBody, Router} from "h3";
import {GameService} from "./game.service";
import {CreateDartGameDto} from "../common/dart.dtos";

export class GameWebService {

  private readonly ROUTE_PREFIX = "/games";

  private gameService: GameService;

  constructor(gameService: GameService, apiRouter: Router) {
    this.gameService = gameService;
    this.initRoutes(apiRouter);
  }

  initRoutes(router: Router) {
    router.get(
      this.ROUTE_PREFIX,
      defineEventHandler(async (event) => {
        return await this.gameService.getAllGames();
      })
    );

    router.post(
      this.ROUTE_PREFIX,
      defineEventHandler(async (event) => {
        try {
          const gameDto: CreateDartGameDto = JSON.parse(await readBody(event));
          return await this.gameService.createGame(gameDto);
        } catch (e) {
          console.error("Error when creating the game :", e);
          throw e;
        }
      })
    );

    router.get(
      this.ROUTE_PREFIX + "/current-active",
        defineEventHandler(async (event) => {
          try {
            return await this.gameService.getActiveGame();
          } catch (e) {
            console.error("Error when getting the current active game", e);
            throw e;
          }
        })
    );

    router.get(
      this.ROUTE_PREFIX + "/stream-current-active",
      defineEventHandler(async (event) => {
        return this.gameService.streamActiveGame(event);
      })
    )
  }
}