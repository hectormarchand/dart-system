import {DartPlayer, DartGameType, DartGame} from "../common/dart.types";

export class GameBusinessService {

    private readonly THROW_PER_ROUND: number = 3;

    constructor() {
    }

    /**
     * Allows to check if a round is finished, if the game is finished ...
     * In general, move the game from one state to the next state based on dart rules
     * @param game
     */
    executeOneStepInGame(game: DartGame): DartGame {
        const currentPlayer: DartPlayer = game.players[game.currentPlayerIndex];

        // Check is the current player has won
        if (currentPlayer.score === 0) {
            game.winner = currentPlayer;
            game.active = false;
            return game;
        }

        // Check if it is time to change the player
        if (currentPlayer.nbOfDartsThrownThisRound >= this.THROW_PER_ROUND) {
            if (game.currentPlayerIndex === game.players.length - 1) {

                // Check if last round and no winner...
                if (game.currentRound === 15) {
                    game.active = false;
                    return game;
                }

                game.currentRound++;
                game.players.forEach(player => {
                    player.nbOfDartsThrownThisRound = 0;
                    player.scoreThisRound = 0;
                });
            }

            game.currentPlayerIndex = (game.currentPlayerIndex + 1) % game.players.length;
        }

        return game;
    }
}