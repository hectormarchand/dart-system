import { DartGameType } from "./dart.types"

export class CreatePlayerDto {
    name?: string;
}

export class CreateDartGameDto {
    players?: CreatePlayerDto[];
    gameType?: DartGameType;
    totalRound?: number;
}

export class DartHitDto {
    x!: number;
    y!: number;
    estimatedPoints!: number;
}

