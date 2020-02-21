<template>
  <div class="home">
    this app is [[.AppName]]
    <input v-model="roomName" placeholder="room name">
    <input v-model="userName" placeholder="user name">
    <router-link :to="`/radar/${roomName}/${userName}`" class="button">Join Room</router-link>
  </div>
</template>

<script>
export default {
  name: 'home',
  data(){
    return {
        roomName:"",
        userName:""
    }
  },
}
</script>
