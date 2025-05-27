<script setup lang="ts">
import { defineProps } from 'vue'
import { useRouter } from 'vue-router'
import type { AttributeDisplayType } from '@/models/attribute-display-type.ts'
import { AttributeType } from '@/models/attribute-type.ts'

// Define props
const props = defineProps<{
  items: Array<Record<string, any>>
  columns: Array<AttributeDisplayType>
  createRoute: string
  editRoute: string
}>()

const router = useRouter()

function handleCreateNew() {
  router.push(props.createRoute)
}

function handleRowClick(item: Record<string, any>) {
  router.push(`${props.editRoute}/${item.id}`)
}
</script>

<template>
  <div class="table-container">
    <table>
      <thead>
        <tr>
          <th v-for="column in props.columns" :key="column.displayName">
            {{ column.displayName }}
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in props.items" :key="item.id" @click="handleRowClick(item)">
          <td v-for="column in props.columns" :key="column.displayName">
            <span v-if="column.type === AttributeType.NUMBER">
              {{ item[column.attributeName] }}
            </span>
            <span v-else-if="column.type === AttributeType.DATE">
              {{ new Date(item[column.attributeName]).toLocaleDateString() }}
            </span>
            <span v-else>{{ item[column.attributeName] }}</span>
          </td>
        </tr>
      </tbody>
    </table>
    <button @click="handleCreateNew">Create New</button>
  </div>
</template>

<style scoped>
.table-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 2rem;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  background-color: var(--vt-c-black);
  text-align: left;
  color: white;
}

tbody tr {
  cursor: pointer;
}

button {
  margin-top: 20px;
  padding: 10px 20px;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>
