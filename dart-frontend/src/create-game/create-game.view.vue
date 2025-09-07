<script setup lang="ts">
import { computed, ref, type Ref } from 'vue'
import Stepper from 'primevue/stepper'
import StepList from 'primevue/steplist'
import StepPanels from 'primevue/steppanels'
import Step from 'primevue/step'
import StepPanel from 'primevue/steppanel'
import { HttpService } from '@/http/http.service.ts'
import type { CreateDartGameDto } from '@/common/dart.dtos.ts'

const httpService: HttpService = new HttpService()

const numberOfPlayers: Ref<number> = ref(1)
const gameType: Ref<'301' | '501' | '701'> = ref('501')
const gameTypeOptions = ['301', '501', '701']

const players: Ref<{ name: string }[]> = computed(() => {
  return Array.from({ length: numberOfPlayers.value }, () => ({ name: '' }))
})

async function createGame() {
  const payload: CreateDartGameDto = {
    players: players.value,
    gameType: gameType.value,
    totalRound: 15, // Fixed to 15 for now
  }

  await httpService.post({ path: '/games', body: payload })
}
</script>

<template>
  <form class="card flex justify-center pt-20" @submit.prevent="createGame">
    <Stepper value="1" linear class="basis-[50rem]">
      <StepList>
        <Step value="1">Sélection du type de partie</Step>
        <Step value="2">Ajout des joueurs</Step>
      </StepList>
      <StepPanels>
        <StepPanel v-slot="{ activateCallback }" value="1">
          <div class="flex flex-col h-60">
            <div class="border-2 border-dashed flex-auto flex flex-col">
              <div class="w-full grid grid-cols-2 px-28 pt-12 items-center">
                <span>Nombre de joueurs</span>
                <InputNumber
                  v-model="numberOfPlayers"
                  inputId="horizontal-buttons"
                  showButtons
                  buttonLayout="horizontal"
                  :step="1"
                  :max="8"
                  :min="1"
                >
                  <template #incrementicon>
                    <span class="text-xl">&plus;</span>
                  </template>
                  <template #decrementicon>
                    <span class="text-xl">&minus;</span>
                  </template>
                </InputNumber>
              </div>

              <div class="w-full grid grid-cols-2 px-28 pt-12 items-center">
                <span>Type de jeu</span>
                <SelectButton v-model="gameType" :options="gameTypeOptions" />
              </div>
            </div>
          </div>
          <div class="flex pt-6 justify-between">
            <RouterLink :to="{ name: 'home' }">
              <Button severity="secondary" label="Annuler" />
            </RouterLink>
            <Button label="Suivant" @click="activateCallback('2')" />
          </div>
        </StepPanel>
        <StepPanel v-slot="{ activateCallback }" value="2">
          <div class="flex flex-col">
            <div
              class="border-2 border-dashed flex flex-auto flex-col gap-3 justify-center items-center py-10"
            >
              <div v-for="(player, index) in players" :key="index" class="flex flex-col gap-1">
                <label :for="'player-' + index">{{ 'Nom du joueur ' + (index + 1) }}</label>
                <InputText
                  :id="'player-' + index"
                  type="text"
                  v-model="player.name"
                  placeholder="Nom du joueur"
                />
              </div>
            </div>
          </div>
          <div class="flex pt-6 justify-between">
            <Button label="Précédent" severity="secondary" @click="activateCallback('1')" />
            <Button label="Créer la partie" type="submit" />
          </div>
        </StepPanel>
      </StepPanels>
    </Stepper>
  </form>
</template>
