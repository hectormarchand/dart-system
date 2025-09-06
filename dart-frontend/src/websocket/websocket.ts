import { useWebSocket } from '@vueuse/core'
import { ref, type Ref, watch } from 'vue'

const WS_URL: string = "/ws";

const { data } = useWebSocket(WS_URL);

const dartGameMessage: Ref = ref();
const dartHitMessage: Ref = ref();

watch(data, (value) => {
  const parsedValue = JSON.parse(value);

  switch (parsedValue.type) {
    case "dart-game":
      dartGameMessage.value = parsedValue.data;
      break;
    case "dart-hit":
      dartHitMessage.value = parsedValue.data;
      break;
    default:
      console.warn(`Websocket message type ${parsedValue.type} not supported`);
  }
});

export { dartGameMessage, dartHitMessage };
