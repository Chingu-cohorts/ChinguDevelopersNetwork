<template>
  <div class="container settings">
    <div class="columns">
      <div class="column is-8-desktop">
        <h1>User Settings</h1>
        <p>Update your preferences</p>

        <div class="field">
          <label class="label">First Name</label>
          <div class="control">
            <input class="input shrink" type="text" placeholder="First name" v-model="user.first_name">
          </div>
        </div>

        <div class="field">
          <label class="label">Last Name</label>
          <div class="control">
            <input class="input shrink" type="text" placeholder="Last name" v-model="user.last_name">
          </div>
        </div>

        <div class="field">
          <label class="label">About</label>
          <div class="control">
            <textarea class="textarea" v-model="user.about"></textarea>
          </div>
        </div>

        <label class="label">GitHub</label>
        <div class="field has-addons">
          <p class="control">
            <a class="button is-static">
              github.com/
            </a>
          </p>
          <p class="control">
            <input class="input" type="text" placeholder="GitHub username" v-model="user.github_username">
          </p>
        </div>

        <label class="label">Medium</label>
        <div class="field has-addons">
          <p class="control">
            <a class="button is-static">
              medium.com/@
            </a>
          </p>
          <p class="control">
            <input class="input" type="text" placeholder="Medium username" v-model="user.medium_username">
          </p>
        </div>

        <label class="label">LinkedIn</label>
        <div class="field has-addons">
          <p class="control">
            <a class="button is-static">
              linkedin.com/
            </a>
          </p>
          <p class="control">
            <input class="input" type="text" placeholder="LinkedIn username" v-model="user.linkedin_username">
          </p>
        </div>

        <label class="label">Twitter</label>
        <div class="field has-addons">
          <p class="control">
            <a class="button is-static">
              twitter.com/
            </a>
          </p>
          <p class="control">
            <input class="input" type="text" placeholder="Twitter username" v-model="user.twitter_username">
          </p>
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
        linkedin_username: '',
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
.settings {
  margin-top: 2em;
  margin-bottom: 2em;
}

.settings h1 {
  font-weight: 700;
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

.settings .input {
  min-width: 400px;
  max-width: 400px;
}

.settings .shrink {
  width: 40%;
}
</style>
