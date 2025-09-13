<script setup lang="ts">
import CameraStream from '@/common/camera/camera-stream.component.vue'
</script>

<template>
  <div class="pt-40 flex justify-center">
    <div class="w-[45rem]">
      <Carousel :value="['left', 'front', 'right']" :num-visible="1" class="w-fit">
        <template #item="camera">
          <CameraStream :camera="camera.data" />
        </template>
      </Carousel>
    </div>
  </div>

  <div class="flex justify-center pt-10">
    <RouterLink to="/">
      <Button label="Retour" />
    </RouterLink>
  </div>
</template>

<style></style>