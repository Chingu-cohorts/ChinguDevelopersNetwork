import * as types from '../mutation-types'
import { http } from '@/api'

// Initial state
const state = {
  users: [],
  currentUser: {},
  loggedUser: {}
}

// Actions
const actions = {
  LOAD_USERS_LIST ({ commit }) {
    http.get('/users').then(res => {
      commit(types.SET_USERS_LIST, { list: res.data })
    }, err => {
      console.error(err)
    })
  },

  LOAD_USER_DATA ({ commit }, username) {
    http.get('/users/' + username).then(res => {
      commit(types.SET_CURRENT_USER_DATA, { user: res.data })
    }, err => {
      console.error(err)
    })
  },

  LOAD_LOGGED_USER_DATA ({ commit }) {
    http.get('/currentuser').then(res => {
      commit(types.SET_LOGGED_USER_DATA, { user: res.data })
    }).catch(err => {
      console.error(err)
    })
  },

  LOGOUT_USER ({ commit }) {
    window.localStorage.removeItem('token')
    commit(types.REMOVE_LOGGED_USER_DATA)
  }
}

// Mutations
const mutations = {
  [types.SET_USERS_LIST] (state, { list }) {
    state.users = list
  },

  [types.SET_CURRENT_USER_DATA] (state, { user }) {
    state.currentUser = user
  },

  [types.SET_LOGGED_USER_DATA] (state, { user }) {
    state.loggedUser = user
  },

  [types.REMOVE_LOGGED_USER_DATA] (state) {
    state.loggedUser = {}
  }
}

export default {
  state,
  actions,
  mutations
}
