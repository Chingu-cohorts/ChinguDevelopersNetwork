<template>
  <nav class="nav has-shadow">
    <div class="container">
      <div class="nav-left">
        <router-link
          :to="{name: 'Hello'}"
          class="nav-item">
          <img src="https://i.imgur.com/CH034bi.png" alt="The Zoo">
        </router-link>
        <router-link
          :to="{ name: 'CohortList'}"
          exact-active-class="is-active"
          class="nav-item is-tab is-hidden-mobile">
          Cohorts
        </router-link>
        <router-link
          :to="{ name: 'UserList'}"
          exact-active-class="is-active"
          class="nav-item is-tab is-hidden-mobile">
          Users
        </router-link>
        <a class="nav-item is-tab is-hidden-mobile">Forums</a>
        <a class="nav-item is-tab is-hidden-mobile">Projects</a>
      </div>
      <span class="nav-toggle">
        <span></span>
        <span></span>
        <span></span>
      </span>
      <div class="nav-right nav-menu">
        <a class="nav-item is-tab is-hidden-tablet is-active">Home</a>
        <a class="nav-item is-tab is-hidden-tablet">Cohorts</a>
        <a class="nav-item is-tab is-hidden-tablet">Forums</a>
        <a class="nav-item is-tab is-hidden-tablet">Projects</a>
        <a class="nav-item is-tab">
          <figure class="image is-32x32" style="margin-right: 8px;">
            <img src="https://i.imgur.com/N8iiBNO.jpg">
          </figure>
          Profile
        </a>
        <a class="nav-item is-tab">Logout</a>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: 'navbar'
}
</script>
