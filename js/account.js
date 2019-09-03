import * as vue from 'https://cdn.jsdelivr.net/npm/vue/dist/vue.js'
import AccountDetails from './account-details.js'

var app = new Vue({
    el: '#app',
    components: {
        'account-details': AccountDetails
    },    
    data: {
        owner: 'Paura'
    }
}) 