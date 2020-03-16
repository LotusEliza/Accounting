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
                                                          type="text"
                                                          :label="labels.productName"
                                                          v-model="form.ProductName"
                                    />
                                </div>

                                <div class="column">
                                    <BInputWithValidation rules="required"
                                                          type="text"
                                                          :label="labels.vendorCode"
                                                          v-model="form.VendorCode"
                                    />
                                </div>

                                <div class="column">
                                    <BSelectWithValidation rules="required"
                                                           :label="labels.categoryName"
                                                           v-model="form.CategoryID"
                                    >
                                        <option v-for="(option, key) in categories"
                                                :key="key"
                                                :value="option.ID"
                                        >
                                            {{option.CategoryName}}
                                        </option>
                                    </BSelectWithValidation>
                                </div>

                            </div>
                            <div class="buttons pt-2" >
                                <b-button  @click="validate().then(addProduct)"
                                           type="is-link is-primary"
                                           class="is-small"
                                >
                                    {{$t("buttons.add")}}
                                </b-button>
                                <b-button tag="router-link"
                                          to="/products"
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
    import { mapGetters } from 'vuex';
    import { ValidationObserver } from 'vee-validate';
    import BSelectWithValidation from '../components/inputs/BSelectWithValidation';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';

    export default {
        name: 'productAdd',
        components: {
            ValidationObserver,
            BSelectWithValidation,
            BInputWithValidation,
        },
        data() {
            return {
                // products: [],
                labels:{
                    productName: this.$t("products.table.productName"),
                    categoryName: this.$t("products.table.categoryName"),
                    companyName: this.$t("products.table.companyName"),
                    buyPrice: this.$t("products.table.buyPrice"),
                    sellPrice: this.$t("products.table.sellPrice"),
                    amount: this.$t("products.table.amount"),
                    vendorCode: this.$t("products.table.vendorCode"),
                },
                form:{
                    ProductName: '',
                    CategoryID: '',
                    SupplierID: '',
                    BuyPrice: null,
                    SellPrice: null,
                    VendorCode: null,
                    Amount: null,
                }
            }
        },
        mounted(){
            this.$store.dispatch('categories/getCategories');
            this.$store.dispatch('suppliers/getSuppliers');
        },
        computed: {
            ...mapGetters("categories", [
                'categories',
            ]),
            ...mapGetters("suppliers", [
                'suppliers',
            ]),
        },
        methods: {
            addProduct(){
                this.$store.dispatch('products/addProduct', this.form).then(() => this.$router.push('/products'));
            },

        },
    }
</script>

<style>
</style>
