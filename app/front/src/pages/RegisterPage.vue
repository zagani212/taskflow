<script setup>
import axios from "axios"
import { ref } from "vue"

const registerForm = ref({
  email: "",
  password: "",
})

const onSubmit = async (event) => {
  event.preventDefault()
  const response = await axios.post(
    "http://localhost:8080/auth/register",
    registerForm.value
  )
  console.log(response)
}
</script>

<template>
  <section class="auth">
    <div class="auth-card">
      <div class="auth-header">
        <p class="eyebrow">Get started</p>
        <h2>Create your account</h2>
        <p class="muted">Join Taskflow to organize work with your team.</p>
      </div>
      <form class="auth-form" @submit="onSubmit">
        <label class="field">
          <span>Email</span>
          <input
            v-model="registerForm.email"
            type="email"
            placeholder="name@company.com"
            required
          />
        </label>
        <label class="field">
          <span>Password</span>
          <input
            v-model="registerForm.password"
            type="password"
            placeholder="Create a password"
            required
          />
        </label>
        <button class="btn primary" type="submit">Create account</button>
        <p class="muted small auth-footer">
          Already have an account?
          <RouterLink class="link" to="/login">Sign in</RouterLink>
        </p>
      </form>
    </div>
  </section>
</template>
