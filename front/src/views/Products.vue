<template>
    <section class="plr-20">
        <confirm ref="conf" @clicked="removeProduct"></confirm>
        <p class="title">{{$t("titles.products")}}</p>
        <div class="buttons pt-2 is-hidden-tablet">
            <b-button @click="confirmRemove()"
                      type="is-light"
                      class="is-small"
                      :disabled="!visible"
            >
                {{$t("buttons.remove")}}
            </b-button>
            <b-button tag="router-link"
                      :to="{ path: `/products/update` }"
                      type="is-link is-primary"
                      class="is-small"
                      v-on:click.native="updateProd"
                      :event="visible ? 'click' : ''"
                      :disabled="!visible"
            >
                {{$t("buttons.update")}}
            </b-button>

            <b-button tag="router-link"
                      :to="{ path: `/products/add` }"
                      type="is-link is-primary"
                      class="is-small"
            >
                {{$t("buttons.add")}}
            </b-button>

        </div>
        <section class="bg-img ">
            <template v-if="items" class="plr-20">
                <b-table
                        :data="isEmpty ? [] : items"
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
            </template>

            <template v-else>
                <td slot="empty" colspan="2">
                    <div class="content has-text-grey has-text-centered">
                        <p>
                            <b-icon
                                    icon="emoticon-sad"
                                    size="is-large">
                            </b-icon>
                        </p>
                        <p>Nothing here.</p>
                    </div>
                </td>
            </template>

            <div class="buttons pt-2">
                <b-button @click="confirmRemove()"
                          type="is-light"
                          size="is-small"
                          :disabled="!visible"
                >
                    {{$t("buttons.remove")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/products/update` }"
                          type="is-link is-primary"
                          class="is-small"
                          v-on:click.native="updateProd"
                          :event="visible ? 'click' : ''"
                          :disabled="!visible"
                >
                    {{$t("buttons.update")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/products/add` }"
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

    export default {
        name: 'products',
        components: {
            'confirm': Confirm,
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
                arr: ['nina', 'alina'],
                columns: [
                    {
                        field: 'ID',
                        label: 'ID',
                        width: '90',
                        numeric: true,
                        centered: true,
                        searchable: true,
                    },
                    {
                        field: 'VendorCode',
                        label: this.$t("products.table.vendorCode"),
                        centered: true,
                        searchable: true,
                    },

                    {
                        field: 'ProductName',
                        label: this.$t("products.table.productName"),
                        centered: true,
                        searchaObjble: true,
                    },
                    {
                        field: 'CategoryName',
                        label: this.$t("products.table.categoryName"),
                        centered: true,
                        searchable: true,
                    },
                ]
            }
        },
        computed: {
            ...mapGetters("products", [
                'items',
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
            },
            columns: function () {
                console.log('hello')
            }
        },
        mounted(){
            this.$store.dispatch('categories/getCategories');
            this.$store.dispatch('products/getProducts');
        },
        methods:{
            async removeProduct(){
                let count = null;
                for (let i = 0; i < this.checkedRows.length; i++) {
                    const response = await this.$store.dispatch('products/removeProduct', this.checkedRows[i].ID);
                    if(response === 'OK'){
                        console.log(response);
                        count++;
                    }
                }
                if(this.checkedRows.length === count){
                    Toast.open({
                        message: this.$t("products.toast.remove"),
                        type: 'is-danger'
                    });
                }
                this.visible = false;
            },
            async confirmRemove(){
                let names = [];
                try {
                    for (let i = 0; i < this.checkedRows.length; i++) {
                        names.push(this.checkedRows[i].ProductName)
                    }
                } catch (error) {
                    console.log('error')
                } finally {
                    this.$refs.conf.confirm(names)
                }
            },
            updateProd(){
                if(this.checkedRows.length > 0){
                    console.log('update');
                    this.$store.dispatch('products/setUpdateProducts', this.checkedRows);
                }
            },
        },
    }
</script>

<style lang="scss">
    @import "../styles/_variables.scss";
    @import "../../node_modules/bulma/sass/utilities/_all.sass";

    table td:nth-child(1) {
        background-color: $violetpale;
    }

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
