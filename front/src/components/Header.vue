<template>
    <b-navbar>
<!--        <a href="#" @click="setLocale('en')"><flag iso="us"></flag></a>-->
        <template slot="start">

            <b-navbar-item tag="router-link" to="/products">
                {{$t("header.nav.products")}}
            </b-navbar-item>

            <b-navbar-item tag="router-link" to="/suppliers">
                {{$t("header.nav.suppliers")}}
            </b-navbar-item>

            <b-navbar-item tag="router-link" to="/cashbox">
                {{$t("header.nav.cashBox")}}
            </b-navbar-item>

            <b-navbar-dropdown :label="add">
                <b-navbar-item tag="router-link" to="/categories">
                    {{$t("header.nav.categories")}}
                </b-navbar-item>

            </b-navbar-dropdown>

        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <slot></slot>
            </b-navbar-item>

            <b-navbar-item tag="div">
                <div class="buttons">
                    <a class="button is-primary">
                        <strong>Sign up</strong>
                    </a>
                    <a class="button is-light">
                        Log in
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>

<script>
    export default {
        name: 'Header',
        data(){
            return{
                id: null,
                add: this.$t("header.nav.add")
            }
        },
        methods:{
            searchById () {
                this.$router.push({ path: '/players/profile/'+this.id });
                this.id = null;
            }
        }
    }
</script>


<style lang="scss">
    @import "../styles/_variables.scss";
    @import "../../node_modules/bulma/sass/utilities/_all.sass";

    .router-link-exact-active{
        background-color: $violetdarkpale !important;
    }

    .navbar-dropdown {
        background-color: $darkgrey !important;
    }

    @import "../../node_modules/bulma";
    @import "../../node_modules/buefy/src/scss/buefy.scss";
</style>
