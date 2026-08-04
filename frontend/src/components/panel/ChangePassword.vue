<script setup lang="ts">
import * as z from 'zod'
import type { FormSubmitEvent } from '@nuxt/ui'
import { reactive } from 'vue'
import { useAuth } from '../../scripts/useAuth'


const { changePassword } = useAuth()

const schema = z.object({
  oldPassword: z.string('Старый пароль обязателен'),
  newPassword: z.string('Новый пароль обязателен').min(4, 'Минимум 4 символа')
})

type Schema = z.output<typeof schema>

const state = reactive<Partial<Schema>>({
  oldPassword: undefined,
  newPassword: undefined
})

const toast = useToast()
async function onSubmit(event: FormSubmitEvent<Schema>) {
  console.log(event.data)
  const result = await changePassword(event.data.oldPassword, event.data.newPassword)
  if (result) {
    toast.add({ title: 'Success', description: 'Password has been changed.', color: 'success' })
  } else {
    toast.add({ title: 'Error', description: 'Failed to change password.', color: 'error' })
  }

}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4 w-full" @submit="onSubmit">
    <UFormField label="Старый пароль" name="oldPassword">
      <UInput v-model="state.oldPassword" type="password" class="w-full" />
    </UFormField>

    <UFormField label="Новый пароль" name="newPassword">
      <UInput v-model="state.newPassword" type="password" class="w-full" />
    </UFormField>

    <UButton type="submit" block>
        Сменить пароль
    </UButton>
  </UForm>
</template>
