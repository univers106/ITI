<script setup lang="ts">
import type { NavigationMenuItem } from '@nuxt/ui'


const items: NavigationMenuItem[] = [
  {
    label: 'GitHub',
    to: 'https://github.com/univers106/ITI',
    target: '_blank'
  },
  {
    label: 'Политика конфиденциальности',
    to: '/policies',
  },
  {
    label: 'Контакт',
    to: '/about',
  },
  {
    label: 'Панель',
    to: '/panel',
  },
]

</script>

<template>
  <UFooter>
    <template #left>
        <p class="text-muted text-sm">Copyright © {{ new Date().getFullYear() }}</p>
    </template>

    <template #right>
        <UNavigationMenu :items="items" variant="link" orientation="vertical"/>

    </template>
  </UFooter>
</template>
