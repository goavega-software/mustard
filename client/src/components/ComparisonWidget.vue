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
  @Prop() private threshold?: number;
  @Prop() private eventId?: string;

  get model(): comparisonModel {
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
        stacked: true,
        toolbar: false
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
      delta = (newValue - oldValue) / oldValue,
      threshold = this.threshold || 1;
    return {
      center: true,
      red: delta > threshold,
      green: delta < threshold
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
