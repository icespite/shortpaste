<template>
  <div>
    <form @submit.prevent="handleSubmit">
      <label>
        Username:
        <input v-model="username" required />
      </label>
      <br />
      <label>
        Password:
        <input type="password" v-model="password" required />
      </label>
      <br />
      <button type="submit">Login</button>
    </form>
    <button @click="mockLogin">Mock Login</button>
    <div v-if="error">{{ error }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import axios from 'axios'
import type { UserProfile } from '@/types'

export default defineComponent({
  name: 'LoginView',
  setup() {
    const username = ref('')
    const password = ref('')
    const error = ref('')

    const router = useRouter()
    const store = useStore()

    const mockLogin = () => {
      const mockUserProfile: UserProfile = {
        id: '1337',
        name: username.value
      }
      store.commit('login', mockUserProfile)
      router.push('/')
    }

    const handleSubmit = async () => {
      try {
        const response = await axios.post('/api/v1/login', {
          username: username.value,
          password: password.value
        })

        if (response.status === 200 && response.data) {
          // Store User-Profile
          store.commit('login', response.data)
          // Redirect to home page after successful login
          router.push('/')
        }
      } catch (err) {
        if (
          axios.isAxiosError<Error>(err) &&
          err.response &&
          err.response.status >= 400 &&
          err.response.status < 500
        ) {
          error.value = 'Invalid credentials or other client error'
        } else {
          error.value = 'Server error or network issue'
        }
      }
    }

    return {
      username,
      password,
      mockLogin,
      error,
      handleSubmit
    }
  }
})
</script>
