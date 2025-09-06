<script setup lang="ts">
import DartboardComponent from '@/common/dartboard/dartboard-component.vue'
import { HttpService } from '@/http/http.service.ts'
import { onMounted, type Ref, ref } from 'vue'

const httpService: HttpService = new HttpService()

const activeGame: Ref = ref(null)

onMounted(async () => {
  activeGame.value = await httpService.get({ path: '/games/current-active' })
})
</script>

<template>
  <div class="flex flex-col">
    <div class="flex justify-between py-4">
      <RouterLink to="create-game">
        <Button label="Créer une partie" />
      </RouterLink>
      <RouterLink to="game" v-if="activeGame">
        <Button label="Regarder la partie active" />
      </RouterLink>
      <RouterLink to="camera-calibration">
        <Button label="Calibrer les caméras" />
      </RouterLink>
    </div>
    <Divider />
  </div>

  <div class="flex justify-center pt-10">
    <DartboardComponent></DartboardComponent>
  </div>
</template>

<style scoped>
.p-divider-horizontal {
  margin: 0;
}
</style>
