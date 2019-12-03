<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-2 is-vertical is-parent">
                {{form}}
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title"> {{$t("suppliers.addMsg")}}</p>
                    <section>
                        <div class="columns is-mobile is-multiline">

                            <div class="column is-narrow">
                                <b-field :label="labels.companyName">
                                    <b-input
                                             v-model="form.CompanyName"
                                             placeholder="Enter company name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.contactName">
                                    <b-input
                                             v-model="form.ContactName"
                                             placeholder="Enter contact name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.contactTitle">
                                    <b-input
                                             v-model="form.ContactTitle"
                                             placeholder="Enter contact title"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.city">
                                    <b-autocomplete
                                            rounded
                                            size="is-small"
                                            v-model="form.City"
                                            :data="filteredDataArray"
                                            placeholder="Select City"
                                            @select="option => selected = option">
                                        <template slot="empty">No results found</template>
                                    </b-autocomplete>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.email">
                                    <b-input type="email"
                                             v-model="form.Email"
                                             placeholder="Enter contact name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.phone">
                                    <vue-tel-input v-model="phone" v-bind="bindProps" @input="addPhone()"></vue-tel-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.address">
                                    <b-input
                                            type="textarea"
                                            v-model="form.Address"
                                            placeholder="Enter contact name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" >
                            <b-button  @click="addSupplier()" type="is-success">{{$t("buttons.add")}}</b-button>
                            <b-button tag="router-link"
                                      to="/suppliers"
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
    // import { mapGetters } from 'vuex';
    import { VueTelInput } from 'vue-tel-input'

    export default {
        name: 'suppliersAdd',
        components: {
            VueTelInput,
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
                    // onlyCountries: ["RU", "US", "CH"],
                    ignoredCountries: [],
                    autocomplete: "off",
                    name: "telephone",
                    maxLen: 25,
                    wrapperClasses: "",
                    inputClasses: "",
                    // dropdownOptions: {
                    //     disabledDialCode: false
                    // },
                    inputOptions: {
                        showDialCode: true
                    }
                },
                phone: null,
                // products: [],
                //autocomplite
                data: [
                    'Kyiv',
                    'Kharkiv',
                    'Odessa',
                    'Dnipro',
                    'Donetsk',
                    'Lviv',
                    'Mykolaiv',
                    'Sevastopol',
                    'Kherson',
                    'Sumy',
                    'Rivne',
                    'Ivano-Frankivsk'
                ],

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
                        .indexOf(this.name.toLowerCase()) >= 0
                })
            },
        },
        methods: {
            addPhone(){
                this.form.Phone = this.phone;
            },
            addSupplier(){
                console.log(this.form)
                this.$store.dispatch('suppliers/addSupplier', this.form);
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
