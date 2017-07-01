import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const jwt = window.localStorage.getItem('token')

const http = axios.create({
  baseURL: 'http://localhost:8081/api',
  timeout: 5000,
  headers: {
    'Authorization': jwt
  }
})

const store = new Vuex.Store({
  state: {
    cohorts: [],
    users: [],
    currentCohort: {},
    currentUser: {},
    collapsed: false
  },

  actions: {
    LOAD_COHORTS_LIST: function ({ commit }) {
      http.get('/cohorts').then(res => {
        commit('SET_COHORTS_LIST', { list: res.data })
      }, err => {
        console.error(err)
      })
    },

    LOAD_COHORT_DATA: function ({ commit }, name) {
      http.get('/cohorts/' + name).then(res => {
        commit('SET_CURRENT_COHORT', { cohort: res.data })
      }, err => {
        console.error(err)
      })
    },

    LOAD_USERS_LIST: function ({ commit }) {
      http.get('/users').then(res => {
        commit('SET_USERS_LIST', { list: res.data })
      }, err => {
        console.error(err)
      })
    },

    LOAD_USER_DATA: function ({ commit }, username) {
      http.get('/users/' + username).then(res => {
        commit('SET_CURRENT_USER_DATA', { user: res.data })
      }, err => {
        console.error(err)
      })
    },

    POST_REGISTRATION_DATA: function ({ commit }, user) {
      return new Promise((resolve, reject) => {
        http.post('/users', {
          username: user.username,
          email: user.email,
          password: user.password
        }).then(res => {
          console.log(res.data.token)
          resolve()
        }).catch(err => {
          if (err.response) {
            console.log(err.response.data)
          } else {
            console.error(err)
          }
          reject()
        })
      })
    },

    POST_LOGIN_DATA: function ({ commit }, user) {
      return new Promise((resolve, reject) => {
        http.post('/users/login', {
          username: user.username,
          password: user.password
        }).then(res => {
          console.log(res)
          localStorage.setItem('token', res.data.token)
          resolve()
        }).catch(err => {
          if (err.response) {
            console.log(err.response)
          } else {
            console.error(err)
          }
          reject()
        })
      })
    },

    COLLAPSE_NAVBAR: function ({ commit }) {
      commit('COLLAPSE_MAIN_NAVBAR')
    }
  },

  mutations: {
    SET_COHORTS_LIST: (state, { list }) => {
      state.cohorts = list
    },

    SET_CURRENT_COHORT: (state, { cohort }) => {
      state.currentCohort = cohort
    },

    SET_USERS_LIST: (state, { list }) => {
      state.users = list
    },

    SET_CURRENT_USER_DATA: (state, { user }) => {
      state.currentUser = user
    },

    COLLAPSE_MAIN_NAVBAR: (state) => {
      state.collapsed = !state.collapsed
    }
  }
})

export default store
