
import {ToastProgrammatic as Toast} from "buefy";
import {i18n} from '../../plugins/i18n';
import Vue from 'vue'

export default {
    namespaced: true,
    state: {
        products: [],
        productsUpdate: [],
        decommissioned: []
    },
    getters: {
        items: state => state.products,
        productsUpdate: state => state.productsUpdate,
        decommissioned: state => state.decommissioned,
    },
    mutations: {
        SET_PRODUCTS(state, products){
            state.products = products;
        },
        REMOVE_PRODUCT(state, id){
            let p = state.products.filter(p => p.ID != id);
            state.products = p;
        },
        SET_PRODUCTS_UPDATE(state, products){
            state.productsUpdate = products;
            const parsed = JSON.stringify(products);
            localStorage.setItem('products', parsed);
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
                console.log(decommissioned[i]);
            }
        },
        UPDATE_NAME(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.ProductName = item.ProductName;
        },
        UPDATE_CATEGORY_ID(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.CategoryID = item.CategoryID;
        },
        UPDATE_COMPANY_ID(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.SupplierID = item.CompanyID;
        },
        UPDATE_SELL_PRICE(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.SellPrice = item.SellPrice;
        },
        UPDATE_BUY_PRICE(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.BuyPrice = item.BuyPrice;
        },
        UPDATE_AMOUNT(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.Amount = item.Amount;
        },
        UPDATE_VENDOR_CODE(state, item){
            let p = state.productsUpdate.find(p => p.ID ==  item.ID);
            p.VendorCode = item.VendorCode;
        },
        UPDATE_AMOUNT_DEC(state, item){
            let d = state.decommissioned.find(d => d.ID ==  item.ID);
            d.Amount = item.Amount;
        },
        UPDATE_COMMENT(state, item){
            let d = state.decommissioned.find(d => d.ID ==  item.ID);
            d.Comment = item.Comment;
        },
        ADD_PRODUCT(state, product){
            if(!state.products){
                state.products = [product];
            }else {
                let p = state.products.concat(product);
                state.products = p;
            }
        }
    },
    actions: {
        async getProducts({commit}){
            let response =  await window.axios.get('/products');
            console.log(response.data.Items);
            // debugger
            commit('SET_PRODUCTS', response.data.Items);
        },
        async removeProduct({commit}, id){
            let response =  await window.axios.post('/products/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                // Toast.open({
                //     message: i18n.t("products.toast.remove"),
                //     type: 'is-danger'
                // });
                commit('REMOVE_PRODUCT', id);
                return 'OK';
            }
        },
        setUpdateProducts({commit}, products){
            commit('SET_PRODUCTS_UPDATE', products);
        },

        setDecommissioned({commit}, decommissioned){
            let arr = [];
            let item;
            for (let i = 0; i < decommissioned.length; i++) {
                item = { ID: decommissioned[i].ID, Amount: null, Comment: '', SellPrice: decommissioned[i].SellPrice };
                arr =  arr.concat(item);
            }
            commit('SET_DECOMMISSIONED', arr);
        },
        async updateProducts({state}){
            let count = null;
            for (let i = 0; i < state.productsUpdate.length; i++) {
            let response =  await window.axios.post('/products/update', {
                ID: state.productsUpdate[i].ID,
                ProductName: state.productsUpdate[i].ProductName,
                SupplierID: state.productsUpdate[i].SupplierID,
                CategoryID: state.productsUpdate[i].CategoryID,
                BuyPrice: state.productsUpdate[i].BuyPrice,
                SellPrice: state.productsUpdate[i].SellPrice,
                Amount: state.productsUpdate[i].Amount,
                VendorCode: state.productsUpdate[i].VendorCode,
            });
            if(response.status == 200 || response.status == 204){
                console.log('YES!');
                count++;
              }
            }
            if(state.productsUpdate.length === count){
                Toast.open({
                    message: i18n.t("products.toast.update"),
                    type: 'is-success'
                });
            }

        },
        async addProduct({commit}, product){
            let time = new Date();
            let response = await window.axios.post('/products/add', {
                ProductName: product.ProductName,
                VendorCode: product.VendorCode,
                SupplierID: product.SupplierID,
                CategoryID: product.CategoryID,
                BuyPrice: product.BuyPrice,
                SellPrice: product.SellPrice,
                Amount: product.Amount,
                DateCreate: time,
            });
            if (response.status == 200 || response.status == 204) {
                commit('ADD_PRODUCT', product);
                Toast.open({
                    message: i18n.t("products.toast.add"),
                    type: 'is-success'
                });
            }
        },
        async addDecommission({state}){
            let time = new Date();
            let count = null;
            for (let i = 0; i < state.decommissioned.length; i++) {
                state.decommissioned[i].ID;
                let p = state.products.find(p => p.ID ==  state.decommissioned[i].ID);
                if(p.Amount < state.decommissioned[i].Amount){
                    Toast.open({
                        message: i18n.t("products.toast.decommissionedTooMuch")+" Id: "+state.decommissioned[i].ID,
                        type: 'is-danger'
                    });
                    console.log('More then needed!')
                }else{
                    console.log('Fine!');

                    let response =  await window.axios.post('/products/decommissioned', {
                        ProductID: state.decommissioned[i].ID,
                        AmountDecom: state.decommissioned[i].Amount,
                        Amount: p.Amount-state.decommissioned[i].Amount,
                        // Decommissioned: state.decommissioned[i].Amount,
                        Comment: state.decommissioned[i].Comment,
                        SellPrice: p.SellPrice,
                        DateCreate: time,
                    });
                    if(response.status == 200 || response.status == 204){
                        console.log('YES!');
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
        }
    }
}
