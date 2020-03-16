<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title">{{$t("supply.addMsg")}}</p>
                        <ValidationObserver ref="observer">
                            <section slot-scope="{ validate }">
                                <div class="columns is-mobile is-multiline">

                                    <div class="column is-narrow">
                                            <BSelectWithValidation rules="required"
                                                                   :label="labels.productName"
                                                                   v-model="form.ProductId"
                                            >
                                                <option v-for="(option, key) in items"
                                                        :key="key"
                                                        :value="option.ID"
                                                >
                                                    {{option.ProductName}}
                                                </option>
                                            </BSelectWithValidation>
                                    </div>

                                    <div class="column is-narrow">

                                        <BSelectWithValidation rules="required"
                                                               :label="labels.companyName"
                                                               v-model="form.SupplierId"
                                        >
                                            <option v-for="(option, key) in suppliers"
                                                    :key="key"
                                                    :value="option.ID"
                                            >
                                                {{option.CompanyName}}
                                            </option>
                                        </BSelectWithValidation>
                                    </div>

                                    <div class="column is-narrow">
                                        <b-field :label="labels.buyPrice">
                                            <b-numberinput
                                                           v-model="form.BuyPrice"
                                                           placeholder="Enter buy price"
                                                           rounded
                                                           size="is-small"
                                                           step="0.05"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>

                                    <div class="column is-narrow">
                                        <b-field :label="labels.surcharge">
                                            <b-numberinput
                                                    :value="form.Surcharge"
                                                    @input="surchargeAdd($event, 2)"
                                                    placeholder="Enter surcharge"
                                                    rounded
                                                    size="is-small"
                                                    step="1"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>

                                    <div class="column is-narrow">
                                        <b-field :label="labels.sellPrice">
                                            <b-numberinput
                                                    :value="form.SellPrice"
                                                    @input="percentAdd($event)"
                                                    placeholder="Enter surcharge"
                                                    rounded
                                                    size="is-small"
                                                    step="0.05"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>

                                    <div class="column is-narrow">
                                        <b-field :label="labels.amount">
                                            <b-numberinput
                                                    v-model="form.Amount"
                                                    placeholder="Enter surcharge"
                                                    rounded
                                                    size="is-small"
                                                    step="1"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>

                                    <div class="column is-narrow">
                                        <BSelectWithValidation rules="required"
                                                               :label="labels.unit"
                                                               v-model="form.UnitId"
                                        >
                                            <option v-for="(option, key) in units"
                                                    :key="key"
                                                    :value="option.ID"
                                            >
                                                {{option.UnitName}}
                                            </option>
                                        </BSelectWithValidation>
                                    </div>

                                    <div class="column is-narrow">
                                        <BInputWithValidation rules="required"
                                                              type="textarea"
                                                              :label="labels.comment"
                                                              v-model="form.Comment"
                                        />
                                    </div>
                                </div>
                                <div class="buttons pt-2" >
                                    <b-button  @click="validate().then(addSupply)"
                                               type="is-link is-primary"
                                               class="is-small"
                                    >
                                        {{$t("buttons.add")}}
                                    </b-button>
                                    <b-button tag="router-link"
                                              to="/supply"
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
    import BInputWithValidation from "../components/inputs/BInputWithValidation";
    import BSelectWithValidation from '../components/inputs/BSelectWithValidation';
    import {ToastProgrammatic as Toast} from "buefy";
    import {i18n} from "../plugins/i18n";

    export default {
        name: 'productAdd',
        components: {
            ValidationObserver,
            BSelectWithValidation,
            BInputWithValidation,
        },
        data() {
            return {
                error: false,
                labels:{
                    productName: this.$t("products.table.productName"),
                    categoryName: this.$t("products.table.categoryName"),
                    companyName: this.$t("products.table.companyName"),
                    buyPrice: this.$t("products.table.buyPrice"),
                    surcharge: this.$t("supply.table.surcharge"),
                    sellPrice: this.$t("products.table.sellPrice"),
                    amount: this.$t("products.table.amount"),
                    vendorCode: this.$t("products.table.vendorCode"),
                    unit: this.$t("supply.table.unit"),
                    comment: this.$t("lables.comment"),
                },
                form:{
                    ProductId: '',
                    SupplierId: '',
                    BuyPrice: null,
                    SellPrice: null,
                    Surcharge: null,
                    Amount: null,
                    UnitId: "",
                    Comment: "",
                },
            }
        },
        mounted(){
            this.$store.dispatch('categories/getCategories');
            this.$store.dispatch('suppliers/getSuppliers');
            this.$store.dispatch('products/getProducts');
            this.$store.dispatch('supply/getUnits');
        },
        computed: {
            ...mapGetters("categories", [
                'categories',
            ]),
            ...mapGetters("suppliers", [
                'suppliers',
            ]),
            ...mapGetters("products", [
                'items',
            ]),
            ...mapGetters("supply", [
                'units',
            ]),
        },
        methods: {
            addSupply(){
                try {
                    if(this.form.Amount && this.form.BuyPrice && this.form.SellPrice && this.form.Surcharge){
                        this.error = false;
                        this.$store.dispatch('supply/addSupply', this.form).
                        then(() => this.$store.dispatch('credit/addCredit', this.form));
                    }else {
                        this.error = true;
                        Toast.open({
                            message: i18n.t("supply.toast.checkInput"),
                            type: 'is-danger'
                        });
                    }

                } catch (error) {
                    console.log('error')
                } finally {
                    if(!this.error){
                        setTimeout(this.$router.push('/supply'), 1000);
                    }
                }
            },
            surchargeAdd(event){
                let percent = event;
                this.form.Surcharge = Number(percent);
                this.form.SellPrice = Number((this.form.BuyPrice + (this.form.BuyPrice/100)*percent).toFixed(2));
            },
            percentAdd(event){
                let sellPrice = event;
                let onePercent = this.form.BuyPrice/100;
                this.form.Surcharge = Math.round((sellPrice - this.form.BuyPrice)/ onePercent);
                console.log(this.form.Surcharge)
            }

        },
    }
</script>

<style>
</style>
