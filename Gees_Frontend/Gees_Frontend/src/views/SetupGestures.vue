<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { fetchGestures } from '@/services/GestureService.ts';
import Table from '@/components/Table.vue';
import type { Gesture } from '@/models/gesture.ts';
import { AttributeDisplayType } from '@/models/attribute-display-type.ts';
import { AttributeType } from '@/models/attribute-type.ts';

const gestures = ref<Gesture[]>([]);

const columns = [
  new AttributeDisplayType(
    'ID',
    AttributeType.NUMBER,
    'id',
    'fingerprint',
    'Unique identifier for the gesture'
  ),
  new AttributeDisplayType(
    'Gesture Name',
    AttributeType.TEXT,
    'name',
    'description',
    'Enter the name of the gesture'
  ),
  new AttributeDisplayType(
    'Description',
    AttributeType.TEXT,
    'description',
    'description',
    'Enter the description of the gesture'
  ),
  new AttributeDisplayType(
    'URL of Image',
    AttributeType.URL,
    'image_url',
    'image',
    'Enter the URL of the image'
  )
];

onMounted(async () => {
  gestures.value = await fetchGestures();
});
</script>

<template>
  <Table :items="gestures" :columns="columns" createRoute="/setup-gestures/create" editRoute="/setup-gestures/edit" />
</template>
