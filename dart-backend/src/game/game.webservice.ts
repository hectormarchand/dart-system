import {createError, defineEventHandler, readBody, Router} from "h3";
import {GameService} from "./game.service";
import {CreateDartGameDto, PatchGameDto} from "../common/dart.dtos";
import {DartGame} from "../common/dart.types";

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
        this.ROUTE_PREFIX + "/current-active",
        defineEventHandler(async (event) => {
            const pacthDto: PatchGameDto = JSON.parse(await readBody(event));
            if (!this.isValidGamePatch(pacthDto)) {
                throw createError({statusCode: 400, statusMessage: "Invalid patch request"});
            }

            const currentActiveGame = await this.gameService.getActiveGame();

            if (currentActiveGame == null) {
                throw createError({statusCode: 400, statusMessage: "No current active game"});
            }

            const patchedGame: DartGame = await this.gameService.patchGame(currentActiveGame, pacthDto);
            this.gameService.pushGameToClients(patchedGame);

            return patchedGame;
      })
    )
  }

    private isValidGamePatch(patchDto: PatchGameDto): boolean {
        return patchDto.type === "next-player" || (patchDto.type === "fix-score" && patchDto.fix !== undefined);
    }
}