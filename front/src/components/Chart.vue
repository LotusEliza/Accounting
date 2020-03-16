<script>
    import { Line } from "vue-chartjs";

    export default {
        extends: Line,
        props: ['labels', 'datasets'],
        data() {
            return {
                gradient: null,
                gradient2: null
            };
        },
        mounted() {
            this.gradient = this.$refs.canvas
                .getContext("2d")
                .createLinearGradient(0, 0, 0, 450);
            this.gradient2 = this.$refs.canvas
                .getContext("2d")
                .createLinearGradient(0, 0, 0, 450);

            this.gradient.addColorStop(0, "rgba(255, 0,0, 0.5)");
            this.gradient.addColorStop(0.5, "rgba(255, 0, 0, 0.25)");
            this.gradient.addColorStop(1, "rgba(255, 0, 0, 0)");

            this.gradient2.addColorStop(0, "rgba(0, 231, 255, 0.9)");
            this.gradient2.addColorStop(0.5, "rgba(0, 231, 255, 0.25)");
            this.gradient2.addColorStop(1, "rgba(0, 231, 255, 0)");

            this.renderChart(
                {
                    labels: this.labels,
                    datasets: [
                        {
                            label: "Credit",
                            borderColor: "#FC2525",
                            pointBackgroundColor: "white",
                            borderWidth: 1,
                            pointBorderColor: "white",
                            backgroundColor: this.gradient,
                            data: this.datasets[0].data
                        },
                        {
                            label: "Debit",
                            borderColor: "#05CBE1",
                            pointBackgroundColor: "white",
                            pointBorderColor: "white",
                            borderWidth: 1,
                            backgroundColor: this.gradient2,
                            data: this.datasets[1].data
                            // data:[40, 39, 10, 40, 39, 80, 40]
                        }
                    ]
                },
                { responsive: true, maintainAspectRatio: false }
            );
        }
    };
</script>
