<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
import type { AttributeDisplayType } from '@/models/attribute-display-type.ts'

const props = defineProps<{
  fields: Record<string, AttributeDisplayType>;
  formName: string;
  formValue: any;
}>();

const emit = defineEmits<{
  (e: 'submit', formData: Record<string, string | number>): void;
}>();

function handleSubmit(event: Event) {
  event.preventDefault();
  emit('submit', props.formValue);
}
</script>

<template>
  <div class="form-container">
    <div class="form-page">
      <h1>{{ formName }}</h1>
      <form @submit="handleSubmit">
        <div v-for="(field, key) in fields" :key="key" class="input-container">
          <hr class="horizontal-line" v-if="field.editable"/>
          <label class="label-container" :for="key" v-if="field.editable">
            <span class="material-symbols-outlined icon">{{ field.icon }}</span>
            {{ field.displayName }}
          </label>
          <input
            v-if="field.editable"
            class="input"
            :id="key"
            v-model="props.formValue[key]"
            :type="field.type"
            :readonly="!field.editable"
            required
          />
          <p class="description" v-if="field.editable">{{ field.description }}</p>
        </div>
        <input class="submit" type="submit" value="Submit" />
      </form>
    </div>
  </div>
</template>

<style scoped>
.horizontal-line {
  border: none;
  border-top: 2px solid black;
  width: 100%;
}

h1 {
  color: #2c3e50 !important;
  text-align: center;
}

.form-page {
  display: flex;
  width: 80%;
  padding-bottom: 2%;
  background-color: white;
  padding-inline: 2%;
  padding-top: 2%;
  flex-direction: column;
}

.form-container {
  display: flex;
  justify-content: center;
  flex-direction: row;
}

form {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
form > div {
  display: flex;
  flex-direction: column;
  gap: 10px;
}


.input-container {
  padding: 9px;
  color: black;
}

.input {
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 1.2em;
}

.submit {
  padding: 10px;
  background-color: #2c3e50;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1.2em;
  cursor: pointer;
}

.label-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
  font-size: 1.6em;
}

.description {
  font-size: 0.9em;
  color: #666;
}</style>
