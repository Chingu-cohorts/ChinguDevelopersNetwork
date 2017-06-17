import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    cohorts: [],
    users: []
  },

  actions: {
    LOAD_COHORTS_LIST: function ({ commit }) {
      axios.get('http://localhost:8081/api/cohorts').then(res => {
        commit('SET_COHORTS_LIST', { list: res.data })
      }, err => {
        console.error(err)
      })
    }
  },

  mutations: {
    SET_COHORTS_LIST: (state, { list }) => {
      state.cohorts = list
    }
  }
})

export default store
