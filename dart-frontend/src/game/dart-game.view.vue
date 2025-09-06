<script setup lang="ts">
import { onMounted, ref, type Ref, watch } from 'vue'
import { HttpService } from '@/http/http.service.ts'
import * as sse from "@/server-sent-events/server-sent-events.ts"

const gameUid = 1; // TODO : Get the uid from the route param

const httpService: HttpService = new HttpService();

const dartGame: Ref = ref(null);

watch(sse.dartGameEvent, (updatedGame) => {
  dartGame.value = updatedGame;
});

onMounted(async () => {
  dartGame.value = await httpService.get({ path: `/games/${gameUid}` });
});
</script>

<template>
  <p>dart game</p>
</template>

<style scoped></style>
