<template>
  
  <div class="column is-one-third">
    <router-link :to="{ name: 'ShowCohort', params: { name: cohort.name }}">
    <div class="card">
      <div class="card-image">
        <figure class="image is-4by3">
          <img src="http://bulma.io/images/placeholders/1280x960.png" alt="Image">
        </figure>
      </div>
      <div class="card-content">
        <div class="media">
          <div class="media-content">
            <p class="title is-4 has-text-centered">{{ cohort.name }}</p>
          </div>
        </div>

        <div class="content has-text-centered">
          {{ cohort.description }}
        </div>
      </div>
    </div>
    </router-link>
  </div>
  
</template>

<script>
export default {
  name: 'individual-cohort',
  props: ['cohort']
}
</script>

<style scoped>
.card .title {
  color: #363636;
}
</style>

