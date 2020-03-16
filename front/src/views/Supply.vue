<template>
    <section class="p-3">
        <confirm ref="conf" @clicked="removeSupply"></confirm>
        <p class="title is-hidden-tablet">{{$t("titles.warehouse")}}</p>

        <div class="buttons pt-2">
            <b-button @click="confirmRemove()"
                      type="is-light"
                      class="is-small is-hidden-tablet"
                      :disabled="!visible"
            >
                {{$t("buttons.remove")}}
            </b-button>
            <b-button tag="router-link"
                      :to="{ path: `/supply/update` }"
                      type="is-link is-primary"
                      class="is-small is-hidden-tablet"
                      v-on:click.native="updateSupply"
                      :event="visible ? 'click' : ''"
                      :disabled="!visible"
            >
                {{$t("buttons.update")}}
            </b-button>

            <b-button tag="router-link"
                      :to="{ path: `/supply/add` }"
                      type="is-link is-primary"
                      class="is-small is-hidden-tablet"
            >
                {{$t("buttons.add")}}
            </b-button>
        </div>
        <section class="bg-img">
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
            <div class="buttons pt-2">
                <b-button @click="confirmRemove()"
                          type="is-light"
                          size="is-small"
                          :disabled="!visible"
                >
                    {{$t("buttons.remove")}}
                </b-button>
                <b-button tag="router-link"
                          :to="{ path: `/supply/update` }"
                          type="is-link is-primary"
                          class="is-small"
                          v-on:click.native="updateSupply"
                          :event="visible ? 'click' : ''"
                          :disabled="!visible"
                >
                    {{$t("buttons.update")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/supply/decommissioned` }"
                          type="is-link is-primary"
                          class="is-small"
                          v-on:click.native="decommissioned"
                          :event="visible ? 'click' : ''"
                          :disabled="!visible"
                >
                    {{$t("buttons.writeoff")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/supply/add` }"
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
                        field: 'DateCreate',
                        label: this.$t("warehouse.table.dateCreate"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'CompanyName',
                        label: this.$t("warehouse.table.companyName"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'ProductName',
                        label: this.$t("warehouse.table.productName"),
                        searchable: true,
                        centered: true
                    },
                    {
                        field: 'BuyPrice',
                        label: this.$t("warehouse.table.buyPrice"),
                        centered: true
                    },
                    {
                        field: 'SellPrice',
                        label: this.$t("warehouse.table.sellPrice"),
                        centered: true
                    },
                    {
                        field: 'Surcharge',
                        label: this.$t("warehouse.table.surcharge"),
                        centered: true

                    },
                    {
                        field: 'SupAmount',
                        label: this.$t("warehouse.table.supAmount"),
                        centered: true
                    },
                    {
                        field: 'Amount',
                        label: this.$t("warehouse.table.amount"),
                        centered: true
                    },
                    {
                        field: 'UnitName',
                        label: this.$t("warehouse.table.unit"),
                        centered: true
                    },
                    {
                        field: 'Comment',
                        label: this.$t("warehouse.table.comment"),
                        centered: true
                    },
                ]
            }
        },
        mounted(){
            this.$store.dispatch('supply/getSupply');
        },
        computed: {
            ...mapGetters("supply", [
                'items',
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

        methods:{
            async removeSupply(){
                let count = null;
                for (let i = 0; i < this.checkedRows.length; i++) {
                    const response = await this.$store.dispatch('supply/removeSupply', this.checkedRows[i].ID);
                    if(response === 'OK'){
                        console.log(response);
                        count++;
                    }
                }
                if(this.checkedRows.length === count){
                    Toast.open({
                        message: this.$t("supply.toast.remove"),
                        type: 'is-danger'
                    });
                }
                this.visible = false;
            },
            async confirmRemove(){
                let names = [];
                try {
                    for (let i = 0; i < this.checkedRows.length; i++) {
                        names.push(this.checkedRows[i].CompanyName + " - " + this.checkedRows[i].ProductName)
                    }
                } catch (error) {
                    console.log('error')
                } finally {
                    this.$refs.conf.confirm(names)
                }
            },
            updateSupply(){
                if(this.checkedRows.length > 0){
                    console.log('update');
                    this.$store.dispatch('supply/setUpdateSupply', this.checkedRows);
                }
            },
            decommissioned(){
                console.log('click');
                if(this.checkedRows.length > 0){
                    this.$store.dispatch('supply/setDecommissioned', this.checkedRows);
                }
            },
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
