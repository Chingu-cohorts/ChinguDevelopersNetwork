<template>
<div>
  <section class="hero is-primary">
    <div class="hero-body">
      <div class="container">
        <h1 class="title">
          Sign In
        </h1>
        <h2 class="subtitle">
          Start interacting with other members
        </h2>
      </div>
    </div>
  </section>

  <section class="section">
    <div class="columns">
      <div class="container">
        <div class="column is-half is-offset-one-quarter">
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="text" placeholder="Username">
              <span class="icon is-small is-left">
                <i class="fa fa-user"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control has-icons-left">
              <input class="input" type="password" placeholder="Password">
              <span class="icon is-small is-left">
                <i class="fa fa-lock"></i>
              </span>
            </p>
          </div>
          <div class="field">
            <p class="control">
              <button class="button is-primary is-outlined" type="submit">
                Login
              </button>
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</div>
</template>

<script>
export default {
  name: 'sign-in'
}
</script>
