<template>
    <div class="p-3">
        <div class="tile is-ancestor" v-for="(item, index) in categories" :key="index">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("categories.updateMsg")}} #{{ item.ID }}</p>
                    <section>

                        <div class="columns is-mobile is-multiline">
                            <div class="column is-narrow">
                                <b-field :label="labels.categoryName">
                                    <b-input :value="item.CategoryName"
                                             @input="updateCategoryName($event, item.ID)"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.сategoryDescription">
                                    <b-input :value="item.CategoryDescription"
                                             @input="updateCategoryDescription($event, item.ID)"
                                             placeholder="Enter name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" v-if="index === categories.length - 1">
                            <b-button  @click="updateCategory()" type="is-success">{{$t("buttons.submit")}}</b-button>
                            <b-button tag="router-link"
                                      to="/categories"
                                      type="is-link is-light"
                            >
                                {{$t("buttons.back")}}
                            </b-button>
                        </div>
                    </section>
                </div>
            </div>
            <div class="tile is-2 is-vertical is-parent">
            </div>
        </div>
    </div>
</template>

<script>
    // import { mapGetters } from 'vuex';

    export default {
        name: 'productsUpdate',
        components: {
        },
        data() {
            return {
                // products: [],
                labels:{
                    categoryName: this.$t("categories.table.categoryName"),
                    сategoryDescription: this.$t("categories.table.сategoryDescription"),
                }
            }
        },
        mounted(){
        },
        computed: {
            categories() {
                let answer = this.$store.getters['categories/categoriesUpdate'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('categories')));
                    this.$store.commit('categories/SET_CATEGORIES_UPDATE', JSON.parse(localStorage.getItem('categories')))
                    return this.$store.getters['categories/categoriesUpdate'];
                }else {
                    return this.$store.getters['categories/categoriesUpdate'];
                }
            }
        },
        methods: {
            updateCategory(){
                this.$store.dispatch('categories/updateCategory');
            },
            updateCategoryName(event, id){
                let item = {'ID': id, 'CategoryName': event};
                this.$store.commit('categories/UPDATE_CATEGORY_NAME', item)
            },
            updateCategoryDescription(event, id){
                let item = {'ID': id, 'CategoryDescription': event};
                this.$store.commit('categories/UPDATE_CATEGORY_DESCRIPTION', item)
            },
        },
    }
</script>

<style>
</style>
