<template>
<div class="box">
  <div class="columns">
    <article class="media">
      <div class="column is-one-quarter-mobile is-2-desktop">
        <div class="media-left">
          <figure class="image is-64x64">
            <img src="https://gravatar.com/avatar/3196ff4ca2acf9464de8a7c5b5e62957?s=512" alt="Image">
          </figure>
        </div>
      </div>

      <div class="column is-half-mobile is-8-desktop">
        <div class="media-content">
          <div class="content">
            <p>
              <a>The purpose of this project</a> <small>@oxyrus</small> <small>31m ago</small>
              <br>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean efficitur sit amet massa fringilla egestas. Nullam condimentum luctus turpis.
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
export default {
  name: 'IndividualPost'
}
</script>

<style scoped>
img {
  border-radius: 50%;
}

figure {
  margin: 0 auto;
}

.media-right h2 {
  font-weight: 700;
}
</style>
