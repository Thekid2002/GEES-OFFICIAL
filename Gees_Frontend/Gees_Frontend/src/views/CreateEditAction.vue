<script setup lang="ts">
import Form from '@/components/Form.vue';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { createAction, fetchActionById, updateActionById } from '@/services/ActionService.ts'
import { Action } from '@/models/action.ts'
import router from '@/router'

const route = useRoute();
const action = ref<Action>(new Action(null, ''));

const fields = {
  id: {
    type: 'number',
    editable: false,
    displayName: 'ID',
    icon: 'fingerprint',
    description: 'Unique identifier for the action',
  },
  name: {
    type: 'text',
    editable: true,
    displayName: 'Name',
    icon: 'description',
    description: 'Enter the name of the gesture',
  }
};

const formName = ref('Create a new action!');

onMounted(async () => {
  const actionId = route.params.id;
  if (actionId) {
    try {
      const existingAction = await fetchActionById(Number(actionId));
      if (existingAction) {
        action.value = existingAction;
        formName.value = 'Edit Action';
      } else {
        console.error('Action not found');
      }
    } catch (error) {
      console.error('Error fetching action:', error);
    }
  }
});

async function saveAction(formData: Record<string, string | number>) {
  try {
    let updatedAction = new Action(
      formData.id as number,
      formData.name as string
    );
    if (action.value.id) {
      // Update existing action
      await updateActionById(action.value.id, updatedAction);
      router.push('/setup-actions');
    } else {
      // Create new action
      await createAction(updatedAction);
      router.push('/setup-actions');
    }
    console.log('Action saved successfully!');
  } catch (error) {
    console.error('Error saving action:', error);
  }
}
</script>

<template>
  <Form
    :fields="fields"
    :formName="formName"
    :form-value="action"
    @submit="saveAction"
  />
</template>
