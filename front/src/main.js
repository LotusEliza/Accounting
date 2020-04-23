import Vue from 'vue'
import * as VeeValidate from 'vee-validate';
import App from './App.vue'
import Buefy from 'buefy'
import 'buefy/dist/buefy.css'
import VueRouter from 'vue-router'
import Routes from './routes'
import Vuex from 'vuex';
import {store} from './store/store'
import JQuery from 'jquery'
import moment from 'moment'
import VueTelInput from 'vue-tel-input'
import {i18n} from './plugins/i18n'
import FlagIcon from 'vue-flag-icon'
// import validationMessagesRu from 'vee-validate/dist/locale/ru';
import validationMessagesEn from 'vee-validate/dist/locale/en';


// import VueCharts from 'vue-chartjs'
// import { Bar, Line } from 'vue-chartjs'
// Vue.use(VueCharts);
// Vue.use(Bar);
// Vue.use(Line);

import { MonthPicker } from 'vue-month-picker'
import { MonthPickerInput } from 'vue-month-picker'
// import moment from 'moment'

Vue.use(MonthPicker)
Vue.use(MonthPickerInput)


Vue.use(Vuex);
Vue.use(VueRouter);
Vue.use(Buefy);

Vue.use(VeeValidate, {
  events: '',
  errorBagName: 'vErrors',
  i18nRootKey: 'validations', // customize the root path for validation messages.
  i18n,
  dictionary: {
    // ru: validationMessagesRu,
    en: validationMessagesEn,
  }
});

window.$ = JQuery;

/* eslint-disable no-new */
/* eslint-disable no-useless-escape*/

//*************AXIOS**********************
window.axios = require("axios");
window.axios.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
window.axios.defaults.baseURL = "http://localhost:3001/";
window.axios.defaults.headers.Authorization = "Bearer NOTOKEN";

const router = new VueRouter({
  routes: Routes,
  linkActiveClass: 'is-active'
});

Vue.filter('formatDate', function(value) {
  if (value) {
    return moment(String(value)).format('YYYY-MM-DD HH:mm');
  }
});

Vue.config.productionTip = false;



Vue.use(VueTelInput);
Vue.use(FlagIcon)

new Vue({
  i18n,
  render: h => h(App),
  store: store,
  router: router,

}).$mount('#app');
