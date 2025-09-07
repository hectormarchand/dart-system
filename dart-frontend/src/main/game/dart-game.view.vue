<script setup lang="ts">
import { computed, onMounted, ref, type Ref, watch } from 'vue'
import * as sse from '@/server-sent-events/server-sent-events.ts'
import { type DartGame, type DartPlayer } from '@/common/dart.dtos.ts'
import { HttpService } from '@/http/http.service.ts'
import DartSvg from '@/common/dartboard/dart-svg.component.vue'

const httpService: HttpService = new HttpService()

const dartGame: Ref<DartGame | null> = ref(null)
const currentPlayer: Ref<DartPlayer | undefined> = computed(() => {
  return dartGame.value?.players[dartGame.value?.currentPlayerIndex]
})

watch(sse.dartGameEvent, (updatedGame) => {
  dartGame.value = updatedGame
})

onMounted(async () => {
  dartGame.value = await httpService.get<DartGame>({ path: '/games/current-active' })
})
</script>

<template>
  <div class="flex flex-col" v-if="dartGame && currentPlayer">
    <div class="text-center text-8xl pb-36 player-score-this-round">
      <span>{{ currentPlayer.scoreThisRound }}</span>
    </div>

    <div class="grid grid-cols-2 gap-4">
      <div
        class="border-2 border-dashed h-30 w-80 p-2 flex flex-col justify-center items-center"
        v-for="(player, index) in dartGame.players"
        :class="{ borderOrange: player === currentPlayer }"
        :key="index"
      >
        <span>{{ player.name.toUpperCase() }}</span>
        <span class="text-3xl" :class="{ textOrange: player === currentPlayer }">
          {{ player.score }}
        </span>
        <div class="grid grid-cols-2 gap-3 items-center">
          <span class="text-center">{{ player.ppd }} PPD</span>
          <div class="flex flex-row">
            <DartSvg
              v-for="(_, playerIndex) in 3"
              :key="playerIndex"
              :class="{ textOrange: player.nbOfDartsThrownThisRound > playerIndex }"
              class="h-9"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.borderOrange {
  border-color: var(--p-primary-color);
}

.textOrange {
  color: var(--p-primary-color);
}

.player-score-this-round {
  color: var(--p-primary-color);
}
</style>