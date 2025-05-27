<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { fetchActions } from '@/services/ActionService.ts'
import Table from '@/components/Table.vue';
import type { Action } from '@/models/action.ts'
import { AttributeDisplayType } from '@/models/attribute-display-type.ts'
import { AttributeType } from '@/models/attribute-type.ts'

const actions = ref<Action[]>([]);


const columns = [
  new AttributeDisplayType(
    'ID',
    AttributeType.NUMBER,
    'id',
    'fingerprint',
    'Unique identifier for the action'
  ),
  new AttributeDisplayType(
    'Name',
    AttributeType.TEXT,
    'name',
    'description',
    'Enter the name of the action'
  ),
];

onMounted(async () => {
  actions.value = await fetchActions();
});
</script>

<template>
  <Table :items="actions" :columns="columns" createRoute="/setup-actions/create" editRoute="/setup-actions/edit" />
</template>
