<template>
    <div class="p-7">
        <div class="tile is-ancestor" v-for="(item, index) in decommissioned" :key="index">
            <div class="tile is-1 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("products.decommissionedMsg")}}</p>
                    <p class="title">"{{ item.ProductName }}" - "{{item.CompanyName}}"</p>

                    <ValidationObserver ref="observer">
                        <section slot-scope="{ validate }">

                            <div class="columns ">
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

                                <div class="column is-6">
                                    <BInputWithValidation rules="required"
                                                          type="textarea"
                                                          :label="labels.comment"
                                                          :value="item.Comment"
                                                          @input="updateComment($event, item.ID)"
                                    />
                                </div>
                            </div>
                            <div class="buttons pt-2" v-if="index === decommissioned.length - 1">
                                <b-button  @click="validate().then(addDecommission)"
                                           type="is-link is-primary"
                                >
                                    {{$t("buttons.submit")}}
                                </b-button>
                                <b-button tag="router-link"
                                          to="/supply"
                                          type="is-link is-light"
                                >
                                    {{$t("buttons.back")}}
                                </b-button>
                            </div>

                        </section>
                    </ValidationObserver>
                </div>
            </div>
            <div class="tile is-1 is-vertical is-parent">
            </div>
        </div>
    </div>
</template>

<script>
    import { ValidationObserver } from 'vee-validate';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';

    export default {
        name: 'decommissioned',
        components: {
            ValidationObserver,
            BInputWithValidation,
        },
        data() {
            return {
                labels:{
                    comment: this.$t("lables.comment"),
                    amount: this.$t("products.table.amount"),
                },
                form: [],
                error: false,
            }
        },
        mounted(){
        },
        computed: {
            decommissioned() {
                let answer = this.$store.getters['supply/decommissioned'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('decommissioned')));
                    this.$store.commit('supply/SET_DECOMMISSIONED',
                        JSON.parse(localStorage.getItem('decommissioned')));
                    return this.$store.getters['supply/decommissioned'];
                }else {
                    return this.$store.getters['supply/decommissioned'];
                }
            }
        },
        methods: {
            async addDecommission(){
                try {
                    await this.$store.dispatch('supply/addDecommission');
                } catch (error) {
                    this.error = true;
                    console.log('error')
                } finally {
                    if(!this.error){
                        await this.$router.push('/supply');
                    }
                }
            },
            updateAmount(event, id){
                let item = {'ID': id, 'Amount': event};
                this.$store.commit('supply/UPDATE_AMOUNT_DEC', item)
            },
            updateComment(event, id){
                let item = {'ID': id, 'Comment': event};
                this.$store.commit('supply/UPDATE_COMMENT_DEC', item)
            },
        },
    }
</script>

<style>
</style>
