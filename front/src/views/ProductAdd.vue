<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("products.addMsg")}}</p>
                    <section>
                        <div class="columns is-mobile is-multiline">

                            <div class="column is-narrow">
                                <b-field :label="labels.productName">
                                    <b-input
                                             v-model="form.ProductName"
                                             placeholder="Enter name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.vendorCode">
                                    <b-input
                                            v-model="form.VendorCode"
                                            placeholder="Enter name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.categoryName">
                                    <b-select placeholder="Select a category"
                                              v-model="form.CategoryID"
                                              rounded
                                              size="is-small"
                                    >
                                        <option v-for="(option, key) in categories"
                                                :key="key"
                                                :value="option.ID"
                                        >
                                            {{option.CategoryName}}
                                        </option>
                                    </b-select>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.companyName">
                                    <b-select placeholder="Select a category"
                                              v-model="form.SupplierID"
                                              rounded
                                              size="is-small"
                                    >
                                        <option v-for="(option, key) in suppliers"
                                                :key="key"
                                                :value="option.ID"
                                        >
                                            {{option.CompanyName}}
                                        </option>
                                    </b-select>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.buyPrice">
                                    <b-numberinput
                                                   v-model="form.BuyPrice"
                                                   placeholder="Enter sockets"
                                                   rounded
                                                   size="is-small"
                                                   step="0.1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.sellPrice">
                                    <b-numberinput
                                                   v-model="form.SellPrice"
                                                   placeholder="Enter sockets"
                                                   rounded
                                                   size="is-small"
                                                   step="0.1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.amount">
                                    <b-numberinput
                                            v-model="form.Amount"
                                            placeholder="Enter sockets"
                                            rounded
                                            size="is-small"
                                            step="1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" >
                            <b-button  @click="addProduct()" type="is-success">{{$t("buttons.add")}}</b-button>
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
    import { mapGetters } from 'vuex';

    export default {
        name: 'productAdd',
        components: {
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
                console.log(this.form);
                this.$store.dispatch('products/addProduct', this.form);
            },

        },
    }
</script>

<style>
</style>
