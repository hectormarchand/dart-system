export type DartGameType = '301' | '501' | '701'

export interface DartPlayer {
  name: string
  score: number
  scoreThisRound: number
  ppd: number
  nbOfDartsThrownThisRound: number
}

export interface DartGame {
  id: never
  players: DartPlayer[]
  gameType: DartGameType
  totalRound: number
  currentRound: number
  active: boolean
  createdDate: string
  currentPlayerIndex: number
}

export class CreatePlayerDto {
  name?: string
}

export class CreateDartGameDto {
  players?: CreatePlayerDto[]
  gameType?: DartGameType
  totalRound?: number
}

export interface DartHitDto {
  x: number
  y: number
  estimatedPoints: number
}
