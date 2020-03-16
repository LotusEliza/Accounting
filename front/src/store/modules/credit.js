// import {ToastProgrammatic as Toast} from "buefy";
// import {i18n} from "../../plugins/i18n";
// import moment from 'moment';

import Vue from "vue";

export default {
    namespaced: true,
    state: {
        credit: [],
        monthCreditTotal: "",
        monthCredit: [],
        creditYear: [],
        dayCredit: [],
        dayCreditTable: [],
        monthCreditTable: []
    },
    getters: {
        monthInfoTotal: (state, getters, rootState) => {return [{
            credit:Number(state.monthCreditTotal)+Number(rootState.new_utility_bills.monthBillAmount),
            bill: Number(rootState.new_utility_bills.monthBills),
            debit: rootState.debit.monthDebitTotal,
            profit:Number(rootState.debit.monthDebitTotal-(Number(state.monthCreditTotal)+Number(rootState.new_utility_bills.monthBillAmount)))
            .toFixed(2)}]},
        monthCredit: state => state.monthCredit,
        monthCreditTotal: state => state.monthCreditTotal,
        monthCreditDebit: (state, getters, rootState) => {return rootState.debit.monthDebitTable.concat(state.monthCreditTable)
                },
        creditYear: state => state.creditYear,
        dayCredit: state => state.dayCredit,
        dayCreditTable: state => state.dayCreditTable,
        // dayCreditTable: (state) => {
        //     let copiedObject = Object.assign([], state.dayCredit);
        //     for (let i = 0; i < copiedObject.length; i++) {
        //         copiedObject[i].Date = Vue.filter('formatDate')(copiedObject[i].Date);
        //     }
        //     // console.log("copiedObject");
        //     // console.log(copiedObject);
        //     return copiedObject},
    },
    mutations: {
        ADD_CREDIT_MONTH_TOTAL(state, monthCreditTotal){
            state.monthCreditTotal = monthCreditTotal;
        },
        ADD_CREDIT_MONTH(state, monthCredit){
            state.monthCredit = monthCredit;
            for (let i = 0; i < monthCredit.length; i++) {
                monthCredit[i].Date = Vue.filter('formatDate')(monthCredit[i].Date);
            }
            state.monthCreditTable = monthCredit;
        },
        SET_CATEGORIES(state, categories){
            state.categories = categories;
        },
        REMOVE_CATEGORY(state, id){
            let c = state.categories.filter(c => c.ID != id);
            state.categories = c;
        },
        SET_CATEGORIES_UPDATE(state, categories){
            state.categoriesUpdate = categories;
            const parsed = JSON.stringify(categories);
            localStorage.setItem('categories', parsed);
        },
        ADD_CREDIT(state, credit){
            if(!state.credit){
                state.credit = [credit];
            }else {
                let c = state.credit.concat(credit);
                state.credit = c;
            }
        },
        UPDATE_CATEGORY_NAME(state, item){
            let p = state.categoriesUpdate.find(p => p.ID ==  item.ID);
            p.CategoryName = item.CategoryName;
        },
        UPDATE_CATEGORY_DESCRIPTION(state, item){
            let p = state.categoriesUpdate.find(p => p.ID ==  item.ID);
            p.CategoryDescription = item.CategoryDescription;
        },
        ADD_CREDIT_YEAR(state, creditYear){
            state.creditYear = creditYear;
        },
        ADD_CREDIT_DAY(state, dayCredit){
            state.dayCredit =  dayCredit;
            for (let i = 0; i < dayCredit.length; i++) {
                dayCredit[i].Date = Vue.filter('formatDate')(dayCredit[i].Date);
            }
            state.dayCreditTable =  dayCredit;
        },
    },
    actions: {
        async addCredit({commit}, credit){
            let response1 =  await window.axios.get('/supply/id');
            console.log(response1.data.Items[0].ID);

            console.log(credit);
            let creditAmount = credit.BuyPrice * credit.Amount;
            let response = await window.axios.post('/credit/add', {
                Credit: creditAmount,
                SupplyID: response1.data.Items[0].ID,
            });
            if (response.status == 200 || response.status == 204) {
                commit('ADD_CREDIT', credit);
            }
        },
        async getMonthCredit({commit}, credit){
            let response = await window.axios.post('/credit/report', {
                From: credit.from,
                To: credit.to,
            });
            if (response.status == 200 || response.status == 204) {
                let credit = 0;
                if(response.data.Items === null) {
                    console.log("less then 0");
                    commit('ADD_CREDIT_MONTH', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_CREDIT_MONTH', response.data.Items);
                    for (let i = 0; i < response.data.Items.length; i++) {
                        credit += response.data.Items[i].Credit;
                    }
                }
                commit('ADD_CREDIT_MONTH_TOTAL', credit);
            }
        },
        async getYearCredit({commit}, arrayStartEndMonths){
            let finalArray = [];
            for (let i = 0; i <= 11; i++) {
            let response = await window.axios.post('/credit/report', {
                From: arrayStartEndMonths[i].from,
                To: arrayStartEndMonths[i].to,
            });
            if (response.status == 200 || response.status == 204) {
                let credit = 0;
                if(response.data.Items === null) {
                    // console.log("less then 0");
                    finalArray.push(0)
                    // commit('ADD_CREDIT_MONTH', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        credit += response.data.Items[i].Credit;
                    }
                    finalArray.push(credit)
                }
            }
            }
            commit('ADD_CREDIT_YEAR', finalArray);
        },
        async getDayCredit({commit}, date){
            let response = await window.axios.post('/credit/report', {
                From: date[0].from,
                To: date[0].to,
            });
            if (response.status == 200 || response.status == 204) {
                // let debit = 0;
                if(response.data.Items === null) {
                    // console.log("less then 0");
                    commit('ADD_CREDIT_DAY', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_CREDIT_DAY', response.data.Items);
                    // for (let i = 0; i < response.data.Items.length; i++) {
                    //     debit += response.data.Items[i].Debit;
                    // }
                }
                // commit('ADD_DEBIT_MONTH_TOTAL', debit);
                // commit('ADD_CREDIT_MONTH', response.data.Items);

            }

        },
    }
}
