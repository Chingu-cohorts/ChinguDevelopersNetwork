<template>
<article class="media">
  <figure class="media-left">
    <p class="image is-64x64">
      <img :src="userGravatar">
    </p>
  </figure>
  <div class="media-content">
    <div class="content">
      <p>
        <strong>{{ comment.user.username }}</strong>
        <br>{{ comment.content }}<br>
        <small><a>Like</a> · <a>Reply</a> · 3 hrs</small>
      </p>
    </div>
  </div>
</article>  
</template>

<script>
import { gravatar } from '@/components/utils'

export default {
  name: 'individual-comment',

  props: ['comment'],

  computed: {
    userGravatar (props) {
      let { email } = props.comment.user
      return gravatar(email)
    }
  }
}
</script>
