import {Collection, Db} from "mongodb";
import { DartGame, DartPlayer } from "../common/dart.types";
import { CreateDartGameDto, CreatePlayerDto } from "../common/dart.dtos";
import {createEventStream, EventStream, H3Event} from "h3";

export class GameService {

    private gamesCollection: Collection<DartGame>;

    private clientEvents: EventStream[];

    constructor(mongoDb: Db) {
        this.gamesCollection = mongoDb.collection<DartGame>("games");
        this.clientEvents = [];
    }

    async getAllGames(): Promise<DartGame[]> {
        return (await this.gamesCollection.find().toArray()) as DartGame[];
    }

    async createGame(createGame: CreateDartGameDto): Promise<DartGame> {
        const today: Date = new Date();

        const players: DartPlayer[] | undefined = createGame.players?.map((playerDto: CreatePlayerDto) => {
            return {
                name: playerDto.name,
                nbOfDartsThrownThisRound: 0,
                ppd: 0,
                score: 0,
                scoreThisRound: 0,
            } as DartPlayer
        });

        const game: DartGame = {
            players: players ? players : [],
            gameType: createGame.gameType ? createGame.gameType : "501",
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
        await this.gamesCollection.replaceOne({ _id: dartGame.id }, dartGame);
        return dartGame;
    }

    async patchGame(patch: any): Promise<DartGame | null> {
        // TODO : implement this method (fix from client or)
        return null;
    }

    /**
     * Send the game to clients through SSE
     * @param dartGame the dart game
     */
    pushGameToClients(dartGame: DartGame): void {
        this.clientEvents.forEach(async ce => await ce.push(JSON.stringify(dartGame)));
    }

}