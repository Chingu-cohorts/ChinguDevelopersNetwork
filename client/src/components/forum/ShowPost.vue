<template>
  <div class="container">
    <template v-if="!currentPost.title">
      <div class="loader">Loading...</div>
    </template>
    <template v-else>
      <div class="forum-post">
        <div class="columns">
          <div class="column is-hidden-mobile is-2-desktop">
            <div class="box has-text-centered user-box">
              <figure class="image is-128x128">
                <img :src="userGravatar">
              </figure>
              <div>
                <a class="username">{{ currentPost.user.username }}</a>
              </div>
            </div>
          </div>

          <div class="column is-12-mobile is-10-desktop">
            <div class="box">
              <div class="content">
                <h1 class="has-text-centered">{{ currentPost.title }}</h1>
                <hr>
                <div v-html="content">{{ currentPost.content }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import MarkdownIt from 'markdown-it'
import { gravatar } from '@/components/utils'

export default {
  name: 'ShowPost',

  methods: {
    stuff () {
      let { id, slug } = this.$route.params
      return id + slug
    }
  },

  computed: {
    currentPost () {
      return this.$store.state.forum.currentPost
    },

    userGravatar (state) {
      let { email } = state.currentPost.user
      return gravatar(email)
    },

    content (state) {
      let md = new MarkdownIt()
      return md.render(state.currentPost.content)
    }
  },

  mounted () {
    let { id } = this.$route.params
    this.$store.dispatch('LOAD_FORUM_POST', id)
  }
}
</script>

<style scoped>
hr {
  width: 40%;
  margin: 0 auto;
}

.forum-post {
  margin-top: 1em;
  margin-bottom: 1em;
}

figure img {
  border-radius: 50%;
  margin: 0.5em 0;
}

.user-box {
  border-radius: 0;
  box-shadow: none;
}

.username {
  font-weight: 700;
}

.loader,
.loader:before,
.loader:after {
  background: #00d1b2;
  -webkit-animation: load1 1s infinite ease-in-out;
  animation: load1 1s infinite ease-in-out;
  width: 1em;
  height: 4em;
}
.loader {
  color: #00d1b2;
  text-indent: -9999em;
  margin: 88px auto;
  position: relative;
  font-size: 11px;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  position: absolute;
  top: 0;
  content: '';
}
.loader:before {
  left: -1.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 1.5em;
}
@-webkit-keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}
@keyframes load1 {
  0%,
  80%,
  100% {
    box-shadow: 0 0;
    height: 4em;
  }
  40% {
    box-shadow: 0 -2em;
    height: 5em;
  }
}
</style>
