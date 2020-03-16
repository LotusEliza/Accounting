import {ToastProgrammatic as Toast} from "buefy";
import {i18n} from "../../plugins/i18n";

export default {
    namespaced: true,
    state: {
        suppliers: [],
        suppliersUpdate: [],
    },
    getters: {
        suppliers: state => state.suppliers,
        suppliersUpdate: state => state.suppliersUpdate,
    },
    mutations: {
        SET_SUPPLIERS(state, suppliers){
            state.suppliers = suppliers;
        },
        REMOVE_SUPPLIER(state, id){
            let s = state.suppliers.filter(s => s.ID != id);
            state.suppliers = s;
        },
        SET_SUPPLIERS_UPDATE(state, suppliers){
            state.suppliersUpdate = suppliers;
            const parsed = JSON.stringify(suppliers);
            localStorage.setItem('suppliers', parsed);
        },
        SET_COMPANIES_NAMES(state, names){
            state.companiesNames = names;
        },
        UPDATE_CITY(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.City = item.City;
        },
        UPDATE_COMPANY_NAME(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.CompanyName = item.CompanyName;
        },
        UPDATE_CONTACT_NAME(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.ContactName = item.ContactName;
        },
        UPDATE_CONTACT_TITLE(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.ContactTitle = item.ContactTitle;
        },
        UPDATE_EMAIL(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.Email = item.Email;
        },
        UPDATE_PHONE(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.Phone = item.Phone;
        },
        UPDATE_ADDRESS(state, item){
            let s = state.suppliersUpdate.find(s => s.ID ==  item.ID);
            s.Address = item.Address;
        },
        ADD_SUPPLIER(state, supplier){
            if(!state.suppliers){
                state.suppliers = [supplier];
            }else {
                let s = state.suppliers.concat(supplier);
                state.suppliers = s;
            }
        }
    },
    actions: {
        async getSuppliers({commit}){
            let response =  await window.axios.get('/suppliers');
            commit('SET_SUPPLIERS', response.data.Items);
        },
        async removeSupplier({commit}, id){
            let response =  await window.axios.post('/suppliers/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                commit('REMOVE_SUPPLIER', id);
                return 'OK';
            }
        },
        setUpdateSuppliers({commit}, suppliers){
            commit('SET_SUPPLIERS_UPDATE', suppliers);
        },
        async updateSupplier({state}){
            let count = null;
            for (let i = 0; i < state.suppliersUpdate.length; i++) {
                let response = await window.axios.post('/suppliers/update', {
                    ID: state.suppliersUpdate[i].ID,
                    CompanyName: state.suppliersUpdate[i].CompanyName,
                    ContactName: state.suppliersUpdate[i].ContactName,
                    ContactTitle: state.suppliersUpdate[i].ContactTitle,
                    Address: state.suppliersUpdate[i].Address,
                    City: state.suppliersUpdate[i].City,
                    Phone: state.suppliersUpdate[i].Phone,
                    Email: state.suppliersUpdate[i].Email
                });
                if (response.status == 200 || response.status == 204) {
                    count++;
                }
            }
            if(state.suppliersUpdate.length === count){
                Toast.open({
                    message: i18n.t("suppliers.toast.update"),
                    type: 'is-success'
                });
            }
        },
        async addSupplier({commit}, supplier){
                let response = await window.axios.post('/suppliers/add', {
                    CompanyName: supplier.CompanyName,
                    ContactName: supplier.ContactName,
                    ContactTitle: supplier.ContactTitle,
                    Address: supplier.Address,
                    City: supplier.City,
                    Phone: supplier.Phone,
                    Email: supplier.Email,
                });
                if (response.status == 200 || response.status == 204) {
                    commit('ADD_SUPPLIER', supplier);
                    Toast.open({
                        message: i18n.t("suppliers.toast.add"),
                        type: 'is-success'
                    });
                }
        }
    }
}
