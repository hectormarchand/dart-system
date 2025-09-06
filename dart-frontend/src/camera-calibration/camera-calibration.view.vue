<script setup lang="ts">

import CameraStream from "@/common/camera/camera-stream.component.vue";
</script>

<template>
  <div class="pt-40 flex justify-between">
    <div>
      <CameraStream camera="left" class="border-2 border-dashed camera-stream" />
    </div>
    <div>
      <CameraStream camera="right" class="border-2 border-dashed camera-stream" />
    </div>
  </div>

  <div class="flex justify-center pt-10">
    <RouterLink to="/">
      <Button label="Retour" />
    </RouterLink>
  </div>
</template>

<style>
</style>
