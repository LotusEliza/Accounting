import Vue from "vue";

export default {
    namespaced: true,
    state: {
        orders: [],
        categoriesUpdate: [],
    },
    getters: {
        orders: state => state.orders,
    },
    mutations: {
        SET_ORDERS(state, orders){
            for (let i = 0; i < orders.length; i++) {
                orders[i].DateCreate = Vue.filter('formatDate')(orders[i].DateCreate);
            }
            state.orders = orders.sort(function (a, b) {
                if (a.DateCreate > b.DateCreate) return -1;
                if (a.DateCreate < b.DateCreate) return 1;
                return 0;
            });
        },
    },
    actions: {
        async getOrders({commit}){
            let response =  await window.axios.get('/orders');
            commit('SET_ORDERS', response.data.Items);
        },
    }
}
