<template>
<div>
  <section class="hero is-primary">
    <div class="hero-body">
      <div class="container has-text-centered">
        <figure class="image is-256x256">
          <img src="http://lorempixel.com/300/300/">
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
            <p class="title">42</p>
          </div>
        </div>
        <div class="level-item has-text-centered">
          <div>
            <p class="heading">Reputation</p>
            <p class="title">Sage</p>
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
            <p>I've been coding since I was 12 years old, six years later, I am changing the world.</p>
          </div>
          <div class="block has-text-centered">
            <a class="button is-primary is-outlined">
              Contact
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
import { mapState } from 'vuex'

export default {
  name: 'show-user',

  computed: mapState([
    'currentUser'
  ]),

  mounted () {
    let { username } = this.$route.params
    this.$store.dispatch('LOAD_USER_DATA', username)
  }
}
</script>

<style scoped>
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
