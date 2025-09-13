<script setup lang="ts">
import { computed, onMounted, ref, type Ref, watch } from 'vue'
import * as sse from '@/server-sent-events/server-sent-events.ts'
import { type DartGame, type DartPlayer, type PatchGameDto } from '@/common/dart.dtos.ts'
import { HttpService } from '@/http/http.service.ts'
import DartSvg from '@/common/dartboard/dart-svg.component.vue'
import CameraStream from '@/common/camera/camera-stream.component.vue'
import { onKeyStroke } from '@vueuse/core'

const httpService: HttpService = new HttpService()

const dartGame: Ref<DartGame | null> = ref(null)
const currentPlayer: Ref<DartPlayer | undefined> = computed(() => {
  return dartGame.value?.players[dartGame.value?.currentPlayerIndex]
})

watch(sse.dartGameEvent, (updatedGame) => {
  dartGame.value = updatedGame
})

// Force the game to move to the next player if the space touch is pressed
onKeyStroke(' ', async () => {
  await nextPlayer()
})

async function nextPlayer() {
  const payload: PatchGameDto = {
    type: 'next-player',
  }
  await httpService.patch({ path: '/games/current-active', body: payload })
}

onMounted(async () => {
  dartGame.value = await httpService.get<DartGame>({ path: '/games/current-active' })
})
</script>

<template>
  <div class="flex flex-col" v-if="dartGame && currentPlayer">
    <div class="grid grid-cols-2 gap-4 pb-16">
      <div class="flex items-center justify-center text-8xl player-score-this-round">
        <span>{{ currentPlayer.scoreThisRound }}</span>
      </div>
      <div class="flex justify-center">
        <CameraStream camera="front" class="w-72 aspect-4/3" />
      </div>
    </div>

    <div class="text-center pb-2">{{ dartGame.currentRound }} / {{ dartGame.totalRound }}</div>

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
    <div class="pt-15">
      <RouterLink :to="{ name: 'home' }">
        <Button label="Retour" />
      </RouterLink>
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