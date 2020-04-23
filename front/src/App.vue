<template>
  <div id="app" class="bg-img">
      <app-header :key="componentKey">
        <a href="#" @click="setLocale('en')" style="padding-right: 3px"><flag iso="us"></flag></a>
        <a href="#" @click="setLocale('ru')"><flag iso="ru"></flag></a>
        <app-search class="is-hidden-touch"></app-search>
      </app-header>
<!--      <Chart/>-->
      <transition name="fade" mode="out-in" appear>
        <router-view class="bg-img" ref="comp" :key="componentKey" v-on:update="updateKey"></router-view>
      </transition>
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Search from './components/SearchPlayer.vue'
// const path = './public/inflicted.mp3';

export default {
  name: 'app',
  components: {
    'app-header': Header,
    'app-search': Search,
      // Chart
  },
  data(){
    return{
      hello:'hello title',
      componentKey: 0,
    }
  },
    created() {
      ///////////////////////////////////////////////////////////////////
        this.interval = setInterval(() => this.getOrders(), 50000);
    },
    computed: {
        count () {
            return this.$store.state.orders.orders.length
        }
    },
    watch: {
        count (newCount, oldCount) {
            if(newCount > oldCount){
                this.playSound()
                console.log("New Order!")
            }
        }
    },
  methods:{
      playSound() {
          const path = '/inflicted.mp3';
          const audio = new Audio(path);
          let playPromise = audio.play();

          if (playPromise !== undefined) {
              playPromise.then(function(){
                  console.log('Did you hear that?');
              })
                  .catch(error => {
                      console.log(`playSound error: ${error}`);
                  });
          }
      },
    setLocale(locale){
        import(`./langs/${locale}.json`).then((msgs) => {
            this.$i18n.setLocaleMessage(locale, msgs);
            this.$i18n.locale = locale;
            this.$validator.locale = locale;
            this.forceRerender();
        })
    },
    forceRerender() {
      this.componentKey += 1;
    },
      updateKey(){
          console.log('yes')
          this.forceRerender()
      },
      getOrders(){
          console.log("function run!");
          this.$store.dispatch('orders/getOrders');
      }
  },


}
</script>

<style>

  [v-cloak] { display: none; }
  .pl-15{
      padding-left: 15px !important;
  }
  .plr-20{
      padding-top: 2% !important;
      padding-left: 20% !important;
      padding-right: 20% !important;
  }

  .plr-30{
      padding-top: 2% !important;
      padding-left: 30% !important;
      padding-right: 30% !important;
  }



  .pt-30{
      padding-top: 30px !important;
  }

  .white{
      background-color: rgba(254, 251, 235, 0.5) !important;
  }

  .plr-15{
      padding-top: 2% !important;
      padding-left: 15% !important;
      padding-right: 15% !important;
  }

  .plr-10{
      padding-top: 2% !important;
      padding-left: 4% !important;
      padding-right: 4% !important;
  }


  .p-3 {
    padding: 20px 50px 0px 50px !important;
  }
  .pd-3 {
    padding: 0px 50px 50px 50px !important;
  }
  .pt-2 {
    padding-top: 20px !important;
  }
  .p-10 {
      padding: 10px !important;
  }
  .pl-2 {
      padding-top: 10px !important;
  }
  .p-4 {
    padding: 10px 200px 0px 200px !important;
  }
  .m-3 {
    margin: 0px 5px 0px 5px !important;
  }
  .pb-3 {
    padding-bottom: 10px !important;
  }

  /*.container{*/
  /*  margin: 0px !important;*/
  /*  padding: 0px !important;*/
  /*}*/


  /*html {*/
  /*  background: url('https://www.wallpaperflare.com/static/472/195/776/planet-sun-digital-art-space-wallpaper.jpg') no-repeat center center fixed;*/
  /*  -webkit-background-size: cover;*/
  /*  -moz-background-size: cover;*/
  /*  -o-background-size: cover;*/
  /*  background-size: cover;*/
  /*}*/


  .bg-img {
    background-image: url("assets/54088765-healthy-food-background-studio-photo-of-different-fruits-on-white-wooden-table-high-resolution-produ.jpg") ;
    background-position: center center;
    background-repeat:  no-repeat;
    background-attachment: fixed;
    background-size:  cover;
    background-color: #999;

  }
  .tile.is-child.box {
    background-color: rgba(254, 251, 235, 0.5) !important;
  }

  body {
    font-size: 14px !important;
    background-color: grey;
  }

  .title {
    /*color: #363636;*/
    font-size: 1.5rem !important;
    line-height: 0.5 !important;
    /*font-weight: 600;*/
    /*line-height: 1.125;*/
  }
  .subtitle {
    font-size: 1rem !important;
    line-height: 0.9 !important;
  }

  /*TRANSITION*/

  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.2s
  }

  .fade-enter, .fade-leave-active {
    opacity: 0
  }

  .buttons:last-child {
       margin-bottom: 0px !important;
  }

  .p-7{
      padding-right: 25%  ;
      padding-left: 25%;
      padding-top: 20px;
  }
  .p-8{
      padding-right: 20%  ;
      padding-left: 20%;
      padding-top: 20px;
  }
  .p-20{
      padding-top: 2em;
      padding-left: 2em;
      padding-right: 2em;
  }
  .pb-20{
      padding-bottom: 20%;
  }
    /*.is-active{*/
    /*    background-color: red;*/
    /*}*/
</style>
