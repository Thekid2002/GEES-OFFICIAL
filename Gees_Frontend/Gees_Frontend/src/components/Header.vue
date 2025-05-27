<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
  <header class="header">
    <nav class="header-nav">
      <RouterLink to="/" class="router-link-nav">
        <span class="material-symbols-outlined icon">home</span>
        Home
      </RouterLink>
      <RouterLink to="/record-gestures" class="router-link-nav">
        <span class="material-symbols-outlined icon">radio_button_checked</span>
        Record Gestures
      </RouterLink>
      <RouterLink to="/view-gestures" class="router-link-nav">
        <span class="material-symbols-outlined icon">motion_blur</span>
        View Gesture
      </RouterLink>
      <RouterLink to="/setup-actions" class="router-link-nav">
        <span class="material-symbols-outlined icon">list</span>
        Create Actions
      </RouterLink>
      <RouterLink to="/setup-gestures" class="router-link-nav">
        <span class="material-symbols-outlined icon">gesture</span>
        Setup Gestures
      </RouterLink>
      <RouterLink to="/map-gestures-to-actions" class="router-link-nav">
        <span class="material-symbols-outlined icon">map</span>
        Map Gestures To Actions
      </RouterLink>
    </nav>
  </header>
</template>

<style scoped>
/* Ensure header navigation is properly aligned */
.header-nav {
  width: 100%;
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  align-items: center;
  padding: 10px;
}

/* Style the header */
.header {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #2c3e50;
  padding: 10px 0;
}

/* Style for navigation links */
.router-link-nav {
  display: flex;
  flex-direction: row;
  align-items: center;
  text-decoration: none;
  color: white;
  font-size: 1.2em;
  padding: 10px 20px;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.router-link-nav:hover {
  background-color: #34495e;
}

/* Active link styling */
.router-link-active {
  background-color: #2980b9;
}

/* Ensure icons have proper spacing */
.icon {
  font-variation-settings: 'FILL' 1;
  font-size: 24px;
  margin-right: 8px;
}
</style>
