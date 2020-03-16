<template>
    <div >
        <div class="columns" >
            <div class="column is-2">
            </div>
            <div class="column box mt-3">
                <p class="title">Отчет за день</p>
                <div class="pl-2">
                    <Datepicker v-model="date" name="uniquename" format="dd MMM yyyy" :language="ru" :monday-first="true"></Datepicker>
                </div>
                <div v-if="date !== ''">
                    <div class="buttons pt-2">
                        <b-button @click="selectedBtn(index)"
                                  v-for="(btn,index) in btns"
                                  v-bind:key="index"
                                  :class="{selBtn: btn.id === selBtn}"
                                  type="is-primary"
                        >
                            {{btn.name}}
                        </b-button>
                    </div>

                    <chart  v-if="visible === 'chart day'" :labels="labels" :datasets="datasets" :key="componentKey"/>
                    <div v-if="visible === 'debit'" class="plr-15">
                        <b-table
                                :data="isEmpty ? [] : dayDebit"
                                :columns="columnsDebit"
                                :striped="isStriped"
                                :narrowed="isNarrowed"
                                :mobile-cards="hasMobileCards"
                        >
                        </b-table>
                        <br>
                        <p class="total subtitle"> {{$t("cashBox.labels.sum2")}} {{debitTotal}}</p>
                    </div>
                    <div v-if="visible === 'credit'" class="plr-15">
                        <b-table
                                :data="isEmpty ? [] : dayCreditTable"
                                :columns="columnsCredit"
                                :striped="isStriped"
                                :narrowed="isNarrowed"
                                :mobile-cards="hasMobileCards"
                        >
                        </b-table>
                    </div>
                    <div v-if="visible === 'decommissioned'" class="plr-15">
                        <b-table
                                :data="isEmpty ? [] : decommissioned_day"
                                :columns="columnsDecom"
                                :striped="isStriped"
                                :narrowed="isNarrowed"
                                :mobile-cards="hasMobileCards"
                        >
                        </b-table>
                    </div>
                </div>
            </div>
            <div class="column is-2">
            </div>
        </div>
    </div>
</template>

