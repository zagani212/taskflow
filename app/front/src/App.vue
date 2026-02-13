<script setup>
import { onBeforeUnmount, onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const isAuthed = ref(false)

const syncAuth = () => {
  isAuthed.value = window.localStorage.getItem('TOKEN') === 'fake token'
}

const onLogout = async () => {
  window.localStorage.removeItem('token')
  syncAuth()
  await router.push('/login')
}

onMounted(() => {
  syncAuth()
  window.addEventListener('storage', syncAuth)
})

onBeforeUnmount(() => {
  window.removeEventListener('storage', syncAuth)
})

watch(
  () => route.fullPath,
  () => syncAuth()
)
</script>

<template>
  <div class="app-shell">
    <header class="top-bar">
      <div class="brand">
        <span class="brand-mark">TF</span>
        <div>
          <p class="brand-name">Taskflow</p>
          <p class="brand-tag">Workspace</p>
        </div>
      </div>
      <nav class="nav-links">
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/dashboard">Dashboard</RouterLink>
        <RouterLink to="/admin">Admin</RouterLink>
        <RouterLink v-if="!isAuthed" to="/login">Login</RouterLink>
        <RouterLink v-if="!isAuthed" to="/register">Register</RouterLink>
        <button v-else class="nav-button" type="button" @click="onLogout">
          Logout
        </button>
      </nav>
    </header>

    <main class="page">
      <RouterView />
    </main>
  </div>
</template>
