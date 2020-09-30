<template>
  <div class="clock">
    <div class="clock1" v-if="clockOneModel">
      <h4>{{ clockOneModel.date }}</h4>
      <h3>{{ clockOneModel.time }}</h3>
      <p class="muted">{{ clockOneTz }}</p>
    </div>
    <div class="clock2">
      <h3>{{ model.date }}</h3>
      <h1>{{ model.time }}</h1>
      <p class="muted text-center">{{ numberTrivia.trivia }}</p>
    </div>
    <div class="clock3" v-if="clockThreeModel">
      <h4>{{ clockThreeModel.date }}</h4>
      <h3>{{ clockThreeModel.time }}</h3>
      <p class="muted">{{ clockThreeTz }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import { getStoreItem, State } from "../store";
type TimeModel = {
  date: string;
  time: string;
};
@Component
export default class Clock extends Vue {
  private model: TimeModel = {
    date: "",
    time: ""
  };
  private clockOneModel?: TimeModel= {
    date: "",
    time: ""
  };
  private clockThreeModel?: TimeModel= {
    date: "",
    time: ""
  };
  private timerID?: number;
  @Prop() private clockOneTz?: string;
  @Prop() private clockThreeTz?: string;
  @Prop() private eventId?: string;

  mounted() {
    this.timerID = setInterval(this.updateTime, 1000);
    this.updateTime();
  }
  get numberTrivia(): { trivia: string } {
    return (
      getStoreItem((this.$store.state as unknown) as State, this.eventId) || {
        trivia: "Hello world"
      }
    );
  }
  updateTime() {
    const now = new Date();
    this.model = this.getModel(now);
    if (this.clockOneTz) {
      this.clockOneModel = this.getModel(now, this.clockOneTz);
    }
    if (this.clockThreeTz) {
      this.clockThreeModel = this.getModel(now, this.clockThreeTz);
    }
  }
  getModel(now: Date, tz?: string): TimeModel {
    const options = {
        weekday: "short",
        year: "numeric",
        month: "short",
        day: "numeric",
        hour12: true,
        hour: "numeric",
        minute: "numeric",
        timeZone: tz ? tz : undefined
      },
      currentTime = now.toLocaleString("en-US", options),
      pieces = currentTime.split(","),
      time = pieces.splice(pieces.length - 1, 1)[0],
      date = pieces.splice(0, pieces.length - 1).join(", ");
    return { time, date };
  }
  beforeDestroy() {
    clearInterval(this.timerID);
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.clock {
  background: #264653;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  & .clock1,
  & .clock3 {
    flex: 1 1 auto;
    text-align: center;
  }
  & .clock2 {
    flex: 2 1 auto;
    align-items: center;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: column;
  }
}
</style>
