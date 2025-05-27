<script setup lang="ts">
import Form from '@/components/Form.vue';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { createGesture, fetchGestureById, updateGestureById } from '@/services/GestureService.ts';
import { Gesture } from '@/models/gesture.ts';
import router from '@/router';

const route = useRoute();
const gesture = ref<Gesture>(new Gesture(null, '', null, null));

const fields = {
  id: {
    type: 'number',
    editable: false,
    displayName: 'ID',
    icon: 'fingerprint',
    description: 'Unique identifier for the gesture',
  },
  name: {
    type: 'text',
    editable: true,
    displayName: 'Name',
    icon: 'gesture',
    description: 'Enter the name of the gesture',
  },
  description: {
    type: 'text',
    editable: true,
    displayName: 'Description',
    icon: 'description',
    description: 'Enter a description for the gesture',
  },
  image_url: {
    type: 'text',
    editable: true,
    displayName: 'Image URL',
    icon: 'image',
    description: 'Provide an image URL for the gesture',
  },
};

const formName = ref('Create a new gesture!');

onMounted(async () => {
  const gestureId = route.params.id;
  if (gestureId) {
    try {
      const existingGesture = await fetchGestureById(Number(gestureId));
      if (existingGesture) {
        gesture.value = existingGesture;
        formName.value = 'Edit Gesture';
      } else {
        console.error('Gesture not found');
      }
    } catch (error) {
      console.error('Error fetching gesture:', error);
    }
  }
});

async function saveGesture(formData: Record<string, string | number | null>) {
  try {
    const updatedGesture = new Gesture(
      formData.id as number | null,
      formData.name as string,
      formData.description as string | null,
      formData.image_url as string | null
    );
    if (gesture.value.id) {
      // Update existing gesture
      await updateGestureById(gesture.value.id, updatedGesture);
      router.push('/setup-gestures');
    } else {
      // Create new gesture
      await createGesture(updatedGesture);
      router.push('/setup-gestures');
    }
    console.log('Gesture saved successfully!');
  } catch (error) {
    console.error('Error saving gesture:', error);
  }
}
</script>

<template>
  <Form
    :fields="fields"
    :formName="formName"
    :form-value="gesture"
    @submit="saveGesture"
  />
</template>
