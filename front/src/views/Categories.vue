<template>
    <section class="p-7">
        <confirm ref="conf" @clicked="removeCategory"></confirm>
        <p class="title">{{$t("titles.categories")}}</p>
        <div class="buttons pt-2 is-hidden-tablet">
            <b-button @click="confirmRemove()"
                      type="is-light"
                      :disabled="!visible"
                      class="is-small"
            >
                {{$t("buttons.remove")}}
            </b-button>
            <b-button tag="router-link"
                      :to="{ path: `/categories/update` }"
                      type="is-link is-primary"
                      v-on:click.native="updateCategories"
                      :event="visible ? 'click' : ''"
                      :disabled="!visible"
                      class="is-small"
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

        <section class="bg-img">

            <template v-if="categories">
                <b-table
                        :data="isEmpty ? [] : categories"
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
                          :to="{ path: `/categories/update` }"
                          type="is-link is-primary"
                          class="is-small"
                          v-on:click.native="updateCategories"
                          :event="visible ? 'click' : ''"
                          :disabled="!visible"
                >
                    {{$t("buttons.update")}}
                </b-button>

                <b-button tag="router-link"
                          :to="{ path: `/categories/add` }"
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
        name: 'categories',
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
                        field: 'CategoryName',
                        label: this.$t("categories.table.categoryName"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'CategoryDescription',
                        label: this.$t("categories.table.сategoryDescription"),
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
            async removeCategory(){
                let count = null;
                for (let i = 0; i < this.checkedRows.length; i++) {
                    const response = await this.$store.dispatch('categories/removeCategory', this.checkedRows[i].ID);
                    if(response === 'OK'){
                        console.log(response);
                        count++;
                    }
                }
                if(this.checkedRows.length === count){
                    Toast.open({
                        message: this.$t("categories.toast.remove"),
                        type: 'is-danger'
                    });
                }
                this.visible = false;
            },
            async confirmRemove(){
                let names = [];
                try {
                    for (let i = 0; i < this.checkedRows.length; i++) {
                        names.push(this.checkedRows[i].CategoryName)
                    }
                } catch (error) {
                    console.log('error')
                } finally {
                    this.$refs.conf.confirm(names)
                }
            },
            updateCategories(){
                if(this.checkedRows.length > 0){
                    console.log('update');
                    this.$store.dispatch('categories/setUpdateCategories', this.checkedRows);
                }
            },
        }
    }
</script>
