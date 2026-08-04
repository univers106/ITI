<script setup lang="ts">
import Header from './components/Header.vue'
</script>

<template>
  <UApp>
      <Header />

      <UMain>
          <UContainer>
        <RouterView />
        </UContainer>
      </UMain>

      <Footer />
  </UApp>
</template>

<style scoped>
</style>
