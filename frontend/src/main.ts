import './assets/main.css'
import ui from '@nuxt/ui/vue-plugin'
import { createApp, onMounted } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'

import App from './App.vue'
import Main from './screens/Main.vue'
import List from './screens/List.vue'
import Login from './screens/Login.vue'
import Panel from './screens/Panel.vue'

import ErrorNotFound from './screens/Error404.vue'
import { useAuth } from './scripts/useAuth.ts'

const app = createApp(App)

const routes = [
  { path: '/', component: Main },
  { path: '/iti', component: List },
  { path: '/login', component: Login },

  { path: '/panel', component: Panel, meta: { requiresAuth: true }},
  // 404 ошибка
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: ErrorNotFound }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

const { isAuthenticated, me } = useAuth()

onMounted(async () => {
  await me()
})

router.beforeEach((to, from) => {
  console.log(to, from, isAuthenticated.value)
  if (to.meta.requiresAuth && !isAuthenticated.value) {
    return {
      path: '/login',
      query: { redirect: to.fullPath },
    }
  }
})

app.use(router)
app.use(ui)

app.mount('#app')
