import { useEventSource } from '@vueuse/core'
import { ref, type Ref, watch } from 'vue'

const {data} = useEventSource(
  '/api/games/stream-current-active',
  [], {  }
  );

const dartGameEvent: Ref = ref();

watch(data, (value) => {
  if (value && value !== "{}") {
    dartGameEvent.value = value;
  }
});

export { dartGameEvent };

