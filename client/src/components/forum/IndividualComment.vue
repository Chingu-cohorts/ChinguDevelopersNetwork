<template>
<div class="forum-comment columns">

  <div class="column is-hidden-mobile is-2-desktop">
    <div class="box has-text-centered user-box">
      <figure class="image is-128x128">
        <img :src="userGravatar" class="avatar" :alt="comment.user.username">
      </figure>
      <div>
        <router-link :to="{ name: 'ShowUser', params: { username: comment.user.username } }" class="username" v-bind:class="{ admin: comment.user.is_admin }">{{ comment.user.username }}</router-link>
      </div>
    </div>
  </div>

  <div class="column is-12-mobile is-10-desktop">
    <div class="box">
      <div class="user-data is-hidden-tablet">
        <p>By <router-link :to="{ name: 'ShowUser', params: { username: comment.user.username } }" class="username" v-bind:class="{ admin: comment.user.is_admin }">{{ comment.user.username }}</router-link></p>
      </div>
      <div class="content">
        <div class="comment-content" v-html="content">{{ comment.content }}</div>
      </div>
      <div class="comment-actions" v-if="loggedUser.id === comment.user.id">
        <hr>
        <div class="block is-clearfix">
          <a class="button is-danger is-outlined is-small is-pulled-right" @click="deleteComment">
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
</template>

<script>
import MarkdownIt from 'markdown-it'
import { gravatar } from '@/utils'
import { http } from '@/api'

export default {
  name: 'individual-comment',

  props: ['comment'],

  computed: {
    loggedUser () {
      return this.$store.state.user.loggedUser
    },

    userGravatar (props) {
      let { email } = props.comment.user
      return gravatar(email)
    },

    content (props) {
      let md = new MarkdownIt()
      return md.render(props.comment.content)
    }
  },

  methods: {
    deleteComment (e) {
      e.preventDefault()

      let postId = this.$route.params.id
      let commentId = this.comment.id

      http.delete(`/posts/${postId}/comments/${commentId}`).then(res => {
        this.$store.dispatch('LOAD_FORUM_POST', postId)
      }).catch(err => {
        if (err.response) {
          console.log(err.response.data)
        } else {
          console.error(err)
        }
      })
    }
  }
}
</script>

<style scoped>
.forum-comment {
  margin-top: 1em;
  margin-bottom: 1em;
}

.forum-comment figure img {
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

.comment-content {
  color: #333;
  text-align: justify;
  line-height: 1.7em;
  letter-spacing: 1px;
}

.comment-actions .block {
  margin-top: 1em;
}

.comment-actions .block .button {
  margin-left: 1em;
}
</style>
