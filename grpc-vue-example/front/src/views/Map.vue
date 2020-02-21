<template>
  <div id="map-wrap">
      aaaaaaa
    <router-link to="/" class="mainMapLink">
      <a>トップページへ移動する</a>
    </router-link>
    <MglMap
      :accessToken="accessToken"
      :mapStyle="mapStyle"
      :zoom.sync="zoom"
      :center="center"
    >
      <MglNavigationControl />
    </MglMap>
  </div>
</template>
<script>
import { RoomHandlerClient } from '@/assets/room_grpc_web_pb.js'
import { RoomRequest } from '@/assets/room_pb.js'
import { User,Coordinate } from '@/assets/user_pb.js'
import "mapbox-gl/dist/mapbox-gl.css"; // mapbox 用の CSS
import{
  MglMap,
  MglMarker,
  MglNavigationControl
} from "vue-mapbox"; 

import { getDistance ,getCenter} from 'geolib';

export default {
  name: "mapBox",
  components: {
    MglMap, // メインの地図
    MglNavigationControl // 拡大縮尺等のコントローラーコンポーネント
  },
  data: function () {
    return {
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
      accessToken: 'pk.eyJ1IjoicnlvbWFrIiwiYSI6ImNrMW9yYWVzbDBwYzkzZHBnOHo0OHo3MGMifQ.mQ54Y49eqhaAuLljonRLJQ',
      zoom: 17,
      mapStyle: "mapbox://styles/mapbox/streets-v10", // 見た目。色々あるが標準のものを採用
      center: { lon: 139.7009177, lat: 35.6580971 } // 地図の中心地
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
    centerCoords(){
        let coords = getCenter();
        return coords?coords:{latitude:14,longitude:14}
    }
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
