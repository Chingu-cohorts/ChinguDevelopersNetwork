import Vue from 'vue'
import Router from 'vue-router'
import Hello from '@/components/Hello'
import CohortList from '@/components/cohort/CohortList'
import ShowCohort from '@/components/cohort/ShowCohort'
import UserList from '@/components/user/UserList'
import ShowUser from '@/components/user/ShowUser'
import PostList from '@/components/forum/PostList'
import ShowPost from '@/components/forum/ShowPost'
import NewPost from '@/components/forum/NewPost'
import SignIn from '@/components/user/SignIn'
import SignUp from '@/components/user/SignUp'
import Settings from '@/components/user/Settings'
import TermsOfService from '@/components/misc/TermsOfService'
import Privacy from '@/components/misc/Privacy'

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
      path: '/signup',
      name: 'SignUp',
      component: SignUp
    },
    {
      path: '/signin',
      name: 'SignIn',
      component: SignIn
    },
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    },
    {
      path: '/forum',
      name: 'PostList',
      component: PostList
    },
    {
      path: '/forum/:id(\\d+)/:slug',
      name: 'ShowPost',
      component: ShowPost
    },
    {
      path: '/new-post',
      name: 'NewPost',
      component: NewPost
    },
    {
      path: '/tos',
      name: 'TermsOfService',
      component: TermsOfService
    },
    {
      path: '/privacy',
      name: 'PrivacyPolicy',
      component: Privacy
    }
  ],
  mode: 'history'
})
