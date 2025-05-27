<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Action } from '@/models/action.ts'
import PlotlyChart from '@/components/PlotlyChart.vue'
import { PlotData } from '@/models/plot-data.ts'
import { MessageType, type WSMessage } from '@/models/WSMessage.ts'
import { fetchGestures } from '@/services/GestureService.ts'

const gestures = ref<Action[]>([])
const selectedGesture = ref<Action | null>(null)
let socket: WebSocket | null = null

enum Training {
  Training = 'training',
  Final = 'final',
  Error = 'error',
  None = 'none',
}

const training = ref(Training.None as Training)
const countdown = ref<number | null>(null)
const messageTraining = ref('')
const gestureData = ref(new PlotData() as PlotData | undefined)

onMounted(async () => {
  gestures.value = await fetchGestures()
  setupWebSocket()
})

function handleMessageUpdate(data: WSMessage, trainingMsg: string, trainingStatus?: Training) {
  messageTraining.value = trainingMsg
  if (trainingStatus) training.value = trainingStatus
}

function setupWebSocket() {
  const endpoint = import.meta.env.VITE_WS_URL + '/ws-record-gestures'
  socket = new WebSocket(endpoint)

  socket.onopen = () => {
    console.log('WebSocket connection established')
  }

  socket.onmessage = (event: MessageEvent) => {
    try {
      const msg = JSON.parse(event.data) as WSMessage

      switch (msg.msgType) {
        case MessageType.COUNTDOWN:
          countdown.value = parseInt(msg.data) || 0
          training.value = Training.Training
          messageTraining.value = 'Wait!'
          break

        case MessageType.START:
          countdown.value = 0
          handleMessageUpdate(msg, 'Perform Action', Training.Training)
          break

        case MessageType.STOP:
          countdown.value = 0
          handleMessageUpdate(msg, 'Action recorded', Training.Final)

          gestureData.value = new PlotData()
          const parsedData = JSON.parse(msg.data)

          parsedData.DataPoints.forEach((point) => {
            gestureData.value?.addPoint(point.AX, point.AY, point.AZ)
          })
          break

        case MessageType.ERROR:
          countdown.value = 0
          handleMessageUpdate(msg, msg.data!, Training.Error)
          break

        default:
          console.warn('Unknown message type:', msg.msgType)
      }

      console.log('Message from server:', msg)
    } catch (err) {
      console.error('Failed to parse WebSocket message', err)
    }
  }

  socket.onerror = (error) => {
    console.error('WebSocket error:', error)
  }

  socket.onclose = () => {
    console.log('WebSocket connection closed')
  }
}

function startTraining() {
  if (!selectedGesture.value) {
    console.error('No gesture selected')
    return
  }

  training.value = Training.None
  messageTraining.value = ''
  gestureData.value = new PlotData()

  const gesture = selectedGesture.value
  console.log(`Starting training for gesture with ID: ${gesture.id}`)
  sendGestureId(gesture.id!)
}

function sendGestureId(id: number) {
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify({ id }))
    console.log('Action ID sent via WebSocket:', id)
  } else {
    console.error('WebSocket is not open')
  }
}

function getStatusClass(status: Training): string {
  if (status === Training.Training) {
    return 'track'
  } else if (status === Training.Final) {
    return 'final'
  } else if (status === Training.Error) {
    return 'error'
  }
  return ''
}

function getGestureData(): PlotData {
  if (gestureData.value) {
    return new PlotData(gestureData.value.getData())
  }
  return new PlotData()
}
</script>

<template>
  <Header />
  <div class="tracking-container">
    <div class="tracking-box left-box">
      <div :class="['track-features', getStatusClass(training)]">
        <h1>Track Action</h1>
        <p>Select a Action to track below.</p>
        <form @submit.prevent="startTraining">
          <select v-model="selectedGesture">
            <option :value="null">None</option>
            <option v-for="gesture in gestures" :key="gesture.id" :value="gesture">
              {{ gesture.name }}
            </option>
          </select>
          <input type="submit" value="Track Action" />
        </form>
        <div>
          <p>Training...</p>
          <p v-if="countdown && training === Training.Training">Countdown: {{ countdown }}</p>
          <p :style="{ color: messageTraining === 'Perform Action' ? 'green' : '#fc0fc0' }">
            {{ messageTraining }}
          </p>
        </div>
      </div>
    </div>
    <div class="tracking-box right-box">
      <plotly-chart :input-data="getGestureData()" />
    </div>
  </div>
</template>

<style scoped>
.tracking-container {
  margin-top: 20px;
  width: 100%;
  min-height: 600px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.tracking-box {
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 20px;
}

.left-box {
  width: 45%;
  height: 600px;
}

.right-box {
  width: 55%;
  height: 600px;
}

@keyframes borderBlink {
  0%,
  100% {
    border-color: #2c3e50;
  }
  50% {
    border-color: yellow;
  }
}

@keyframes borderError {
  0%,
  100% {
    border-color: #2c3e50;
  }
  50% {
    border-color: #ca1212;
  }
}

@keyframes borderTurnGreen {
  0% {
    border-color: yellow;
  }
  100% {
    border-color: green;
  }
}

.track-features {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  margin: 0 auto;
  color: white;
  width: 400px;
  height: 400px;
  border-radius: 50%;
  border: 20px solid #2c3e50;
  background-color: var(--vt-c-black);
}

.track {
  animation: borderBlink 1s linear infinite;
}
.final {
  animation: borderTurnGreen 1s linear forwards;
}
.error {
  animation: borderError 1.5s linear infinite;
}
</style>
