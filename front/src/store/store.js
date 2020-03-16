import Vue from 'vue'
import Vuex from 'vuex'
import products from './modules/products'
import categories from './modules/categories'
import suppliers from './modules/suppliers'
import orders from './modules/orders'
// import warehouse from './modules/warehouse'
import supply from './modules/supply'
import credit from './modules/credit'
import debit from './modules/debit'
import new_utility_bills from './modules/new_utility_bills'

Vue.use(Vuex);

export const store = new Vuex.Store({
    strict: true,
    modules:{
        products,
        categories,
        suppliers,
        // warehouse,
        supply,
        credit,
        debit,
        // utility_bills,
        new_utility_bills,
        orders,
    },
    state: {

    },
    mutations: {

    },
    getters : {
        // supply() { state => state.suppliers, } // -> getters['account/isAdmin']

    }
});
