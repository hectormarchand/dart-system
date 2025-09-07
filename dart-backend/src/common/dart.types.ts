import { ObjectId } from "mongodb";

export type DartGameType = "301" | "501" | "701";

export class DartPlayer {
    name!: string;
    score!: number;
    scoreThisRound!: number;
    ppd!: number;
    nbOfDartsThrownThisRound!: number;
}

export class DartGame {
    _id?: ObjectId;
    players!: DartPlayer[];
    gameType!: DartGameType;
    totalRound!: number;
    currentRound!: number;
    active!: boolean;
    createdDate!: string;
    currentPlayerIndex!: number;
}