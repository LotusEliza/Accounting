<template>
    <div class="columns p-10">
        <div class="column is-half p-10" >
            <div class="tile is-child box" style="text-align: center">
<!--_______________________________BUTTONS_______________________________________-->
                <div class="buttons pt-2">
                    <b-button @click="selectedBtn(index)"
                              v-for="(btn,index) in btns"
                              v-bind:key="index"
                              :class="{selBtn: btn.id === selBtn}"
                              type="is-primary"
                              class="is-small"
                    >
                        {{btn.label}}
                    </b-button>
                </div>

                <div class="content p-20" style="text-align: center">
<!--_______________________________MONTH_PICKER_______________________________________-->
                    <month-picker
                            @change="showDate"
                            @change-year="(v) => year = v"
                            @clear="showClearText"
                            :lang="lang"
                            :clearable="isClearable"
                            :editable-year="isEditableYear"
                            :variant="selectedVariant"
                            :show-year="showYear"
                            :default-month = "default_month"
                    >
                    </month-picker>
                </div>
                <hr>
<!--_______________________________ TABLE "MONTH TOTAL INFO" _______________________________________-->
                <b-table
                        :data="isEmpty ? [] : monthInfoTotal"
                        :striped="isStriped"
                        :narrowed="isNarrowed"
                        :mobile-cards="hasMobileCards"
                >
                    <template slot-scope="props">
                        <b-table-column field="Tg" :label="tableLabels[0].credit" width="40" numeric>
                            {{ props.row.credit }}
                        </b-table-column>

                        <b-table-column field="Name" :label="tableLabels[0].debit">
                            {{ props.row.debit }}
                        </b-table-column>

                        <b-table-column field="Name" :label="tableLabels[0].profit">
                            <p v-bind:class="{ sign: sign === 'negative'}">{{ props.row.profit }}</p>
                        </b-table-column>
                    </template>

                    <template slot="empty">
                        <section class="section">
                            <div class="content has-text-grey has-text-centered">
                                <p>
                                    <b-icon
                                            icon="emoticon-sad"
                                            size="is-large">
                                    </b-icon>
                                </p>
                                <p>Nothing here.</p>
                            </div>
                        </section>
                    </template>
                </b-table>
            </div>

        </div>
        <div class="column is-half p-10" >
            <div class="tile is-child box" style="text-align: center">
<!--_______________________________ "DEBIT/CREDIT" _______________________________________-->
                <div v-if="visible === 'debit/credit'">
                    <p class="title">Кредит/Дебит</p>
                    <b-table
                            :data="isEmpty ? [] : monthCreditDebit"
                            :columns="columns"
                            :striped="isStriped"
                            :narrowed="isNarrowed"
                            :mobile-cards="hasMobileCards"
                    >
                    </b-table>
                    <br>
                    <p class="total subtitle"> {{$t("report.creditTotal")}}  {{monthCreditTotal}}</p>
                    <p class="total subtitle"> {{$t("report.debitTotal")}}  {{monthDebitTotal}}</p>
                </div>
<!--_______________________________ "UTILITY BILLS" _______________________________________-->
                <div v-if="visible === 'utility bills'">
                    <p class="title">Коммунальные платежи</p>
                    <b-table
                            :data="isEmpty ? [] : monthBillTable"
                            :columns="columnsBill"
                            :striped="isStriped"
                            :narrowed="isNarrowed"
                            :mobile-cards="hasMobileCards"
                    >
                    </b-table>
                    <br>
                    <p class="total subtitle"> {{$t("cashBox.labels.sum2")}}  {{totalForMonth}}</p>
                </div>
<!--_______________________________ "DECOMMISSIONED" _______________________________________-->
                <div v-if="visible === 'decommissioned'">
                    <p class="title">Списанный товар</p>
                    <b-table
                            :data="isEmpty ? [] : decommissioned_month"
                            :columns="columns_decom"
                            :striped="isStriped"
                            :narrowed="isNarrowed"
                            :mobile-cards="hasMobileCards"
                    >
                    </b-table>
                    <br>
                    <p class="total subtitle"> {{$t("cashBox.labels.sum2")}}  {{totalForMonthDecom}}</p>
                </div>
                <hr>
<!--_______________________________ CHART "MONTH" _______________________________________-->
                <div  v-if="visible === 'chart month'">
                    <p class="title">График за месяц</p>
                    <chart :labels="labels" :datasets="datasets" :key="componentKey"/>
                </div>

