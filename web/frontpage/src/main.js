import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import './plugins/http'
import moment from 'moment'


Vue.filter('dateformat', function (inDate, outDate) {
  return moment(inDate).format(outDate)
})
Vue.config.productionTip = false

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