<script>
    import Datepicker from 'vuejs-datepicker';
    import moment from "moment";
    import {mapGetters} from "vuex";
    import chart from "../components/Chart";
    import {en, ru} from 'vuejs-datepicker/dist/locale'

    export default {
        name: 'dayReport',
        components: {
            Datepicker,
            chart,
        },
        data() {
            return {
                ru: ru,
                en: en,
                date: '',
                labels: ["00:00", "01:00", "02:00", "03:00", "04:00", "05:00", "06:00", "07:00",
                         "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00",
                         "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00", "00:00"],
                datasets:[
                    {data: []},
                    {data: []}
                ],
                componentKey: 0,

                //--------buttons----------
                selBtn: 0,
                btns: [
                    {id: 0, name: "debit"},
                    {id: 1, name: "credit"},
                    {id: 2, name: "chart day"},
                    {id: 3, name: "decommissioned"},
                ],
                visible: "debit",

                //------table--------
                hasMobileCards: true,
                isEmpty: false,
                narrowed: true,
                isStriped: false,
                isNarrowed: false,
                columnsDebit: [
                        {
                            field: 'UnicID',
                            label: "ID",
                            centered: true,
                        },
                        {
                            field: 'Debit',
                            label: this.$t("report.debit"),
                            searchable: true,
                            centered: true,
                        },
                        {
                            field: 'DateTable',
                            label: this.$t("warehouse.table.dateCreate"),
                            searchable: true,
                            centered: true,
                        },
                        {
                            field: 'Comment',
                            label: this.$t("warehouse.table.comment"),
                            searchable: true,
                            centered: true,
                        },
                    ],
                columnsCredit: [
                    {
                        field: 'UnicID',
                        label: "ID",
                        centered: true,
                    },
                    {
                        field: 'Credit',
                        label: this.$t("report.credit"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'Date',
                        label: this.$t("warehouse.table.dateCreate"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'Comment',
                        label: this.$t("warehouse.table.comment"),
                        searchable: true,
                        centered: true,
                    },
                ],
                columnsDecom: [
                    {
                        field: 'UnicID',
                        label: "ID",
                        centered: true,
                    },
                    {
                        field: 'ProductName',
                        label: this.$t("products.table.productName"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'AmountDecom',
                        label: this.$t("products.table.amount"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'Comment',
                        label: this.$t("warehouse.table.comment"),
                        searchable: true,
                        centered: true,
                    },
                ]
            }
        },
        computed: {
            ...mapGetters("debit", [
                'dayDebit', 'dayDebitTable', 'debitTotal'
            ]),
            ...mapGetters("credit", [
                'dayCredit', 'dayCreditTable'
            ]),
            ...mapGetters("supply", [
                'decommissioned_day'
            ]),
        },
        mounted(){
        },
        filters: {
            moment: function (date) {
                return moment(date).format('HH:mm');
            }
        },
        watch: {
            date(val) {
                this.datasets[1].data = [];
                this.forceRerender();

                let date = moment(val).format('DD/MM/YYYY');
                let from = moment(date, 'DD/MM/YYYY', true).format();
                let to = moment(from).add('days', 1).format();
                let array = [{from: from, to: to}];
                this.getDayDebit(array);
                this.getDayCredit(array);
                this.$store.dispatch('supply/decommissionedDayGet', array);
            },
            dayDebit(dayTime){
                if(dayTime.length){
                let arrayTimeDebit = [];
                let arrayTime = [];
                for (let i = 0; i < dayTime.length; i++) {
                    let time = moment(moment.utc(dayTime[i].Date)).format('HH.mm.ss');
                    arrayTimeDebit.push({time: time, debit: dayTime[i].Debit})
                    arrayTime.push(time)
                }
                this.getDebitHour(arrayTime, arrayTimeDebit)
            }else {
                    console.log('no debit')
                }
            },
        },
        methods: {
            forceRerender() {
                this.componentKey += 1;
            },
            formArrayDebit(arrayTimes, arrayTimeDebit){
                let debit = 0;
                for(let j = 0; j < arrayTimes.length; j++){
                    for(let k = 0; k < arrayTimeDebit.length; k++){
                        if(arrayTimes[j] === arrayTimeDebit[k].time){
                            debit += arrayTimeDebit[k].debit;
                        }
                    }
                }
                if(debit > 0){
                    return debit;
                }else {
                    return 0;
                }
            },
            getDebitHour(arrayTime, arrayTimeDebit){
                let finalArrayTime = [];
                let minStart = "23:00";
                let min, max;
                for (let i = 0; i < 25; i++) {
                    if(i===0){
                        max = moment(minStart, "HH:mm").add(1, 'hours').format('HH:mm');
                        let result = arrayTime.filter(a => a > minStart && a < max);
                        let debit = this.formArrayDebit(result, arrayTimeDebit);
                        finalArrayTime.push(debit);
                    }else {
                        min = moment(minStart, "HH:mm").add(1, 'hours').format('HH:mm');
                        max = moment(minStart, "HH:mm").add(2, 'hours').format('HH:mm');
                        let result = arrayTime.filter(a => a > min && a < max);
                        let debit = this.formArrayDebit(result, arrayTimeDebit);
                        finalArrayTime.push(debit);
                        minStart = min;
                    }
                }
                this.datasets[1].data = finalArrayTime;
                this.forceRerender();
            },
            getDayDebit(array){
                this.$store.dispatch('debit/getDayDebit', array);
            },
            getDayCredit(array){
                this.$store.dispatch('credit/getDayCredit', array);
            },
            selectedBtn(index){
                this.toggle(this.btns[index].name);
                this.selBtn = index;
            },
            toggle(data){
                this.visible = data;
            },
        }
    }
</script>

<style scoped>
    .selBtn{
        background-color: deeppink;
    }
    p.title {
        text-align: center;
    }
    .pl-2{
        padding-left: 42%;
    }
    .pt-2{
        padding-left: 27%;
    }
    .mt-3{
        margin-top: 3%;
    }
</style>

