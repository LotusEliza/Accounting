<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("products.addMsg")}}</p>
                    <section>

                        {{form}}
                        <div class="columns is-mobile is-multiline">
                            <div class="column is-narrow">
                                <b-field :label="labels.categoryName">
                                    <b-input
                                            v-model="form.categoryName"
                                            placeholder="Enter name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.сategoryDescription">
                                    <b-input
                                            v-model="form.сategoryDescription"
                                            placeholder="Enter name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>


                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" >
                            <b-button  @click="addProduct()" type="is-success">{{$t("buttons.add")}}</b-button>
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
        name: 'productAdd',
        components: {
        },
        data() {
            return {
                labels:{
                    categoryName: this.$t("categories.table.categoryName"),
                    сategoryDescription: this.$t("categories.table.сategoryDescription"),
                },
                form:{
                    сategoryDescription: '',
                    categoryName: '',
                }
            }
        },
        mounted(){
        },
        computed: {
        },
        methods: {
            addProduct(){
                console.log(this.form);
                this.$store.dispatch('categories/addCategory', this.form);
            },

        },
    }
</script>

<style>
</style>
