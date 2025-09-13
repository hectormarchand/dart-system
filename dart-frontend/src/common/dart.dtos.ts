export type DartGameType = '301' | '501' | '701'

export interface DartPlayer {
  name: string
  score: number
  scoreThisRound: number
  ppd: number
  nbOfDartsThrownThisRound: number
}

export interface DartGame {
  _id: string
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
  r: number
  theta: number
  estimatedPoints: number
}

export interface PatchGameDto {
  type: 'next-player' | 'fix-score'
  fix?: number // TO BE IMPLEMENTED : fix a wrong score from the client
}
