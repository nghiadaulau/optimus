<template>
  <fieldset>
    <legend v-if="label">{{ label }}</legend>
    <slot></slot>
  </fieldset>
</template>

<script setup>
defineProps({
  label: {
    type: String,
    default: null,
  },
});
</script>

<style scoped>
fieldset {
  border: 1px solid #7c7c7c;
  border-radius: 5px;
  padding: 20px;
  margin: 10px;
}

.dark fieldset {
  border: 1px solid #7c7c7c;
  border-radius: 5px;
  padding: 20px;
  margin: 10px;
}

legend {
  font-weight: bold;
  padding: 5px;
}
</style>
