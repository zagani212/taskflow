<script setup>
import axios from "axios"
import { ref } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()

const loginForm = ref({
    email: "",
    password: ""
})

const onSubmit = async (event) => {
  event.preventDefault()
  const response = await axios.post("http://localhost:8080/auth/login", loginForm.value)
  if(response.status == 200){
    localStorage.setItem('TOKEN', response.data.token)
    router.push('/')
  }
}
</script>

<template>
  <section class="auth">
    <div class="auth-card">
      <div class="auth-header">
        <p class="eyebrow">Welcome back</p>
        <h2>Log in to Taskflow</h2>
        <p class="muted">Access your dashboard and keep tasks moving.</p>
      </div>
      <form class="auth-form" @submit="onSubmit">
        <label class="field">
          <span>Email</span>
          <input type="email" v-model="loginForm.email" placeholder="name@company.com" required />
        </label>
        <label class="field">
          <span>Password</span>
          <input type="password" v-model="loginForm.password" placeholder="••••••••" required />
        </label>
        <button class="btn primary" type="submit">Sign in</button>
        <p class="muted small">Forgot your password? Contact your admin.</p>
        <p class="muted small auth-footer">
          New to Taskflow?
          <RouterLink class="link" to="/register">Create an account</RouterLink>
        </p>
      </form>
    </div>
  </section>
</template>
