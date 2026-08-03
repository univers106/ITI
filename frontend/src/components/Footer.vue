<script setup lang="ts">
import { useRouter } from 'vue-router'
import type { NavigationMenuItem } from '@nuxt/ui'
import { useAuth } from '../scripts/useAuth'

const router = useRouter()
const { isAuthenticated, logout } = useAuth()

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

async function handleLogout() {
  await logout()
  router.push('/')
}
</script>

<template>
  <UFooter>
    <template #left>
        <p class="text-muted text-sm">Copyright © {{ new Date().getFullYear() }}</p>
    </template>

        <UButton
            v-if="isAuthenticated"
            label="Выйти"
            color="neutral"
            variant="outline"
            size="sm"
            @click="handleLogout"
        />
    <template #right>
        <UNavigationMenu :items="items" variant="link" orientation="vertical"/>

    </template>
  </UFooter>
</template>
