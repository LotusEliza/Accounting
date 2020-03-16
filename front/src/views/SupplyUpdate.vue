<template>
    <div class="p-3">
        <div class="tile is-ancestor" v-for="(item, index) in supply" :key="index">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title"> {{$t("supply.updateMsg")}} #{{ item.ID }}</p>
                    <section>
                        <div class="columns is-mobile is-multiline">

                            <div class="column is-narrow">
                                <b-field :label="labels.productName">
                                    <b-select placeholder="Select a product"
                                              :value="item.ProductID"
                                              @input="updateProductID($event, item.ID)"
                                              rounded
                                              size="is-small"
                                    >
                                        <option v-for="(option, key) in items"
                                                :key="key"
                                                :value="option.ID"
                                        >
                                            {{option.ProductName}}
                                        </option>
                                    </b-select>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.companyName">
                                    <b-select placeholder="Select a category"
                                              :value="item.SupplierID"
                                              @input="updateSupplierID($event, item.ID)"
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
                                            :value="item.BuyPrice"
                                            @input="updateBuyPrice($event, item.ID)"
                                            placeholder="Enter buy price"
                                            rounded
                                            size="is-small"
                                            step="0.1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.surcharge">
                                    <b-input
                                            type="number"
                                            :value="item.Surcharge"
                                            @input="updateSurcharge($event, item.ID)"
                                            placeholder="Enter contact name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.sellPrice">
                                    <b-numberinput
                                            :value="item.SellPrice"
                                            @input="updateSellPrice($event, item.ID)"
                                            placeholder="Enter surcharge"
                                            rounded
                                            size="is-small"
                                            step="0.1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.supAmount">
                                    <b-numberinput
                                            :value="item.SupAmount"
                                            @input="updateSupAmount($event, item.ID)"
                                            placeholder="Enter surcharge"
                                            rounded
                                            size="is-small"
                                            step="1"
                                    ></b-numberinput>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.amount">
                                    <b-numberinput
                                            :value="item.Amount"
                                            @input="updateAmount($event, item.ID)"
                                            placeholder="Enter surcharge"
                                            rounded
                                            size="is-small"
                                            step="1"
                                    ></b-numberinput>
                                </b-field>
                            </div>


                            <div class="column is-narrow">
                                <b-field :label="labels.unit">
                                    <b-select placeholder="Select a unit"
                                              :value="item.UnitID"
                                              @input="updateUnitID($event, item.ID)"
                                              rounded
                                              size="is-small"
                                    >
                                        <option v-for="(option, key) in units"
                                                :key="key"
                                                :value="option.ID"
                                        >
                                            {{option.UnitName}}
                                        </option>
                                    </b-select>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.comment">
                                    <b-input
                                            type="textarea"
                                            :value="item.Comment"
                                            @input="updateComment($event, item.ID)"
                                            placeholder="Enter contact name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" v-if="index === supply.length - 1">
                            <b-button  @click="updateSupply()"
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
        name: 'suppliersUpdate',
        components: {
        },
        data() {
            return {
                name: '',
                selected: null,
                phone: '12345',
                labels:{
                    productName: this.$t("products.table.productName"),
                    categoryName: this.$t("products.table.categoryName"),
                    companyName: this.$t("products.table.companyName"),
                    buyPrice: this.$t("products.table.buyPrice"),
                    surcharge: this.$t("supply.table.surcharge"),
                    sellPrice: this.$t("products.table.sellPrice"),
                    amount: this.$t("warehouse.table.amount"),
                    supAmount: this.$t("warehouse.table.supAmount"),
                    vendorCode: this.$t("products.table.vendorCode"),
                    unit: this.$t("supply.table.unit"),
                    comment: this.$t("lables.comment"),
                },
            }
        },
        mounted(){
            this.$store.dispatch('suppliers/getSuppliers');
            this.$store.dispatch('products/getProducts');
            this.$store.dispatch('supply/getUnits');
            // this.$store.dispatch('categories/getCategories');
            // this.$store.dispatch('suppliers/getSuppliers');
        },
        computed: {
            filteredDataArray() {
                return this.data.filter((option) => {
                    return option
                        .toString()
                        .toLowerCase()
                        .indexOf(this.name.toLowerCase()) >= 0
                })
            },
            ...mapGetters("suppliers", [
                'suppliers',
            ]),
            ...mapGetters("suppliers", [
                'companiesNames',
            ]),
            ...mapGetters("products", [
                'items',
            ]),
            ...mapGetters("supply", [
                'units',
            ]),
            supply() {
                let answer = this.$store.getters['supply/supplyUpdate'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('supply')));
                    this.$store.commit('supply/SET_SUPPLY_UPDATE', JSON.parse(localStorage.getItem('supply')))
                    return this.$store.getters['supply/supplyUpdate'];
                }else {
                    return this.$store.getters['supply/supplyUpdate'];
                }
            }
        },
        methods: {
            async updateSupply(){
                try {
                    await   this.$store.dispatch('supply/updateSupply');
                } catch (error) {
                    console.log('error')
                } finally {
                    await this.$router.push('/supply');
                }
                // this.$store.dispatch('supply/updateSupply')
            },
            updateSurcharge(event, id){
                let item = { 'ID': id, 'Surcharge': event };
                this.$store.commit('supply/UPDATE_SURCHARGE', item);
            },
            updateProductID(event, id){
                // console.log(event)
                // console.log(id)
                // this.name = event;
                let item = { 'ID': id, 'ProductID': event };
                this.$store.commit('supply/UPDATE_PRODUCT_ID', item);
            },
            updateSupplierID(event, id){
                // console.log('update name comp');
                let item = { 'ID': id, 'SupplierID': event };
                this.$store.commit('supply/UPDATE_SUPPLIER_ID', item);
            },
            updateBuyPrice(event, id){
                // console.log('update name contact'+event+id);
                let item = {'ID': id, 'BuyPrice': event};
                this.$store.commit('supply/UPDATE_BUY_PRICE', item)
            },
            updateSellPrice(event, id){
                // console.log('update name contact'+event+id);
                let item = {'ID': id, 'SellPrice': event};
                this.$store.commit('supply/UPDATE_SELL_PRICE', item)
            },
            updateUnitID(event, id){
                let item = {'ID': id, 'UnitID': event};
                this.$store.commit('supply/UPDATE_UNIT_ID', item)
            },
            updateAmount(event, id){
                let item = {'ID': id, 'Amount': event};
                this.$store.commit('supply/UPDATE_AMOUNT', item)
            },
            updateSupAmount(event, id){
                let item = {'ID': id, 'SupAmount': event};
                this.$store.commit('supply/UPDATE_SUP_AMOUNT', item)
            },
            updateComment(event, id){
                let item = {'ID': id, 'Comment': event};
                this.$store.commit('supply/UPDATE_COMMENT', item)
            },

        },
    }
</script>

<style>
    .vti__input {
        border-radius: 0 20px 20px 0 !important;
    }
    .vue-tel-input {
        border-radius: 0 20px 20px 0 !important;
    }
    .textarea:not([rows]){
        border-radius: 20px 20px 0 20px !important;
    }
</style>
