<template>
    <section>
        <button class="button is-primary is-medium"
                @click="isComponentModalActive = true">
            Remove
        </button>

        <b-modal :active.sync="isComponentModalActive"
                 has-modal-card
                 trap-focus
                 aria-role="dialog"
                 aria-modal>
            <modal-form v-bind="formProps"></modal-form>
        </b-modal>
    </section>
</template>

<script>
    const ModalForm = {
        props: ['email', 'password', 'checkbox'],
        template: `
                <div class="modal-card" style="width: auto">
                    <header class="modal-card-head">
                        <p class="modal-card-title">Are you sure you want to delete</p>
                    </header>

                    <section class="modal-card-body">
                        {{checkbox}}
                        <b-checkbox native-value="yes" v-model="checkbox">Delete all</b-checkbox>
                    </section>

                    <footer class="modal-card-foot">
                        <button class="button" type="button" @click="$parent.close()">Cancel</button>
                        <button class="button is-primary" type="button" @click="$parent.hello()">Delete</button>
                    </footer>
                </div>
        `
    }

    export default {
        components: {
            ModalForm
        },
        data() {
            return {
                isComponentModalActive: false,
                formProps: {
                    email: 'evan@you.com',
                    password: 'testing',
                    checkbox: [],
                }
            }
        },
        methods:{
            hello(){
                console.log('hekki');
            }
        }
    }
</script>
