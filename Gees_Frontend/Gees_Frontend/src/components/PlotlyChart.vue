<template>
  <div ref="plotContainer" class="plot-container"></div>
</template>

<script setup lang="ts">
import { ref, watch, onMounted, onUnmounted } from 'vue'
import Plotly from 'plotly.js-dist-min';
import { PlotData } from '@/models/plot-data.ts'

const props = defineProps({
  inputData: {
    type: PlotData,
    required: true,
  },
});

const plotContainer = ref<HTMLDivElement | null>(null);

const layout = {
  title: "3D GEES Plot",
  autosize: true,
  scene: {
    xaxis: { title: "X Axis" },
    yaxis: { title: "Y Axis" },
    zaxis: { title: "Z Axis" },
  },
  margin: { l: 0, r: 0, b: 0, t: 50 },
  responsive: true,
  paper_bgcolor: "rgba(166,161,156,0.1)",
};

const resizePlot = () => {
  if (plotContainer.value) {
    const container = plotContainer.value.getBoundingClientRect();
    Plotly.relayout(plotContainer.value, {
      width: container.width,
      height: container.height,
    });
  }
};

onMounted(() => {
  if (plotContainer.value) {
    Plotly.newPlot(
      plotContainer.value,
      [
        {
          x: [],
          y: [],
          z: [],
          mode: "markers+lines",
          marker: { size: 8, color: "blue" },
          line: {width: 2, color: "red"},
          type: "scatter3d",
        },
      ],
      layout
    );
  }
  window.addEventListener('resize', resizePlot);
});

onUnmounted(() => {
  window.removeEventListener('resize', resizePlot);
});

watch(
  () => props.inputData,
  (newData) => {
    if (plotContainer.value) {
      const plotData = newData.getData();
      Plotly.react(plotContainer.value, [
        {
          x: [...plotData.x],
          y: [...plotData.y],
          z: [...plotData.z],
          mode: "markers+lines",
          marker: { size: 8, color: "blue" },
          line: {width: 2, color: "red"},
          type: "scatter3d",
        },
      ], layout);
    }
    console.log('Updated data:', newData.getData());
  },
  { deep: true, immediate: true }
);
</script>

<style scoped>
.plot-container {
  width: 80%;
  height: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>
