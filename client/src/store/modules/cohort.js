import * as types from '../mutation-types'
import { http } from '@/api'

// Initial state
const state = {
  cohorts: [],
  currentCohort: {}
}

// Actions
const actions = {
  LOAD_COHORTS_LIST ({ commit }) {
    http.get('/cohorts').then(res => {
      commit(types.SET_COHORTS_LIST, { list: res.data })
    }, err => {
      console.error(err)
    })
  },

  LOAD_COHORT_DATA: function ({ commit }, name) {
    http.get('/cohorts/' + name).then(res => {
      commit(types.SET_CURRENT_COHORT, { cohort: res.data })
    }, err => {
      console.error(err)
    })
  }
}

const mutations = {
  [types.SET_COHORTS_LIST] (state, { list }) {
    state.cohorts = list
  },

  [types.SET_CURRENT_COHORT] (state, { cohort }) {
    state.currentCohort = cohort
  }
}

export default {
  state,
  actions,
  mutations
}
