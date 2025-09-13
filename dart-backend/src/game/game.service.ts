import {Collection, Db} from "mongodb";
import {DartGame, DartGameType, DartPlayer} from "../common/dart.types";
import {CreateDartGameDto, CreatePlayerDto, PatchGameDto} from "../common/dart.dtos";
import {createEventStream, EventStream, H3Event} from "h3";
import {GameBusinessService} from "../game-business/game-business.service";

export class GameService {

    private gamesCollection: Collection<DartGame>;

    private gameBusinessService: GameBusinessService;

    private clientEvents: EventStream[];

    constructor(mongoDb: Db, gameBusinessService: GameBusinessService) {
        this.gamesCollection = mongoDb.collection<DartGame>("games");
        this.gameBusinessService = gameBusinessService;
        this.clientEvents = [];
    }

    async getAllGames(): Promise<DartGame[]> {
        return (await this.gamesCollection.find().toArray()) as DartGame[];
    }

    async createGame(createGame: CreateDartGameDto): Promise<DartGame> {
        // Unactive all other games
        await Promise.all(
          (await this.getAllGames()).map(async otherGame => {
              console.log(otherGame)
              otherGame.active = false;
              return this.updateGame(otherGame);
          })
        )

        const today: Date = new Date();

        const totalScoreToDo = this.gameTypeToScore(createGame.gameType);

        const players: DartPlayer[] | undefined = createGame.players?.map((playerDto: CreatePlayerDto) => {
            return {
                name: playerDto.name,
                nbOfDartsThrownThisRound: 0,
                nbOfDartsThrownTotal: 0,
                ppd: 0,
                score: totalScoreToDo,
                scoreIncludingBust: 0,
                scoreThisRound: 0,
            } as DartPlayer
        });

        const game: DartGame = {
            players: players ? players : [],
            winner: undefined,
            gameType: createGame.gameType,
            totalRound: createGame.totalRound ? createGame.totalRound : 15,
            currentRound: 1,
            active: true,
            createdDate: `${today.getFullYear()}-${today.getMonth() + 1}-${today.getDate()} ${today.getHours()}:${today.getMinutes()}`,
            currentPlayerIndex: 0
        }

        await this.gamesCollection.insertOne(game);
        return game;
    }

    async getActiveGame(): Promise<DartGame | null> {
        const activeGames: DartGame[] = (await this.gamesCollection.find({
            active: true
        }).toArray()) as DartGame[];

        if (activeGames.length > 1) {
            console.warn("WARNING : More than one game is active");
        }

        return activeGames.length ? activeGames[0] : null;
    }

    async streamActiveGame(event: H3Event): Promise<void> {
        const eventStream: EventStream = createEventStream(event);

        eventStream.onClosed(() => {
            this.clientEvents = this.clientEvents.filter(ev => ev !== eventStream);
        })

        this.clientEvents.push(eventStream);

        eventStream.send();

        const activeGame: DartGame | null = await this.getActiveGame();
        await eventStream.push(activeGame ? JSON.stringify(activeGame): "{}");
    }

    async updateGame(dartGame: DartGame): Promise<DartGame> {
        await this.gamesCollection.replaceOne({_id: dartGame._id}, dartGame);
        return dartGame;
    }

    async patchGame(currentActiveGame: DartGame, patch: PatchGameDto): Promise<DartGame> {
        switch (patch.type) {
            case "next-player":
                currentActiveGame = this.patchGameNextPlayer(currentActiveGame);
                break;
            case "fix-score":
                // currentActiveGame = this.patchGameFixScore(currentActiveGame);
                break;
        }

        currentActiveGame = await this.updateGame(currentActiveGame);

        return currentActiveGame;
    }

    private patchGameNextPlayer(game: DartGame): DartGame {
        const currentPlayer: DartPlayer = game.players[game.currentPlayerIndex];

        // Update player ppd and player score
        currentPlayer.nbOfDartsThrownThisRound = 3;
        currentPlayer.nbOfDartsThrownTotal = Math.ceil((currentPlayer.nbOfDartsThrownTotal + 1) / 3) * 3;
        currentPlayer.ppd = currentPlayer.scoreIncludingBust / currentPlayer.nbOfDartsThrownTotal;

        return this.gameBusinessService.executeOneStepInGame(game);
    }

    // TODO
    private patchGameFixScore(game: DartGame): void {
    }

    /**
     * Send the game to clients through SSE
     * @param dartGame the dart game
     */
    pushGameToClients(dartGame: DartGame): void {
        this.clientEvents.forEach(async ce => await ce.push(JSON.stringify(dartGame)));
    }

    private gameTypeToScore(gameType: DartGameType): number {
        switch (gameType) {
            case "301":
                return 301;
            case "501":
                return 501;
            case  "701":
                return 701;
        }
    }

}