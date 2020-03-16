import {ToastProgrammatic as Toast} from "buefy";
import {i18n} from '../../plugins/i18n';
import Vue from "vue";

export default {
    namespaced: true,
    state: {
        supply: [],
        supplyUpdate: [],
        decommissioned: [],
        units:[],
        decommissioned_month: [],
        decommissioned_day: [],
        totalForMonth: null
    },
    getters: {
        items: state => state.supply,
        units: state => state.units,
        supplyUpdate: state => state.supplyUpdate,
        decommissioned: state => state.decommissioned,
        decommissioned_month: state => state.decommissioned_month,
        decommissioned_day: state => state.decommissioned_day,
        totalForMonthDecom: state => state.totalForMonth,
    },
    mutations: {
        ADD_DECOMMISSIONED_DAY(state, data){
            for (let i = 0; i < data.length; i++) {
                data[i].DateCreate = Vue.filter('formatDate')(data[i].DateCreate);
            }
            state.decommissioned_day = data;
        },
        ADD_DECOMMISSIONED_MONTH(state, data){
            let totalForMonth = null;
            for (let i = 0; i < data.length; i++) {
                totalForMonth += Number(data[i].AmountDecom * data[i].SellPrice);

            }
            state.totalForMonth = totalForMonth;
            for (let i = 0; i < data.length; i++) {
                data[i].DateCreate = Vue.filter('formatDate')(data[i].DateCreate);
            }
            state.decommissioned_month = data;
        },
        SET_SUPPLY(state, supply){
            state.supply = supply;
        },
        REMOVE_SUPPLY(state, id){
            let s = state.supply.filter(s => s.ID != id);
            state.supply = s;
        },
        SET_SUPPLY_UPDATE(state, supply){
            state.supplyUpdate = supply;
            const parsed = JSON.stringify(supply);
            localStorage.setItem('supply', parsed);
        },
        UPDATE_PRODUCT_ID(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.ProductID = item.ProductID;
        },
        UPDATE_SUPPLIER_ID(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.SupplierID = item.SupplierID;
        },
        UPDATE_SELL_PRICE(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.SellPrice = item.SellPrice;
            let onePercent = s.BuyPrice/100;
            s.Surcharge = Math.round((s.SellPrice - s.BuyPrice)/onePercent);
        },
        UPDATE_BUY_PRICE(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.BuyPrice = item.BuyPrice;
            let sellPrice = Math.round(s.BuyPrice+(s.BuyPrice/100)*s.Surcharge);
            s.SellPrice = sellPrice;
        },
        UPDATE_AMOUNT(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.Amount = item.Amount;
        },
        UPDATE_SUP_AMOUNT(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.SupAmount = item.SupAmount;
        },
        UPDATE_UNIT_ID(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.UnitID = item.UnitID;
        },
        UPDATE_SURCHARGE(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.Surcharge = item.Surcharge;
            let sellPrice = Number((s.BuyPrice + (s.BuyPrice/100)*item.Surcharge).toFixed(2));
            s.SellPrice = sellPrice;
        },
        UPDATE_COMMENT(state, item){
            let s = state.supplyUpdate.find(s => s.ID ==  item.ID);
            s.Comment = item.Comment;
        },
        ADD_SUPPLY(state, supply){
            if(!state.supply){
                state.supply = [supply];
            }else {
                let s = state.supply.concat(supply);
                state.supply = s;
            }
        },
        SET_DECOMMISSIONED(state, decommissioned){

            state.decommissioned = decommissioned;
            const parsed = JSON.stringify(decommissioned);
            localStorage.setItem('decommissioned', parsed);
        },
        PREPARE_DECOMMISSIONED(decommissioned){
            for (let i = 0; i < decommissioned.length; i++) {
                Vue.set(decommissioned[i], "AmountDecommissioned", "");
                Vue.set(decommissioned[i], "Comment", "");
            }
        },
        UPDATE_AMOUNT_DEC(state, item){
            let d = state.decommissioned.find(d => d.ID ==  item.ID);
            d.Amount = item.Amount;
        },
        UPDATE_COMMENT_DEC(state, item){
            let d = state.decommissioned.find(d => d.ID ==  item.ID);
            d.Comment = item.Comment;
        },
        SET_UNITS(state, units){
            state.units = units;
        },
        UPDATE_AMOUNT_SUPPLY(state, array){
            // console.log(array);
            let s = state.supply.filter(res=>res.ID == array.id);
            // console.log(s)
            // s[0].Amount = array.inStock;
            s[0].Amount = array.amount;
        },
        UPDATE_SUPPLY_AMOUNT(state, array){
            console.log('Here is update sup amount');
            let s = state.supply.filter(res=>res.ID == array.SupplyID);
            s.Amount = array.Amount;
        }
    },
    actions: {
        async getSupply({commit}){
            let response =  await window.axios.get('/supply');
            commit('SET_SUPPLY', response.data.Items);
        },
        async getUnits({commit}){
            let response =  await window.axios.get('/units');
            commit('SET_UNITS', response.data.Items);
        },
        async removeSupply({commit}, id){
            let response =  await window.axios.post('/supply/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                commit('REMOVE_SUPPLY', id);
                return 'OK';
            }
            let response2 =  await window.axios.post('/credit/remove', {
                SupplyID: id,
            });
            if(response2.status == 200 || response2.status == 204){
                console.log("yes")
            }
        },
        setUpdateSupply({commit}, supply){
            commit('SET_SUPPLY_UPDATE', supply);
        },
        async updateSupply({state}){
            let count = null;
            for (let i = 0; i < state.supplyUpdate.length; i++) {
                let response =  await window.axios.post('/supply/update', {
                    ID: state.supplyUpdate[i].ID,
                    ProductID: state.supplyUpdate[i].ProductID,
                    SupplierID: state.supplyUpdate[i].SupplierID,
                    BuyPrice: state.supplyUpdate[i].BuyPrice,
                    SellPrice: state.supplyUpdate[i].SellPrice,
                    Surcharge: Number(state.supplyUpdate[i].Surcharge),
                    UnitID: state.supplyUpdate[i].UnitID,
                    Amount: state.supplyUpdate[i].Amount,
                    SupAmount: state.supplyUpdate[i].SupAmount,
                    Comment: state.supplyUpdate[i].Comment,
                });
                if(response.status == 200 || response.status == 204){
                    count++;
                }
                let creditAmount = state.supplyUpdate[i].BuyPrice * state.supplyUpdate[i].Amount;
                let response2 =  await window.axios.post('/credit/update', {
                    SupplyID: state.supplyUpdate[i].ID,
                    Credit: creditAmount,
                });
                if(response2.status == 200 || response2.status == 204){
                    console.log("yes")
                }
            }

            if(state.supplyUpdate.length === count){
                Toast.open({
                    message: i18n.t("products.toast.update"),
                    type: 'is-success'
                });
            }

        },
        async addSupply({commit}, supply){
            let time = new Date();
            let response = await window.axios.post('/supply/add', {
                ProductId: supply.ProductId,
                SupplierId: supply.SupplierId,
                BuyPrice: supply.BuyPrice,
                SellPrice: supply.SellPrice,
                Surcharge: supply.Surcharge,
                SupAmount: supply.Amount,
                Amount: supply.Amount,
                UnitID: supply.UnitId,
                Comment: supply.Comment,
                DateCreate: time,
            });
            if (response.status == 200 || response.status == 204) {
                commit('ADD_SUPPLY', supply);
                Toast.open({
                    message: i18n.t("products.toast.add"),
                    type: 'is-success'
                });
            }
        },
        setDecommissioned({commit}, decommissioned){
            let arr = [];
            let item;
            for (let i = 0; i < decommissioned.length; i++) {
                item = { ID: decommissioned[i].ID, Amount: null, Comment: '', SellPrice: decommissioned[i].SellPrice,
                    ProductName: decommissioned[i].ProductName, CompanyName: decommissioned[i].CompanyName  };
                arr =  arr.concat(item);
            }
            commit('SET_DECOMMISSIONED', arr);
        },
        async addDecommission({state}){
            let time = new Date();
            let count = null;
            for (let i = 0; i < state.decommissioned.length; i++) {
                let p = state.supply.find(p => p.ID ==  state.decommissioned[i].ID);
                if(p.Amount < state.decommissioned[i].Amount){
                    Toast.open({
                        message: i18n.t("products.toast.decommissionedTooMuch")+" Id: "+state.decommissioned[i].ID,
                        type: 'is-danger'
                    });
                }else{
                    let response =  await window.axios.post('/supply/decommissioned', {
                        ID: state.decommissioned[i].ID,
                        ProductID: p.ProductID,
                        AmountDecom: state.decommissioned[i].Amount,
                        Amount: p.Amount-state.decommissioned[i].Amount,
                        // Decommissioned: state.decommissioned[i].Amount,
                        Comment: state.decommissioned[i].Comment,
                        SellPrice: p.SellPrice,
                        DateCreate: time,
                    });
                    if(response.status == 200 || response.status == 204){
                        count++;
                    }

                }
            }
            if(state.decommissioned.length === count){
                Toast.open({
                    message: i18n.t("products.toast.decommissioned"),
                    type: 'is-success'
                });
            }
        },
        async decommissionedMonthGet({commit}, data){
            let response = await window.axios.post('/supply/month_decommissioned', {
                From: data.from,
                To: data.to,
            });
            if (response.status == 200 || response.status == 204) {
                if(response.data.Items === null) {
                    console.log("less then 0");
                    commit('ADD_DECOMMISSIONED_MONTH', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_DECOMMISSIONED_MONTH', response.data.Items);
                }
            }
        },
        async updateSupplyAmountTest({commit, state}, data){
            let arrayOfItems = state.supply.filter(item=>item.ProductName == data.name && item.Amount !== 0);

            if(arrayOfItems.length > 1){
                if(data.amount >= arrayOfItems[0].Amount){
                    console.log("More then in first");
                    let rest = data.amount - arrayOfItems[0].Amount;

                    if(rest){
                        let amount = Number(arrayOfItems[1].Amount - rest);
                        let id = arrayOfItems[1].ID;
                        let response = await window.axios.post('/supply/update_amount', {
                            Amount: amount,
                            ID: id,
                        });
                        if (response.status == 200 || response.status == 204) {
                            // Vue.set(data, "amountCurrent", amount);
                            // console.log(response);

                            commit('UPDATE_AMOUNT_SUPPLY',  {id, amount});
                        }
                    }
                    let amount = 0;
                    let id = arrayOfItems[0].ID;
                    let response = await window.axios.post('/supply/update_amount', {
                        Amount: amount,
                        ID: id,
                    });
                    if (response.status == 200 || response.status == 204) {
                        commit('UPDATE_AMOUNT_SUPPLY',  {id, amount});
                    }
                }else{
                    let amount = arrayOfItems[0].Amount - data.amount;
                    let id = arrayOfItems[0].ID;
                    let response = await window.axios.post('/supply/update_amount', {
                        Amount: amount,
                        ID: id,
                    });
                    if (response.status == 200 || response.status == 204) {
                        commit('UPDATE_AMOUNT_SUPPLY',  {id, amount});
                    }
                }
            }else {
                let amount = data.inStock;
                let id = arrayOfItems[0].ID;
                let response = await window.axios.post('/supply/update_amount', {
                    Amount: amount,
                    ID: id,
                });
                if (response.status == 200 || response.status == 204) {
                    console.log(response);
                    let amount = data.inStock;
                    commit('UPDATE_AMOUNT_SUPPLY',  {id, amount});
                }
            }
        },

        async decommissionedDayGet({commit}, data){
            console.log("data")
            console.log(data)
            let response = await window.axios.post('/supply/month_decommissioned', {
                From: data[0].from,
                To: data[0].to,
            });
            if (response.status == 200 || response.status == 204) {
                if(response.data.Items === null) {
                    console.log("less then 0");
                    commit('ADD_DECOMMISSIONED_DAY', []);
                }else {
                    for (let i = 0; i < response.data.Items.length; i++) {
                        Vue.set(response.data.Items[i], "UnicID", i+1);
                    }
                    commit('ADD_DECOMMISSIONED_DAY', response.data.Items);
                }
            }
        },

        // {commit}, data
//         async updateSupplyAmount({commit}, data){
//             console.log("data")
//             console.log(data);
//             let response = await window.axios.post('/supply/update_amount', {
//                 Amount: data.Amount,
//                 ID: data.SupplyID,
//             });
//
//             if (response.status == 200 || response.status == 204) {
//                 console.log(response)
//                 commit('UPDATE_AMOUNT_SUPPLY');
//                 // if(response.data.Items === null) {
//                 //     console.log("less then 0");
//                 //     commit('ADD_DECOMMISSIONED_MONTH', []);
//                 // }else {
//                 //     for (let i = 0; i < response.data.Items.length; i++) {
//                 //         Vue.set(response.data.Items[i], "UnicID", i+1);
//                 //     }
//                 //     commit('ADD_DECOMMISSIONED_MONTH', response.data.Items);
//                 // }
//                 // // commit('ADD_DECOMMISSIONED_MONTH', response.data.Items);
//             }
//         },
    }
}
