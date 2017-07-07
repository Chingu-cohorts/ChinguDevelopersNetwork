import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import CohortList from '@/components/cohort/CohortList'
import ShowCohort from '@/components/cohort/ShowCohort'
import UserList from '@/components/user/UserList'
import ShowUser from '@/components/user/ShowUser'
import PostList from '@/components/forum/PostList'
import SignIn from '@/components/user/SignIn'
import SignUp from '@/components/user/SignUp'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Hello',
      component: Hello
    },
    {
      path: '/cohorts',
      name: 'CohortList',
      component: CohortList
    },
    {
      path: '/cohorts/:name',
      name: 'ShowCohort',
      component: ShowCohort
    },
    {
      path: '/users',
      name: 'UserList',
      component: UserList
    },
    {
      path: '/users/:username',
      name: 'ShowUser',
      component: ShowUser
    },
    {
      path: '/forum',
      name: 'PostList',
      component: PostList
    },
    {
      path: '/signup',
      name: 'SignUp',
      component: SignUp
    },
    {
      path: '/signin',
      name: 'SignIn',
      component: SignIn
    }
  ],
  mode: 'history'
})
