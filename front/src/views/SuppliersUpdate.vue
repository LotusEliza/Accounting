<template>
    <div class="p-3">
<!--        {{suppliers}}-->
        <div class="tile is-ancestor" v-for="(item, index) in suppliers" :key="index">
            <div class="tile is-2 is-vertical is-parent">
            </div>
            <div class="tile is-vertical is-parent">
                <div class="tile is-child box">
                    <p class="title"> {{$t("suppliers.updateMsg")}} #{{ item.ID }}</p>
                    <section>

                        <div class="columns is-mobile is-multiline">
<!--                            companiesNames-->

                            <div class="column is-narrow">
                                <b-field :label="labels.companyName">
                                    <b-input :value="item.CompanyName"
                                             @input="updateCompanyName($event, item.ID)"
                                             placeholder="Enter company name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.contactName">
                                    <b-input :value="item.ContactName"
                                             @input="updateContactName($event, item.ID)"
                                             placeholder="Enter contact name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.contactName">
                                    <b-input :value="item.ContactTitle"
                                             @input="updateContactTitle($event, item.ID)"
                                             placeholder="Enter contact name"
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
                                            :value="item.City"
                                            @input="updateCity($event, item.ID)"
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
                                             :value="item.Email"
                                             @input="updateEmail($event, item.ID)"
                                             placeholder="Enter contact name"
                                             rounded
                                             size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field label="Phone">
                                    <vue-tel-input v-model="phone"></vue-tel-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.phone">
                                    <b-input
                                            type="number"
                                            :value="item.Phone"
                                            @input="updatePhone($event, item.ID)"
                                            placeholder="Enter contact name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                            <div class="column is-narrow">
                                <b-field :label="labels.address">
                                    <b-input
                                            type="textarea"
                                            :value="item.Address"
                                            @input="updateAddress($event, item.ID)"
                                            placeholder="Enter contact name"
                                            rounded
                                            size="is-small"
                                    ></b-input>
                                </b-field>
                            </div>

                        </div>
                    </section>
                    <section>
                        <div class="buttons pt-2" v-if="index === suppliers.length - 1">
                            <b-button  @click="updateSuppliers()" type="is-success">{{$t("buttons.submit")}}</b-button>
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
    import { VueTelInput } from 'vue-tel-input'
    // import {ToastProgrammatic as Toast} from "buefy";

    export default {
        name: 'suppliersUpdate',
        components: {
            VueTelInput,
        },
        data() {
            return {
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
                phone: '12345',
                labels:{
                    companyName: this.$t("suppliers.table.companyName"),
                    contactName: this.$t("suppliers.table.contactName"),
                    contactTitle: this.$t("suppliers.table.contactTitle"),
                    address: this.$t("suppliers.table.address"),
                    city: this.$t("suppliers.table.city"),
                    phone: this.$t("suppliers.table.phone"),
                    email: this.$t("suppliers.table.email"),
                },
            }
        },
        mounted(){
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
                'companiesNames',
            ]),
            suppliers() {
                let answer = this.$store.getters['suppliers/suppliersUpdate'].length;
                if(!answer){
                    console.log(JSON.parse(localStorage.getItem('suppliers')));
                    this.$store.commit('suppliers/SET_SUPPLIERS_UPDATE', JSON.parse(localStorage.getItem('suppliers')))
                    return this.$store.getters['suppliers/suppliersUpdate'];
                }else {
                    return this.$store.getters['suppliers/suppliersUpdate'];
                }
            }
        },
        methods: {
            updateSuppliers(){
                this.$store.dispatch('suppliers/updateSupplier')
                // let count = null;
                // for (let i = 0; i < this.$store.getters['suppliers/suppliersUpdate'].length; i++) {
                //     const response = await this.$store.dispatch('suppliers/updateSupplier');
                //     if(response === 'OK'){
                //         count++;
                //     }
                // }
                // if(this.$store.getters['suppliers/suppliersUpdate'].length === count){
                //     Toast.open({
                //         message: this.$t("suppliers.toast.update"),
                //         type: 'is-success'
                //     });
                // }
            },
            updateCity(event, id){
                this.name = event;
                let item = { 'ID': id, 'City': this.name };
                this.$store.commit('suppliers/UPDATE_CITY', item);
            },
            updateCompanyName(event, id){
                console.log('update name comp');
                let item = { 'ID': id, 'CompanyName': event };
                this.$store.commit('suppliers/UPDATE_COMPANY_NAME', item);
            },
            updateContactName(event, id){
                console.log('update name contact'+event+id);
                let item = {'ID': id, 'ContactName': event};
                this.$store.commit('suppliers/UPDATE_CONTACT_NAME', item)
            },
            updateContactTitle(event, id){
                console.log('update cont title'+event+id);
                let item = {'ID': id, 'ContactTitle': event};
                this.$store.commit('suppliers/UPDATE_CONTACT_TITLE', item)
            },
            updateEmail(event, id){
                console.log('update cont title'+event+id);
                let item = {'ID': id, 'Email': event};
                this.$store.commit('suppliers/UPDATE_EMAIL', item)
            },
            updatePhone(event, id){
                console.log('update cont title'+event+id);
                let item = {'ID': id, 'Phone': event};
                this.$store.commit('suppliers/UPDATE_PHONE', item)
            },
            updateAddress(event, id){
                console.log('update cont title'+event+id);
                let item = {'ID': id, 'Address': event};
                this.$store.commit('suppliers/UPDATE_ADDRESS', item)
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
