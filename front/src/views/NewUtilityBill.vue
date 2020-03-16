<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-3 is-vertical is-parent">
            </div>
            <div class="tile is-6 is-vertical is-parent">

                <p class="title">Коммунальные платежи</p>
<!--_______________________________ TABLE _______________________________________-->
                <b-table
                        :data="isEmpty ? [] : bills"
                        :striped="isStriped"
                        :narrowed="isNarrowed"
                        :hoverable="isHoverable"
                        :loading="isLoading"
                        :mobile-cards="hasMobileCards">

                    <template slot-scope="props">

                        <b-table-column field="Date" label="ID">
                            {{ props.row.ID }}
                        </b-table-column>

                        <b-table-column field="Date" :label="labels.date">
                            {{ props.row.DateCreate | formatDate }}
                        </b-table-column>

                        <b-table-column field="Date" :label="labels.amount"  style="text-align: center">
                            {{ props.row.Amount }}
                        </b-table-column>

                        <b-table-column field="Date" label="Название"  style="text-align: center">
                            {{ props.row.BillName }}
                        </b-table-column>

                        <b-table-column field="" label="  ">
                            <b-button @click="removeBill(props.row.ID)"
                                      type="is-light"
                                      :disabled="visible"
                                      v-show="visible !== props.row.Date"
                                      size="is-small"
                                      expanded
                            >
                                Remove
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
<!--_______________________________ BUTTONS ADD CANCEL_______________________________________-->
                <b-button @click="show = true"
                          type="is-primary"
                          expanded
                          v-show="!show"
                >
                    Добавить платеж
                </b-button>
                <b-button @click="cancelAddNorm"
                          type="is-primary"
                          expanded
                          v-show="show"
                >
                    Отмена
                </b-button>
<!--_______________________________ FORM ADD UTILITY BILL _______________________________________-->
                <transition name="fade">
                    <div class="tile is-ancestor" v-if="show">
                        <div class="tile is-vertical is-parent">
                            <div class="tile is-child box">
                                <ValidationObserver ref="observer">
                                    <section slot-scope="{ validate }">
                                        <div class="columns is-mobile is-multiline">

                                            <div class="column">
                                                <b-field label="Дата">
                                                    <Datepicker v-model="newBill.DateCreate" :language="ru" :monday-first="true"></Datepicker>
                                                </b-field>
                                            </div>

                                            <div class="column">
                                                <b-field label="Сумма">
                                                    <BInputWithValidation rules="required"
                                                                          type="text"
                                                                          v-model="newBill.Amount"
                                                    />
                                                </b-field>
                                            </div>

                                            <div class="column">
                                                <b-field label="Назначение">
                                                    <b-autocomplete
                                                            rounded
                                                            size="is-small"
                                                            v-model="newBill.BillName"
                                                            :data="filteredDataArray"
                                                            placeholder="вода"
                                                            @select="option => selected = option"
                                                    >
                                                        <template slot="empty">No results found</template>
                                                    </b-autocomplete>
                                                </b-field>
                                            </div>
                                        </div>
<!--_______________________________ BUTTONS ADD RESET _______________________________________-->
                                        <div class="buttons pt-2">
                                            <b-button  @click="validate().then(saveNewBill)"
                                                       type="is-primary"
                                            >
                                                Add
                                            </b-button>
                                            <b-button @click="reset()"
                                                      type="is-light"
                                            >
                                                Reset
                                            </b-button>
                                        </div>
                                    </section>
                                </ValidationObserver>
                            </div>
                        </div>
                    </div>
                </transition>
            </div>
            <div class="tile is-3 is-vertical is-parent">
            </div>
        </div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    import { ValidationObserver } from 'vee-validate';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';
    import Datepicker from 'vuejs-datepicker';
    import {en, ru} from 'vuejs-datepicker/dist/locale'
    import {ToastProgrammatic as Toast} from "buefy";


    export default {
        name: 'norms',
        components: {
            ValidationObserver,
            BInputWithValidation,
            Datepicker
        },
        data() {
            return {

                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
                options: {
                    "resourceType": [
                        {text: 'gas', value: 1},
                        {text: 'metal', value: 2},
                        {text: 'oil', value: 3},
                        {text: 'send', value: 4},
                    ],
                },
                labels:{
                    date: this.$t("warehouse.table.dateCreate"),
                    amount: this.$t("products.table.amount"),
                    contactTitle: this.$t("suppliers.table.contactTitle"),
                    address: this.$t("suppliers.table.address"),
                    city: this.$t("suppliers.table.city"),
                    phone: this.$t("suppliers.table.phone"),
                    email: this.$t("suppliers.table.email"),
                },
                show: false,
                visible: false,
                newBill: {
                    'DateCreate': '',
                    'BillName': '',
                    'Amount': null,
                },
                objAmount: null,
                amount: null,
                objResource: null,
                //dataPicker
                ru: ru,
                en: en,
                //auto select dropdown
                selected: null,
            }
        },
        mounted() {
            this.$store.dispatch('new_utility_bills/getBills');
            this.$store.dispatch('new_utility_bills/getBillsNames');
        },
        computed: {
            ...mapGetters("new_utility_bills", [
                'bills', 'billsNames'
            ]),

            filteredDataArray() {
                return this.billsNames.filter((option) => {
                    return option
                        .toString()
                        .toLowerCase()
                        .indexOf(this.newBill.BillName.toLowerCase()) >= 0
                })
            },
        },

        methods:{
            removeBill(date){
                this.$store.dispatch('new_utility_bills/removeBill', date);
            },

            cancelAddNorm(){
                this.show = false;
                this.cleanFields();
            },
            cleanFields(){
                this.newBill.Amount = null;
                this.newBill.DateCreate = null;
                this.newBill.BillName = "";
            },
            reset(){
               this.cleanFields();
                requestAnimationFrame(() => {
                    this.$refs.observer.reset();
                });
            },

            async saveNewBill(){
                if(this.newBill.DateCreate && this.newBill.BillName){
                    let copiedObject = Object.assign({}, this.newBill);
                    await this.$store.dispatch('new_utility_bills/addBill', copiedObject).then(() =>  this.reset());
                }else {
                    Toast.open({
                        message: "Пожалуйста заполните все поля!",
                        type: 'is-danger'
                    });
                }

                // this.show = false;
            }
        }
    }
</script>

<style>
    .vdp-datepicker * {
        border-radius: 15px;
    }

    .vdp-datepicker {
        padding-top: 3px;
    }

    .vdp-datepicker input:focus{
        box-shadow: 0 0 0.5pt 1pt #3273dc !important;
        outline: none !important;
        /*border-radius: 15px !important;*/
    }
        /*************Transition******************/
    .fade-enter-active, .fade-leave-active {
        transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
        opacity: 0;
    }

</style>
