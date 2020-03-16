import {ToastProgrammatic as Toast} from "buefy";
import Vue from "vue";
// import {i18n} from "../../plugins/i18n";

export default {
    namespaced: true,
    state: {
        bills: [],
        billsNames: [],
        monthBill:[],
        monthBillAmount: null,
        monthBillAmountTest: 3,
        monthBillTable: [],
        totalForMonth: null,
    },
    getters: {
        bills: state => state.bills,
        billsNames: state => state.billsNames.map(a => a.BillName),
        monthBills: state => state.monthBill,
        monthBillsAmount: state => state.monthBillAmount,
        monthBillTable: state => state.monthBillTable,
        totalForMonth: state => state.totalForMonth,
    },
    mutations: {
        SET_BILLS(state, bills){
            state.bills = bills;

        },
        REMOVE_BILL(state, id){
            let b = state.bills.filter(s => s.ID != id);
            state.bills = b;
            console.log("remove bill mutation!")
        },
        SET_BILLS_NAMES(state, billsNames){
            state.billsNames = billsNames;
        },
        ADD_BILL(state, bill){
            let id = state.bills[state.bills.length - 1].ID;
            console.log(id);
            Vue.set(bill, "ID", id+1);
            if(!state.bills){
                state.bills = [bill];
            }else {
                let b = state.bills.concat(bill);
                state.bills = b;
            }
        },
        MONTH_BILL_AMOUNT_ZERO(state){
            state.monthBillAmount = 0;
        },
        ADD_BILL_MONTH(state, data){
            if(data.items.length){
                state.monthBill = data.items;
                state.monthBillAmount = data.amount;
                // console.log('yes');
                for (let i = 0; i < data.items.length; i++) {
                    Vue.set(data.items[i], "UnicID", i+1);
                    data.items[i].Date = Vue.filter('formatDate')(data.items[i].Date);
                }

                let totalForMonth = null;
                for (let i = 0; i < data.items.length; i++) {
                    totalForMonth += Number(data.items[i].Amount);
                }
                state.totalForMonth = totalForMonth;
                state.monthBillTable = data.items;

            }else {
                state.monthBill = [];
                state.totalForMonth = null;
                console.log('no')
            }
        }
    },
    actions: {
        async getBills({commit}){
            let response =  await window.axios.get('/utility/get');
            commit('SET_BILLS', response.data.Items);
        },

        async getBillsNames({commit}){
            let response =  await window.axios.get('/utility/names');
            commit('SET_BILLS_NAMES', response.data.Items);
        },

        async removeBill({commit}, id){
            let response =  await window.axios.post('/utility/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                commit('REMOVE_BILL', id);
                console.log("yes")
            }
        },


        async addBill({commit}, bill){
            // let time = new Date();
            // let id = state.bills[state.bills.length - 1].ID;
            let response = await window.axios.post('/utility/set', {
                Amount: Number(bill.Amount),
                BillName: bill.BillName,
                DateCreate: bill.DateCreate,
            });
            if (response.status == 200 || response.status == 204) {
                // bill = bill.concat({DateCreate: time});
                // Vue.set(bill, "DateCreate", time);
                commit('ADD_BILL', bill);

                Toast.open({
                    message: 'yes',
                    type: 'is-success'
                });
            }
        },

        async getMonthBill({commit}, date){
            let response = await window.axios.post('/utility/month_bills', {
                From: date.from,
                To: date.to,
            });
            if (response.status == 200 || response.status == 204) {
                // let credit = 0;
                commit('MONTH_BILL_AMOUNT_ZERO');
                if(response.data.Items === null) {
                    let amount=0;
                    let items = [];
                    console.log("less then 0");

                    commit('ADD_BILL_MONTH', {items, amount});
                }else {
                    let amount=0;
                    let items = response.data.Items;
                    for (let i = 0; i < response.data.Items.length; i++) {
                        amount+=response.data.Items[i].Amount;
                    }
                    console.log("amount");
                    console.log(amount);
                    commit('ADD_BILL_MONTH', {items, amount});
                }
            }
        },

        // async removeBill({commit}, id){
        //     let response =  await window.axios.post('/utility/remove', {
        //         ID: id,
        //     });
        //     if(response.status == 200 || response.status == 204){
        //         commit('REMOVE_BILL', id);
        //         console.log("yes")
        //     }
        // },
        // async getMonthBill({commit, state}, date){
        //     let response = await window.axios.post('/utility_bills', {
        //         From: date.from,
        //         To: date.to,
        //     });
        //     if (response.status == 200 || response.status == 204) {
        //         // let credit = 0;
        //         state.monthBillAmount = 0;
        //         if(response.data.Items === null) {
        //             console.log("less then 0");
        //             commit('ADD_BILL_MONTH', []);
        //         }else {
        //             commit('ADD_BILL_MONTH', response.data.Items);
        //         }
        //     }
        // },
        // async getBillNames({commit}){
        //     let response =  await window.axios.get('/utility/bill_names');
        //     commit('SET_BILLS', response.data.Items);
        // },

    }
}
