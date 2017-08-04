<template>
  <div class="container settings">
    <div class="columns">
      <div class="column is-8-desktop">
        <h1>User Settings</h1>
        <p>Update your preferences</p>

        <h2>Name</h2>
        <div class="field is-grouped">
          <div class="control">
            <input class="input" type="text" placeholder="First name" v-model="user.first_name">
          </div>

          <div class="control">
            <input class="input shrink" type="text" placeholder="Last name" v-model="user.last_name">
          </div>
        </div>

        <div class="field">
          <label class="label">About</label>
          <p class="control is-expanded">
            <textarea class="textarea" v-model="user.about"></textarea>
          </p>
        </div>

        <h2>Social Networks</h2>
        <p>Only include your username.</p>
        <div class="field is-grouped is-grouped-multiline">
          
          <div class="field has-addons">
            <p class="control">
              <a class="button is-static">
                <i class="fa fa-github"></i>
              </a>
            </p>
            <p class="control">
              <input class="input" type="text" placeholder="GitHub username" v-model="user.github_username">
            </p>
          </div>

          <div class="field has-addons">
            <p class="control">
              <a class="button is-static">
                <i class="fa fa-medium"></i>
              </a>
            </p>
            <p class="control">
              <input class="input" type="text" placeholder="Medium username" v-model="user.medium_username">
            </p>
          </div>

          <div class="field has-addons">
            <p class="control">
              <a class="button is-static">
                <i class="fa fa-twitter"></i>
              </a>
            </p>
            <p class="control">
              <input class="input" type="text" placeholder="Twitter username" v-model="user.twitter_username">
            </p>
          </div>

        </div>

        <button class="button is-primary is-outlined" @click="updateUser">Save</button>
      
      </div>

      <div class="column is-4-desktop">
        <div class="user-data">

          <figure class="image is-128x128">
            <img :src="userGravatar" class="avatar">
          </figure>

          <h2 class="has-text-centered">{{ loggedUser.first_name }} {{ loggedUser.last_name }}</h2>
          <p class="has-text-centered">{{ loggedUser.username }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { gravatar } from '@/utils'
import { http } from '@/api'

export default {
  name: 'settings',

  data () {
    return {
      user: {
        first_name: '',
        last_name: '',
        about: '',
        github_username: '',
        medium_username: '',
        twitter_username: ''
      }
    }
  },

  computed: {
    loggedUser () {
      return this.$store.state.user.loggedUser
    },

    userGravatar (state) {
      let { email } = state.loggedUser
      return gravatar(email)
    }
  },

  methods: {
    updateUser (e) {
      let { user } = this

      http.put('/users', {
        ...user
      }).then(res => {
        this.$router.push({ name: 'ShowUser', params: { username: this.loggedUser.username } })
      }, err => {
        console.error(err)
      })
    }
  }
}
</script>

<style scoped>
h1, h2, h3 {
  font-weight: 700;
}

.settings {
  margin-top: 2em;
  margin-bottom: 2em;
}

.settings h1 {
  font-size: 2em;
}

.settings .user-data {
  padding: 2em 0;
  color: #fff;
  background-color: #15df89;
}

.settings .user-data figure {
  margin: 0 auto;
}

.settings .user-data h2 {
  font-weight: 700;
}
</style>
