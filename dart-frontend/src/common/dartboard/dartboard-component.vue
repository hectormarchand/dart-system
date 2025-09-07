<script setup lang="ts">
import { ref, type Ref, watch } from 'vue'
import DartBoardSvg from '@/common/dartboard/dart-board-svg.component.vue'
import * as sse from '@/server-sent-events/server-sent-events.ts'
import DartHitSvg from '@/common/dartboard/dart-hit-svg.component.vue'

const lastThreeDartHits: Ref<{ r: number; theta: number }[]> = ref([
  { r: 23, theta: 3.14 / 2 },
  { r: 12, theta: 3.14 },
])

watch(sse.dartHitEvent, (dartHit) => {
  if (!dartHit) {
    return
  }

  lastThreeDartHits.value.push(dartHit)

  while (lastThreeDartHits.value.length > 3) {
    lastThreeDartHits.value.shift()
  }
})

// TODO : add the last three targets into the html template
</script>

<template>
  <div class="dartboard-svg">
    <DartBoardSvg />
    <DartHitSvg
      class="h-6 dart-hit"
      v-for="(dartHit, index) in lastThreeDartHits"
      :key="index"
      :style="{
        left: `${50 + dartHit.r * Math.cos(dartHit.theta)}%`,
        top: `${50 - dartHit.r * Math.sin(dartHit.theta)}%`,
      }"
    />
  </div>
</template>

<style scoped>
.dartboard-svg {
  position: relative;
  height: 80vh;
  width: fit-content;
}

.dart-hit {
  position: absolute;
  transform: translate(-50%, -50%);
  pointer-events: none;
}
</style>
