import Vue from 'vue'
import Vuex from 'vuex'
import products from './modules/products'
import categories from './modules/categories'
import suppliers from './modules/suppliers'

Vue.use(Vuex);

export const store = new Vuex.Store({
    strict: true,
    modules:{
        products,
        categories,
        suppliers,
    },
    state: {

    },
    mutations: {

    },
    getters : {

    }
});
