<template>
  <div>
    <div>
      <s-button style="margin: 10px">Primary</s-button>
      <s-button size="small">Primary small</s-button>
    </div>

    <div>
      <s-button style="margin: 10px" type="secondary">Secondary</s-button>
      <s-button type="secondary" size="small">Secondary small</s-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();
</script>

<style scoped></style>
