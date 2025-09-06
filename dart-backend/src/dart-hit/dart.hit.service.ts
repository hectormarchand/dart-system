
import { DartGame, DartPlayer } from "../common/dart.types";
import { DartHitDto } from "../common/dart.dtos";
import { GameService } from "../game/game.service";
import {WebsocketService} from "../websocket/websocket.service";

export class DartHitService {

    private gameService: GameService;

    private websocketService: WebsocketService;

    constructor(gameService: GameService,
                websocketService: WebsocketService) {
        this.gameService = gameService;
        this.websocketService = websocketService;
    }

    async addDartHit(dartHitDto: DartHitDto): Promise<DartGame | null> {
        // if game active, update game
        let currentActiveGame: DartGame | null = await this.gameService.getActiveGame();
        if (currentActiveGame) {
            currentActiveGame = await this.updateGameAfterDartHit(dartHitDto, currentActiveGame);
            await this.gameService.updateGame(currentActiveGame);
            this.websocketService.sendDartGameToPeers(currentActiveGame);
        }

        this.websocketService.sendDartHitToPeers(dartHitDto);

        return currentActiveGame;
    }

    // TODO : maybe move this method to buisinessDartService
    private async updateGameAfterDartHit(dartHitDto: DartHitDto, dartGame: DartGame): Promise<DartGame> {
        const THROW_PER_ROUND = 3;

        const currentPlayer: DartPlayer = dartGame.players[dartGame.currentPlayerIndex];
        currentPlayer.nbOfDartsThrownThisRound ++;

        // Update player ppd and player score
        currentPlayer.score += dartHitDto.estimatedPoints;
        const totalDartThrown: number = (dartGame.currentRound - 1) * THROW_PER_ROUND + currentPlayer.nbOfDartsThrownThisRound;
        currentPlayer.ppd = currentPlayer.score / totalDartThrown;

        // Check if it is time to change the player
        if (currentPlayer.nbOfDartsThrownThisRound >= THROW_PER_ROUND) {
            if (dartGame.currentPlayerIndex === dartGame.players.length - 1) {
                dartGame.currentRound ++;
                dartGame.players.forEach(player => player.nbOfDartsThrownThisRound = 0);
            }

            dartGame.currentPlayerIndex = (dartGame.currentPlayerIndex + 1) % dartGame.players.length;
        }

        return dartGame;
    }


}