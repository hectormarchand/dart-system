<script setup lang="ts">
import { HttpService } from '@/http/http.service.ts'
import { onMounted, type Ref, ref } from 'vue'

const httpService: HttpService = new HttpService()

const activeGame: Ref = ref(null)

onMounted(async () => {
  activeGame.value = await httpService.get({ path: '/games/current-active' })
})
</script>

<template>
  <div class="flex flex-col gap-12 items-center">
    <RouterLink :to="{ name: 'create-game' }">
      <Button label="Creer une partie" />
    </RouterLink>
    <RouterLink :to="{ name: 'camera-calibration' }">
      <Button label="Calibrer les cameras" />
    </RouterLink>
    <RouterLink :to="{ name: 'game' }">
      <Button label="Regarder la partie active" :disabled="!activeGame" />
    </RouterLink>
  </div>
</template>

<style scoped></style>