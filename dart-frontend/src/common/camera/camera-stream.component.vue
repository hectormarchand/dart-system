<script setup lang="ts">
import type { PropType } from 'vue'

const props = defineProps({
  camera: {
    type: String as PropType<'left' | 'front' | 'right'>,
    required: true,
  },
})

const cameraStreamUrl = `/api/stream/${props.camera}-camera`
</script>

<template>
  <img :src="cameraStreamUrl" alt="Camera stream video" class="border-2 border-dashed" />
</template>
