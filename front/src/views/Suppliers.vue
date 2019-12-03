<template>
    <section class="p-3">
        <confirm ref="conf" @clicked="removeSupplier"></confirm>
        <p class="title is-hidden-tablet">{{$t("titles.suppliers")}}</p>
        <div class="buttons pt-2 is-hidden-tablet">
            <b-button @click="confirmRemove()"
                      type="is-light"
                      class="is-small"
                      :disabled="!visible"
            >
                {{$t("buttons.remove")}}
            </b-button>
            <b-button tag="router-link"
                      :to="{ path: `/suppliers/update` }"
                      type="is-link is-primary"
                      class="is-small"
                      v-on:click.native="updateSuppl"
                      :event="visible ? 'click' : ''"
                      :disabled="!visible"
            >
                {{$t("buttons.update")}}
            </b-button>

            <b-button tag="router-link"
                      :to="{ path: `/suppliers/add` }"
                      type="is-link is-primary"
                      class="is-small"
            >
                {{$t("buttons.add")}}
            </b-button>
        </div>
        <section class="bg-img">
            <b-table
                    :data="isEmpty ? [] : suppliers"
                    :columns="columns"
                    :checked-rows.sync="checkedRows"
                    checkable
                    :checkbox-position="checkboxPosition"
                    :striped="isStriped"
                    :narrowed="isNarrowed"
                    :hoverable="isHoverable"
                    :loading="isLoading"
                    :mobile-cards="hasMobileCards"
            >
            </b-table>
            <div class="buttons pt-2">
                <b-button @click="confirmRemove()"
                          type="is-light"
                          size="is-small"
                          :disabled="!visible"
                >
                    {{$t("buttons.remove")}}
                </b-button>
                <b-button tag="router-link"
                          :to="{ path: `/suppliers/update` }"
                          type="is-link is-primary"
                          class="is-small"
                          v-on:click.native="updateSuppl"
                          :event="visible ? 'click' : ''"
                          :disabled="!visible"
                >
                    {{$t("buttons.update")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/suppliers/add` }"
                          type="is-link is-primary"
                          class="is-small"
                >
                    {{$t("buttons.add")}}
                </b-button>
            </div>

        </section>
    </section>
</template>

<script>
    import { mapGetters } from 'vuex';
    import Confirm from "../components/Confirm";
    import {ToastProgrammatic as Toast} from "buefy";

    // import Modal from "../components/ModalMultipleRemove";

    export default {
        name: 'products',
        components: {
            'confirm': Confirm,
            // 'modal': Modal,
        },
        data() {
            return {
                visible: false,
                //------table--------
                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
                checkedRows: [],
                checkboxPosition: 'right',
                columns: [
                    {
                        field: 'ID',
                        label: 'ID',
                        width: '90',
                        numeric: true,
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'CompanyName',
                        label: this.$t("suppliers.table.companyName"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'ContactName',
                        label: this.$t("suppliers.table.contactName"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'ContactTitle',
                        label: this.$t("suppliers.table.contactTitle"),
                        centered: true
                    },
                    {
                        field: 'Address',
                        label: this.$t("suppliers.table.address"),
                        centered: true
                    },
                    {
                        field: 'City',
                        label: this.$t("suppliers.table.city"),
                        centered: true
                    },
                    {
                        field: 'Phone',
                        label: this.$t("suppliers.table.phone"),
                        centered: true

                    },
                    {
                        field: 'Email',
                        label: this.$t("suppliers.table.email"),
                        centered: true
                    },
                ]
            }
        },
        computed: {
            ...mapGetters("suppliers", [
                'suppliers',
            ]),
            ...mapGetters("categories", [
                'categories',
            ]),
        },
        watch: {
            checkedRows: function() {
                if(this.checkedRows.length > 0){
                    this.visible = true;
                }else {
                    this.visible = false;
                }
            }
        },
        mounted(){
            this.$store.dispatch('suppliers/getSuppliers');
            this.$store.dispatch('products/getProducts');
        },
        methods:{
            async removeSupplier(){
                let count = null;
                for (let i = 0; i < this.checkedRows.length; i++) {
                    const response = await this.$store.dispatch('suppliers/removeSupplier', this.checkedRows[i].ID);
                    if(response === 'OK'){
                        console.log(response);
                        count++;
                    }
                }
                if(this.checkedRows.length === count){
                    Toast.open({
                        message: this.$t("suppliers.toast.remove"),
                        type: 'is-danger'
                    });
                }
                this.visible = false;


                // let id = 9;
                // for (let i = 0; i < this.checkedRows.length; i++) {
                //     this.$store.dispatch('suppliers/removeSupplier', this.checkedRows[i].ID);
                // }
            },
            async confirmRemove(){
                let names = [];
                try {
                    for (let i = 0; i < this.checkedRows.length; i++) {
                        names.push(this.checkedRows[i].CompanyName)
                    }
                } catch (error) {
                    console.log('error')
                } finally {
                    this.$refs.conf.confirm(names)
                }
            },
            updateSuppl(){
                if(this.checkedRows.length > 0){
                    console.log('update');
                    this.$store.dispatch('suppliers/setUpdateSuppliers', this.checkedRows);
                }
            },
            // afterEnter () {
            //     console.log('after enter in transition object') // works
            // },
        }
    }
</script>

<style lang="scss">
    @import "../styles/_variables.scss";
    @import "../../node_modules/bulma/sass/utilities/_all.sass";

    table td:nth-child(1) {
        background-color: $violetpale;
    }
    /*.table thead th {*/
    /*    background-color: $pink;*/
    /*}*/

    $table-row-hover-background-color: $hover;
    $narbar-hover-background-color: $hover;

    .b-table .table {
        background-color: $table;
    };
    /*// Import Bulma and Buefy styles*/
    $colors: (
    "light": ($pinkstrong, white),
    "primary": ($violet, white),
    "info": ($peachy, $white),
    "success": ($green, $success-invert),
    "warning": ($darkyellow, $white),
    "danger": ($pinkstrong, white),
    );
    .title{
        color: $violet !important;
    }
    .navbar{
        background-color: $darkgrey !important;
    }
    .navbar-item, .navbar-link{
        color: $green !important;
        font-size: 20px;
    }
    $navbar-item-hover-background-color:  $violetpale;

    .navbar-burger{
        color: $green !important;
    }
    $pagination-margin: 0.25rem !default;

    $pagination-color: $green !default;
    $pagination-border-color: $green !default;
    $pagination-active-color:  $peachy !default;
    $pagination-current-background-color: $violet !default;
    .navbar-menu{
        background-color: $darkgrey !important;
    }
    @import "../../node_modules/bulma";
    @import "../../node_modules/buefy/src/scss/buefy.scss";
</style>
