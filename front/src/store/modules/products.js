
import {ToastProgrammatic as Toast} from "buefy";
import {i18n} from '../../plugins/i18n';

export default {
    namespaced: true,
    state: {
        products: [],
        productsUpdate: [],
    },
    getters: {
        items: state => state.products,
        products: state => state.products,
        productsUpdate: state => state.productsUpdate,
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
            // debugger
            commit('SET_PRODUCTS', response.data.Items);
        },
        async removeProduct({commit}, id){
            let response =  await window.axios.post('/products/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                commit('REMOVE_PRODUCT', id);
                return 'OK';
            }
        },
        setUpdateProducts({commit}, products){
            commit('SET_PRODUCTS_UPDATE', products);
        },
        async updateProducts({state}){
            let count = null;
            for (let i = 0; i < state.productsUpdate.length; i++) {
            let response =  await window.axios.post('/products/update', {
                ID: state.productsUpdate[i].ID,
                ProductName: state.productsUpdate[i].ProductName,
                CategoryID: state.productsUpdate[i].CategoryID,
                VendorCode: state.productsUpdate[i].VendorCode,
            });
            if(response.status == 200 || response.status == 204){
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
                CategoryID: product.CategoryID,
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
    }
}
