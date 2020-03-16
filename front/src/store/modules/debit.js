
import {i18n} from "../../plugins/i18n";
import {ToastProgrammatic as Toast} from "buefy";
import Vue from "vue";

export default {
    namespaced: true,
    state: {
        debit:[],
        lastId: '',
        monthDebit: [],
        dayDebit: [],
        dayDebitTable: [],
        debitYear: [],
        monthDebitTotal: "",
        monthDebitTable: [],
        debitTotal: null,
    },
    getters: {
        monthDebit: state => state.monthDebit,
        monthDebitTable: state => state.monthDebitTable,
        monthDebitTotal: state => state.monthDebitTotal,
        dayDebit: state => state.dayDebit,
        dayDebitTable: state => state.dayDebitTable,
        debitYear: state => state.debitYear,
        debit: state => state.debit,
        debitTotal: state => state.debitTotal,
    },
    mutations: {
        ADD_DEBIT(state, debit){
           state.debit =  debit;
        },
        SET_LAST_ID(state, lastId){
            state.lastId =  lastId;
        },
        ADD_DEBIT_MONTH(state, monthDebit){
            for (let i = 0; i < monthDebit.length; i++) {
                Vue.set(monthDebit[i], "Comment", "trans #"+monthDebit[i].ID);
            }
            state.monthDebit =  monthDebit;

            for (let i = 0; i < monthDebit.length; i++) {
                monthDebit[i].Date = Vue.filter('formatDate')(monthDebit[i].Date);
            }
            state.monthDebitTable =  monthDebit;
        },

        ADD_DEBIT_DAY(state, dayDebit){
            let debitTotal = null;
            for (let i = 0; i < dayDebit.length; i++) {
                console.log("dayDebitTotal")
                console.log(debitTotal)
                debitTotal += dayDebit[i].Debit;
                console.log(debitTotal)
            }
            state.debitTotal = debitTotal;
            state.dayDebit =  dayDebit;
            // for (let i = 0; i < dayDebit.length; i++) {
            //     dayDebit[i].Date = Vue.filter('formatDate')(dayDebit[i].Date);
            // }
            // state.dayDebitTable = dayDebit;
        },
        ADD_DEBIT_DAY_TABLE(state, dayDebitTable){
            // for (let i = 0; i < dayDebitTable.length; i++) {
            //     Vue.set(dayDebitTable[i], "Comment", "trans #"+dayDebitTable[i].ID);
            //     dayDebitTable[i].Date = Vue.filter('formatDate')(dayDebitTable[i].Date);
            // }
            state.dayDebitTable =  dayDebitTable;
        },
        ADD_DEBIT_MONTH_TOTAL(state, monthDebitTotal){
            state.monthDebitTotal = monthDebitTotal;
        },
        ADD_DEBIT_YEAR(state, debitYear){
            state.debitYear = debitYear;
        }
    },
    actions: {
        async getMonthDebit({commit}, debit){
            let response = await window.axios.post('/debit/report', {
                From: debit.from,
                To: debit.to,
            });
            if (response.status == 200 || response.status == 204) {
                let debit = 0;
                if(response.data.Items === null) {
                    console.log("less then 0");
                    commit('ADD_DEBIT_MONTH', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_DEBIT_MONTH', response.data.Items);
                    for (let i = 0; i < response.data.Items.length; i++) {
                        debit += response.data.Items[i].Debit;
                    }
                }
                commit('ADD_DEBIT_MONTH_TOTAL', Number(debit).toFixed(2));
            }
        },
        async addTransaction({commit}, debit){
            let time = new Date();
            let response = await window.axios.post('/debit/add', {
                Debit: debit,
                DateCreate: time,
            });
            if (response.status == 200 || response.status == 204) {
                commit('ADD_DEBIT', debit);
            }
        },
        async lastDebitId({commit}){
            let response =  await window.axios.get('/debit/last_id');
            commit('SET_LAST_ID', response.data.Items);
        },

        async update({commit}, array){
            for(let j = 0; j < array.length; j++) {
                let response = await window.axios.post('/supply/update_amount', {
                    Amount: Number(array[j].Amount),
                    ID: array[j].SupplyID,
                });
                if (response.status == 200 || response.status == 204) {
                    console.log("updated!");
                    commit('supply/UPDATE_SUPPLY_AMOUNT', array[j], { root: true })
                }
            }
        },
        async getYearDebit({commit}, arrayStartEndMonths){
            let finalArray = [];
            for (let i = 0; i <= 11; i++) {
                let response = await window.axios.post('/debit/report', {
                    From: arrayStartEndMonths[i].from,
                    To: arrayStartEndMonths[i].to,
                });
                if (response.status == 200 || response.status === 204) {
                    let debit = 0;
                    if(response.data.Items === null) {
                        // console.log("less then 0");
                        finalArray.push(0)
                    }else {
                        for (let i = 0; i < response.data.Items.length; i++) {
                            debit += response.data.Items[i].Debit;
                        }
                        finalArray.push(debit)
                    }
                }
            }
            commit('ADD_DEBIT_YEAR', finalArray);
            console.log(finalArray);
        },
        async getDayDebit({commit}, date){
            let response = await window.axios.post('/debit/report', {
                From: date[0].from,
                To: date[0].to,
            });
            if (response.status == 200 || response.status == 204) {
                // let debit = 0;
                if(response.data.Items === null) {
                    console.log("less then 0");
                    commit('ADD_DEBIT_DAY', []);
                    // commit('ADD_DEBIT_DAY_TABLE', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "DateTable", date = Vue.filter('formatDate')(response.data.Items[i].Date));
                        Vue.set(response.data.Items[i], "Comment", "trans #"+response.data.Items[i].ID);
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_DEBIT_DAY', response.data.Items);
                }
            }
        },

        async addDebitsProducts({state, rootState, dispatch}, items){
            let count = 0;
            let itemsLength = items.length;
            for (let i = 0; i < items.length; i++) {
                let supply = rootState.supply.supply.filter(res=>res.ProductName == items[i].name);
                supply = supply.filter(res=>res.Amount !== 0);
                supply = supply.sort((a, b) => (a.SellPrice > b.SellPrice) ? 1 : -1);
                console.log('supply');
                console.log(supply);
                let array = [];
                let amount = items[i].amount;

                //creating array of items { amount: 120, SupplyID: 12 }
                for(let j = 0; j < supply.length; j++) {
                    console.log(supply[j].ID);
                    console.log(supply[j].Amount);
                    console.log(amount);
                    if(supply[j].Amount > 0){
                        if (amount > 0) {
                            if (amount < supply[j].Amount) {
                                array.push({Amount: supply[j].Amount - amount, SupplyID: supply[j].ID});
                                break
                            } else {
                                amount = amount - supply[j].Amount;
                                array.push({Amount: 0, SupplyID: supply[j].ID});
                            }
                        } else {
                            console.log('Amount is Null')
                        }
                    }
                }
                await dispatch('update', array);
                let response = await window.axios.post('/debits_products/add', {
                    DebitID: state.lastId[0].ID,
                    Amount: Number(items[i].amount),
                    ProductID: items[i].id,
                    Price: items[i].price,
                });
                if (response.status == 200 || response.status == 204) {
                    count++
                }
            }
            if(count === itemsLength){
                Toast.open({
                    message: i18n.t("cashBox.toast.add"),
                    type: 'is-success'
                });
            }
        }
    }
}
