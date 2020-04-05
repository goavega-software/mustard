<template>
  <div v-bind:class="{ ...classes }">
    <div>
      <h3>Covid-19 India</h3>
      <apexchart
        width="100%"
        type="bar"
        height="50%"
        :options="chartOptions"
        :series="series"
      ></apexchart>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
type comparisonModel = {
  newVal?: number;
  oldVal?: number;
  delta?: number;
};
@Component
export default class ComparisonWidget extends Vue {
  @Prop() private new?: number;
  @Prop() private old?: number;
  private threshold = 0.1;
  @Prop() private eventId?: string;
  private model: comparisonModel;
  private classes;
  private series;
  private chartOptions;

  get model() {
    return (
      this.$store.state[this.eventId || "undefined"] || {
        newValue: this.new,
        oldValue: this.old,
        delta: (this.new || 0) - (this.old || 0)
      }
    );
  }

  get chartOptions() {
    return {
      chart: {
        type: "bar",
        height: 350,
        stacked: true
      },
      plotOptions: {
        bar: {
          horizontal: true
        }
      },
      stroke: {
        width: 1,
        colors: ["#fff"]
      },
      xaxis: {
        categories: [],
        labels: {
          show: false
        }
      },
      fill: {
        opacity: 1
      },
      legend: {
        position: "top",
        horizontalAlign: "left",
        offsetX: 40
      },
      yaxis: {
        show: false
      }
    };
  }

  get series() {
    return [
      {
        name: "yesterday",
        data: [this.model.oldVal || 0]
      },
      {
        name: "today",
        data: [this.model.newVal || 0]
      }
    ];
  }

  get classes() {
    const newValue = this.model.newVal || 1,
      oldValue = this.model.oldVal || 1,
      delta = (newValue - oldValue) / oldValue;
    return {
      center: true,
      red: delta > this.threshold,
      green: delta < this.threshold
    };
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.red {
  background: #ef2d56;
}
.blue {
  background: #1446a0;
}
.green {
  background: #2d936c;
}
</style>
