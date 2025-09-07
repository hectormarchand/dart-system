import { DartGameType } from "./dart.types"

export class CreatePlayerDto {
    name?: string;
}

export interface CreateDartGameDto {
    players?: CreatePlayerDto[];
    gameType: DartGameType;
    totalRound?: number;
}

export class DartHitDto {
    r!: number;
    theta!: number; // Theta in rad
    estimatedPoints!: number;
}

