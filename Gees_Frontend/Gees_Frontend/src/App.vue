<script setup lang="ts">
import { RouterView } from 'vue-router'
import Header from '@/components/Header.vue'
</script>

<template>
  <Header style="height: 10%">

  </Header>

  <main>
    <RouterView />
  </main>
</template>

<style scoped>
</style>
