<template>
    <div class="p-3">
            <div class="tile is-ancestor" v-for="(item, index) in products" :key="index">
                <div class="tile is-2 is-vertical is-parent">
                </div>
                <div class="tile is-vertical is-parent">
                    <div class="tile is-child box">
                        <p class="title">{{$t("products.updateMsg")}} #{{ item.ID }}</p>
                        <section>

                            <div class="columns is-mobile is-multiline">
                                <div class="column is-narrow">
                                    <b-field :label="labels.productName">
                                        <b-input :value="item.ProductName"
                                                 @input="updateName($event, item.ID)"
                                                 placeholder="Enter name"
                                                 rounded
                                                 size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.vendorCode">
                                        <b-input :value="item.VendorCode"
                                                 @input="updateVendorCode($event, item.ID)"
                                                 placeholder="Enter name"
                                                 rounded
                                                 size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.categoryName">
                                        <b-select placeholder="Select a category"
                                                  :value="item.CategoryID"
                                                  @input="updateCategoryID($event, item.ID)"
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
                                                  :value="item.SupplierID"
                                                  @input="updateCompanyID($event, item.ID)"
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
                                        <b-numberinput :value="item.BuyPrice"
                                                       @input="updateBuyPrice($event, item.ID)"
                                                       placeholder="Enter sockets"
                                                       rounded
                                                       size="is-small"
                                                       step="0.1"
                                        ></b-numberinput>
                                    </b-field>
                                </div>


                                <div class="column is-narrow">
                                    <b-field :label="labels.sellPrice">
                                        <b-numberinput :value="item.SellPrice"
                                                       @input="updateSellPrice($event, item.ID)"
                                                       placeholder="Enter sockets"
                                                       rounded
                                                       size="is-small"
                                                       step="0.1"
                                        ></b-numberinput>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.amount">
                                        <b-numberinput :value="item.Amount"
                                                       @input="updateAmount($event, item.ID)"
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
                            <div class="buttons pt-2" v-if="index === products.length - 1">
                                <b-button  @click="updateProduct()" type="is-success">{{$t("buttons.submit")}}</b-button>
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
    // import {ToastProgrammatic as Toast} from "buefy";

    export default {
        name: 'productsUpdate',
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
            products() {
                let answer = this.$store.getters['products/productsUpdate'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('products')));
                    this.$store.commit('products/SET_PRODUCTS_UPDATE', JSON.parse(localStorage.getItem('products')))
                    return this.$store.getters['products/productsUpdate'];
                }else {
                    return this.$store.getters['products/productsUpdate'];
                }
            }
        },
        methods: {
            updateProduct(){
                this.$store.dispatch('products/updateProducts');
            },
            updateName(event, id){
                let item = {'ID': id, 'ProductName': event};
                this.$store.commit('products/UPDATE_NAME', item)
            },
            updateCategoryID(event, id){
                let item = {'ID': id, 'CategoryID': event};
                this.$store.commit('products/UPDATE_CATEGORY_ID', item)
            },
            updateCompanyID(event, id){
                let item = {'ID': id, 'CompanyID': event};
                this.$store.commit('products/UPDATE_COMPANY_ID', item)
            },
            updateSellPrice(event, id){
                let item = {'ID': id, 'SellPrice': event};
                this.$store.commit('products/UPDATE_SELL_PRICE', item)
            },
            updateBuyPrice(event, id){
                let item = {'ID': id, 'BuyPrice': event};
                this.$store.commit('products/UPDATE_BUY_PRICE', item)
            },
            updateAmount(event, id){
                let item = {'ID': id, 'Amount': event};
                this.$store.commit('products/UPDATE_AMOUNT', item)
            },
            updateVendorCode(event, id){
                let item = {'ID': id, 'VendorCode': event};
                this.$store.commit('products/UPDATE_VENDOR_CODE', item)
            },
        },
    }
</script>

<style>
</style>
