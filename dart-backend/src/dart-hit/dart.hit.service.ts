
import { DartGame, DartPlayer } from "../common/dart.types";
import { DartHitDto } from "../common/dart.dtos";
import { GameService } from "../game/game.service";
import {createEventStream, EventStream, H3Event} from "h3";
import {GameBusinessService} from "../game-business/game-business.service";

export class DartHitService {

    private gameService: GameService;

    private gameBusinessService: GameBusinessService;

    private clientEvents: EventStream[];

    constructor(gameService: GameService, gameBusinessService: GameBusinessService) {
        this.gameService = gameService;
        this.gameBusinessService = gameBusinessService;
        this.clientEvents = [];
    }

    async addDartHit(dartHitDto: DartHitDto): Promise<DartGame | null> {
        // if game active, update game
        let currentActiveGame: DartGame | null = await this.gameService.getActiveGame();
        if (currentActiveGame) {
            currentActiveGame = this.updateGameAfterDartHit(dartHitDto, currentActiveGame);
            await this.gameService.updateGame(currentActiveGame);
            this.gameService.pushGameToClients(currentActiveGame);
        }

        this.pushDartHitToClient(dartHitDto);

        return currentActiveGame;
    }

    private updateGameAfterDartHit(dartHitDto: DartHitDto, dartGame: DartGame): DartGame {
        const currentPlayer: DartPlayer = dartGame.players[dartGame.currentPlayerIndex];
        currentPlayer.nbOfDartsThrownThisRound ++;

        // Update player ppd and player score
        currentPlayer.score -= dartHitDto.estimatedPoints;
        currentPlayer.scoreThisRound += dartHitDto.estimatedPoints;
        currentPlayer.scoreIncludingBust += dartHitDto.estimatedPoints;
        currentPlayer.nbOfDartsThrownTotal++;
        currentPlayer.ppd = currentPlayer.scoreIncludingBust / currentPlayer.nbOfDartsThrownTotal;

        // Update the game in general
        dartGame = this.gameBusinessService.executeOneStepInGame(dartGame);

        return dartGame;
    }

    public streamDartHits(event: H3Event): void {
        const eventStream: EventStream = createEventStream(event);

        eventStream.onClosed(() => {
            this.clientEvents = this.clientEvents.filter(ev => ev !== eventStream);
        })

        this.clientEvents.push(eventStream);

        eventStream.send();

        eventStream.push("{}")
    }

    private pushDartHitToClient(dartHit: DartHitDto): void {
        this.clientEvents.forEach(async ce => await ce.push(JSON.stringify(dartHit)));
    }


}