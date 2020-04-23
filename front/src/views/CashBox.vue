<template>
    <div class="columns">
        <div class="column is-2">
        </div>
        <div class="column pt-30 plr-10">
            <div class="tile is-ancestor plr-10 ">
                <div class="tile is-vertical box">
                    <div class="tile">
                        <p class="title is-hidden-tablet">{{$t("titles.cashBox")}}</p>
                        <div class="tile is-parent is-vertical ">
<!--_______________________________ INPUT PRODUCT_NAME/VENDOR_CODE _______________________________________-->
                            <b-field :label="labels.product">
                                <b-autocomplete
                                        rounded
                                        size="is-small"
                                        v-model="form.productName"
                                        :data="filteredDataArray"
                                        ref="productName"
                                        @keyup.enter.native="addTransaction()"
                                >
                                    <template slot="empty">No results found</template>
                                </b-autocomplete>
                            </b-field>
                        </div>
                        <div class="tile is-parent">
<!--_______________________________ INPUT AMOUNT _______________________________________-->
                            <b-field :label="labels.amount">
                                <b-input v-model="form.amount"
                                         rounded
                                         size="is-small"
                                         ref="amount"
                                         @keyup.enter.native="addTransaction()"
                                ></b-input>
                            </b-field>
                        </div>
                        <div class="tile is-parent">
                            <div class="buttons pt-30" >
                                <b-button  @click="addTransaction()"
                                           type="is-light"
                                           size="is-small"
                                >
                                    {{$t("buttons.add")}}
                                </b-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="tile is-ancestor plr-10" v-if="showTrans">
                <div class="tile is-vertical">
                    <div class="tile is-child box" >
                        <div class="tile is-parent is-vertical">
                            <p class="title"> {{$t("cashBox.transaction")}} </p>
