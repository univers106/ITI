<script setup lang="ts">
  import { useAuth } from "../scripts/useAuth";
  import { ref, onMounted } from 'vue'

  const authStore = useAuth();
  const user = ref<any>(null)
  onMounted(async () => {
      // Запрос выполнится строго на клиенте после загрузки страницы
      user.value = await authStore.me()
      console.log(user.value)
  })
</script>

<template>
    <UPageHeader
    title="Панель управления"
    :description="`Вы вошли как ${user?.name}`"
    class="border-b-0"
    />

      <UPageGrid>
          <Account />
          <Account />
          <Account />
          <Account />
          <Account />
          <Account />
      </UPageGrid>
</template>



<style scoped>
</style>
