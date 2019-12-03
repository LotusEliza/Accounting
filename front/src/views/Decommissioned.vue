<template>
    <div class="p-7">
        <div class="tile is-ancestor" v-for="(item, index) in decommissioned" :key="index">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("products.decommissionedMsg")}} #{{ item.ID }}</p>
                    <section>
                        <div class="columns ">

                            <div class="column is-6">
                                    <b-field :label="labels.comment">
                                        <b-input type="textarea"
                                                :value="item.Comment"
                                                @input="updateComment($event, item.ID)"
                                                placeholder="Enter contact name"
                                                rounded
                                                size="is-small"
                                        ></b-input>
                                    </b-field>
                            </div>

                            <div class="column is-6">
                                <b-field :label="labels.amount">
                                    <b-numberinput :value="item.Amount"
                                                   @input="updateAmount($event, item.ID)"
                                                   placeholder="Enter amount"
                                                   rounded
                                                   size="is-small"
                                                   step="1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" v-if="index === decommissioned.length - 1">
                            <b-button  @click="addDecommission()" type="is-success">{{$t("buttons.submit")}}</b-button>
                            <b-button tag="router-link"
                                      to="/products"
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
                labels:{
                    comment: this.$t("lables.comment"),
                    amount: this.$t("products.table.amount"),
                },
                form: []
            }
        },
        mounted(){
        },
        computed: {
            decommissioned() {
                let answer = this.$store.getters['products/decommissioned'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('decommissioned')));
                    this.$store.commit('products/SET_DECOMMISSIONED', JSON.parse(localStorage.getItem('decommissioned')))
                    return this.$store.getters['products/decommissioned'];
                }else {
                    return this.$store.getters['products/decommissioned'];
                }
            }
        },
        methods: {
            addDecommission(){
                this.$store.dispatch('products/addDecommission');
            },
            updateAmount(event, id){
                let item = {'ID': id, 'Amount': event};
                this.$store.commit('products/UPDATE_AMOUNT_DEC', item)
            },
            updateComment(event, id){
                let item = {'ID': id, 'Comment': event};
                this.$store.commit('products/UPDATE_COMMENT', item)
            },
        },
    }
</script>

<style>
</style>
