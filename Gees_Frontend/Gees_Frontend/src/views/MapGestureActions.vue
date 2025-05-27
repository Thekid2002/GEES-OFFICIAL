<script setup lang="ts">
import { ref, onMounted } from 'vue';
import type { Action } from '@/models/action.ts';
import type { Gesture } from '@/models/gesture.ts';
import {
  fetchActions
} from '@/services/ActionService.ts';
import {
  fetchGestures,
} from '@/services/GestureService.ts';
import {
  fetchGestureActionMappings,
  updateGestureActionMappings,
} from '@/services/GestureActionMappingService.ts';

const gestureMap = ref<Map<Gesture, Action>>(new Map());
const possibleGestures = ref<Gesture[]>([]);
const possibleActions = ref<Action[]>([]);
const selectedActions = ref<Record<number, number>>({});

onMounted(async () => {
  try {
    possibleGestures.value = await fetchGestures();
    possibleActions.value = await fetchActions();

    const mappings = await fetchGestureActionMappings();
    const map = new Map<Gesture, Action>();

    if (mappings != null && mappings.length > 0) {
      for (const mapping of mappings) {
        const gesture = possibleGestures.value.find((g) => g.id === mapping.gesture_id);
        const action = possibleActions.value.find((a) => a.id === mapping.action_id);

        if (gesture && action) {
          map.set(gesture, action);
          selectedActions.value[gesture.id] = action.id;
        }
      }
    }

    gestureMap.value = map;
  } catch (error) {
    console.error('Error initializing data:', error);
  }
});

async function save() {
  const mappings = Array.from(gestureMap.value.entries()).map(([gesture, action]) => ({
    gesture_id: gesture.id,
    action_id: action.id,
  }));

  try {
    await updateGestureActionMappings(mappings);
    console.log('Mappings saved successfully!');
  } catch (error) {
    console.error('Error saving mappings:', error);
  }
}
</script>

<template>
  <form @submit.prevent="save">
    <div class="mapping-container">
      <div
        class="movement-container"
        v-for="gesture in possibleGestures"
        :key="gesture.id"
      >
        <h2 class="gesture-title">Map Gesture to Action</h2>
        <div class="map-pair">
          <img
            :src="gesture.image_url"
            alt="Tutorial Image"
            class="tutorial-image"
          />

          <!-- Dropdown for gestures -->
          <select
            v-model="selectedActions[gesture.id]"
            class="gesture-select"
            @change="gestureMap.set(gesture, possibleActions.find(g => g.id === selectedActions[gesture.id])!)"
          >
            <option value="" disabled>Select an action</option>
            <option
              v-for="action in possibleActions"
              :key="action.id"
              :value="action.id"
            >
              {{ action.name }}
            </option>
          </select>
        </div>
      </div>

      <div class="save-button-container">
        <button type="submit" class="i-material-symbols:magic-button">
          <span class="material-symbols-outlined icon">save</span>
          Save
        </button>
      </div>
    </div>
  </form>
</template>

<style>
.mapping-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2rem;
  padding-top: 2rem;
}

.map-pair {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 2rem;
  padding-bottom: 2rem;
  border-bottom: 3px solid white;
}

.movement-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2rem;
}
.tutorial-image {
  width: 40%;
  height: auto;
}

.gesture-title {
  padding-bottom: 0;
  font-weight: 800;
  color: white;

}

.save-button-container {
  padding-bottom: 10rem;
}
</style>
