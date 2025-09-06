<script setup lang="ts">
import type { Ref } from 'vue'

const numberValue: Ref<number> = defineModel({ default: 0 });

const props = defineProps<{
  min: number,
  max: number,
}>()

function increment() {
  if (numberValue.value < props.max) {
    numberValue.value++;
  }
}

function decrement() {
  if (numberValue.value > props.min) {
    numberValue.value--;
  }
}
</script>

<template>
  <div class="quantity">
    <button class="minus" @click="decrement" type="button">&minus;</button>
    <input type="text" class="input-box" v-model="numberValue" :min="min" :max="max">
    <button class="plus" @click="increment" type="button">&plus;</button>
  </div>
</template>

<style scoped>
.quantity {
  display: flex;
  border: 2px solid var(--vt-c-royal-blue);
  border-radius: 4px;
  width: fit-content;
}

.quantity button {
  background-color: var(--vt-c-royal-blue);
  color: #fff;
  border: none;
  cursor: pointer;
  font-size: 20px;
  width: 30px;
  height: auto;
  text-align: center;
}

.input-box {
  width: 40px;
  text-align: center;
  border: none;
  padding: 8px 10px;
  outline: none;
}
</style>
