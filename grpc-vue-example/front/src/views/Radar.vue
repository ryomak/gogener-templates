<template>
  <div class="radar">
    <radar-image :userList="listWithoutMe" :me="meData"/>
  </div>
</template>
<script>
import { RoomHandlerClient } from '@/assets/room_grpc_web_pb.js'
import { RoomRequest } from '@/assets/room_pb.js'
import { User,Coordinate } from '@/assets/user_pb.js'
import radarImage from '@/components/radarImage.vue'

export default {
  name: 'radar',
  components: {
    radarImage
  },
  data: function () {
    return {
      inputField: '',
      distanse:0.0,
      client:null,
      roomData:{
        name:'',
        userList:[],
      },
      meData:{
        coordinate:{
          latitude:0.0,
          longitude:0.0
        },
        heading:0.0
      },
    }
  },
  created: function () {
    // eslint-disable-next-line
    this.client = new RoomHandlerClient('http://localhost:8080', null, null)
    this.updateRoom()
  },
  mounted() {
    this.interval();
  },
  computed:{
    listWithoutMe(){
      return this.roomData.usersList?this.roomData.usersList.filter(u => u.name != this.$route.params.nickname):[];
    },
  },
  methods: {
    interval: function(){
      setInterval(() => {
        this.getGeo();
        console.log(this.roomData)
      }, 2000);
    },
    updateRoom: function (latitude,longitude) {
      // eslint-disable-next-line
      let coordinate = new Coordinate();
      coordinate.setLatitude(latitude);
      coordinate.setLongitude(longitude);
      let user = new User();
      user.setName(this.$route.params.nickname)
      user.setCoordinate(coordinate);
      let request = new RoomRequest()
      request.setName(this.$route.params.roomname);
      request.setUser(user);
      // eslint-disable-next-line
      this.client.getRoom(request, {}, (err, response) => {
        this.roomData = response.toObject();
      })
    },
    getGeo(){
      new Promise((resolve, reject) => {
        if (window.navigator.geolocation) {
          window.navigator.geolocation.getCurrentPosition(
            (response) => {
              resolve(response);
            },
            (e) => {
              console.log(e)
              //alert('データの取得中にエラーが発生しました。');
            },
            {
              enableHighAccuracy: true,
              timeout: 10000,
              maximumAge: 10000,
            },
          );
        } else {
          reject('このブラウザでは位置情報を取得できません');
        }
      })
      .then((response) => {
        
        this.meData.heading = response.coords.heading
        this.meData.coordinate.longitude = response.coords.longitude
        this.meData.coordinate.latitude = response.coords.latitude
        this.updateRoom(response.coords.latitude,response.coords.longitude)
      })
      .catch((e) => {
        console.log(e)
        alert(
          'データが取得できませんでした。電波の届きやすい場所で再度お試しください。',
        );
      });
    }
  }
}
</script>
