<template>
  <div class="column is-2">
    <router-link :to="{ name: 'ShowUser', params: { username: user.username }}">
      <img class="user-image" :src="userGravatar">
    </router-link>
  </div>
</template>

<script>
import md5 from 'blueimp-md5'
export default {
  name: 'individual-user',
  props: ['user'],
  computed: {
    userGravatar (props) {
      let hash = md5(props.user.email)
      let gravatarUrl = 'https://gravatar.com/avatar/' + hash + '?s=512'
      return gravatarUrl
    }
  }
}
</script>

<style scoped>
.user-image {
  display: block;
  width: 130px;
  height: 130px;
  margin: 0em auto;
  border-radius: 99em;
  -webkit-border-radius: 99em;
  -moz-border-radius: 99em;
  border: 0px solid gray;
  box-shadow: 0 3px 2px rgba(0, 0, 0, 0.3);
  transition: all 0.3s ease-in-out;
}
.user-image:hover {
  opacity:0.6;
}
</style>
