<script setup lang="ts">
import * as z from 'zod'
import { useRouter, useRoute } from 'vue-router'
import type { FormSubmitEvent, AuthFormField } from '@nuxt/ui'
import { useAuth } from '../scripts/useAuth'

const router = useRouter()
const route = useRoute()
const toast = useToast()
const { login } = useAuth()

const fields: AuthFormField[] = [{
  name: 'login',
  type: 'login',
  label: 'Логин',
  placeholder: 'Введите ваш логин',
  required: true
}, {
  name: 'password',
  label: 'Пароль',
  type: 'password',
  placeholder: 'Введите ваш пароль',
  required: true
  }
]


const schema = z.object({
  login: z.string('Логин обязателен'),
  password: z.string('Пароль обязателен').min(4, 'Минимум 4 символа')
})

type Schema = z.output<typeof schema>

async function onSubmit(payload: FormSubmitEvent<Schema>) {
  const ok = await login(payload.data.login, payload.data.password)
  if (ok) {
    const redirect = (route.query.redirect as string) || '/panel'
    router.push(redirect)
  } else {
    toast.add({ title: 'Ошибка', description: 'Неверный логин или пароль', color: 'error' })
  }
}
</script>

<template>
  <div class="flex flex-col items-center justify-center gap-4 p-4 w-full mt-10">
    <UPageCard class="w-full max-w-md">
      <UAuthForm
        :schema="schema"
        title="Вход"
        icon="i-tabler-user-key"
        :fields="fields"
        @submit="onSubmit"
      />
    </UPageCard>
  </div>
</template>