<!--_______________________________ TABLE TRANSACTION _______________________________________-->
                            <b-table
                                    :data="isEmpty ? [] : transaction"
                                    :striped="isStriped"
                                    :narrowed="isNarrowed"
                                    :hoverable="isHoverable"
                                    :loading="isLoading"
                                    :mobile-cards="hasMobileCards"
                            >

                                <template slot-scope="props">
                                    <b-table-column field="ID" label="ID" width="40" numeric>
                                        {{ props.row.id }}
                                    </b-table-column>

                                    <b-table-column field="Name" :label="labels.product" >
                                        {{ props.row.name }}
                                    </b-table-column>

                                    <b-table-column field="Price" :label="labels.price">
                                        {{ props.row.price }}
                                    </b-table-column>

                                    <b-table-column field="Amount" :label="labels.amount" v-if="props.row.unit !== 'kg'">
                                        {{ props.row.amount }}
                                    </b-table-column>

                                    <b-table-column field="Amount" :label="labels.amount" v-if="props.row.unit === 'kg'">
                                        {{ (props.row.amount/1000).toFixed(3) }}
                                    </b-table-column>

                                    <b-table-column field="Total" :label="labels.total">
                                        {{ props.row.total }}
                                    </b-table-column>

                                    <b-table-column field="" label="" >
                                        <b-button @click="removeItem(props.row.unicId)"
                                                  type="is-primary"
                                                  size="is-small"
                                        >
                                            X
                                        </b-button>
                                    </b-table-column>
                                </template>
                                <template slot="empty">
                                    <section class="section">
                                        <div class="content has-text-grey has-text-centered">
                                            <p>
                                                <b-icon
                                                        icon="emoticon-sad"
                                                        size="is-large">
                                                </b-icon>
                                            </p>
                                            <p>Nothing here.</p>
                                        </div>
                                    </section>
                                </template>
                            </b-table>
                            <hr>
                            <p class="subtitle"> {{$t("cashBox.labels.sum")}} </p>
                            {{totalForTrans}}
                            <div class="buttons pt-2" >
                                <b-button  @click="makeTransaction()"
                                           type="is-primary"
                                           class="is-small"
                                >
                                    {{$t("buttons.buy")}}
                                </b-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="column is-2">
        </div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    import {ToastProgrammatic as Toast} from "buefy";
    import {i18n} from "../plugins/i18n";

    export default {
        name: 'cashBox',
        components: {
        },
        data() {
            return {
                componentKey: 0,
                form:{
                    productName: '',
                    amount: '',
                },
                transaction: [],
                visible: false,
                showTrans: false,
                totalForTrans: null,
                //************table******************
                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
                unicId: 0,
                labels: {
                    product: this.$t("warehouse.table.productName"),
                    amount: this.$t("products.table.amount"),
                    price: this.$t("cashBox.labels.price"),
                    total: this.$t("cashBox.labels.total"),
                }
            }
        },
        computed: {
            ...mapGetters("products", [
                'products',
            ]),
            ...mapGetters("supply", [
                'items',
            ]),
            ...mapGetters("debit", [
                'debit',
            ]),
            names(){
                let moreThenOne = this.items.filter(p => p.Amount > 0);
                let names = [...new Set(moreThenOne.map((p) => p.ProductName))];
                return names;
            },
            codes(){
                //what products we have in our warehouse (ids)
                let moreThenOne = this.items.filter(p => p.Amount > 0);
                let ids = [...new Set(moreThenOne.map((p) => p.ProductID))];
                let array = [];
                for (let i = 0; i < ids.length; i++) {
                    let productsInStock = this.products.filter(p => p.ID == ids[i]);
                    array.push(productsInStock)
                }
                let codes = [];
                for (let i = 0; i < array.length; i++) {
                    let code = array[i].map((p) => p.VendorCode);
                    codes.push(code[0])
                }
                return codes;
            },
            filteredDataArray() {
                if(isNaN(this.form.productName)){
                    return this.names.filter((option) => {
                        return option
                            .toString()
                            .toLowerCase()
                            .indexOf(this.form.productName.toLowerCase()) >= 0
                    })
                }else {
                    return this.codes.filter((option) => {
                        return option
                            .toString()
                            .toLowerCase()
                            .indexOf(this.form.productName.toLowerCase()) >= 0
                    })
                }
            },
        },
        watch: {
            transaction(){
                if(!this.transaction.length){
                    this.unicId = 0;
                    this.showTrans = false;
                }
            }
        },
        mounted(){
            this.focusInput();
            this.$store.dispatch('categories/getCategories');
            this.$store.dispatch('products/getProducts');
            this.$store.dispatch('supply/getSupply');
        },
        methods:{
            focusInput(){
                this.$refs.productName.focus();
            },
            /**
             * generates an array of dates
             * @param arra | array [amount: 4, id: 129, inStock: 82, name: "молоко2", price: 374, supplyId: total: ,unicId: ,unit: ]
             * @param conditionFn - item => item.id == existId[0].id
             * @return index | number (1)
             */
            indexWhere(array, conditionFn) {
                const item = array.find(conditionFn);
                return array.indexOf(item);
            },
            /**
             * runs when entered amount is more then in stock
             */
            tooMuch(inStock){
                Toast.open({
                    message: i18n.t("cashBox.toast.notEnough")+inStock,
                    type: 'is-danger'
                });
            },
            /**Highest price of all supplies
             * @param array | array [ 0: { Amount: BuyPrice: Comment: SellPrice:... }, 1: {...}]
             * @return max | number (22.50)
             */
            highest(array){
                let prices = array.map((p) => p.SellPrice);
                return  Math.max.apply(Math, prices);
            },
            pushArray(id, supplyId, name, amount, price, total, unit, inStock){
                this.transaction.push({
                    unicId: this.unicId,
                    id: id,
                    supplyId: supplyId,
                    name: name,
                    amount: amount,
                    price: price,
                    total: total,
                    unit: unit,
                    inStock: inStock-amount,
                });
                this.unicId += 1;
            },
            logic2(result2){
                let inStock = this.inStock(result2);
                let highestPrice = this.highest(result2);

                if (this.form.amount === '') {
                    if(result2[0].UnitName === 'kg') {
                        console.log('kg');
                        Toast.open({
                            message: i18n.t("cashBox.toast.noWeight"),
                            type: 'is-danger'
                        });
                    }else {
                        this.pushArray(result2[0].ProductID, result2[0].ID, result2[0].ProductName,
                            1, highestPrice, highestPrice,'', inStock);
                        this.showTrans = true;
                    }
                } else {
                    if(inStock >= this.form.amount) {
                        if(result2[0].UnitName === 'kg') {
                            let priceForGrm = highestPrice/1000;
                            this.pushArray(result2[0].ProductID, result2[0].ID, result2[0].ProductName,
                                this.form.amount, highestPrice,
                                (priceForGrm * this.form.amount).toFixed(2), result2[0].UnitName, inStock)
                        }else {
                            this.pushArray(result2[0].ProductID, result2[0].ID, result2[0].ProductName,
                                this.form.amount, highestPrice,
                                highestPrice * this.form.amount,'', inStock);
                        }
                        this.showTrans = true;
                    }else {
                        console.log("Func inStock >= this.form.amount")
                        this.tooMuch(inStock);
                    }
                }
            },
            /**
             * Distinguish if input field has a Vendor Code or Product Name
             */
            logic1(){
                if(isNaN(this.form.productName)) {
                    // console.log('Not a number');
                    const result = this.products.filter(res=>res.ProductName == this.form.productName);
                    const result2 = this.items.filter(res=>res.ProductID == result[0].ID);
                    this.logic2(result2);
                }else {
                    let result = this.products.filter(res=>res.VendorCode == this.form.productName);
                    let result2 = this.items.filter(res=>res.ProductID == result[0].ID);
                    this.logic2(result2);
                }
                this.clearInput();
                this.focusInput();
            },
            /**
             * Current amount in stock (all supplies)
             * @param array | array [ 0: { Amount: BuyPrice: Comment: SellPrice:... }, 1: {...}]
             * @return inStock | number (2)
             */
            inStock(array){
                let inStock = 0;
                for (let i = 0; i < array.length; i++) {
                    inStock += Number(array[i].Amount);
                }
                return inStock;
            },
            /**
             * Add transaction (check if this item already exist,
             * @param array | array [ 0: { Amount: BuyPrice: Comment: SellPrice:... }, 1: {...}]
             * @return inStock | number (2)
             */
            addTransaction(){
                this.showTrans = true;
                if(this.transaction.length > 0){
                        let existId;
                        let productName;
                        //if entered info is prodName  or vendorCode
                        if(!isNaN(this.form.productName)) {
                            let product = this.products.filter(p => p.VendorCode == this.form.productName);
                            productName = product[0].ProductName
                        }else {
                            productName = this.form.productName;
                        }
                        //check if we already have this prod in transaction array
                        existId = this.transaction.filter(function(elem){
                            if(elem.name == productName) return elem.id;
                        });
                        //if there is already such item in transaction array
                        if(existId.length > 0){
                            const result = this.transaction.filter(res=>res.id == existId[0].id);
                            let totalAmount = Number(result[0].amount) + Number(this.form.amount);
                            let result2 = this.items.filter(res=>res.ProductID == result[0].id);
                            let inStock = this.inStock(result2);
                            if(totalAmount >= inStock){
                                this.tooMuch(inStock-Number(result[0].amount));
                            }else if(totalAmount <= inStock){
                                if(this.form.amount === ''){
                                    this.form.amount = 1;
                                }
                                result[0].total = (Number(result[0].total) + (Number(result[0].total)/Number(result[0].amount)) * this.form.amount).toFixed(2);
                                result[0].amount = Number(result[0].amount) + Number(this.form.amount);
                                result[0].inStock = Number(result[0].inStock) - Number(this.form.amount);

                                this.transaction[this.indexWhere(this.transaction, item => item.id == existId[0].id)] = result[0];
                                this.clearInput();
                                this.focusInput();
                            }
                        }else {
                            //if no such item in transaction array
                            this.logic1();
                        }
                    this.totalTrans();
                }else{
                    //transaction array is empty
                    this.logic1();
                    this.totalTrans()
                }
            },
            /**
             * Counts the total amount of money for transaction
             */
            totalTrans(){
                let totalForTrans = null;
                for (let i = 0; i < this.transaction.length; i++) {
                    totalForTrans += Number(this.transaction[i].total);
                }
                this.totalForTrans = totalForTrans;
            },
            clearInput(){
                this.form.amount = '';
                this.form.productName = '';
                return true
            },
            removeItem(key){
                this.totalForTrans -= this.transaction[key].total;
                this.transaction.splice(key, 1);
                this.unicId -= 1;
            },
            makeTransaction(){
                this.$store.dispatch('supply/getSupply');
                this.$store.dispatch('debit/addTransaction',  this.totalForTrans);
                this.$store.dispatch('debit/lastDebitId',  this.transaction).
                then(() => this.$store.dispatch('debit/addDebitsProducts',  this.transaction));
                for (let i = 0; i < this.transaction.length; i++) {
                    this.$store.dispatch('supply/updateSupplyAmountTest',  this.transaction[i]);
                }
                this.showTrans = false;
                this.transaction = [];
                this.$emit('update')
            },
            forceRerender() {
                this.componentKey += 1;
            },
        }
    }
</script>

<style lang="scss">

    li:first-child {
        color: Red !important;
        background-color: aquamarine;
    }
    .box{
        background-color: rgba(254, 251, 235, 0.5) !important;
    }

</style>
