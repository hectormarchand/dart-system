import { ObjectId } from "mongodb";

export type DartGameType = "301" | "501" | "701";

export class DartPlayer {
    name!: string;
    score!: number;
    scoreIncludingBust!: number;
    scoreThisRound!: number;
    ppd!: number;
    nbOfDartsThrownThisRound!: number;
    nbOfDartsThrownTotal!: number
}

export class DartGame {
    _id?: ObjectId;
    players!: DartPlayer[];
    winner: DartPlayer | undefined;
    gameType!: DartGameType;
    totalRound!: number;
    currentRound!: number;
    active!: boolean;
    createdDate!: string;
    currentPlayerIndex!: number;
}