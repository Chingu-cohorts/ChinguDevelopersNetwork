import Vue from 'vue'
import Vuex from 'vuex'

import cohort from './modules/cohort'
import forum from './modules/forum'
import user from './modules/user'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    cohort,
    forum,
    user
  }
})

export default store
