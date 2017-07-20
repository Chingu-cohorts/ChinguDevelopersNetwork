<template>
<div class="forum-comment">
  <div class="columns">
    <div class="column is-hidden-mobile is-2-desktop">
      <div class="box has-text-centered user-box">
        <figure class="image is-128x128">
          <img :src="userGravatar">
        </figure>
        <div>
          <a class="username" v-bind:class="{ admin: comment.user.is_admin }">{{ comment.user.username }}</a>
        </div>
      </div>
    </div>

    <div class="column is-12-mobile is-10-desktop">
      <div class="box">
        <div class="content">
          <div v-html="content">{{ comment.content }}</div>
        </div>
      </div>
    </div>
  </div>
</div>  
</template>

<script>
import MarkdownIt from 'markdown-it'
import { gravatar } from '@/components/utils'

export default {
  name: 'individual-comment',

  props: ['comment'],

  computed: {
    userGravatar (props) {
      let { email } = props.comment.user
      return gravatar(email)
    },

    content (props) {
      let md = new MarkdownIt()
      return md.render(props.comment.content)
    }
  }
}
</script>

<style scoped>
.forum-comment {
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
  background-color: transparent;
}

.username {
  font-weight: 700;
}

.admin {
  color: #28a0dc;
}

.content {
  line-height: 1.7em;
}
</style>
