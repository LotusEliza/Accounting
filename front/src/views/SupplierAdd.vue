<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title"> {{$t("suppliers.addMsg")}}</p>

                    <ValidationObserver ref="observer">
                        <section slot-scope="{ validate }">

                            <div class="columns is-mobile is-multiline">

                                <div class="column is-narrow">
                                    <b-field :label="labels.companyName">
                                        <b-input
                                                 v-model="form.CompanyName"
                                                 rounded
                                                 size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <BInputWithValidation rules="required"
                                                          type="text"
                                                          :label="labels.contactName"
                                                          v-model="form.ContactName"
                                    />
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.city">
                                        <b-autocomplete
                                                rounded
                                                size="is-small"
                                                v-model="form.City"
                                                :data="filteredDataArray"
                                                placeholder="Николаев"
                                                @select="option => selected = option"
                                        >
                                            <template slot="empty">No results found</template>
                                        </b-autocomplete>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.phone">
                                        <vue-tel-input v-model="phone" v-bind="bindProps" @input="addPhone()"></vue-tel-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.email">
                                        <b-input type="email"
                                                 v-model="form.Email"
                                                 rounded
                                                 size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.address">
                                        <b-input
                                                type="textarea"
                                                v-model="form.Address"
                                                rounded
                                                size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                                <div class="column is-narrow">
                                    <b-field :label="labels.contactTitle">
                                        <b-input
                                                type="textarea"
                                                v-model="form.ContactTitle"
                                                rounded
                                                size="is-small"
                                        ></b-input>
                                    </b-field>
                                </div>

                            </div>
                            <div class="buttons pt-2" >
                                <b-button  @click="validate().then(addSupplier)"
                                           type="is-link is-primary"
                                           class="is-small"
                                >
                                    {{$t("buttons.add")}}
                                </b-button>
                                <b-button tag="router-link"
                                          to="/suppliers"
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
    import { VueTelInput } from 'vue-tel-input'
    import { data } from '../utils/variables.js'
    import { ValidationObserver } from 'vee-validate';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';

    export default {
        name: 'suppliersAdd',
        components: {
            VueTelInput,
            ValidationObserver,
            BInputWithValidation,
        },
        data() {
            return {
                bindProps: {
                    mode: "international",
                    defaultCountry: "UA",
                    disabledFetchingCountry: false,
                    disabled: false,
                    disabledFormatting: false,
                    required: false,
                    enabledCountryCode: false,
                    enabledFlags: true,
                    preferredCountries: ["UA", "RU", "CN", "US"],
                    ignoredCountries: [],
                    autocomplete: "off",
                    name: "telephone",
                    maxLen: 25,
                    wrapperClasses: "",
                    inputClasses: "",
                    inputOptions: {
                        showDialCode: true
                    }
                },
                phone: null,
                //autocomplite
                data: data,
                name: '',
                selected: null,
                labels:{
                    companyName: this.$t("suppliers.table.companyName"),
                    contactName: this.$t("suppliers.table.contactName"),
                    contactTitle: this.$t("suppliers.table.contactTitle"),
                    address: this.$t("suppliers.table.address"),
                    city: this.$t("suppliers.table.city"),
                    phone: this.$t("suppliers.table.phone"),
                    email: this.$t("suppliers.table.email"),
                },
                form: {
                    CompanyName: '',
                    ContactName: '',
                    ContactTitle: '',
                    Address: '',
                    City: '',
                    Phone: '',
                    Email: '',
                }
            }
        },
        computed: {
            filteredDataArray() {
                return this.data.filter((option) => {
                    return option
                        .toString()
                        .toLowerCase()
                        .indexOf(this.form.City.toLowerCase()) >= 0
                })
            },
        },
        methods: {
            addPhone(){
                this.form.Phone = this.phone;
            },
            addSupplier(){
                this.$store.dispatch('suppliers/addSupplier', this.form).then(() => this.$router.push('/suppliers'));
            }
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
