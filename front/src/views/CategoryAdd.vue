<template>
    <div class="p-3">
        <div class="tile is-ancestor plr-15">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("products.addMsg")}}</p>
                    <ValidationObserver ref="observer">
                        <section slot-scope="{ validate }">
                                <div class="columns is-mobile is-multiline">

                                    <div class="column">
                                        <BInputWithValidation rules="required"
                                                              :label="labels.categoryName"
                                                              v-model="form.categoryName"
                                        />
                                    </div>

                                    <div class="column">
                                        <BInputWithValidation rules="required"
                                                              :label="labels.сategoryDescription"
                                                              v-model="form.сategoryDescription"
                                        />
                                    </div>

                                </div>
                                <div class="buttons pt-2" >
                                    <b-button  @click="validate().then(addCategory)"
                                               type="is-link is-primary"
                                               class="is-small"
                                    >
                                        {{$t("buttons.add")}}
                                    </b-button>
                                    <b-button tag="router-link"
                                              to="/categories"
                                              type="is-link is-light"
                                              class="is-small"
                                    >
                                        {{$t("buttons.back")}}
                                    </b-button>
                                </div>
                        </section>
                    </ValidationObserver>
                </div>
            </div>
            <div class="tile is-2 is-vertical is-parent">
            </div>
        </div>
    </div>
</template>

<script>
    import { ValidationObserver } from 'vee-validate';
    import BInputWithValidation from "../components/inputs/BInputWithValidation";

    export default {
        name: 'productAdd',
        components: {
            ValidationObserver,
            BInputWithValidation,
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
            addCategory(){
                this.$store.dispatch('categories/addCategory', this.form).then(() => this.$router.push('/categories'));
            },
        },
    }
</script>

<style>
</style>
