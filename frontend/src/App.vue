<script setup lang="ts">
import Header from './components/Header.vue'
</script>

<template>
  <UApp>
      <Header />

      <main class="min-h-[80vh]">
        <RouterView />
      </main>

      <Footer />
  </UApp>
</template>

<style scoped>
</style>
