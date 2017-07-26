<template>
  <div class="container">
    <div class="post-container">
      <template v-if="!currentPost.title">
        <spinner></spinner>
      </template>
      <template v-else>
        <div class="forum-post">
          <div class="columns">
            <div class="column is-hidden-mobile is-2-desktop">
              <div class="box has-text-centered user-box">
                <figure class="image is-128x128">
                  <img :src="userGravatar" class="avatar" :alt="currentPost.user.username">
                </figure>
                <div>
                  <router-link 
                    :to="{ name: 'ShowUser', params: { username: currentPost.user.username } }"
                    class="username"
                    v-bind:class="{ admin: currentPost.user.is_admin }">{{ currentPost.user.username }}</router-link>
                </div>
              </div>
            </div>

            <div class="column is-12-mobile is-10-desktop">
              <div class="box" style="padding-top: 3.5em">
                <div class="content">
                  <div class="has-text-centered">
                    <h1 class="post-title">{{ currentPost.title }}</h1>
                    <p class="is-hidden-tablet">
                      <router-link
                        :to="{ name: 'ShowUser', params: { username: currentPost.user.username } }"
                        class="username"
                        v-bind:class="{ admin: currentPost.user.is_admin }">
                        {{ currentPost.user.username }}
                      </router-link>
                    </p>
                  </div>
                  <hr>
                  <div class="post-content" v-html="content">{{ currentPost.content }}</div>
                </div>
                <div class="post-actions" v-if="loggedUser.id === currentPost.user.id">
                  <div class="block is-clearfix">
                    <a class="button is-danger is-outlined is-small is-pulled-right">
                      <i class="fa fa-trash"></i>
                    </a>
                    <a class="button is-primary is-outlined is-small is-pulled-right">
                      <i class="fa fa-pencil"></i>
                    </a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </template>

      <individual-comment
        v-for="comment in currentPost.comments"
        :comment="comment"
        :key="comment.id">
      </individual-comment>

      <div class="new-comment" v-if="!currentPost.isClosed">
        <new-comment></new-comment>
      </div>
      <div class="post-closed" v-else>
        <p>This post is closed.</p>
      </div>
    
    </div>  
  </div>
</template>

<script>
import MarkdownIt from 'markdown-it'
import IndividualComment from '@/components/forum/IndividualComment'
import NewComment from '@/components/forum/NewComment'
import Spinner from '@/components/Spinner'
import { gravatar } from '@/components/utils'

export default {
  name: 'ShowPost',

  data () {
    return {
      interval: null
    }
  },

  components: {
    NewComment,
    Spinner,
    IndividualComment
  },

  methods: {
    loadPost () {
      let { id } = this.$route.params
      this.$store.dispatch('LOAD_FORUM_POST', id)
    }
  },

  computed: {
    currentPost () {
      return this.$store.state.forum.currentPost
    },

    loggedUser () {
      return this.$store.state.user.loggedUser
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

    this.interval = setInterval(() => this.loadPost(), 5000)
  },

  beforeDestroy () {
    clearInterval(this.interval)
  }
}
</script>

<style scoped>
.container {
  margin-top: 2em;
  margin-bottom: 2em;
}

.post-container {
  background-color: #f9f9f9;
  padding: 0 0.5em 1em;
  border-radius: 0.2em;
}

hr {
  width: 40%;
  margin: 1.5em auto;
}

.forum-post {
  margin-top: 1em;
  margin-bottom: 1em;
}

.forum-post .post-title {
  text-transform: uppercase;
}

.post-container figure img {
  margin-top: -1em;
}

.user-box {
  border-radius: 0;
  box-shadow: none;
  background-color: transparent;
}

.username {
  font-weight: 700;
}

.admin {
  color: #000;
}

.post-content {
  color: #333;
  text-align: justify;
  line-height: 1.7em;
  letter-spacing: 1px;
}

.post-actions .block {
  margin-top: 1em;
}

.post-actions .block .button {
  margin-left: 1em;
}
</style>
