import {createApp} from 'vue'
import vuetify from './plugins/vuetify'
import App from './App.vue'
import router from './router'
import store from './store'
import VueAxios from "vue-axios"
import axios from "axios"

createApp(App)
    .use(router)
    .use(VueAxios, axios)
    .use(store)
    .use(vuetify)
    .mount('#app')
