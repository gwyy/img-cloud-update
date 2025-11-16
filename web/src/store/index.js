import {defineStore} from 'pinia'
import { ref } from 'vue'


export const useTokenStore = defineStore('token',()=>{
    const token = ref('')
    const isLoginState = ref(false)
    const login = (token_string) => {
        token.value = token_string
        isLoginState.value = true
    }
    const isLogin = () => {
        return isLoginState.value
    }
    const logout = () => {
        token.value = ''
        isLoginState.value = false
    }
    return {
        token,
        isLoginState,
        isLogin,
        login,
        logout
    }
}, {
    persist: true,
  })
