import {ToastProgrammatic as Toast} from "buefy";
import {i18n} from "../../plugins/i18n";

export default {
    namespaced: true,
    state: {
        categories: [],
        categoriesUpdate: [],
    },
    getters: {
        categories: state => state.categories,
        categoriesUpdate: state => state.categoriesUpdate,
    },
    mutations: {
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
        ADD_CATEGORY(state, category){
            if(!state.categories){
                state.categories = [category];
            }else {
                let c = state.categories.concat(category);
                state.categories = c;
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
    },
    actions: {
        async getCategories({commit}){
            let response =  await window.axios.get('/categories');
            // debugger
            commit('SET_CATEGORIES', response.data.Items);
        },
        async removeCategory({commit}, id){
            let response =  await window.axios.post('/categories/remove', {
                ID: id,
            });
            if(response.status == 200 || response.status == 204){
                commit('REMOVE_CATEGORY', id);
                return 'OK';
            }
        },
        setUpdateCategories({commit}, categories){
            commit('SET_CATEGORIES_UPDATE', categories);
        },
        async updateCategory({state}){
            let count = null;
            for (let i = 0; i < state.categoriesUpdate.length; i++) {
                let response =  await window.axios.post('/categories/update', {
                    ID: state.categoriesUpdate[i].ID,
                    CategoryName: state.categoriesUpdate[i].CategoryName,
                    CategoryDescription: state.categoriesUpdate[i].CategoryDescription,
                });
                if(response.status == 200 || response.status == 204){
                    count++;
                }
            }
            if(state.categoriesUpdate.length === count){
                Toast.open({
                    message: i18n.t("categories.toast.update"),
                    type: 'is-success'
                });
            }
        },
        async addCategory({commit}, category){
            let response = await window.axios.post('/categories/add', {
                CategoryName: category.categoryName,
                CategoryDescription: category.сategoryDescription,
            });
            if (response.status == 200 || response.status == 204) {
                commit('ADD_CATEGORY', category);
                // router.push('/categories');
                Toast.open({
                    message: i18n.t("categories.toast.add"),
                    type: 'is-success'
                });

            }
        }

    }
}
