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
    posts: [],
    currentCohort: {},
    currentUser: {},
    loggedUser: {}
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

    LOAD_POSTS_LIST: function ({ commit }) {
      http.get('/posts').then(res => {
        commit('SET_POSTS_LIST', { list: res.data })
      }, err => {
        console.error(err)
      })
    },

    CREATE_FORUM_POST: function ({ commit }, post) {
      return new Promise((resolve, reject) => {
        http.post('/posts', {
          title: post.title,
          content: post.content
        }).then(res => {
          console.log(res.data)
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

    LOAD_LOGGED_USER_DATA: function ({ commit }) {
      http.get('/currentuser').then(res => {
        commit('SET_LOGGED_USER_DATA', { user: res.data })
      }).catch(err => {
        console.error(err)
      })
    },

    LOGOUT_USER: function ({ commit }) {
      window.localStorage.removeItem('token')
      commit('REMOVE_LOGGED_USER_DATA')
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

    SET_POSTS_LIST: (state, { list }) => {
      state.posts = list
    },

    SET_LOGGED_USER_DATA: (state, { user }) => {
      state.loggedUser = user
    },

    REMOVE_LOGGED_USER_DATA: (state) => {
      state.loggedUser = {}
    }
  }
})

export default store
