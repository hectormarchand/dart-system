<script setup lang="ts">
import * as ws from "@/websocket/websocket.ts"
import { ref, type Ref, watch } from 'vue'
import DartBoardSvg from "@/common/dartboard/dart-board-svg.component.vue"

const lastThreeDartHits: Ref<{ r: number, theta: number }[]> = ref([]);

watch(ws.dartHitMessage, (value) => {
  lastThreeDartHits.value.push({ r: value.r, theta: value.theta });

  while (lastThreeDartHits.value.length > 3) {
    lastThreeDartHits.value.shift();
  }
})

// TODO : add the last three targets into the html template
</script>

<template>
  <div class="dartboard-svg">
    <DartBoardSvg></DartBoardSvg>
  </div>
</template>

<style scoped>
.dartboard-svg {
  height: 80vh;
  width: fit-content;
}
</style>
