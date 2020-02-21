<template>
    <canvas class="canvas" :width="minLen" :height="minLen"></canvas>
</template>

<script>
import { getDistance ,getRhumbLineBearing} from 'geolib';
export default {
    name: 'radarImage',
    props: {
        userList: Array,
        me:Object
    },
    data(){
        return {
            ctx:{},
            max:400000
        }
    },
    mounted() {
        this.ctx = this.$el.getContext('2d')
        this.ctx.translate(this.$el.width/2,this.$el.height/2);
        this.ctx.font = "18px arial white";
        this.interval();
    },
    computed:{
        dots(){
            let dd = this.directionAndDistanseWithoutMe(this.userList);
            return {users:dd}
        },
        minLen(){
            //windowの小さい方
            return Math.min(window.innerWidth,window.innerHeight)
        },
        rate(){
            return this.$el.width/this.max/2;//座標 width 200 -> -100~100のために/2
        }
    },
    methods:{
        interval(){
            setInterval(() => {
                this.ctx.clearRect(-1/2*this.$el.width, -1/2* this.$el.height, this.$el.width, this.$el.height );
                this.draw();
            }, 2000);
        },
        draw(){
            this.dots.users.forEach(v=>{
                console.log(v.name)
                this.drawDot(v.name, this.rate*v.zahyo.x,this.rate*v.zahyo.y)
            })
            this.drawEquidistantLine()
        },
        drawDot(name, x, y){
            this.ctx.beginPath();
            this.ctx.arc(x,y , 3,0 , 2 * Math.PI);
            this.ctx.fillStyle = "white";
            this.ctx.fill();
            this.ctx.font = "10px Arial";
            this.ctx.fillStyle = "white";
            this.ctx.fillText(name,x+3, y-8);
        },
        clean(){
            //this.ctx.clearRect(0, 0, 200, 200)
        },
        drawEquidistantLine(){
            this.ctx.beginPath();
            this.ctx.arc(0,0,this.$el.width/2/2, 0 , 2 * Math.PI);
            this.ctx.strokeStyle = "white"
            this.ctx.strokeCircleText = this.max/2 
            this.ctx.stroke();
        },
        directionAndDistanseWithoutMe(list) {
            let users = [];
            list.forEach(user => {
                const direction = this.bearing(
                    this.me.coordinate,
                    user.coordinate,
                )
                const distanse = getDistance(this.me.coordinate,user.coordinate);
                users.push({name:user.name,distanse:distanse,direction:direction,zahyo:
                this.getPointByDistanceAndDegree(distanse,direction)
                });
            });
            return users
        },
        bearing(o,d){
            let brng = getRhumbLineBearing(o,d)
            return brng;
        },
        getPointByDistanceAndDegree(distance, degree){
            const vx = degree <= 180 ? 1:-1;
            const vy = (degree <= 90) || degree >=270? -1:1;// y座標はcanvas場で逆になっている(yの値が大きいと下に行く)
            // 角度をラジアンに変換
            var _radian = Math.PI / 180 * degree;
            // 座標を取得
            var _x = vx * Math.cos(_radian * (Math.PI / 180)) * distance;
            var _y = vy * Math.sin(_radian * (Math.PI / 180)) * distance;
            return{
                x:_x,
                y:_y
            }
        }
    }

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.canvas {
  background: black;
}
</style>
