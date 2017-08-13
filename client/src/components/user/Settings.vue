<template>
  <div class="container">
    
    <h1 class="has-text-centered">Settings</h1>
    <p class="has-text-centered">Change your preferences and update your profile.</p>

    <div class="columns">

      <div class="column is-4">

        <div class="user-data">

          <figure class="image is-128x128">
            <img :src="userGravatar" class="avatar">
          </figure>

          <h2 class="has-text-centered">{{ loggedUser.first_name }} {{ loggedUser.last_name }}</h2>
          <p class="has-text-centered">{{ loggedUser.username }}</p>
          <p class="has-text-centered">{{ loggedUser.about }}</p>

          <p class="has-text-centered">
            <a><i class="fa fa-twitter"></i></a>
            <a><i class="fa fa-github"></i></a>
            <a><i class="fa fa-medium"></i></a>
          </p>

        </div> <!-- user-data -->

      </div> <!-- col-4 -->

      <div class="column is-8">

        <div class="form-container">

          <div class="field is-grouped">
            <div class="control">
              <input class="input" type="text" placeholder="First name" v-model="user.first_name">
            </div>

            <div class="control">
              <input class="input" type="text" placeholder="Last name" v-model="user.last_name">
            </div>
          </div>

          <div class="field">
            <p class="control is-expanded">
              <textarea class="textarea" v-model="user.about"></textarea>
            </p>
          </div>

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

          <button class="button is-primary is-outlined" @click="updateUser">Update</button>

        </div> <!-- form-container -->

      </div> <!-- col-8 -->

    </div> <!-- columns -->

  </div> <!-- container -->
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
.form-container {
  width: 90%;
  margin: 0 auto;
}

.columns {
  margin: 0;
}

.is-grouped .control {
  max-width: 48%;
}

h1, h2, h3 {
  font-weight: 700;
}

.user-data {
  padding: 2em 0;
  color: #fff;
  background-color: #15df89;
}

.user-data .fa {
  color: #ffffff;
  margin: 1rem;
}

.user-data figure {
  margin: 0 auto;
}

.user-data h2 {
  font-weight: 700;
}

.settings .field:not(:last-child) {
  margin-right: 0.8rem;
}
</style>
