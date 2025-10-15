<script setup lang="ts">
defineProps<{
  msg: string
}>()
</script>

<template>
    <h3>
      You’ve successfully created a project

      {{msg}}
    </h3>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

</style>
