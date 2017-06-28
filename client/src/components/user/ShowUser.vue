<template>
<div>
  <section class="hero user-profile-background">
    <div class="hero-body">
      <div class="container has-text-centered">
        <figure class="image is-256x256">
          <img :src="userGravatar">
        </figure>
        <h1 class="title">
          {{ currentUser.first_name }} {{ currentUser.last_name }}
        </h1>
        <h2 class="subtitle">
          {{ currentUser.username }}
        </h2>
        <a href="https://twitter.com/im_oxy">
          <i class="fa fa-twitter"></i>
        </a>
        <a href="https://github.com/Oxyrus">
          <i class="fa fa-github"></i>
        </a>
        <a href="https://twitter.com/im_oxy">
          <i class="fa fa-linkedin"></i>
        </a>
      </div>
    </div>
  </section>

  <section class="section">
    <div class="container">
      
      <nav class="level">
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Projects Completed</p>
            <p class="title">{{ userCompletedProjects }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Reputation</p>
            <p class="title">{{ userReputationLevel }}</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Recommended by</p>
            <p class="title">243</p>
          </div>
        </div>
      </nav>

      <hr>

      <div class="columns feed">

        <div class="column is-one-quarter">
          <div class="content">
            <h2>About</h2>
            <p>{{ userAbout }}</p>
          </div>
          <div class="block has-text-centered">
            <a class="button is-primary is-outlined">
              Contact
            </a>
          </div>
          <div class="block has-text-centered">
            <a class="button is-danger is-outlined">
              Recommend
            </a>
          </div>
        </div>

        <div class="column">

          <article class="media">
            <figure class="media-left">
              <p class="image is-64x64">
                <img src="http://lorempixel.com/128/128/">
              </p>
            </figure>
            <div class="media-content">
              <div class="content">
                <p>
                  <strong>John Smith</strong> <small>@johnsmith</small> <small>31m</small>
                  <br>
                  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
                </p>
              </div>
              <nav class="level is-mobile">
                <div class="level-left">
                  <a class="level-item">
                    <span class="icon is-small"><i class="fa fa-reply"></i></span>
                  </a>
                  <a class="level-item">
                    <span class="icon is-small"><i class="fa fa-retweet"></i></span>
                  </a>
                  <a class="level-item">
                    <span class="icon is-small"><i class="fa fa-heart"></i></span>
                  </a>
                </div>
              </nav>
            </div>
          </article>

        </div>

      </div>
    </div>
  </section>
</div>
</template>

<script>
import md5 from 'blueimp-md5'

export default {
  name: 'show-user',

  computed: {
    currentUser () {
      return this.$store.state.currentUser
    },

    userAbout (state) {
      if (state.currentUser.about) {
        return state.currentUser.about
      }
      return 'No about'
    },

    userCompletedProjects (state) {
      if (state.currentUser.projects) {
        return state.currentUser.projects
      }

      return '0'
    },

    userReputationLevel (state) {
      let experience = state.currentUser.experience
      let reputation

      let assignRep = (name) => {
        reputation = name
      }

      switch (true) {
        case experience >= 1000:
          assignRep('Legend')
          break
        case experience >= 500:
          assignRep('Sage')
          break
        case experience >= 100:
          assignRep('Senior')
          break
        case experience >= 10:
          assignRep('Member')
          break
        default:
          assignRep('Neutral')
          break
      }

      return reputation
    },

    userGravatar (state) {
      let hash = md5(state.currentUser.email)
      let gravatarUrl = 'https://gravatar.com/avatar/' + hash + '?s=512'
      return gravatarUrl
    }
  },

  mounted () {
    let { username } = this.$route.params
    this.$store.dispatch('LOAD_USER_DATA', username)
  }
}
</script>

<style scoped>
.user-profile-background {
  background-image: linear-gradient(135deg, #81FBB8 0%, #28C76F 100%)
}

.hero-body .title, .subtitle {
  color: #fff;
}

.hero-body .fa {
  color: #fff;
}

.is-256x256 {
  height: 256px;
  width: 256px;
}

.hero-body figure {
  display: inline-flex;
  justify-content: center;
}

.hero-body figure img {
  border-radius: 50%;
  border: 1px solid gray;
  box-shadow: 0 3px 2px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease-in-out;
}

.hero-body figure img:hover {
  border-radius: 0%;
}

a .fa {
  margin: 0 0.3em;
}

hr {
  margin: 0 auto;
  width: 80%;
}

.feed {
  margin-top: 2em;
}

.media .image img {
  border-radius: 50%;
}
</style>
