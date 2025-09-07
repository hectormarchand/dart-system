import { useEventSource } from '@vueuse/core'
import { ref, type Ref, watch } from 'vue'
import { type DartGame, type DartHitDto } from '@/common/dart.dtos.ts'

const dartGameEvent: Ref<DartGame | null> = ref(null)
const dartHitEvent: Ref<DartHitDto | null> = ref(null)

{
  const { data } = useEventSource<[], string>('/api/games/stream-current-active', [], {
    autoReconnect: true,
  })

  watch(data, (value) => {
    if (value && value !== '{}') {
      dartGameEvent.value = JSON.parse(value)
    }
  })
}

{
  const { data } = useEventSource<[], string>('/api/dart-hit/stream-hits', [], {
    autoReconnect: true,
  })

  watch(data, (value) => {
    if (value && value !== '{}') {
      dartHitEvent.value = JSON.parse(value)
    }
  })
}


export { dartGameEvent, dartHitEvent }

