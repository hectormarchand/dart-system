<script setup lang="ts">
import DartboardComponent from '@/common/dartboard/dartboard-component.vue'
</script>

<template>
  <div class="flex flex-row h-full">
    <div class="w-1/2 overflow-hidden flex items-center justify-center">
      <RouterView />
    </div>
    <div class="w-1/2 overflow-hidden flex items-center justify-center">
      <DartboardComponent />
    </div>
  </div>
</template>

<style scoped></style>
