<template>
<article class="media">
  <figure class="media-left">
    <p class="image is-64x64">
      <img src="http://bulma.io/images/placeholders/128x128.png">
    </p>
  </figure>
  <div class="media-content">
    <div class="content">
      <p>
        <strong>Barbara Middleton</strong>
        <br>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis porta eros lacus, nec ultricies elit blandit non. Suspendisse pellentesque mauris sit amet dolor blandit rutrum. Nunc in tempus turpis.
        <br>
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
