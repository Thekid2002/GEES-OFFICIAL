<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { fetchActions } from '@/services/ActionService.ts'
import { fetchGestures } from '@/services/GestureService.ts'
import { fetchGestureActionMappings } from '@/services/GestureActionMappingService.ts'
import type { GestureActionMapping } from '@/models/gestureActionMapping.ts'
import type { Gesture } from '@/models/gesture.ts'
import type { Action } from '@/models/action.ts'

class PropagateResult {
  prediction: number;
  confidence: number;
  timeStamp: string;
  constructor(prediction: number, confidence: number, timeStamp: string) {
    this.prediction = prediction;
    this.confidence = confidence;
    this.timeStamp = timeStamp;
  }
}

let gestureActionMaps: GestureActionMapping[] = [];
let gestures: Gesture[] = [];
let actions: Action[] = [];
let socket: WebSocket | null = null;
const propagateResults = ref<PropagateResult[]>([]);

function setupWebSocket() {
  const endpoint = import.meta.env.VITE_WS_URL + '/ws-validate-gestures';
  socket = new WebSocket(endpoint);

  socket.onmessage = (event: MessageEvent) => {
    const data: PropagateResult = JSON.parse(event.data);
    propagateResults.value.push(data);
    propagateResults.value.sort((a, b) => b.timeStamp - a.timeStamp); // Sort by timestamp (descending)
  };

  socket.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  socket.onclose = () => {
    console.log('WebSocket connection closed');
  };
}

onMounted(async () => {
  setupWebSocket();
  gestureActionMaps = await fetchGestureActionMappings();
  gestures = await fetchGestures();
  actions = await fetchActions();
});

function getGestureFromGestureId(gestureId: number): string {
  const gesture = gestures.find(g => g.id === gestureId);
  return gesture ? gesture.name : gestureId.toString();
}

function getActionFromActionId(actionId: number): string {
  const action = actions.find(a => a.id === actionId);
  return action ? action.name : actionId.toString();
}

function getGestureAndActionFromPrediction(prediction: number): { gesture: string, action: string } {
  const mapping = gestureActionMaps.find(mapping => mapping.gesture_id === prediction);
  if (mapping) {
    return { gesture: getGestureFromGestureId(mapping.gesture_id), action: getActionFromActionId(mapping.action_id) };
  }
  return { gesture: prediction, action: "No mapping found, so no action" };
}

const sortedPropagateResults = computed(() =>
  propagateResults.value.slice().sort((a, b) => new Date(b.timeStamp).getTime() - new Date(a.timeStamp).getTime())
);
</script>

<template>
  <div class="tracking-container">
    <div v-for="result in sortedPropagateResults" :key="result.timeStamp">
      Gesture: {{ getGestureAndActionFromPrediction(result.prediction).gesture }},
      Action: {{ getGestureAndActionFromPrediction(result.prediction).action }},
      Confidence: {{ result.confidence }},
      Time: {{ new Date(result.timeStamp).toLocaleTimeString() }}
    </div>
  </div>
</template>

<style>
.tracking-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.tracking-container > div:first-child {
  font-size: 1.6em;
  font-weight: bold;
  background-color: var(--vt-c-indigo);
  color: white;
}

.tracking-container > div {
  font-size: 1.2em;
  font-weight: bold;
  color: white;
  margin-bottom: 10px;
  border-radius: 5px;
}
</style>
