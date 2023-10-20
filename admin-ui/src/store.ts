import { createStore } from 'vuex'
import { type UserProfile } from '@/types'

interface State {
  userProfile: UserProfile | null
}

export const store = createStore<State>({
  state: {
    userProfile: null
  },
  mutations: {
    login(state: State, profile: UserProfile) {
      state.userProfile = profile
    },
    logout(state: State) {
      state.userProfile = null
    }
  },
  getters: {
    isAuthenticated: (state: State) => state.userProfile !== null,
    userProfile: (state: State) => state.userProfile
  }
})
