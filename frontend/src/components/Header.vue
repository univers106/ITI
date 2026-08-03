<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

import type { NavigationMenuItem } from '@nuxt/ui'

const route = useRoute()

const items = computed<NavigationMenuItem[]>(() => [
  {
    label: 'Главная',
    to: '/',
    active: route.path === '/'
  },
  {
    label: 'Список всех ИТИ',
    to: '/iti',
    active: route.path.startsWith('/iti')
  },
  {
    label: 'Про ИТИ',
    to: '/about',
    active: route.path === '/about'
  },

])
</script>

<template>
  <UHeader>
    <template #title>
        <img src="/favicon.svg" alt="logo" class="w-6 h-6" />
        ИТИ
    </template>

    <template #right>
        <UNavigationMenu :items="items" class="hidden sm:flex" />
    </template>

    <template #body>
          <UNavigationMenu :items="items" orientation="vertical" class="-mx-2.5" />
    </template>
  </UHeader>
</template>
