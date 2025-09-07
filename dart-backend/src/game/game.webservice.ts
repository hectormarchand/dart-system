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
      defineEventHandler((event) => {
        return this.gameService.streamActiveGame(event);
      })
    )

    router.patch(
      this.ROUTE_PREFIX + "/:id",
      defineEventHandler((event) => {
        // TODO, game patch from the client (wrong score)
      })
    )

    router.patch(
      this.ROUTE_PREFIX + "/:id/next-player",
      defineEventHandler((event) => {
        // TODO, patch from the client (press space to change player)
      })
    )
  }
}