<!--_______________________________ CHART "YEAR" _______________________________________-->
                <div v-if="visible === 'chart year'">
                    <p class="title">График за год</p>
                    <chart :labels="labelsYear" :datasets="datasetsYear" :key="componentKey"/>
                </div>

            </div>
        </div>
     </div>
</template>

<script>
    import { MonthPicker } from 'vue-month-picker'
    import {mapGetters} from "vuex";
    import chart from "../components/Chart";
    import moment from 'moment';

    export default {
        name: 'report',
        components: {
            MonthPicker,
            chart
        },
        data() {
            return {
                //------buttons--------
                selBtn: 0,
                visible: "debit/credit",
                btns: [
                    {id: 0, name: "debit/credit",  label: this.$t("report.buttons.debit/credit")},
                    {id: 1, name: "decommissioned", label: this.$t("report.buttons.decommissioned")},
                    {id: 2, name: "chart month", label: this.$t("report.buttons.chartMonth")},
                    {id: 3, name: "chart year", label: this.$t("report.buttons.chartYear")},
                    {id: 4, name: "utility bills", label: this.$t("report.buttons.utilityBills")},
                ],


                default_month: new Date().getMonth() +1,
                lang:  this.$i18n.locale,

                //------table--------
                tableLabels: [{
                    credit: this.$t("report.credit"),
                    debit: this.$t("report.debit"),
                    profit: this.$t("report.profit"),
                }],

                hasMobileCards: true,
                isEmpty: false,
                narrowed: true,
                isStriped: false,
                isNarrowed: false,
                columnsBill: [
                    {
                        field: 'UnicID',
                        label: "ID",
                        centered: true,
                    },
                    {
                        field: 'BillName',
                        label: this.$t("report.table.bill"),
                        centered: true,
                    },
                    {
                        field: 'Amount',
                        label: this.$t("products.table.amount"),
                        centered: true,
                    },
                    {
                        field: 'Date',
                        label: this.$t("report.table.date"),
                        centered: true,
                    },
                ],
                columns: [
                    {
                        field: 'UnicID',
                        label: "ID",
                        centered: true,
                    },
                    {
                        field: 'Comment',
                        label: this.$t("warehouse.table.comment"),
                        centered: true
                    },
                    {
                        field: 'Credit',
                        label: this.$t("report.credit"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'Debit',
                        label: this.$t("report.debit"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'Date',
                        label: this.$t("warehouse.table.dateCreate"),
                        searchable: true,
                        centered: true,
                    },
                ],
                columns_decom: [
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
                        centered: true
                    },

                    {
                        field: 'SellPrice',
                        label: this.$t("cashBox.labels.price2"),
                        centered: true
                    },
                    {
                        field: 'Comment',
                        label: this.$t("warehouse.table.comment"),
                        searchable: true,
                        centered: true,
                    },
                    {
                        field: 'DateCreate',
                        label: this.$t("warehouse.table.dateCreate"),
                        searchable: true,
                        centered: true,
                    },

                ],
                debit: null,
                year: 0,
                //------MonthPicker--------
                clearEmittedText: null,
                isMonthPickerVisible: false,
                isClearable: true,
                isEditableYear: false,
                showYear: true,
                selectedDate: null,
                variants: ['default', 'dark'],
                selectedVariant: 'default',
                date: {
                    from: null,
                    to: null,
                    month: null,
                    year: null
                },
                //------CHARTS--------
                labels:[],
                labelsYear:[
                    'January',
                    'February',
                    'March',
                    'April',
                    'May',
                    'June',
                    'July',
                    'August',
                    'September',
                    'October',
                    'November',
                    'December'
                ],
                datasets:[
                    {data: []},//credit
                    {data: []}
                ],
                datasetsYear:[
                    {data: []},//credit
                    {data: []},
                ],
                componentKey: 0,
                currentYear: '',
                sign: '',
                // totalForTrans: null
            }
        },

        computed: {
            ...mapGetters("credit", [
                'monthInfoTotal', 'monthCredit', 'creditYear', 'monthCreditDebit', 'monthCreditTotal'
            ]),
            ...mapGetters("debit", [
                'monthDebit', 'debitYear', 'monthDebitTotal'
            ]),
            ...mapGetters("supply", [
                'decommissioned_month', 'totalForMonthDecom'
            ]),
            ...mapGetters("new_utility_bills", [
                'monthBills', 'monthBillsAmount', 'monthBillTable', 'totalForMonth'
            ]),
        },
        watch:{
            currentYear(){
                //building a new chart when the year is changed
                this.datasetsYear[0].data = [];
                this.getLastFirstInMonth();
            },

            monthInfoTotal(val){
                if(val.length){
                    let result = Math.sign(val[0].debit - val[0].credit);
                    if(result === -1){
                        this.sign = 'negative';
                    }else{
                        this.sign = 'positive';
                    }
                }else {
                    this.sign = 'positive';
                }
            }
        },
        mounted() {
            this.showDate(this.date);
            // this.totalTrans();
        },
        methods: {
            /**
             * month-picker @clear
             */
            showClearText () {
                this.clearEmittedText = 'emitted';
                window.setTimeout(() => {
                    this.clearEmittedText = null
                }, 1000)
            },
            /**
             * generates an array of dates
             * @param start, end | string
             * @return arr | array [Sun Dec 01 2019 00:00:00 GMT+0200]
             */
            getDateArray(start, end){
                let  arr = [];
                let dt = new Date(start);

                while (dt < end) {
                    arr.push(new Date(dt));
                    dt.setDate(dt.getDate() + 1);
                }
                return arr;
            },
            /**
             * generates an array of days in month
             * @param year, monthIndex | string
             * @return array | array [1,2,3,...31]
             */
            arrayOfDays(year, monthIndex){
                // let days = this.daysInMonth(year, monthIndex); //31
                let days = new Date(year, monthIndex, 0).getDate(); //31
                let array = [];
                for (let i = 0; i < days; i++) {
                    array.push(i+1);
                }
                return array;
            },
            clearChart(){
                //clean chart before render
                this.datasets[0].data = [];
                this.datasets[1].data = [];
            },
            /**
             * Month-Picker @change
             * @param date | array [from, to, ...]
             */
            showDate (date) {

                this.date = date;
                this.$store.dispatch('new_utility_bills/getMonthBill', date);
                this.$store.dispatch('credit/getMonthCredit', date);
                this.$store.dispatch('debit/getMonthDebit', date);
                this.$store.dispatch('supply/decommissionedMonthGet', date);

                //to trigger watcher currentYear only when year is really changed
                if(this.currentYear !== date.year){
                    this.currentYear = date.year;
                }
                this.clearChart();
                let arrayOfDaysX = this.arrayOfDays(date.year, date.monthIndex); //[1,2,3,...]
                this.labels = arrayOfDaysX;
                let dateArray = this.getDateArray(date.from, date.to);//[Sun Dec 01 2019 00:00:00 GMT+0200 (Eastern European Standard Time),...]
                setTimeout(()=>{
                    this.creatingArray(dateArray, this.monthCredit);//Y coordinat in "Month credit" - datasets[0].data : [100,12,15,190...]
                    this.creatingArray(dateArray, this.monthDebit);//Y coordinat in "Month debit" - datasets[1].data : [100,12,15,190...]
                },300);
            },
            /**
             * Get all dates that have debit in format DD/MM/YYYY
             * @param array | array { 0:[Comment: buy, Date: 2020-01-09T15:03:59.806Z, Debit: 3.6, ID: 178], 1:[...]}
             * @return dates | array ["09/01/2020", "10/01/2020", "12/01/2020"...]
             */
            formatDateArray(array){
                let dates = [];
                for (let i = 0; i < array.length; i++) {
                    let date = moment(array[i].Date).format('DD/MM/YYYY');
                    dates.push(date);
                }
                return dates;
            },
            formatDate(date){
                return moment(date).format('DD/MM/YYYY')
            },
            /**
             * Get all dates that have debit in format DD/MM/YYYY
             * @param dateArray | array  [Wed Jan 01 2020 00:00:00 GMT+0200 (Eastern European Standard Time),...](31)
             * @param arrayOfItems | array { 0:[Comment: buy, Date: 2020-01-09T15:03:59.806Z, Debit: 3.6, ID: 178], 1:[...]}
             * @return dates | array ["09/01/2020", "10/01/2020", "12/01/2020"...]
             */
            creatingArray(dateArray, arrayOfItems){
                let dates = this.formatDateArray(arrayOfItems); //get all dates that have debit in format DD/MM/YYYY
                let unique_array = [...new Set(dates)];//select only unique dates from array
                let final_array = [];//[ 0: {date: Sun Dec 01 2019 00:00:00 GMT+0200 (Eastern European Standard Time), total: 0},...]

                for (let i = 0; i < dateArray.length; i++) {
                    for(let y = 0; y < unique_array.length; y++){
                        if(this.formatDate(dateArray[i]) == unique_array[y]){
                            let total = null;
                            for (let z = 0; z < arrayOfItems.length; z++) {
                                if(unique_array[y] == this.formatDate(arrayOfItems[z].Date)){
                                    if(arrayOfItems[z].Debit){
                                        total += arrayOfItems[z].Debit;
                                    }else {
                                        total += arrayOfItems[z].Credit;
                                    }
                                }
                            }
                            final_array.push({ date: dateArray[i], total: total });
                        }
                    }
                    if(final_array.length > 0){
                        if(dateArray[i] !== final_array[final_array.length - 1].date){
                            final_array.push({ date: dateArray[i], total: 0 });
                        }
                    }else {
                        final_array.push({ date: dateArray[i], total: 0 });
                    }
                }
                let result = this.mapItems(final_array);
                if(arrayOfItems.length && arrayOfItems[0].Debit){
                    this.datasets[1].data = result;
                }else if(arrayOfItems.length && arrayOfItems[0].Credit){
                    this.datasets[0].data = result;
                }
                this.forceRerender();
            },
            /**
             * Get all dates that have debit in format DD/MM/YYYY
             * @param array | array [0: {date: Wed Jan 01 2020 00:00:00 GMT+0200 (Eastern European Standard Time), total: 0}, 1:{...}] (31)
             * @return result | array [0, 0, 0, 32.20, 45.50,...]
             */
            mapItems(array){
                let result = array.map((item) => {
                    return  item.total;
                });
                return result;
            },
            /**
             * Change template view according to index of pressed button
             * @param index | number 2
             */
            selectedBtn(index){
                this.visible = this.btns[index].name;
                this.selBtn = index;
                if(index === 3){
                    this.datasetsYear[0].data = this.creditYear;
                    this.datasetsYear[1].data = this.debitYear;
                    this.forceRerender();
                }
                if(index === 2){
                    this.forceRerender()
                }
            },
            forceRerender() {
                this.componentKey += 1;
            },
            getLastFirstInMonth(){
                let day = 1;
                let arrayStartEndMonths = [];
                for (let i = 1; i <= 12; i++) {
                    let month = i;
                    let date = this.currentYear+'-'+month+'-'+day;
                    const startOfMonth = moment(date).startOf('month').format('YYYY-MM-DD hh:mm');
                    const endOfMonth   = moment(date).endOf('month').format('YYYY-MM-DD hh:mm');
                    arrayStartEndMonths.push({from: startOfMonth, to: endOfMonth});
                }
                this.$store.dispatch('credit/getYearCredit', arrayStartEndMonths).then(() =>  this.datasetsYear[0].data = this.creditYear);
                this.$store.dispatch('debit/getYearDebit', arrayStartEndMonths).then(() =>  this.datasetsYear[1].data = this.debitYear);
                this.forceRerender();
            }
        }
    }
</script>

<style>
    body {
        background-color: #FEFEFE;
    }
</style>

<style scoped>
    .month-picker__container {
        width: 100% !important;
    }
    .sign {
        color: red;
    }
    @import url('https://fonts.googleapis.com/css?family=Oxygen&display=swap');
    .content {
        max-width: 800px;
        margin: 0 auto;
        font-family: 'Oxygen', sans-serif;
    }
    .content h1 {
        margin-bottom: 2em;
        letter-spacing: 2px;
    }
    .github-corner:hover .octo-arm {
        animation: octocat-wave 560ms ease-in-out;
    }
    .form__label {
        display: block;
    }
    @keyframes octocat-wave {
        0% {
            transform: rotate(0deg);
        }

        20% {
            transform: rotate(-25deg);
        }

        40% {
            transform: rotate(10deg);
        }

        60% {
            transform: rotate(-25deg);
        }

        80% {
            transform: rotate(10deg);
        }

        100% {
            transform: rotate(0deg);
        }
    }

    @media (max-width: 500px) {
        .github-corner:hover .octo-arm {
            animation: none;
        }
        .github-corner .octo-arm {
            animation: octocat-wave 560ms ease-in-out;
        }
    }
    .p-10{
        padding-left: 10%;
        padding-top: 5%;
    }

    .selBtn{
        background-color: deeppink;
    }
    p.total{
        text-align: left;
    }
</style>
