<template>
<div class="box">
  <div class="columns">
    <article class="media">
      <div class="column is-one-quarter-mobile is-2-desktop">
        <div class="media-left">
          <figure class="image is-64x64">
            <img :src="userGravatar" alt="User avatar">
          </figure>
        </div>
      </div>

      <div class="column is-half-mobile is-8-desktop">
        <div class="media-content">
          <div class="content">
            <p>
              <router-link
                :to="{name: 'ShowPost', params: { id: post.id, slug: postSlug }}">
                {{ post.title }}
                </router-link> <small>@{{ post.user.username }}</small> <small>{{ timeAgo }}</small>
              <br>{{ post.content }}
            </p>
          </div>
        </div>
      </div>

      <div class="column is one-quarter-mobile is-2-desktop">
        <div class="media-right has-text-centered">
          <h2>2K</h2>
          <span>Comments</span>
        </div>
      </div>
    </article>
  </div>
</div>
</template>

<script>
import md5 from 'blueimp-md5'
import moment from 'moment'

export default {
  name: 'IndividualPost',

  props: ['post'],

  computed: {
    userGravatar (props) {
      let hash = md5(props.post.user.email)
      let gravatarUrl = 'https://gravatar.com/avatar/' + hash + '?s=512'
      return gravatarUrl
    },

    timeAgo (props) {
      let date = new Date(props.post.updated_at)
      return moment(date, 'YYYYMMDD').fromNow()
    },

    postSlug (props) {
      let { title } = props.post
      return title.toLowerCase().split(' ').join('-')
    }
  }
}
</script>

<style scoped>
img {
  border-radius: 50%;
}

figure {
  margin: 0 auto;
}

.media {
  width: 100%;
}

.media-right h2 {
  font-weight: 700;
}
</style>
