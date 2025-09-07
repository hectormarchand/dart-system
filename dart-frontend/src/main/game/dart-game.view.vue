<script setup lang="ts">
import { onMounted, ref, type Ref, watch } from 'vue'
import * as sse from '@/server-sent-events/server-sent-events.ts'
import { type DartGame } from '@/common/dart.dtos.ts'
import { HttpService } from '@/http/http.service.ts'

const httpService: HttpService = new HttpService()

const dartGame: Ref<DartGame | null> = ref(null)

watch(sse.dartGameEvent, (updatedGame) => {
  dartGame.value = updatedGame
})

onMounted(async () => {
  dartGame.value = await httpService.get<DartGame>({ path: '/games/current-active' })
})
</script>

<template>
  <div class="flex flex-col" v-if="dartGame">
    <div class="text-center text-8xl pb-36 player-score-this-round">
      <span>{{ dartGame.players[dartGame.currentPlayerIndex].scoreThisRound }}</span>
    </div>

    <div class="grid grid-cols-2 gap-4">
      <div
        class="border-2 border-dashed h-30 w-80 p-2"
        v-for="(player, index) in dartGame.players"
        :class="{ borderOrange: index === dartGame?.currentPlayerIndex }"
        :key="index"
      >
        <span>{{ player.name.toUpperCase() }}</span>
        <span>Score : {{ player.score }}</span>
        <span>PPD : {{ player.ppd }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.borderOrange {
  border-color: var(--p-primary-color);
}

.player-score-this-round {
  color: var(--p-primary-color);
}
</style>