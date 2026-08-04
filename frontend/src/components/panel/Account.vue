<script setup lang="ts">
import { useAuth } from '../../scripts/useAuth'
import { useRouter } from 'vue-router'


const { isAuthenticated, logout } = useAuth()
const router = useRouter()

const handleLogout = () => {
  router.push('/')
  logout()
}

const handleChangePassword = () => {
  router.push('/')
  logout()
}
</script>

<template>
  <UCard title="Аккаунт" class="w-full">
      <div class="flex gap-4 flex-col">
        <UButton
          label="Выйти из аккаунта"
          color="error"
          variant="outline"
          block
          @click="handleLogout" />

        <UModal title="Смена пароля">
          <UButton
            label="Сменить пароль"
            color="neutral"
            variant="outline"
            block />

        <template #body>
            <ChangePassword />
        </template>
      </UModal>
    </div>
  </UCard>
</template>
